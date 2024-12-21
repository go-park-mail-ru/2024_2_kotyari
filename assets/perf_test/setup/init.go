package setup

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	unauthorizedErr = errors.New("unauthorized")
)

const (
	handlers         = "handlers"
	handlerLogin     = "login"
	handlerSignup    = "signup"
	sessionId        = "session-id"
	getCsrf          = "get_csrf"
	csrfToken        = "X-Csrf-Token"
	values           = "products"
	handlerAddToCart = "add_to_cart"
)

type Attacker struct {
	client *http.Client
	viper  *viper.Viper
	url    string
	creds  Credentials
}

func NewAttacker(url string, viper *viper.Viper) *Attacker {
	client := &http.Client{}

	log.Println(url)

	return &Attacker{
		url:    url,
		viper:  viper,
		client: client,
	}
}
