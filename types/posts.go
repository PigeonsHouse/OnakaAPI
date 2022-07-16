package types

import (
	"onaka-api/db"
	"time"
)

type PostsResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"-"`
	User       db.User   `gorm:"foreignKey:UserID;reference:ID" json:"user"`
	Content    string    `json:"content"`
	ImageUrl   string    `json:"image_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	FunnyUsers []db.User `gorm:"many2many:funnies;" json:"funny_users"`
	YummyUsers []db.User `gorm:"many2many:yummies;" json:"yummy_users"`
}
