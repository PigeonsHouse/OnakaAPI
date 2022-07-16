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

// func PostImages(file *multipart.FileHeader) (string, error){
// 	temp := db.Psql.Create(&file)
// 	cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
// 	ctx := context.Background()
// 	resp, err := cld.Upload.Upload(ctx, "apple.png", uploader.UploadParams{PublicID: "docs/sdk/go/apple",
//     Transformation: "c_crop,g_center/q_auto/f_auto", Tags: []string{"fruit"}})
// 	my_image, err := cld.Image("docs/sdk/go/apple")
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	url, err := my_image.String()
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	// temp := db.Psql.Create(&file)
// 	// n, err := rand.Int(rand.Reader, big.NewInt(100))
// 	// url := "./static/"+ n.String() + "_"+ time.Now().String() +".png"
// 	// u := &url.URL{}
// 	// f, _ := os.Create(url)
	
// }

func PostPosts(content string, url string, userId string) (db.Posts, error){
	p := db.Posts{Content: content, ImageUrl: url, UserID: userId}
	db.Psql.Create(&p)
	//res_posts := db.Posts{ID: p.ID, Content: p.Content, ImageUrl: p.ImageUrl, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
	if err := db.Psql.First(&p, "id = ?", p.ID).Error; err != nil{
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
