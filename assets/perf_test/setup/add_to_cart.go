package setup

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func (a *Attacker) addToCart() error {
	products := viper.GetIntSlice(values)

	handle := a.viper.GetStringMap(handlers)

	endpoint := a.url + handle[handlerAddToCart].(string)

	for _, product := range products {
		url := endpoint + strconv.Itoa(product)
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

		resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("unexpected status code %d \n \tcsrf: %s \n\t session: %s ",
				resp.StatusCode, a.creds.CSRFToken, a.creds.AuthToken)
		}
	}

	return nil
}
