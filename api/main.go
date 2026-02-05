package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
)

type Task struct {
    ID     string json:"id"
    Prompt string json:"prompt"
    Status string json:"status"
}

var (
    tasks = make(map[string]*Task)
    mu    sync.Mutex
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var t Task
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    mu.Lock()
    t.Status = "accepted"
    tasks[t.ID] = &t
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "accepted",
        "task_id": t.ID,
    })
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    mu.Lock()
    task, ok := tasks[id]
    mu.Unlock()

    if !ok {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(task)
}

func main() {
    http.HandleFunc("/task", taskHandler)
    http.HandleFunc("/status", statusHandler)

    log.Println("PIXN API running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
