package cruds

import (
	"onaka-api/db"
	"onaka-api/types"
	"onaka-api/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email string, name string, password string) (types.UserResponse, error) {
	if err := db.Psql.Where("email = ?", email).First(&db.User{}).Error; err == nil {
		return types.UserResponse{}, err
	}

	hash_pass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u := db.User{Email: email, Name: name, PasswordHash: string(hash_pass)}
	db.Psql.Create(&u)
	res_user := types.UserResponse{ID: u.ID, Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt, UpdatedAt: u.UpdatedAt}
	return res_user, nil
}

func GenerateJWT(email string, password string) (jwtInfo types.JWTInfo, err error) {
	var (
		u     db.User
		token string
	)

	if err = db.Psql.Where("email = ?", email).First(&u).Error; err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return
	}

	token, err = generateToken(u.ID)
	if err != nil {
		return
	}

	jwtInfo = types.JWTInfo{JWT: token}
	return
}

func generateToken(userID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(7 * 24 * time.Hour).Unix(),
	})
	return token.SignedString(utils.SigningKey)
}
