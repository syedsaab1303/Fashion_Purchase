package controllers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "Fashion-Purchase/models"
)

// User Signup function
func UserSignup(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for User Signup")

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Database connection
    db, err := models.ConnectDB()
    if err != nil {
        log.Println(err)
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = user.Register(db)
    if err != nil {
        log.Println(err)
        http.Error(w, "Registration failed", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "User registered successfully")
}

// User Login function
func UserLogin(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for User Login")

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Database connection
    db, err := models.ConnectDB()
    if err != nil {
        log.Println(err)
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    success, err := user.Login(db)
    if err != nil {
        log.Println(err)
        http.Error(w, "Login failed", http.StatusInternalServerError)
        return
    }

    if !success {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(w, "User logged in successfully")
}
