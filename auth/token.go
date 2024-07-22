package auth

import (
	"fmt"
	"net/http"
	"time"
)

var Session = map[string]int64{}

func CreateSession(userId int64) string {
	token := fmt.Sprintf("yo-%d", userId)
	Session[token] = userId
	return token
}

func GetUserIdFromToken(token string) (int64, bool) {
	userID, ok := Session[token]
	return userID, ok
}

func CreateCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
}

func DeleteCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	}
}
