package csrf

import (
	"errors"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCscfUsecase_IsValidCSRFToken(t *testing.T) {
	type want struct {
		valid bool
		err   error
	}

	var tests = []struct {
		name    string
		session model.Session
		token   string
		secret  string
		want    want
	}{
		{
			name: "Валидный токен",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			token:  "87ed91f66860957531f52a1e6da08c844c0a65ed75f2034d6ad62c6744f70928:1734532582",
			secret: "secret",
			want: want{
				valid: true,
				err:   nil,
			},
		},
		{
			name: "Время жизни токена закончилось",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			token:  "87ed91f66860957531f52a1e6da08c844c0a65ed75f2034d6ad62c6744f70928:0",
			secret: "secret",
			want: want{
				valid: false,
				err:   ErrTokenExpired,
			},
		},
		{
			name: "Невалидный формат токена",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			token:  "invalid_token",
			secret: "secret",
			want: want{
				valid: false,
				err:   errors.New("невалидный токен"),
			},
		},
		{
			name: "Невалидное время токена",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			token:  "87ed91f66860957531f52a1e6da08c844c0a65ed75f2034d6ad62c6744f70928:invalid_time",
			secret: "secret",
			want: want{
				valid: false,
				err:   errors.New("невалидный токен"),
			},
		},
		{
			name: "Невалидный MAC",
			session: model.Session{
				SessionID: "test",
				UserID:    1,
			},
			token:  "invalid_mac:1734532582",
			secret: "secret",
			want: want{
				valid: false,
				err:   errors.New("невалидный токен"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csrf := NewCscfUsecase()
			csrf.secret = tt.secret

			valid, err := csrf.IsValidCSRFToken(tt.session, tt.token)
			assert.Equal(t, tt.want.valid, valid)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
