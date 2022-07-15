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
	ID        string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Base
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type Posts struct {
	Base
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	ImageUrl  string `json:"image_url"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	FunnyUser []User `gorm:"many2many:funnies;"`
	YummyUser []User `gorm:"many2many:yummies;"`
}

type Funny struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	PostID string `gorm:"primaryKey" json:"post_id"`
}

type Yummy struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	PostID string `gorm:"primaryKey" json:"post_id"`
}
