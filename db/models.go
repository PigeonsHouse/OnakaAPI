package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}

type Base struct {
	ID        string    `gorm:"primaryKey" json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Base
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type Posts struct {
	Base
	UserID     string `json:"-"`
	Content    string `json:"content"`
	ImageUrl   string `json:"image_url"`
	User       User   `gorm:"foreignKey:UserID;reference:ID" json:"user"`
	FunnyUsers []User `gorm:"many2many:funnies;" json:"funny_users"`
	YummyUsers []User `gorm:"many2many:yummies;" json:"yummy_users"`
}

type Funny struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	PostID string `gorm:"primaryKey" json:"post_id"`
}

type Yummy struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	PostID string `gorm:"primaryKey" json:"post_id"`
}
