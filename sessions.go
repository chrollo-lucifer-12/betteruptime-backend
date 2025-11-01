package main

import (
	"crypto/subtle"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/chrollo-lucifer-12/betteruptime/db"
)

const inactivityTimeoutSeconds = 60 * 60 * 24 * 10
const activityCheckIntervalSeconds = 60 * 60

type SessionWithToken struct {
	Id         string
	SecretHash string
	CreatedAt  time.Time
	Token      string
}

func (s *Server) createSession(user_id uint) *SessionWithToken {
	id, _ := GenerateSecureRandomString()
	secret, _ := GenerateSecureRandomString()
	secretHash := HashSecret(secret)
	secretHashHex := hex.EncodeToString(secretHash)

	token := id + "." + secret
	session := SessionWithToken{
		Id:         id,
		SecretHash: secretHashHex,
		Token:      token,
		CreatedAt:  time.Now(),
	}

	newSession := &db.Session{
		UserID:     user_id,
		SecretHash: secretHashHex,
	}

	if err := s.db.Create(newSession).Error; err != nil {
		return nil
	}

	return &session
}

func (s *Server) validateSession(token string) *db.Session {
	now := time.Now()
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) < 2 {
		return nil
	}
	sessionId := tokenParts[0]
	sessionSecret := tokenParts[1]
	sessionIdUint, _ := strconv.ParseUint(sessionId, 10, 64)
	session := s.getSession(uint(sessionIdUint))
	if session == nil {
		return nil
	}
	secretHashBytes, _ := hex.DecodeString(session.SecretHash)
	expectedHash := HashSecret(sessionSecret)

	if subtle.ConstantTimeCompare(secretHashBytes, expectedHash) != 1 {
		return nil
	}
	if now.Sub(session.UpdatedAt) >= time.Duration(activityCheckIntervalSeconds)*time.Second {
		s.db.Model(&db.Session{}).
			Where("id = ?", session.ID).
			Update("updated_at", now)
	}

	return session
}

func (s *Server) getSession(sessionId uint) *db.Session {
	session := &db.Session{}
	if err := s.db.Where("id = ?", sessionId).Model(session).Error; err != nil {
		return nil
	}
	if time.Since(session.UpdatedAt) >= inactivityTimeoutSeconds*1000 {
		s.deleteSession(sessionId)
		return nil
	}
	return session
}

func (s *Server) deleteSession(sessionId uint) {
	s.db.Where("id = ?", sessionId).Delete(&db.Session{})
}
