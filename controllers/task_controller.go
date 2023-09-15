package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "your_project/models"
)

var tasks []models.Task
var taskID = 1

// CreateTask membuat tugas baru
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var newTask models.Task
    _ = json.NewDecoder(r.Body).Decode(&newTask)
    newTask.ID = taskID
    taskID++
    tasks = append(tasks, newTask)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newTask)
}

// GetTasks mengambil semua tugas
func GetTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// ... Tambahkan fungsi lainnya untuk mengedit, menghapus, dan mengambil satu tugas.
