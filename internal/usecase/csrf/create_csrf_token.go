package csrf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"strconv"
	"time"
)

const lifeTime = time.Hour * 5

func (csrf *CscfUsecase) CreateCsrfToken(session model.Session, now time.Time) (string, error) {
	var (
		h    = hmac.New(sha256.New, []byte(csrf.secret))
		exp  = now.Add(lifeTime).Unix()
		data = fmt.Sprintf("%s:%d:%d", session.SessionID, session.UserID, exp)
	)

	h.Write([]byte(data))
	token := hex.EncodeToString(h.Sum(nil)) + ":" + strconv.FormatInt(exp, 10)

	return token, nil
}
