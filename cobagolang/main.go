package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID             int            `db:"id"`
	Name           string         `db:"name"`
	Email          string         `db:"email"`
	Occupation     string         `db:"occupation"`
	PasswordHash   string         `db:"password_hash"`
	AvatarFilename string         `db:"avatar_file_name"`
	Role           string         `db:"role"`
	CreatedAt      sql.NullTime   `db:"created_at"` // Use sql.NullTime for nullable time fields
	UpdatedAt      sql.NullTime   `db:"updated_at"` // Use sql.NullTime for nullable time fields
	Token          sql.NullString `db:"token"`
}

func main() {
	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}

func handler(c *gin.Context) {
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

	rows, err := db.Query("SELECT id, name, email, occupation, password_hash, role, token, created_at, updated_at, avatar_file_name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Occupation, &user.PasswordHash, &user.Role, &user.Token, &user.CreatedAt, &user.UpdatedAt, &user.AvatarFilename)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}
