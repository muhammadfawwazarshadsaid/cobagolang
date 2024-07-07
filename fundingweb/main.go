package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
    connStr := "postgresql://fundingdb_owner:1UAC3qyWuYLf@ep-morning-cake-a1mxst3a.ap-southeast-1.aws.neon.tech/fundingdb?sslmode=require"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    var version string
    if err := db.QueryRow("SELECT version()").Scan(&version); err != nil {
        panic(err)
    }

    fmt.Printf("version=%s\n", version)


	var name string
    if err := db.QueryRow("SELECT name FROM Users WHERE email='admin@gmail.com'").Scan(&name); err != nil {
        panic(err)
    }

	fmt.Printf("name=%s\n", name)


    emails, err := db.Query("SELECT email FROM Users")
    if err != nil {
        panic(err)
    }
    defer emails.Close()

	for emails.Next() {
        var email string
        // Scan the 'name' column value into the 'name' variable
        if err := emails.Scan(&email); err != nil {
            panic(err)
        }
        // Process each row's data as needed
        fmt.Println("Email:", email)
    }

    // Check for errors during row iteration
    if err = emails.Err(); err != nil {
        panic(err)
    }
}
