package setup

import "errors"

func (a *Attacker) getAuthToken() (Credentials, error) {
	creds, err := a.loginUser()
	if err != nil {
		if errors.Is(err, unauthorizedErr) {
			creds, err = a.registerUser()
			if err != nil {
				return creds, err
			}
		}

		return creds, err
	}

	return creds, nil
}
