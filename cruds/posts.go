package cruds

import (
	"onaka-api/db"
)

func GetTimeLine() (timeline []db.Posts, err error) {
	err = db.Psql.Model(&db.Posts{}).Find(&timeline).Error
	if err != nil {
		return
	}
	for i, post := range timeline {
		var user []db.User
		err = db.Psql.Model(&post).Association("User").Find(&user)
		if err != nil {
			return
		}
		timeline[i].User = user[0]
		err = db.Psql.Model(&post).Association("FunnyUsers").Find(&user)
		if err != nil {
			return
		}
		timeline[i].FunnyUsers = user
		err = db.Psql.Model(&post).Association("YummyUsers").Find(&user)
		if err != nil {
			return
		}
		timeline[i].YummyUsers = user
	}

	return
}

func GetPost(postId string) (post db.Posts, err error) {
	err = db.Psql.First(&post, "id = ?", postId).Error
	if err != nil {
		return
	}
	var user []db.User
	err = db.Psql.Model(&post).Association("User").Find(&user)
	if err != nil {
		return
	}
	post.User = user[0]
	err = db.Psql.Model(&post).Association("FunnyUsers").Find(&user)
	if err != nil {
		return
	}
	post.FunnyUsers = user
	err = db.Psql.Model(&post).Association("YummyUsers").Find(&user)
	if err != nil {
		return
	}
	post.YummyUsers = user
	return
}

func DeletePost(postId string, userId string) (err error) {
	db.Psql.Where("posts_id = ?", postId).Delete(&db.Yummy{})
	db.Psql.Where("posts_id = ?", postId).Delete(&db.Funny{})

	err = db.Psql.Where("id = ? AND user_id = ?", postId, userId).Delete(&db.Posts{}).Error
	return
}
