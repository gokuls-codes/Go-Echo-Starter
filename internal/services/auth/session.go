package auth

import (
	"log"
	"time"

	"github.com/gokuls-codes/go-echo-starter/types"
	"github.com/google/uuid"
)

func GenerateSessionCookie(user *types.User, s types.UserStore) (*types.Session, error) {
	sess := new(types.Session)
	sess.UserId = user.ID
	sess.ExpiresAt = time.Now().Add(time.Hour * 1)

	sessionToken := uuid.New().String()
	sess.SessionToken = sessionToken

	err := s.CreateSessionForUser(sess)

	if err != nil {
		return nil, err
	}
	return sess, nil
}

func CheckIfLoggedIn(sessionToken string, s types.UserStore) bool {
	sess, err := s.FindSessionBySessionId(sessionToken)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return !sess.ExpiresAt.Before(time.Now())
}