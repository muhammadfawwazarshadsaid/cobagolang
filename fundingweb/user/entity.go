package user

import (
	"time"
	
	_ "github.com/lib/pq"

)


type User struct {
	ID int 
	Name string
	Email string
	Occupation string
	PasswordHash string
	AvatarFileName string
	Role string
	Token string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(newUser *User) error {
	now := time.Now()
	newUser.CreatedAt = now
	newUser.UpdatedAt = now
if _, err := db.Exec(`; err != nil {
	return nil, err
}

        INSERT INTO users (name, email, occupation, password_hash, avatar_file_name, role, token, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
        newUser.Name, newUser.Email, newUser.Occupation, newUser.PasswordHash, newUser.AvatarFileName, newUser.Role, newUser.Token, now, now)
    
    return err
}