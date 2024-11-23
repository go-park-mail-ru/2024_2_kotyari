package main

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/configs/postgres"
	"log"
)

func main() {
	_, err := postgres.LoadPgxPool(postgres.MainDBCFG)
	if err != nil {
		log.Fatal(err)
	}
}
