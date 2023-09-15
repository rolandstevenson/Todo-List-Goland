package routes

import (
    "github.com/gorilla/mux"
    "your_project/controllers"
)

// SetRoutes menentukan semua rute API
func SetRoutes(router *mux.Router) {
    router.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
    router.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")
    // ... Tambahkan rute lainnya untuk mengedit, menghapus, dan mengambil satu tugas.
}
