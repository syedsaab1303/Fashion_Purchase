package controllers



import (
    "encoding/json"
    "net/http"
    "time"
    "Fashion-Purchase/models"
    "Fashion-Purchase/config" 
    _ "github.com/go-sql-driver/mysql"
)



type CartItem struct {
    UserID    int `json:"user_id"`
    ProductID int `json:"product_id"`
    Quantity  int `json:"quantity"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
    var cartItem CartItem
    err := json.NewDecoder(r.Body).Decode(&cartItem)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Database connection establish
    db := config.ConnectDB()
    defer db.Close()

    var price float64
    var stockQuantity int
    err = db.QueryRow("SELECT price, quantity FROM products WHERE id = ?", cartItem.ProductID).Scan(&price, &stockQuantity)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    if cartItem.Quantity > stockQuantity {
        http.Error(w, "Not enough stock available", http.StatusBadRequest)
        return
    }

    totalPrice := price * float64(cartItem.Quantity)

    _, err = db.Exec("INSERT INTO carts (user_id, product_id, quantity, total_price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
        cartItem.UserID, cartItem.ProductID, cartItem.Quantity, totalPrice, time.Now(), time.Now())
    if err != nil {
        http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("UPDATE products SET quantity = quantity - ? WHERE id = ?", cartItem.Quantity, cartItem.ProductID)
    if err != nil {
        http.Error(w, "Failed to update product quantity", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Product added to cart successfully"))
}


// DeleteFromCart function
func DeleteFromCart(w http.ResponseWriter, r *http.Request) {
    productID := r.URL.Query().Get("product_id")
    userID := r.URL.Query().Get("user_id")

    // Database connection
    db, err := models.ConnectDB()
    if err != nil {
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    query := `DELETE FROM carts WHERE user_id = ? AND product_id = ?`
    _, err = db.Exec(query, userID, productID)
    if err != nil {
        http.Error(w, "Failed to delete product from cart", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Product removed from cart")
}





func ViewCart(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")

    db, err := models.ConnectDB()
    if err != nil {
        http.Error(w, "Database connection failed", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    query := `SELECT product_id, quantity, total_price, created_at, updated_at FROM carts WHERE user_id = ?`
    rows, err := db.Query(query, userID)
    if err != nil {
        http.Error(w, "Failed to retrieve cart items", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var cartItems []models.Cart // Assuming Cart is defined in models with fields matching the SELECT statement
    for rows.Next() {
        var cartItem models.Cart
        err := rows.Scan(&cartItem.ProductID, &cartItem.Quantity, &cartItem.TotalPrice, &cartItem.CreatedAt, &cartItem.UpdatedAt)
        if err != nil {
            http.Error(w, "Failed to scan cart item", http.StatusInternalServerError)
            return
        }
        cartItems = append(cartItems, cartItem)
    }

   if err = rows.Err(); err != nil {
        http.Error(w, "Failed to retrieve cart items", http.StatusInternalServerError)
        return
    }

   w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cartItems)
}
