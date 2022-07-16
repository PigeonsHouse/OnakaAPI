package cruds

import (
	"onaka-api/db"
	"onaka-api/types"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email string, name string, password string) types.UserResponse {
	hash_pass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u := db.User{Email: email, Name: name, PasswordHash: string(hash_pass)}
	db.Psql.Create(&u)
	return types.UserResponse{ID: u.ID, Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt, UpdatedAt: u.UpdatedAt}
}
