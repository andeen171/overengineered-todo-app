package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "todo-api/internal/routes"
)

func main() {
    router := mux.NewRouter()

    routes.RegisterRoutes(router)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

