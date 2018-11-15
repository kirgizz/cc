package auth

import (
	"golang.org/x/crypto/bcrypt"
	"app/models"
	"math/rand"
	"app/services"
)

func Login(nickname, email, password string, rememberMe bool) (int, string) {
	var c models.Credentials
	if len(nickname) == 0 {
		services.GetInstanceDB().Raw("SELECT password from credentials WHERE user_id in (select id from users where email = ?)", email).Scan(&c)
	} else if len(email) == 0 {
		services.GetInstanceDB().Raw("SELECT password from credentials WHERE user_id in (select id from users where nickname = ?)", nickname).Scan(&c)
	}

	if len(c.Password) == 0 {
		return 404, "User not found"
	}
	//create password hash and check it with database

	err := CheckPasswordHash(password, c.Password)
	if err == false {
		return 403, "Wrong password, Acces denied"
	} else {
		if rememberMe {
			//createSession()
		}
		return 200, "Success"
	}
	return 403, "Execute access forbidden"
}

func logout(nickname string) int{
	var c models.Credentials
	services.GetInstanceDB().Raw("SELECT * from credentials WHERE user_id in (select id from users where nickname = ?)", nickname).Scan(&c)
	if len(c.Password) == 0 {
		return 404
	}
	c.Ssid = ""
	return 200
}


func checkSession(c models.Credentials, session string) bool {
	if c.Ssid == session {
		return true
	}
	return false
}


func createSsId(c models.Credentials) error {
	var err error
	c.Ssid, err  = CreateHash(RandString(16), 32)
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

func CreateHash(pattern string, length int) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pattern), length)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

