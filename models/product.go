package models

import (
    "database/sql"
)

// Product struct
type Product struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Quantity    int       `json:"quantity"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}

func (p *Product) Create(db *sql.DB) (int, error) {
    query := "INSERT INTO products (name, description, price, quantity) VALUES (?, ?, ?, ?)"
    result, err := db.Exec(query, p.Name, p.Description, p.Price, p.Quantity)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return int(id), nil
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
    rows, err := db.Query("SELECT * FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var product Product
        err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}

func (p *Product) Update(db *sql.DB) error {
    query := "UPDATE products SET name = ?, description = ?, price = ?, quantity = ? WHERE id = ?"
    _, err := db.Exec(query, p.Name, p.Description, p.Price, p.Quantity, p.ID)
    return err
}


func DeleteProduct(db *sql.DB, id int) error {
    query := "DELETE FROM products WHERE id = ?"
    _, err := db.Exec(query, id)
    return err
}
