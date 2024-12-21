package setup

type User struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	key      = "user"
	login    = "login"
	email    = "email"
	password = "password"
)

func (a *Attacker) newUser() User {
	u := a.viper.GetStringMap(key)

	return User{
		Login:    u[login].(string),
		Email:    u[email].(string),
		Password: u[password].(string),
	}
}
