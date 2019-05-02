package server

import (
	"app/auth"
	"app/models"
	"app/services"
	"app/utils"
	"encoding/json"
	"github.com/ivahaev/go-logger"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func setCORS(w http.ResponseWriter) {
	//	cors := cors.New(cors.Options{
	//		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
	//		//http://localhost:8081/message/
	//		//AllowedOrigins:   []string{"*"},
	//		AllowedOrigins:   []string{"http://localhost:3005/"},
	//		//AllowedOrigins:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	//		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//		ExposedHeaders:   []string{"Link"},
	//		AllowCredentials: true,
	//		MaxAge:           300, // Maximum value not ignored by any of major browsers
	//	})
	//Access-Control-Allow-Headers: Content-Type

	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Methods", "GET POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3005")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

}

func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var c models.Credentials
		ssid, err := r.Cookie(SSID)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
			return
		}

		services.GetInstanceDB().Where("ssid = ?", ssid.Value).First(&c)
		if len(c.Ssid) == 0 {
			w.WriteHeader(404)
			w.Write([]byte("ssid empty"))
			return
		}
		//ctx := context.WithValue(r.Context())
		next.ServeHTTP(w,r)
	})
}

func CheckSession2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ctx := context.WithValue(r.Context())
		next.ServeHTTP(w,r)
	})
}

var CheckSessionHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	setCORS(w)
	var statusCode = 403
	var ssid struct {
		Ssid string    `json:"ssid"`
	}

	body := json.NewDecoder(r.Body)
	err := body.Decode(&ssid)
	if err != nil {
		logger.Error(err)
	}

	result := auth.TempCheckSession(ssid.Ssid)
	if result {
		statusCode = 200
	}
	logger.Info(statusCode)
	w.WriteHeader(statusCode)
	//w.Write([]byte("result"))
})

//next.ServeHTTP(w, r.WithContext(ctxGroup))

//https://habr.com/post/323714/
func createSession(ssid string) {

}

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	setCORS(w)
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
	status, result := auth.CheckUserPass(r.Form.Get(EMAIL), r.Form.Get(PASS), ssid)
	if status == 200 {
		ssidCookie := http.Cookie{
			Name:   SSID,
			Value:  ssid,
			Domain: DOMAIN,
			Path:   PATH,
			MaxAge: maxAge}
		http.SetCookie(w, &ssidCookie)


		//http.Redirect(w,r,"http://c-c.ru", http.StatusMovedPermanently)
		//return
	}
	w.WriteHeader(status)
	w.Write([]byte(result))
})


var Logout = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	logger.Info("this is logout page")
	http.Redirect(w,r,"http://c-c.ru", http.StatusMovedPermanently)
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
	//should send text letter on email to check user registration
	if status == 200 {
		ssid := utils.CreateHash(r.UserAgent() + r.Host)
		ssidCookie := http.Cookie{
			Name:   SSID,
			Value:  ssid,
			Domain: DOMAIN,
			Path:   PATH,
			MaxAge: 3600}
		//http.SetCookie(w, &ssidCookie)
		logger.Info(ssidCookie)
		logger.Info("aalaafdasdfsdf")
		http.Redirect(w, r, "http://google.ru", http.StatusMovedPermanently)
		//w.WriteHeader(status)
		w.Write([]byte(message))
		return
	}
	logger.Info("lolololololo")
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
	setCORS(w)
	w.Write([]byte(b))
})

//getArticleById

var GetArticle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	parsedRequestURI := strings.Split(r.RequestURI, "/")
	rawId := parsedRequestURI[len(parsedRequestURI)-1]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		//maybe redirect will be better???
		logger.Error(err)
		w.WriteHeader(404)
		w.Write([]byte("not found"))
		return
	}
	article := models.GetArticleById(id)
	//a := models.GetArticleById()
	//this is baaaaad stuff!!!!
	jsonArticle, err := json.Marshal(article)
	if err != nil {
		logger.Error(err)
	}
	//lw.WriteHeader(200)
	logger.Info(jsonArticle)
	//w.Write([]byte(jsonArticle))
	http.Redirect(w,r, "http://c-c.ru/post.html", http.StatusMovedPermanently)
	//w.Write([]byte(jsonArticle))
})

var CalculateRating = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	//is it correct error answer??
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("wrong data in request"))
	}
	var msg struct{
		Article string `json: "article"`
		Direction int `json: direction`
	}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("cant parse json"))
	}

	err, rating := models.CalculateRating(msg.Article, msg.Direction)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("something goes wrong"))
	}
	w.WriteHeader(200)
	logger.Info(rating)
	w.Write([]byte(strconv.Itoa(rating)))
	//w.Write([]byte(string(rating)))

})

func Cleaner(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	logger.Info(msg)

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}


