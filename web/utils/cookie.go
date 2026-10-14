package utils

import (
	"net/http"
	"time"

	"github.com/arthurcommodore/components/utils"
)

func CreateCookie(cookie http.Cookie, w http.ResponseWriter, r *http.Request) {
	cookie.HttpOnly = isHTTPS(r)
	http.SetCookie(w, &cookie)
}

func SetSessionCookie(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		http.Error(w, "Error token generate", http.StatusInternalServerError)
		return
	}
	CreateCookie(http.Cookie{
		Name:     "session",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(48 * time.Hour),
		Secure:   isHTTPS(r),
		Path:     "/",
	}, w, r)
}

func DestroySessionCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
