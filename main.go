package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":4000",
		Handler: routes(),
	}
	fmt.Println("Запуск сервера на :4000")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
