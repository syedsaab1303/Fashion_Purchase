package routes

import (
    "Fashion-Purchase/controllers"
    "net/http"
)

func RegisterUserRoutes() {
    http.HandleFunc("/user/signup", controllers.UserSignup)
    http.HandleFunc("/user/login", controllers.UserLogin)
}
