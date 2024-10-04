package controllers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "Fashion-Purchase/models"
)

// Admin signup function
func AdminSignup(w http.ResponseWriter, r *http.Request) {
    var admin models.Admin

    err := json.NewDecoder(r.Body).Decode(&admin)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Database conection
    db, err := models.ConnectDB()
    if err != nil {
        log.Println(err)
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    id, err := admin.Register(db)
    if err != nil {
        log.Println(err)
        http.Error(w, "Registration failed", http.StatusInternalServerError)
        return
    }


    fmt.Fprintf(w, "Admin registered successfully with ID: %d", id)
}



func AdminLogin(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for Admin Login")

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var admin models.Admin
    err := json.NewDecoder(r.Body).Decode(&admin)
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

   success, err := admin.Login(db)
    if err != nil {
        log.Println(err)
        http.Error(w, "Login failed", http.StatusInternalServerError)
        return
    }

    if !success {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    fmt.Fprintf(w, "Admin logged in successfully")
}



