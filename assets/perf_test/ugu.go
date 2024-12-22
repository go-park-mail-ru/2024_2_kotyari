package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

const (
	handlerLogin     = "login"
	handlerSignup    = "sign_up"
	getCsrf          = "get_csrf"
	handlers         = "handlers"
	handlerAddToCart = "add_to_cart"
	handlerMakeOrder = "make_order"
	sessionId        = "session-id"
	csrfToken        = "X-Csrf-Token"
	values           = "products"
)

type Credentials struct {
	CSRFToken string
	AuthToken string
}

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type Attacker struct {
	client *http.Client
	viper  *viper.Viper
	url    string
	creds  Credentials
}

func NewAttacker(url string, viper *viper.Viper) *Attacker {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	fmt.Println(viper.GetStringMap(handlers))

	return &Attacker{
		url:    url,
		viper:  viper,
		client: client,
	}
}

func (a *Attacker) getAuthToken() (Credentials, error) {
	user := a.newUser()
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return Credentials{}, fmt.Errorf("error when json.Marshal: %v", err)
	}

	handle := a.viper.GetStringMapString(handlers)
	loginEndpoint := a.url + handle[handlerLogin]

	req, err := http.NewRequest(http.MethodPost, loginEndpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return Credentials{}, fmt.Errorf("error when http.NewRequest: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return Credentials{}, fmt.Errorf("error when client.Do: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return Credentials{}, fmt.Errorf("unexpected status code %d during login", resp.StatusCode)
	}

	creds := Credentials{}

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == sessionId {
			creds.AuthToken = cookie.Value
		}
	}

	return creds, nil
}

func (a *Attacker) registerUser() (Credentials, error) {
	user := a.newUser()
	jsonBytes, err := json.Marshal(user)

	if err != nil {
		return Credentials{}, fmt.Errorf("error when json.Marshal: %v", err)
	}

	handle := a.viper.GetStringMapString(handlers)
	signupEndpoint := a.url + handle[handlerSignup]

	req, err := http.NewRequest(http.MethodPost, signupEndpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return Credentials{}, fmt.Errorf("error when http.NewRequest: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return Credentials{}, fmt.Errorf("error when client.Do: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return Credentials{}, fmt.Errorf("error when io.ReadAll: %v", readErr)
	}

	if resp.StatusCode != http.StatusOK {
		return Credentials{}, fmt.Errorf("unexpected status code %d during signup\n\t%s\n\t%s", resp.StatusCode, signupEndpoint, string(bodyBytes))
	}

	creds := Credentials{}

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == sessionId {
			creds.AuthToken = cookie.Value
		}
	}

	return creds, nil
}

func (a *Attacker) getCreds() error {
	creds, err := a.getAuthToken()
	if err != nil {
		log.Printf("Username failed: %v. Attempting to register.", err)
		creds, err = a.registerUser()
		if err != nil {
			return fmt.Errorf("registration failed: %v", err)
		}
	}

	handle := a.viper.GetStringMap(handlers)

	req, err := http.NewRequest(http.MethodGet, a.url+handle[getCsrf].(string), nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:     sessionId,
		Value:    creds.AuthToken,
		Domain:   a.url,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	})

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}

	csrf := resp.Header.Get(csrfToken)
	if csrf == "" {
		return errors.New("no CSRF token found")
	}

	creds.CSRFToken = csrf

	a.creds = creds

	return nil
}

func (a *Attacker) newUser() User {
	u := a.viper.GetStringMapString("user")

	user := User{
		Username: u["username"],
		Email:    u["email"],
		Password: u["password"],
	}

	user.RepeatPassword = user.Password
	return user
}

func (a *Attacker) addToCart(productId int) error {
	handle := a.viper.GetStringMapString(handlers)
	endpoint := a.url + handle[handlerAddToCart]

	url := fmt.Sprintf(endpoint, productId)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:     sessionId,
		Value:    a.creds.AuthToken,
		Domain:   a.url,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	})

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(csrfToken, a.creds.CSRFToken)

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusConflict {

		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("error when io.ReadAll: %v", readErr)
		}

		return fmt.Errorf("unexpected status code %d when adding to cart\t error: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

type CreateOrderRequest struct {
	Address   string `json:"address"`
	PromoCode string `json:"promocode"`
}

func (a *Attacker) makeOrder() error {
	handle := a.viper.GetStringMapString(handlers)
	endpoint := a.url + handle[handlerMakeOrder]

	data := CreateOrderRequest{
		Address: "тест",
	}
	jsomBytes, err := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsomBytes))
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:     sessionId,
		Value:    a.creds.AuthToken,
		Domain:   a.url,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		HttpOnly: true,
	})

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(csrfToken, a.creds.CSRFToken)

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("error when io.ReadAll: %v", readErr)
		}

		return fmt.Errorf("unexpected status code %d when creating order\t msg: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func runLoadTest(a *Attacker, products []int, concurrency int, iterations int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	start := time.Now()
	successfulOrders := atomic.Int32{}
	failedOrders := atomic.Int32{}

	for i := 0; i < iterations; i++ {
		for _, product := range products {
			wg.Add(1)
			sem <- struct{}{}

			go func(iteration int, product int) {
				defer wg.Done()
				defer func() { <-sem }()

				if err := a.addToCart(product); err != nil {
					log.Printf("Iteration %d: Error adding product %d to cart: %v", iteration+1, product, err)
					failedOrders.Add(1)
					return
				}

				if err := a.makeOrder(); err != nil {
					log.Printf("Iteration %d: Error creating order for product %d: %v", iteration+1, product, err)
					failedOrders.Add(1)
					return
				}

				successfulOrders.Add(1)
				log.Printf("Iteration %d: Successfully created order for product %d", iteration, product)
			}(i, product)
		}
	}

	wg.Wait()
	duration := time.Since(start)

	log.Printf("Load Test Completed")
	log.Printf("Total Iterations: %d", iterations)
	log.Printf("Products per Iteration: %d", len(products))
	log.Printf("Total Orders Attempted: %d", iterations*len(products))
	log.Printf("Successful Orders: %d", successfulOrders.Load())
	log.Printf("Failed Orders: %d", failedOrders.Load())
	log.Printf("Total Duration: %v", duration)
	log.Printf("Throughput: %.2f orders/sec", float64(successfulOrders.Load())/duration.Seconds())
}

func main() {
	viper.AddConfigPath("assets/perf_test")
	viper.SetConfigName("setup_test")

	v := viper.GetViper()
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	url := v.GetString("domain")
	a := NewAttacker(url, v)

	if err := a.getCreds(); err != nil {
		log.Fatalf("Error getting credentials: %v", err)
	}

	products := v.GetIntSlice(values)
	if len(products) == 0 {
		log.Fatalf("No products found in configuration")
	}

	concurrency := v.GetInt("load_test.concurrency")
	iterations := v.GetInt("load_test.iterations")

	log.Printf("Starting Load Test with %d concurrency and %d iterations", concurrency, iterations)
	runLoadTest(a, products, concurrency, iterations)
}
