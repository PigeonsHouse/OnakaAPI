package cruds

import (
	"errors"
	"fmt"
	"onaka-api/db"
	"onaka-api/types"
	"onaka-api/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, email string, password string) (res_user db.User, err error) {
	if err = db.Psql.Where("email = ?", email).First(&db.User{}).Error; err == nil {
		err = errors.New("email is already exist")
		return
	}
	hash_pass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	res_user = db.User{Email: email, Name: name, PasswordHash: string(hash_pass)}
	err = db.Psql.Create(&res_user).Error
	return
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
	if err = db.Psql.First(&db.User{}, "id = ?", userId).Error; err != nil {
		return
	}

	posts := []db.Post{}
	if err = db.Psql.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return
	}
	for _, post := range posts {
		err = DeletePost(post.ID, userId)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	err = db.Psql.Where("id = ?", userId).Delete(&db.User{}).Error
	return
}

func UpdateName(userId string, name string) (u db.User, err error) {
	if err = GetUserByID(&u, userId); err != nil {
		return
	}
	u.Name = name
	err = db.Psql.Save(&u).Error
	return
}
