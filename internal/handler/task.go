package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
	"github.com/Akakazkz/go-task-manager-api/internal/service"
)
type TaskHandler interface{
	Create(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct{
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) TaskHandler{
	return &taskHandler{
		service: service,
	}
}

func (h *taskHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err !=nil{
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&task); err !=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}