package auth

import (
	"app/models"
	"app/services"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)



func checkUserPass(nickname, email, password string, rememberMe bool) (int, string) {
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

	result := checkPasswordHash(password, c.Password)
	if result == false {
		return 403, "Forbidden"
	} else {
		if rememberMe {
			var err error
			c.Ssid, err = createHash(randString(16), 32)
			if err != nil {
				return 503, "Can'n create session id"
				//what should be here for correct work???
			}
		}
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
	c.Ssid, err  = createHash(randString(16), 32)
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

func createHash(pattern string, length int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pattern), length)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}




//HTTP handlers
func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c models.Credentials
		ssid, err := r.Cookie("ssid")
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
		}
		services.GetInstanceDB().Where("ssid = ?", ssid).First(&c)
		if len(c.Ssid) == 0 {
			w.WriteHeader(404)
			w.Write([]byte("Empty"))
		}
		c.Ssid = ""
		w.WriteHeader(200)

		w.Write([]byte("Success"))
	})
}

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//Check form values?
	status, result := checkUserPass(r.Form.Get("name"),r.Form.Get("email"), r.Form.Get("password"), false )
	w.WriteHeader(status)
	w.Write([]byte(result))
})

var Logout = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var c models.Credentials
	ssid, err := r.Cookie("ssid")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Cookie not found"))
	}
	services.GetInstanceDB().Where("ssid = ?", ssid).First(&c)
	if len(c.Ssid) == 0 {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	}
	c.Ssid = ""
	w.WriteHeader(200)
	w.Write([]byte("Success"))
})