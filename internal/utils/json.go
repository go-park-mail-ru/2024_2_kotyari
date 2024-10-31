package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
)

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		response := Response{
			Status: status,
			Body:   data,
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, errs.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func WriteErrorJSON(w http.ResponseWriter, status int, errResponse error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := errs.HTTPErrorResponse{
		ErrorMessage: errResponse.Error(),
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, errs.InternalServerError.Error(), http.StatusInternalServerError)
		return
	}
}
