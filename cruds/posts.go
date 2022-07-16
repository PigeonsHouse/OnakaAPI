package cruds

import (
	"onaka-api/db"
)

func GetTimeLine() ([]db.Posts, error){
	var timeline []db.Posts
	err := db.Psql.Find(&timeline).Error
	return timeline, err
}
