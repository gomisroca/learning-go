package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt string
}

func InsertUser(db *sql.DB, name, email string) (int64, error) {
    query := `INSERT INTO users (name, email) VALUES (?, ?)`
    res, err := db.Exec(query, name, email)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

func GetUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query("SELECT id, name, email, created_at FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }

    return users, rows.Err()
}

func main() {
	// Connect to DB
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/testdb?parseTime=true")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    fmt.Println("Database connection successful!")
	
    // Create table if not exists
    createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`
    if _, err = db.Exec(createTable); err != nil {
        panic(err)
    }

    // Insert a user
    userID, err := InsertUser(db, "Bob", "bob@example.com")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Inserted user with ID %d\n", userID)

    // Fetch users
    users, err := GetUsers(db)
    if err != nil {
        panic(err)
    }

    fmt.Println("Users in database:")
    for _, u := range users {
        fmt.Printf("ID=%d, Name=%s, Email=%s, CreatedAt=%s\n",
            u.ID, u.Name, u.Email, u.CreatedAt)
    }

}