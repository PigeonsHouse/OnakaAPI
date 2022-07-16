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
