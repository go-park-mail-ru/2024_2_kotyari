package setup

import (
	"errors"
	"net/http"
)

func (a *Attacker) getCreds() error {
	creds, err := a.getAuthToken()
	if err != nil {
		return err
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
