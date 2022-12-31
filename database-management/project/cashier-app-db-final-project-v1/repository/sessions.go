package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	if err := u.db.Create(&session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	var session model.Session
	if err := u.db.Where("token = ?", tokenTarget).Delete(&session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	var sessions model.Session
	if err := u.db.Model(&sessions).Where("username = ?", session.Username).Update("token", session.Token).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token) // TODO: replace this
	if err != nil {
		return model.Session{}, err
	}
	if u.TokenExpired(session) {
		err := u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}
	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	var session model.Session
	if err := u.db.Where("token is NOT NULL and username = ?", name).First(&session).Error; err != nil {
		return model.Session{}, err 
	}
	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	if err := u.db.Where("token = ?", token).First(&session).Error; err != nil {
		return model.Session{}, err 
	}
	return session, nil // TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
