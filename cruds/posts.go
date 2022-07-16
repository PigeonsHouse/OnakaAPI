package cruds

import (
	"mime/multipart"
	"net/textproto"
	"onaka-api/db"
)

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64
	content  []byte
	tmpfile  string
}

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

func PostImages(file *multipart.FileHeader) (string, error){
	temp := db.Psql.Create(&file)
	
}
