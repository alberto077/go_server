package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/alberto077/go_server/internal/schemas"
)

type TaskStorage struct {
	mu      *sync.Mutex
	storage map[string]schemas.TaskIn
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		mu:      &sync.Mutex{},
		storage: make(map[string]schemas.TaskIn),
	}
}

func (ts *TaskStorage) Health(w http.ResponseWriter, req *http.Request) {

	res := schemas.HealthOut{
		HealthMessage: "Server up",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Could not encode message", http.StatusInternalServerError)
		return
	}
}

func (ts *TaskStorage) CreateTask(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	var task schemas.TaskIn
	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		http.Error(w, "Could not read inputed task", http.StatusBadRequest)
		return
	}
	// need logic to insert into a database and if succesful return a TaskOut after successful insertion

	_, ok := ts.storage[task.Name]
	if ok {
		http.Error(w, "Task Already Exists", http.StatusAlreadyReported)
		return
	}

	ts.mu.Lock()
	ts.storage[task.Name] = task
	ts.mu.Unlock()

	res := struct {
		Message string
	}{
		Message: fmt.Sprintf("Task %s created successful", task.Name),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Could not encode return message", http.StatusInternalServerError)
		return
	}

}

func (ts *TaskStorage) ReadTask(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var task schemas.TaskIn

	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		http.Error(w, "Malformed Request", http.StatusBadRequest)
		return
	}

	val, ok := ts.storage[task.Name]
	if !ok {
		http.Error(w, "Task Already Exists", http.StatusAlreadyReported)
		return
	}

	if err := json.NewEncoder(w).Encode(val); err != nil {
		http.Error(w, "Trouble retuning encoding", http.StatusInternalServerError)
		return
	}

}
