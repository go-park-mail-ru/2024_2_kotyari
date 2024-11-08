package utils

import "time"

func IsExpired(exp int64) bool {
	return exp < time.Now().Unix()
}
