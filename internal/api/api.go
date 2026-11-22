package api

import (
	"encoding/json"
	"net/http"

	"github.com/alberto077/go_server/internal/schemas"
)

func Health(w http.ResponseWriter, req *http.Request) {

	res := schemas.HealthOut{
		HealthMessage: "Server up",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Could not encode message", http.StatusInternalServerError)
		return
	}
}
