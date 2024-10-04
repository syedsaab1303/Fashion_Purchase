package models

import (
    "database/sql"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"` // Password ko hash karna chahiye
}

func (u *User) Register(db *sql.DB) (int, error) {
    query := "INSERT INTO users (username, password) VALUES (?, ?)"
    result, err := db.Exec(query, u.Username, u.Password)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return int(id), nil
}


func (u *User) Login(db *sql.DB) (bool, error) {
    query := "SELECT id FROM users WHERE username = ? AND password = ?"
    row := db.QueryRow(query, u.Username, u.Password)
    err := row.Scan(&u.ID)
    if err != nil {
        return false, err
    }
    return true, nil
}
