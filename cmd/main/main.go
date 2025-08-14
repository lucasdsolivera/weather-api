package main

import (
	"fmt"
	"net/http"

	"github.com/lucasdsolivera/weather-api/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server started on localhost%s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
