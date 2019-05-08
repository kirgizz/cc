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
	m.CreateTablePublications()
	m.CreateTableCredentials()
	m.CreateTableTags()
	m.CreateTableRubrics()
}


func UpdateDbStruct() {
	db := services.GetInstanceDB()
	db.AutoMigrate(&models.User{})
}


