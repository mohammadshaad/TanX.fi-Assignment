package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "price-alert-app/cmd/internal/models"
    "price-alert-app/cmd/internal/services"
    "time"
    "github.com/dgrijalva/jwt-go"
)

// Static user store for demonstration purposes
var users = map[string]string{
    "user1": "password1",
    "user2": "password2",
}

func CreateAlert(w http.ResponseWriter, r *http.Request) {
    var alert models.Alert
    if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := services.CreateAlert(alert); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if err := services.DeleteAlert(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func FetchAlerts(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    alerts, err := services.FetchAlerts(status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(alerts)
}

var secretKey = []byte("shaad")

func Login(w http.ResponseWriter, r *http.Request) {
    var user map[string]string
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    username := user["username"]
    password := user["password"]

    // Check the username and password against the static user store
    if storedPassword, ok := users[username]; ok && storedPassword == password {
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "username": username,
            "exp":      time.Now().Add(time.Hour * 1).Unix(),
        })
        tokenString, err := token.SignedString(secretKey)
        if err != nil {
            http.Error(w, "Error generating token", http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
    } else {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}
