package models

import (
	"app/services"
	"time"
)

type Comment struct {
	Id        int       `gorm:"primary_key:true"`
	UserId    int       `gorm:"column:user_id"`
	ArticleId int       `gorm:"column:article_id"`
	Body      string    `gorm:"column:body"`
	Likes     int       `gorm:"column:like"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (m Model) CreateTableComments() {
	services.GetInstanceDB().CreateTable(&Comment{})
	services.GetInstanceDB().Model(&Comment{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	services.GetInstanceDB().Model(&Comment{}).AddForeignKey("article_id", "articles(id)", "RESTRICT", "RESTRICT")
}

func createComment(article_id, user_id int, body string) error {
	a := &Comment{
		UserId:    user_id,
		ArticleId: article_id,
		Body:      body,
		Likes:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	db := services.GetInstanceDB()

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func updateComment(commentId int, body string) error {
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&Comment{}).Where("id = ?", commentId).Update("body", body).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func calculateLikes1(commentId, like int) error {
	db := services.GetInstanceDB().Begin()
	var c Comment
	db.Where("id = ?", commentId).First(&c)
	c.Likes += like

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Where("id = ?", commentId).Update("likes", c.Likes).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func deleteComment(commentId int) error {
	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(&Comment{}).Where("id = ?", commentId).Update("status", "deleted").Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
