package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

func CreateHash(pattern string) string {
	h := sha256.New()
	h.Write([]byte(pattern))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func CheckPasswordHash(password, hash string) bool {
	return password == hash
	//pass := CreateHash(password)
	//if pass != hash {
	//	return false
	//}
	//return true
	//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//return err == nil
}


