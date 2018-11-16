package models

import (
	"app/services"
	"encoding/json"
	"fmt"
	"github.com/ivahaev/go-logger"
	"net/http"
	"time"
)

type Article struct {
	Id         int       `gorm:"primary_key:true"`
	UserId     int       `gorm:"column:user_id"`
	Name       string    `gorm:"column:name;not null;unique"`
	Body       string    `gorm:"column:body"`
	Picture    string    `gorm:"column:picture;not null;unique"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UptdatedAt time.Time `gorm:"column:update_at"`
	Likes      int       `gorm:"column:likes"`
	Rating     int       `gorm:"column:rating"`
	ViewCount  int       `gorm:"column:view_count"`
	//	Status			string 		`gorm:"column:status"`
	//tags 			[]tags 		`gorm:"foreignkey:id"`
	//Created
	//Updatedb
	//Deleted
}

func CreateTableArticles() {
	services.GetInstanceDB().CreateTable(&Article{})
	services.GetInstanceDB().Model(&Article{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func getArticleByName(name string) *Article {
	var a Article
	services.GetInstanceDB().Where("name = ?", name).First(&a)
	return &a
}

func getArticleById(articleId int) *Article {
	var a Article
	services.GetInstanceDB().Where("id = ?", articleId).First(&a)
	return &a
}

func getArticlesByDate(date time.Time) *[]Article {
	var a []Article
	services.GetInstanceDB().Where("created_td = ?", date).Find(&a)
	return &a
}

func getArticlesByUserNickname(nickname string) *[]Article {
	var a []Article
	services.GetInstanceDB().Raw("SELECT * from articles WHERE user_id in (select id from users where nickname = ?)", nickname).Scan(&a)
	fmt.Println(a)
	return &a
}

func getArticlesByRating(rating int, duration string) *[]Article {
	var a []Article
	services.GetInstanceDB().Where(fmt.Sprintf("rating %s ?", duration), rating).Find(&a)
	return &a
}

func getArticlesFromDatabase() *[]Article {
	var a []Article
	services.GetInstanceDB().Limit(5).Find(&a)
	return &a
}

func saveArticlenDb(name, body, picturePath string, userId int) (int, error) {
	a := &Article{
		UserId:    userId,
		Likes:     0,
		Rating:    0,
		ViewCount: 0,
		Name:      name,
		Body:      body,
		CreatedAt: time.Now(),
		Picture:   picturePath}
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return 500, tx.Error
	}
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return 500, err
	}
	if tx.Commit().Error != nil {
		return 200, tx.Commit().Error
	}
	return 500, tx.Commit().Error
}

func updateArticle(name, body, picturePath string, articleId int) error {
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&Article{}).Where("id = ?", articleId).Update(map[string]interface{}{"name": name, "body": body, "picturePath": picturePath}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func deleteArticle(articleId int) error {
	a := &Article{Id: articleId}
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Delete(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func calculateRating(name string) {

}

func calculateLikes(name string, like int) error {
	db := services.GetInstanceDB().Begin()
	var a Article
	db.Where("name = ?", name).First(&a)
	a.Likes += like

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Where("name = ?", name).Update("likes", a.Likes).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func incrimentViewCounts(name string) error {
	db := services.GetInstanceDB().Begin()
	var a Article
	db.Where("name = ?", name).First(&a)
	a.ViewCount += 1
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Where("name = ?", name).Update("view_count", a.ViewCount).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func addTags() {

}

func addComment() {

}



//HTTP hadnlers
var AddArticle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var u User
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//where get current user???
	services.GetInstanceDB().Where("email = ?", "la-la-la").Find(&u)
	picturePath := "testPath"
	status, result := saveArticlenDb(r.Form.Get("name"), r.Form.Get("body"), picturePath, u.Id)
	w.WriteHeader(status)
	w.Write([]byte(result.Error()))
})

var GetArticles = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	a := getArticlesFromDatabase()
	b, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
	}
	w.Write([]byte(b))
})
