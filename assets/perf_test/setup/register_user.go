package setup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (a *Attacker) registerUser() (Credentials, error) {
	handle := a.viper.GetStringMap(handlers)

	user := a.newUser()
	jsomBytes, err := json.Marshal(user)
	if err != nil {
		return Credentials{},
			fmt.Errorf("[ Attacker.Signup ] error when json.Marshal: %v ]", err)
	}

	req, err := http.NewRequest(http.MethodPost, a.url+handle[handlerSignup].(string), bytes.NewBuffer(jsomBytes))
	if err != nil {
		return Credentials{},
			fmt.Errorf("[ Attacker.Signup ] error when http.NewRequest: %v ]", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return Credentials{},
			fmt.Errorf("[ Attacker.Signup ] error when client.Do: %v ]", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusUnauthorized {
		return Credentials{},
			fmt.Errorf("[ Attacker.Signup ] error when client.Do: %d , error: %v ]",
				resp.StatusCode, unauthorizedErr)
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
