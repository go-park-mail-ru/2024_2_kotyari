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

// Credentials хранит CSRF токен и Auth токен пользователя
type Credentials struct {
	CSRFToken string
	AuthToken string
}

// User представляет структуру пользователя
type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

// Attacker представляет отдельного пользователя
type Attacker struct {
	userID string // Уникальный идентификатор пользователя
	client *http.Client
	viper  *viper.Viper
	url    string
	creds  Credentials
}

// NewAttacker создает нового Attacker с уникальным userID
func NewAttacker(url string, v *viper.Viper, userID string) *Attacker {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//fmt.Println(v.GetStringMap(handlers))

	return &Attacker{
		userID: userID, // Инициализация userID
		url:    url,
		viper:  v,
		client: client,
	}
}

// Metrics собирает и хранит метрики нагрузки
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

// NewMetrics инициализирует новую структуру Metrics
func NewMetrics() *Metrics {
	return &Metrics{
		StatusCodes: make(map[int]int64),
		ErrorSet:    make(map[string]struct{}),
		LatencyMin:  time.Hour, // Инициализируем большим значением
	}
}

// RecordRequest записывает данные о запросе в метрики
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

// PrintReport выводит отчет по собранным метрикам
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

// getAuthToken выполняет попытку входа пользователя и получает Auth токен
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

// registerUser регистрирует нового пользователя и получает Auth токен
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

// getCreds получает учетные данные пользователя, регистрируя его при необходимости
func (a *Attacker) getCreds(metrics *Metrics) error {
	creds, err := a.getAuthToken(metrics)
	if err != nil {
		//("Username %s failed: %v. Attempting to register.", a.userID, err)
		creds, err = a.registerUser(metrics)
		if err != nil {
			return fmt.Errorf("registration failed for %s: %v", a.userID, err)
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

// newUser генерирует нового пользователя с уникальными данными
func (a *Attacker) newUser() User {
	u := a.viper.GetStringMapString("user")

	user := User{
		Username: fmt.Sprintf("%s_%s", u["username"], a.userID),          // Генерация уникального имени пользователя
		Email:    fmt.Sprintf("%s_%s@example.com", u["email"], a.userID), // Генерация уникального email
		Password: u["password"],
	}

	user.RepeatPassword = user.Password

	//("username: %s, email %s password: %s\n", user.Username, user.Email, user.Password)

	return user
}

// addToCart добавляет продукт в корзину пользователя
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

// CreateOrderRequest представляет структуру запроса для создания заказа
type CreateOrderRequest struct {
	Address   string `json:"address"`
	PromoCode string `json:"promocode"`
}

// makeOrder создает заказ для пользователя
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

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code %d when creating order\t msg: %s", resp.StatusCode, string(bodyBytes))
		metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, err)
		return err
	}

	metrics.RecordRequest(resp.StatusCode, bytesIn, bytesOut, latency, nil)
	return nil
}

// runLoadTest выполняет нагрузочное тестирование для всех пользователей
func runLoadTest(attackers []*Attacker, products []int, concurrency int, iterations int, metrics *Metrics, rps int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	// Инициализируем ограничитель скорости
	limiter := rate.NewLimiter(rate.Limit(rps), rps)

	startTime := time.Now()

	for i := 0; i < iterations; i++ {
		for _, product := range products {
			for _, attacker := range attackers {
				wg.Add(1)
				sem <- struct{}{}

				// Перед запуском горутины проверяем ограничитель скорости
				if err := limiter.Wait(context.Background()); err != nil {
					//("Limiter error: %v", err)
					<-sem
					wg.Done()
					continue
				}

				go func(a *Attacker, product int, iteration int) {
					defer wg.Done()
					defer func() { <-sem }()

					// Добавить в корзину
					if err := a.addToCart(product, metrics); err != nil {
						// Можно раскомментировать для детального логирования ошибок
						//("Iteration %d: Error adding product %d to cart for user %s: %v", iteration+1, product, a.userID, err)
						return
					}

					// Создать заказ
					if err := a.makeOrder(product, metrics); err != nil {
						// Можно раскомментировать для детального логирования ошибок
						//log.Printf("Iteration %d: Error creating order for product %d for user %s: %v", iteration+1, product, a.userID, err)
						return
					}

					// Можно добавить логирование успешных операций
					//log.Printf("Iteration %d: Successfully created order for product %d for user %s", iteration, product, a.userID)
				}(attacker, product, i)
			}
		}
	}

	wg.Wait()
	totalDuration := time.Since(startTime)
	metrics.PrintReport(totalDuration)
}

func main() {
	// Настройка Viper для чтения конфигурации
	viper.AddConfigPath("assets/perf_test")
	viper.SetConfigName("setup_test")

	v := viper.GetViper()
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	url := v.GetString("domain")

	// Чтение количества пользователей из конфигурации
	numUsers := v.GetInt("load_test.users")
	if numUsers <= 0 {
		numUsers = 1 // По умолчанию 1 пользователь, если не задано
	}

	// Создание списка Attacker для каждого пользователя
	attackers := make([]*Attacker, numUsers)
	for i := 0; i < numUsers; i++ {
		userID := fmt.Sprintf("user%d", i+1)
		attackers[i] = NewAttacker(url, v, userID)
	}

	metrics := NewMetrics()

	var wg sync.WaitGroup
	// Получение учетных данных для всех пользователей параллельно
	for _, attacker := range attackers {
		wg.Add(1)
		go func(a *Attacker) {
			defer wg.Done()
			if err := a.getCreds(metrics); err != nil {
				//log.Printf("Error getting credentials for %s: %v", a.userID, err)
			}
		}(attacker)
	}
	wg.Wait()

	products := v.GetIntSlice(values)
	if len(products) == 0 {
		log.Fatalf("No products found in configuration")
	}

	concurrency := v.GetInt("load_test.concurrency")
	iterations := v.GetInt("load_test.iterations")
	rps := v.GetInt("load_test.rps")

	log.Printf("Starting Load Test with %d users, %d concurrency, %d iterations and %d RPS", numUsers, concurrency, iterations, rps)
	runLoadTest(attackers, products, concurrency, iterations, metrics, rps)
}
