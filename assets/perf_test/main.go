package main

import (
	"bytes"
	"github.com/go-park-mail-ru/2024_2_kotyari/assets/perf_test/setup"
	viperLib "github.com/spf13/viper"
	veg "github.com/tsenart/vegeta/lib"
	"log"
	"mime/multipart"
)

const (
	domain = "domain"
)

func main() {
	viperLib.AddConfigPath("assets/perf_test")
	viperLib.SetConfigName("setup_test")

	viper := viperLib.GetViper()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	url := viper.GetString(domain)

	a := setup.NewAttacker(url, viper)

	_ = a

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := writer.SetBoundary("vegetaboundary")
	if err != nil {
		log.Fatal(err)
	}

	attacker := veg.NewAttacker()

	_ = attacker
}
