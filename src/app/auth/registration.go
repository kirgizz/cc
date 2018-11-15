package auth

import (
	"app/models"
	"app/services"
	"time"
)


//THIS IS A NEW SHIT

func RegisterUser(nickname, email, password, name string, birthday time.Time) (int, error) {
	var u models.User
	var c models.Credentials
	u.NickName = nickname
	u.Email = email
	u.Name = name
	u.BirthDate = birthday

	db := services.GetInstanceDB()
	tx := db.Begin()
	if tx.Error != nil {
		return 500, tx.Error
	}
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return 500, err
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return 500, tx.Commit().Error
	}

	err := db.Where("nickname = ?", nickname).First(&u).Error
	if err != nil {
		return 500, err
	}

	c.Password = password
	c.UserId = u.Id

	tx = db.Begin()
	if tx.Error != nil {
		return 500, tx.Error
	}
	if err := tx.Create(&c).Error; err != nil {
		tx.Rollback()
		return 500, err
	}
	if tx.Commit().Error != nil {
		tx.Rollback()
		return 500, tx.Commit().Error
	}
	return 200, tx.Commit().Error
}



//Id        	int      	`gorm:"primary_key:true"`
//BirthDate 	time.Time 	`gorm:"column:birthdate"`
//CreatedAt 	time.Time	`gorm:"column:created_at`
//Rating   	int 		`gorm:"column:rating"`
//Name     	string 		`gorm:"column:name"`
//NickName 	string 		`gorm:"column:nickname;not null;unique"`
//Email    	string 		`gorm:"column:email;not null;unique"`
//Avatar   	string 		`gorm:"column:avatar"`
//Status   	string 		`gorm:"column:status"`