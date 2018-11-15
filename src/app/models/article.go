package models

import (
	"app/services"
	"fmt"
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

func GetArticleByName(name string) *Article {
	var a Article
	services.GetInstanceDB().Where("name = ?", name).First(&a)
	return &a
}

func GetArticleById(articleId int) *Article {
	var a Article
	services.GetInstanceDB().Where("id = ?", articleId).First(&a)
	return &a
}

func GetArticlesByDate(date time.Time) *[]Article {
	var a []Article
	services.GetInstanceDB().Where("created_td = ?", date).Find(&a)
	return &a
}

func GetArticlesByUserNickname(nickname string) *[]Article {
	var a []Article
	services.GetInstanceDB().Raw("SELECT * from articles WHERE user_id in (select id from users where nickname = ?)", nickname).Scan(&a)
	fmt.Println(a)
	return &a
}

func GetArticlesByRating(rating int, duration string) *[]Article {
	var a []Article
	services.GetInstanceDB().Where(fmt.Sprintf("rating %s ?", duration), rating).Find(&a)
	return &a
}

func GetArticles() *[]Article {
	var a []Article
	services.GetInstanceDB().Limit(5).Find(&a)
	return &a
}

func AddArticle(name, body, picturePath string, userId int) (int, error) {
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

func UpdateArticle(name, body, picturePath string, articleId int) error {
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

func DeleteArticle(articleId int) error {
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

func CalculateRating(name string) {

}

func CalculateLikes(name string, like int) error {
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

func IncrimentViewCounts(name string) error {
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

func AddTags() {

}

func AddComment() {

}
