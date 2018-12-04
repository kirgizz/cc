package auth

import (
	"app/models"
	"app/services"
	"app/utils"
)



func CheckUserPass(email, password, ssid string) (int, string) {
	var c models.Credentials
		services.GetInstanceDB().Raw("SELECT * from credentials WHERE user_id in (select id from users where email = ?)", email).Scan(&c)

	if len(c.Password) == 0 {
		return 404, "User not found"
	}

	//create password hash and check it with database
	result := utils.CheckPasswordHash(utils.CreateHash(password), c.Password)
	if result == false {
		return 403, "Forbidden"
	} else {
		//Transaction??
		services.GetInstanceDB().Model(&c).Where("id = ?", c.Id).Update("ssid", ssid)
		//Where("name = ?", "jinzhu").First(&user)


		return 200, "Success"
	}
	return 403, "Execute access forbidden"
}

func checkSession(c models.Credentials, session string) bool {
	if c.Ssid == session {
		return true
	}
	return false
}


func createSsId(c models.Credentials) error {
	var err error
	c.Ssid  = utils.CreateHash(utils.RandString(16))
	if err != nil {
		return err
	}
	//save ssid to database
	return nil
}

func removeSsid(user_id int) {
	//find credential with this is
	//remove ssid
}