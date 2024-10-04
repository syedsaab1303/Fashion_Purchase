package routes

import (
    "Fashion-Purchase/controllers"
    "net/http"
)





func RegisterAdminRoutes(authKey string) {
    http.HandleFunc("/admin/signup", controllers.AdminSignup)
    http.HandleFunc("/admin/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") != authKey {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        controllers.AdminLogin(w, r)
    }) 
}



































