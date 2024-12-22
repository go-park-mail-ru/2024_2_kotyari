package main

import (
	"bytes"
	"context"
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
	"golang.org/x/time/rate"
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

func NewAttacker(url string, v *viper.Viper) *Attacker {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	fmt.Println(v.GetStringMap(handlers))

	return &Attacker{
		url:    url,
		viper:  v,
		client: client,
	}
}

type Metrics struct {
	TotalRequests   int64
	SuccessRequests int64
	FailedRequests  int64
	TotalBytesIn    int64
	TotalBytesOut   int64
	StatusCodes     map[int]int64
	LatencySum      time.Duration
	LatencyCount    int64
	LatencyMin      time.Duration
	LatencyMax      time.Duration
	ErrorSet        map[string]struct{}
	mutex           sync.Mutex
}

func NewMetrics() *Metrics {
	return &Metrics{
		StatusCodes: make(map[int]int64),
		ErrorSet:    make(map[string]struct{}),
		LatencyMin:  time.Hour, // Инициализируем большим значением
	}
}

func (m *Metrics) RecordRequest(statusCode int, bytesIn, bytesOut int64, latency time.Duration, err error) {
	atomic.AddInt64(&m.TotalRequests, 1)
	atomic.AddInt64(&m.TotalBytesIn, bytesIn)
	atomic.AddInt64(&m.TotalBytesOut, bytesOut)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.StatusCodes[statusCode]++
	m.LatencySum += latency
	m.LatencyCount++
	if latency < m.LatencyMin {
		m.LatencyMin = latency
	}
	if latency > m.LatencyMax {
		m.LatencyMax = latency
	}

	if err != nil {
		atomic.AddInt64(&m.FailedRequests, 1)
		m.ErrorSet[err.Error()] = struct{}{}
	} else {
		atomic.AddInt64(&m.SuccessRequests, 1)
	}
}

func (m *Metrics) PrintReport(duration time.Duration) {
	fmt.Println("----- Load Test Report -----")
	fmt.Printf("Requests\t[total, rate, throughput]\t%d, %.2f, %.2f\n",
		m.TotalRequests,
		float64(m.TotalRequests)/duration.Seconds(),
		float64(m.SuccessRequests)/duration.Seconds(),
	)
	fmt.Printf("Duration\t[total, attack]\t%s, %s\n",
		duration.String(),
		duration.String(),
	)
	meanLatency := time.Duration(0)
	if m.LatencyCount > 0 {
		meanLatency = m.LatencySum / time.Duration(m.LatencyCount)
	}
	fmt.Printf("Latencies\t[mean, min, max]\t%s, %s, %s\n",
		meanLatency,
		m.LatencyMin,
		m.LatencyMax,
	)
	fmt.Printf("Bytes In\t[total, mean]\t%d, %.2f\n",
		m.TotalBytesIn,
		float64(m.TotalBytesIn)/float64(m.TotalRequests),
	)
	fmt.Printf("Bytes Out\t[total, mean]\t%d, %.2f\n",
		m.TotalBytesOut,
		float64(m.TotalBytesOut)/float64(m.TotalRequests),
	)
	successRatio := 0.0
	if m.TotalRequests > 0 {
		successRatio = (float64(m.SuccessRequests) / float64(m.TotalRequests)) * 100
	}
	fmt.Printf("Success\t[ratio]\t%.2f%%\n", successRatio)
	fmt.Printf("Status Codes\t[code:count]\t")
	for code, count := range m.StatusCodes {
		fmt.Printf("%d:%d ", code, count)
	}
	fmt.Println()
	fmt.Printf("Error Set:\n")
	for err := range m.ErrorSet {
		fmt.Printf("\t%s\n", err)
	}
	fmt.Println("----------------------------")
}

func (a *Attacker) getAuthToken(metrics *Metrics) (Credentials, error) {
	user := a.newUser()
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return Credentials{}, fmt.Errorf("error when json.Marshal: %v", err)
	}

	handle := a.viper.GetStringMapString(handlers)
	loginEndpoint := a.url + handle[handlerLogin]

	start := time.Now()
	req, err := http.NewRequest(http.MethodPost, loginEndpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return Credentials{}, fmt.Errorf("error when http.NewRequest: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	latency := time.Since(start)
	if err != nil {
		metrics.RecordRequest(0, 0, 0, latency, err)
		return Credentials{}, fmt.Errorf("error when client.Do: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	bytesIn := int64(len(bodyBytes))
	bytesOut := int64(len(jsonBytes))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("unexpected status code %d during login", resp.StatusCode)
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, err)
		return Credentials{}, err
	}

	creds := Credentials{}

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == sessionId {
			creds.AuthToken = cookie.Value
		}
	}

	metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, nil)
	return creds, nil
}

func (a *Attacker) registerUser(metrics *Metrics) (Credentials, error) {
	user := a.newUser()
	jsonBytes, err := json.Marshal(user)

	if err != nil {
		return Credentials{}, fmt.Errorf("error when json.Marshal: %v", err)
	}

	handle := a.viper.GetStringMapString(handlers)
	signupEndpoint := a.url + handle[handlerSignup]

	start := time.Now()
	req, err := http.NewRequest(http.MethodPost, signupEndpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return Credentials{}, fmt.Errorf("error when http.NewRequest: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	latency := time.Since(start)
	if err != nil {
		metrics.RecordRequest(0, 0, 0, latency, err)
		return Credentials{}, fmt.Errorf("error when client.Do: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	bytesIn := int64(len(bodyBytes))
	bytesOut := int64(len(jsonBytes))

	if readErr != nil {
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, readErr)
		return Credentials{}, fmt.Errorf("error when io.ReadAll: %v", readErr)
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code %d during signup\n\t%s\n\t%s", resp.StatusCode, signupEndpoint, string(bodyBytes))
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, err)
		return Credentials{}, err
	}

	creds := Credentials{}

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == sessionId {
			creds.AuthToken = cookie.Value
		}
	}

	metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, nil)
	return creds, nil
}

