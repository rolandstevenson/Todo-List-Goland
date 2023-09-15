package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "your_project/routes"
)

func main() {
    router := mux.NewRouter()
    routes.SetRoutes(router)

    log.Fatal(http.ListenAndServe(":8080", router))
}
