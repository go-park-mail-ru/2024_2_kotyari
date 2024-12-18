package csrf

import (
	"testing"
	"time"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCscfUsecase_CreateCsrfToken(t *testing.T) {
	type want struct {
		token string
		err   error
	}

	var tests = []struct {
		name    string
		session model.Session
		now     time.Time
		secret  string
		want    want
	}{
		{
			name: "Успешное создание токена",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			now:    time.Now(),
			secret: "secret",
			want: want{
				token: "some_token",
				err:   nil,
			},
		},
		{
			name: "Секретный ключ не задан",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			now:    time.Now(),
			secret: "",
			want: want{
				token: "some_token",
				err:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csrf := NewCscfUsecase()
			csrf.secret = tt.secret

			token, err := csrf.CreateCsrfToken(tt.session, tt.now)
			assert.NotNil(t, token)
			assert.NoError(t, err)
		})
	}
}
