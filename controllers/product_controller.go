package controllers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "Fashion-Purchase/models"

)


// Create Product function
func CreateProduct(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for Create Product")

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var product models.Product
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    db, err := models.ConnectDB()
    if err != nil {
        log.Println(err)
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = product.Create(db)
    if err != nil {
        log.Println(err)
        http.Error(w, "Product creation failed", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Product created successfully")
}


//  GetAllProducts function

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for all products")

   
    db, err := models.ConnectDB()
    if err != nil {
        log.Println(err)
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, name, description, price, quantity, created_at, updated_at FROM products")
    if err != nil {
        log.Println(err)
        http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        var createdAt, updatedAt string // Temporary variables for scanning

        err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &createdAt, &updatedAt)
        if err != nil {
            log.Println(err)
            http.Error(w, "Failed to scan product", http.StatusInternalServerError)
            return
        }

        product.CreatedAt = createdAt
        product.UpdatedAt = updatedAt

        products = append(products, product)
    }

   w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

// Update Product function
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for Update Product")

    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var product models.Product
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    db, err := models.ConnectDB()
    if err != nil {
        log.Println(err)
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

   err = product.Update(db)
    if err != nil {
        log.Println(err)
        http.Error(w, "Product update failed", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Product updated successfully")
}


func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for Delete Product")

    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing product ID", http.StatusBadRequest)
        return
    }

    productId, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
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

   err = models.DeleteProduct(db, productId)
    if err != nil {
        log.Println(err)
        http.Error(w, "Product deletion failed", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Product deleted successfully")
}









































