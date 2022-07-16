package cruds

import (
	"onaka-api/db"
)

func GetTimeLine() (timeline []db.Post, err error) {
	err = db.Psql.Model(&db.Post{}).Find(&timeline).Error
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

func GetPost(postId string) (post db.Post, err error) {
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

func PostPosts(content string, url string, userId string) (db.Post, error) {
	p := db.Post{Content: content, ImageUrl: url, UserID: userId}
	db.Psql.Create(&p)
	if err := db.Psql.First(&p, "id = ?", p.ID).Error; err != nil {
		return p, err
	}
	var user []db.User
	err := db.Psql.Model(&p).Association("User").Find(&user)
	if err != nil {
		return p, err
	}
	p.User = user[0]
	err = db.Psql.Model(&p).Association("FunnyUsers").Find(&user)
	if err != nil {
		return p, err
	}
	p.FunnyUsers = user
	err = db.Psql.Model(&p).Association("YummyUsers").Find(&user)
	if err != nil {
		return p, err
	}
	p.YummyUsers = user
	return p, nil
}

func DeletePost(postId string, userId string) (err error) {
	if err = db.Psql.Where("id = ? AND user_id = ?", postId, userId).First(&db.Post{}).Error; err != nil {
		return
	}
	db.Psql.Where("posts_id = ?", postId).Delete(&db.Yummy{})
	db.Psql.Where("posts_id = ?", postId).Delete(&db.Funny{})

	err = db.Psql.Where("id = ? AND user_id = ?", postId, userId).Delete(&db.Post{}).Error
	return
}
