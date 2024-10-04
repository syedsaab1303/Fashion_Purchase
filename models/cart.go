package models

// Cart model
type Cart struct {
    ID         int     `json:"id"`
    UserID     int     `json:"user_id"`
    ProductID  int     `json:"product_id"`
    Quantity   int     `json:"quantity"`
    TotalPrice float64 `json:"total_price"`
    CreatedAt  string  `json:"created_at"`
    UpdatedAt  string  `json:"updated_at"`
}
