package cruds

import (
	"errors"
	"onaka-api/db"
)

func GiveFunny(postId string, userId string) (post db.Posts, err error) {
	if err = db.Psql.Where("posts_id = ? AND user_id = ?", postId, userId).First(&db.Funny{}).Error; err == nil {
		err = errors.New("funny is already exist")
		return
	}

	fn := db.Funny{
		PostsID: postId,
		UserID:  userId,
	}
	if err = db.Psql.Create(fn).Error; err != nil {
		return
	}
	post, err = GetPost(postId)
	return
}

func DeleteFunny(postId string, userId string) (post db.Posts, err error) {
	if err = db.Psql.Where("posts_id = ? AND user_id = ?", postId, userId).First(&db.Funny{}).Error; err != nil {
		err = errors.New("funny is not found")
		return
	}

	fn := db.Funny{
		PostsID: postId,
		UserID:  userId,
	}
	if err = db.Psql.Delete(fn).Error; err != nil {
		return
	}
	post, err = GetPost(postId)
	return
}
