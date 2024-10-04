package routes

import (
    "Fashion-Purchase/controllers"
    "net/http"
)


func RegisterProductRoutes() {
    http.HandleFunc("/admin/products", controllers.CreateProduct) // Create Product
    http.HandleFunc("/admin/products/all", controllers.GetAllProducts) // Get All Products
    http.HandleFunc("/admin/products/update", controllers.UpdateProduct) // Update Product
    http.HandleFunc("/admin/products/delete", controllers.DeleteProduct) // Delete Product
}
