package cruds

import (
	"errors"
	"onaka-api/db"
	"onaka-api/types"
	"onaka-api/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, email string, password string) (types.UserResponse, error) {
	if err := db.Psql.Where("email = ?", email).First(&db.User{}).Error; err == nil {
		return types.UserResponse{}, errors.New("email is already exist")
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

func GetUserByID(user *db.User, userID string) (err error) {
	usrTmp := db.User{}
	err = db.Psql.First(&usrTmp, "id = ?", userID).Error
	*user = usrTmp
	return
}

func DeleteUser(userId string) (err error) {
	if err = db.Psql.First(&db.User{}).Error; err != nil {
		return
	}

	posts := []db.Posts{}
	if err = db.Psql.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return
	}
	for _, post := range posts {
		err = DeletePost(post.ID, userId)
	}

	err = db.Psql.Where("id = ?", userId).Delete(&db.User{}).Error
	return
}
