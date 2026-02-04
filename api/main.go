package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TaskRequest struct {
	ID     string `json:"id"`
	Prompt string `json:"prompt"`
}

type TaskResponse struct {
	Status string `json:"status"`
	TaskID string `json:"task_id"`
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req TaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := TaskResponse{
		Status: "accepted",
		TaskID: req.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/task", taskHandler)

	log.Println("PIXN API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
