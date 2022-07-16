package cruds

import (
	"errors"
	"onaka-api/db"
)

func GiveYummy(postId string, userId string) (post db.Post, err error) {
	if err = db.Psql.Where("posts_id = ? AND user_id = ?", postId, userId).First(&db.Yummy{}).Error; err == nil {
		err = errors.New("yummy is already exist")
		return
	}

	ym := db.Yummy{
		PostID: postId,
		UserID: userId,
	}
	if err = db.Psql.Create(ym).Error; err != nil {
		return
	}
	post, err = GetPost(postId)
	return
}

func DeleteYummy(postId string, userId string) (post db.Post, err error) {
	if err = db.Psql.Where("posts_id = ? AND user_id = ?", postId, userId).First(&db.Yummy{}).Error; err != nil {
		err = errors.New("yummy is not found")
		return
	}

	ym := db.Yummy{
		PostID: postId,
		UserID: userId,
	}
	if err = db.Psql.Delete(ym).Error; err != nil {
		return
	}
	post, err = GetPost(postId)
	return
}
