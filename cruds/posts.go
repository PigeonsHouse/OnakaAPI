package cruds

import (
	"fmt"
	"onaka-api/db"
)

func GetTimeLine(limit int, page int) (timeline []db.Post, err error) {
	err = db.Psql.Model(&db.Post{}).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&timeline).Error
	if err != nil {
		return
	}
	for i, post := range timeline {
		var user []db.User
		err = db.Psql.Model(&post).Association("User").Find(&user)
		fmt.Println(user)
		timeline[i].User = user[0]
		err = db.Psql.Model(&post).Association("FunnyUsers").Find(&user)
		fmt.Println(user)
		timeline[i].FunnyUsers = user
		err = db.Psql.Model(&post).Association("YummyUsers").Find(&user)
		fmt.Println(user)
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
	db.Psql.Where("post_id = ?", postId).Delete(&db.Yummy{})
	db.Psql.Where("post_id = ?", postId).Delete(&db.Funny{})

	err = db.Psql.Where("id = ? AND user_id = ?", postId, userId).Delete(&db.Post{}).Error
	return
}

func GetPostsByUserId(userId string, limit int, page int) (ps []db.Post, err error) {
	if err = db.Psql.First(&db.User{}, "id = ?", userId).Error; err != nil {
		return
	}
	db.Psql.Where("user_id = ?", userId).Order("created_at desc").Limit(limit).Offset((page - 1) * limit).Find(&ps)
	for i, post := range ps {
		var user []db.User
		err = db.Psql.Model(&post).Association("User").Find(&user)
		if err != nil {
			return
		}
		ps[i].User = user[0]
		err = db.Psql.Model(&post).Association("FunnyUsers").Find(&user)
		if err != nil {
			return
		}
		ps[i].FunnyUsers = user
		err = db.Psql.Model(&post).Association("YummyUsers").Find(&user)
		if err != nil {
			return
		}
		ps[i].YummyUsers = user
	}
	return
}
