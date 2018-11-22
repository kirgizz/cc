package server

import (
	"app/auth"
	"app/models"
	"app/services"
	"app/utils"
	"encoding/json"
	"github.com/ivahaev/go-logger"
	"net/http"
)

var H string

const (
	SSID = "ssid"
	REMEMBER = "remember"
	EMAIL = "email"
	PASS = "password"
	DOMAIN = "c-c.ru"
	PATH = "/"
)

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	logger.Info(r)
	logger.Info(r.Cookies())
	err := r.ParseForm()
	if err != nil {
		logger.Crit(err)
	}
	logger.Info(r.Form)
	w.Write([]byte("Hello from not implemented function"))
})


func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c models.Credentials
		logger.Info(r.Cookies())
		ssid, err := r.Cookie(SSID)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
			return
		}

		services.GetInstanceDB().Where("ssid = ?", ssid).First(&c)
		if len(c.Ssid) == 0 {
			w.WriteHeader(404)
			w.Write([]byte("ssid empty"))
			return
		}
		//ctx := context.WithValue(r.Context())
		next.ServeHTTP(w,r)
	})
}

//next.ServeHTTP(w, r.WithContext(ctxGroup))

//https://habr.com/post/323714/
var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	var maxAge int
	//Check form values?
	if len(r.Form.Get(REMEMBER)) != 0 {
		maxAge = 2592000
	} else {
		maxAge = 3600
	}
	ssid := utils.CreateHash(r.UserAgent() + r.Host)

	//not good to save ssid in this function
	status, _ := auth.CheckUserPass(r.Form.Get(EMAIL), r.Form.Get(PASS), ssid)
	if status == 200 {
		ssidCookie := http.Cookie{
			Name:   SSID,
			Value:  ssid,
			Domain: DOMAIN,
			Path:   PATH,
			MaxAge: maxAge}
		http.SetCookie(w, &ssidCookie)
	}
	w.WriteHeader(status)
	http.Redirect(w, r, "http://c-c.ru", http.StatusSeeOther)


	//w.Write([]byte(result))
})


var Logout = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var c models.Credentials
	ssid, err := r.Cookie(SSID)
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


var Register = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//Check form values
	//VALIDATE EMAIL
	status, message := auth.RegisterUser(r.Form.Get(EMAIL), r.Form.Get("nickname"), r.Form.Get("psw"))
	w.WriteHeader(status)
	w.Write([]byte(message))

})


//HTTP hadnlers
var AddArticle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var u models.User
	err := r.ParseForm()
	if err != nil {
		logger.Error(err)
		w.Write([]byte(err.Error()))
	}
	logger.Info("coo")
	logger.Info(r.Cookies())
	ssid, err := r.Cookie("ssid")
	if err != nil {
		logger.Error(err)
		//Is this status code is correct???
		w.WriteHeader(403)
		w.Write([]byte("Wrong auth token"))
		return
	}

	services.GetInstanceDB().Raw("SELECT * from users WHERE id in (select user_id from credentials where ssid = ?)", ssid.Value).Scan(&u)
	status, result := models.SaveArticlenDb(r.Form.Get("name"), r.Form.Get("body"), u.Id)
	w.WriteHeader(status)
	w.Write([]byte(result))
})

var GetArticles = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	a := models.GetArticlesFromDatabase()
	b, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
	}
	w.Write([]byte(b))
})


