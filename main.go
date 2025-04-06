package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/book-wise/config"
	"github.com/book-wise/router"
)

func main() {
	config.Load()
	r := router.NewRouter()

	fmt.Printf("Starting api in port: %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
