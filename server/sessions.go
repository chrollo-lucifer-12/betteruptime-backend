package server

import (
	"crypto/subtle"
	"encoding/hex"
	"strings"
	"time"

	"github.com/chrollo-lucifer-12/betteruptime/db"
	"github.com/google/uuid"
)

const inactivityTimeoutSeconds = 60 * 60 * 24 * 10
const activityCheckIntervalSeconds = 60 * 60

type SessionWithToken struct {
	SessionId  string
	SecretHash string
	CreatedAt  time.Time
	Token      string
}

func (s *Server) createSession(user_id uint) *SessionWithToken {
	id := uuid.New().String()
	secret, _ := GenerateSecureRandomString()
	secretHash := HashSecret(secret)
	secretHashHex := hex.EncodeToString(secretHash)

	token := id + "." + secret
	session := SessionWithToken{
		SessionId:  id,
		SecretHash: secretHashHex,
		Token:      token,
		CreatedAt:  time.Now(),
	}

	sessionId_uuid, _ := uuid.Parse(session.SessionId)

	newSession := &db.Session{
		UserID:     user_id,
		SecretHash: secretHashHex,
		SessionID:  sessionId_uuid,
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
	sessionId_uuid, _ := uuid.Parse(sessionId)
	session := s.getSession(sessionId_uuid)
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

func (s *Server) getSession(sessionId uuid.UUID) *db.Session {
	session := &db.Session{}
	if err := s.db.Where("session_id = ?", sessionId).First(session).Error; err != nil {
		return nil
	}
	if time.Since(session.UpdatedAt) >= time.Duration(inactivityTimeoutSeconds)*time.Second {
		s.deleteSession(sessionId)
		return nil
	}
	return session
}

func (s *Server) deleteSession(sessionId uuid.UUID) {
	s.db.Where("session_id = ?", sessionId).Delete(&db.Session{})
}
