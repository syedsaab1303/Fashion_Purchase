package main

import (
    "Fashion-Purchase/routes"
    "log"
    "net/http"
)

const authKey = "tam123098alisyed0981303"

func main() {
    routes.RegisterAdminRoutes(authKey)
    routes.RegisterUserRoutes()
    routes.RegisterProductRoutes()
    routes.SetupRoutes() 

    log.Println("Server is running on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
