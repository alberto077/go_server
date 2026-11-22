package main

import (
	"fmt"
	"net/http"

	"github.com/alberto077/go_server/internal/api"
)

func main() {
	http.HandleFunc("/health", api.Health)

	fmt.Println("Serving listening on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
