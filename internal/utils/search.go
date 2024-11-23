package utils

import "net/http"

const (
	searchQueryParamKey = "q"
)

func GetSearchQuery(r *http.Request) string {
	return r.URL.Query().Get(searchQueryParamKey)
}
