package auth

import (
	"app/models"
	"app/services"
	"net/http"
	"time"
)


//THIS IS A NEW SHIT

func registerUser(nickname, email, password, name string, birthday time.Time) (int, error) {
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

//HTTP handler
var Register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var birthday time.Time
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//Check form values
	birthday, err = time.Parse("2006-01-02T15:04:05.000Z", r.Form.Get("birthday"))
	status, result := registerUser(r.Form.Get("nickname"),r.Form.Get("email"), r.Form.Get("password"),r.Form.Get("name"), birthday)
	w.WriteHeader(status)
	w.Write([]byte(result.Error()))
})