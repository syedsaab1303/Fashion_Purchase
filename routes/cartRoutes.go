package routes

import (
    "github.com/gorilla/mux"
    "Fashion-Purchase/controllers" 
)

// SetupRoutes function
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Cart routes
    router.HandleFunc("/cart/add", controllers.AddToCart).Methods("POST")
    router.HandleFunc("/cart/delete", controllers.DeleteFromCart).Methods("DELETE")
    router.HandleFunc("/cart/view", controllers.ViewCart).Methods("GET") 

    return router
}
