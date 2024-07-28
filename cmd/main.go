package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "price-alert-app/cmd/internal/handlers"
    "price-alert-app/cmd/internal/middleware"
    "price-alert-app/cmd/internal/config"
    "price-alert-app/pkg/websocket"
)

func main() {
    config.InitDB()
    go websocket.ConnectToBinance()

    r := mux.NewRouter()
    r.HandleFunc("/alerts/create", handlers.CreateAlert).Methods("POST")
    r.HandleFunc("/alerts/delete/{id}", handlers.DeleteAlert).Methods("DELETE")
    r.HandleFunc("/alerts/fetch", handlers.FetchAlerts).Methods("GET")
    r.HandleFunc("/login", handlers.Login).Methods("POST")

    r.Use(middleware.JWTMiddleware) // Apply middleware after defining routes

    log.Fatal(http.ListenAndServe(":8080", r))
}