func (a *Attacker) getCreds(metrics *Metrics) error {
	creds, err := a.getAuthToken(metrics)
	if err != nil {
		log.Printf("Username failed: %v. Attempting to register.", err)
		creds, err = a.registerUser(metrics)
		if err != nil {
			return fmt.Errorf("registration failed: %v", err)
		}
	}

	handle := a.viper.GetStringMapString(handlers)

	start := time.Now()
	req, err := http.NewRequest(http.MethodGet, a.url+handle[getCsrf], nil)
	if err != nil {
		metrics.RecordRequest(0, 0, 0, 0, err)
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
	latency := time.Since(start)
	if err != nil {
		metrics.RecordRequest(0, 0, 0, latency, err)
		return err
	}
	defer resp.Body.Close()

	csrf := resp.Header.Get(csrfToken)
	bytesIn := int64(len(csrf))
	bytesOut := int64(0)

	if csrf == "" {
		err = errors.New("no CSRF token found")
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, err)
		return err
	}

	creds.CSRFToken = csrf

	a.creds = creds

	metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, nil)
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

func (a *Attacker) addToCart(productId int, metrics *Metrics) error {
	handle := a.viper.GetStringMapString(handlers)
	endpoint := a.url + handle[handlerAddToCart]

	url := fmt.Sprintf(endpoint, productId)

	start := time.Now()
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		metrics.RecordRequest(0, 0, 0, 0, err)
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
	latency := time.Since(start)
	bytesIn := int64(0)
	bytesOut := int64(0)
	if err != nil {
		metrics.RecordRequest(0, bytesIn, bytesOut, latency, err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	bytesIn = int64(len(bodyBytes))

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusConflict {
		err = fmt.Errorf("unexpected status code %d when adding to cart\t error: %s", resp.StatusCode, string(bodyBytes))
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, err)
		return err
	}

	metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, nil)
	return nil
}

type CreateOrderRequest struct {
	Address   string `json:"address"`
	PromoCode string `json:"promocode"`
}

func (a *Attacker) makeOrder(productId int, metrics *Metrics) error {
	handle := a.viper.GetStringMapString(handlers)
	endpoint := a.url + handle[handlerMakeOrder]

	data := CreateOrderRequest{
		Address: "тест",
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		metrics.RecordRequest(0, 0, 0, 0, err)
		return err
	}

	start := time.Now()
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		metrics.RecordRequest(0, 0, 0, 0, err)
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
	latency := time.Since(start)
	bytesIn := int64(0)
	bytesOut := int64(len(jsonBytes))
	if err != nil {
		metrics.RecordRequest(0, bytesIn, bytesOut, latency, err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	bytesIn = int64(len(bodyBytes))

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		err = fmt.Errorf("unexpected status code %d when creating order\t msg: %s", resp.StatusCode, string(bodyBytes))
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, err)
		return err
	}

	metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, nil)
	return nil
}

func runLoadTest(a *Attacker, products []int, concurrency int, iterations int, metrics *Metrics, rps int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	// Инициализируем ограничитель скорости
	limiter := rate.NewLimiter(rate.Limit(rps), rps)

	startTime := time.Now()
	for i := 0; i < iterations; i++ {
		for _, product := range products {
			wg.Add(1)
			sem <- struct{}{}

			// Перед запуском горутины проверяем ограничитель скорости
			if err := limiter.Wait(context.Background()); err != nil {
				log.Printf("Limiter error: %v", err)
				<-sem
				wg.Done()
				continue
			}

			go func(iteration int, product int) {
				defer wg.Done()
				defer func() { <-sem }()

				// Добавить в корзину
				if err := a.addToCart(product, metrics); err != nil {
					//log.Printf("Iteration %d: Error adding product %d to cart: %v", iteration+1, product, err)
					return
				}

				// Создать заказ
				if err := a.makeOrder(product, metrics); err != nil {
					//log.Printf("Iteration %d: Error creating order for product %d: %v", iteration+1, product, err)
					return
				}

				//log.Printf("Iteration %d: Successfully created order for product %d", iteration, product)
			}(i, product)
		}
	}

	wg.Wait()
	totalDuration := time.Since(startTime)
	metrics.PrintReport(totalDuration)
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

	metrics := NewMetrics()

	if err := a.getCreds(metrics); err != nil {
		log.Fatalf("Error getting credentials: %v", err)
	}

	products := v.GetIntSlice(values)
	if len(products) == 0 {
		log.Fatalf("No products found in configuration")
	}

	concurrency := v.GetInt("load_test.concurrency")
	iterations := v.GetInt("load_test.iterations")
	rps := v.GetInt("load_test.rps")

	log.Printf("Starting Load Test with %d concurrency, %d iterations and %d RPS", concurrency, iterations, rps)
	runLoadTest(a, products, concurrency, iterations, metrics, rps)
}
