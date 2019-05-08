package models

import (
	"app/services"
	"fmt"
	"github.com/ivahaev/go-logger"
	"strings"
	"time"
)

type Publication struct {
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
	Tags       []*tags   `gorm:"many2many:article_tags;"`
	Rubrics       []*Rubrics   `gorm:"many2many:article_rubrics;"`
	//	Status			string 		`gorm:"column:status"`
	//tags 			[]tags 		`gorm:"foreignkey:id"`
	//Created
	//Updatedb
	//Deleted
}

func (m *Model) CreateTablePublications() {
	services.GetInstanceDB().CreateTable(&Publication{})
	services.GetInstanceDB().Model(&Publication{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func (a *Publication) saveArticles(userid int, name, body string, ) (error){

	return services.GetInstanceDB().Save(&a).Error
}










func GetArticleByName(name string) *Publication {
	var a Publication
	services.GetInstanceDB().Where("name = ?", name).First(&a)
	return &a
}

func GetArticleById(articleId int) *Publication {
	var a Publication
	services.GetInstanceDB().Where("id = ?", articleId).First(&a)
	return &a
}

func getArticlesByDate(date time.Time) *[]Publication {
	var a []Publication
	services.GetInstanceDB().Where("created_td = ?", date).Find(&a)
	return &a
}

func getArticlesByUserNickname(nickname string) *[]Publication {
	var a []Publication
	services.GetInstanceDB().Raw("SELECT * from articles WHERE user_id in (select id from users where nickname = ?)", nickname).Scan(&a)
	fmt.Println(a)
	return &a
}

func getArticlesByRating(rating int, duration string) *[]Publication {
	var a []Publication
	services.GetInstanceDB().Where(fmt.Sprintf("rating %s ?", duration), rating).Find(&a)
	return &a
}

func GetArticlesFromDatabase() interface{} {
	//var g []Article
	type article struct {
		Id         int
		Email      string
		Name       string
		Rating     int
		Body       string
		View_count int
	}
	var a []article
	services.GetInstanceDB().Raw("select articles.id, users.email, articles.name, articles.rating, articles.body, articles.view_count from users, articles where users.id = articles.user_id").Scan(&a)
	//services.GetInstanceDB().Limit(5).Find(&a)
	for _, v := range a {

		v.Body = strings.Split(v.Body, "<cut>")[0]
	}
	return a
	//select users.email, articles.name, articles.rating, articles.body from users, articles where users.id = articles.user_id
}

func SaveArticlenDb(name, body string, userId int) (int, string) {
	//check tag <cut> in article body
	a := &Publication{
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
	if err := tx.Model(&Publication{}).Where("id = ?", articleId).Update(map[string]interface{}{"name": name, "body": body, "picturePath": picturePath}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func deleteArticle(articleId int) error {
	a := &Publication{Id: articleId}
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

func CalculateRating(name string, rating int) (error, int) {

	var a Publication
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error, -1
	}
	//db.Model(&user).Where("active = ?", true).Update("name", "hello")
	if err := tx.Where("name = ?", name).First(&a).Error; err != nil {
		logger.Error(err)
		tx.Rollback()
		return err, -1

	}
	a.Rating = a.Rating + rating
	if a.Rating < 0 {
		a.Rating = 0
	}
	if err := tx.Model(&a).Update("rating", a.Rating).Error; err != nil {
		logger.Error(err)
		tx.Rollback()
		return err, -1
	}
	return tx.Commit().Error, a.Rating
}

func calculateLikes(name string, like int) error {
	db := services.GetInstanceDB().Begin()
	var a Publication
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
	var a Publication
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
