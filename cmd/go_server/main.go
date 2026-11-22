package main

import (
	"fmt"
	"net/http"

	"github.com/alberto077/go_server/internal/api"
)

func main() {

	storage := api.NewTaskStorage()

	http.HandleFunc("/health", storage.Health)
	http.HandleFunc("POST /create", storage.CreateTask)
	http.HandleFunc("GET /read", storage.ReadTask)

	fmt.Println("Serving listening on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
