package migrations

import (
	"app/models"
	"app/services"
)

type migrations struct {

}
func CreateDBStruct() {
	var m models.Model
	m.CreateTableUsers()
	m.CreateTableArticles()
	m.CreateTableComments()
	m.CreateTableCredentials()
	m.CreateTableTags()
}


func UpdateDbStruct() {
	db := services.GetInstanceDB()
	db.AutoMigrate(&models.User{})
}


