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
		return 404, "not found"
	}

	//create password hash and check it with database
	result := utils.CheckPasswordHash(utils.CreateHash(password), c.Password)
	if result == false {
		return 403, "wrong credentials"
	} else {
		var u models.User
		//Transaction??
		//wrong behavior change id every time when user logged in
		services.GetInstanceDB().Model(&c).Where("id = ?", c.Id).Update("ssid", ssid)
		//Where("name = ?", "jinzhu").First(&user)
		services.GetInstanceDB().Where("email = ?", email).First(&u)

		return 200, ssid
	}
	return 403, "wrong credentials"
}

func checkSession(c models.Credentials, session string) bool {
	if c.Ssid == session {
		return true
	}
	return false
}

func TempCheckSession(ssid string) bool {
	var s models.Credentials
	services.GetInstanceDB().Where("ssid = ?", ssid).First(&s)
	if len(s.Password) == 0 {
		return false
	}
	return true
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
