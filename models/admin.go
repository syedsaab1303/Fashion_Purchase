package models

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Admin model 
type Admin struct {
    ID       int
    Username string
    Password string
}

// Database connection
func ConnectDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "experiment:experiment@tcp(127.0.0.1:3306)/fashion_purchase")
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Admin register 
func (admin *Admin) Register(db *sql.DB) (int64, error) {
    query := "INSERT INTO admins (username, password) VALUES (?, ?)"
    result, err := db.Exec(query, admin.Username, admin.Password)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}


func (admin *Admin) Login(db *sql.DB) (bool, error) {
    var password string
    query := "SELECT password FROM admins WHERE username = ?"
    err := db.QueryRow(query, admin.Username).Scan(&password)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil 
        }
        return false, err 
    }
    
    if admin.Password == password {
        return true, nil
    }
    return false, nil
}




































































