package auth

import "fmt"

var session = map[string]int64{}

func CreateSession(userId int64) string {
	token := fmt.Sprintf("yo-%d", userId)
	session[token] = userId
	return token
}
