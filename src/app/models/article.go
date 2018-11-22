package models

import (
	"app/services"
	"fmt"
	"github.com/ivahaev/go-logger"
	"time"
)

type Article struct {
	Id         int       `gorm:"primary_key:true"`
	UserId     int       `gorm:"column:user_id"`
	Name       string    `gorm:"column:name;not null;unique" sql:"DEFAULT:NULL"`
	Body       string    `gorm:"column:body"`
	Picture    string    `gorm:"column:picture;unique" sql:"DEFAULT:NULL"`
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

func GetArticlesFromDatabase() *[]Article {
	var a []Article
	services.GetInstanceDB().Limit(5).Find(&a)
	return &a
}

func SaveArticlenDb(name, body string, userId int) (int, string) {
	a := &Article{
		UserId:    userId,
		Likes:     0,
		Rating:    0,
		ViewCount: 0,
		Name:      name,
		Body:      body,
		CreatedAt: time.Now()}

	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		logger.Error(tx.Error.Error())
		return 500, tx.Error.Error()
	}
	if err := tx.Create(&a).Error; err != nil {
		logger.Error(err.Error())
		tx.Rollback()
		logger.Error(err.Error())
		return 500, err.Error()
	}
	//IS IT CORRECT BEHAVIOR?
	//if err := tx.Commit().Error; err == nil {
	//	return 200, "Success"
	//}
	if err := tx.Commit().Error; err == nil {
		return 200, "success"
	}
	return 500, tx.Commit().Error.Error()
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