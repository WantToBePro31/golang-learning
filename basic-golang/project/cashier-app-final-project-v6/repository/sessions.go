package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"errors"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	// Select target token and delete from listSessions
	// TODO: answer here
	for i := 0; i < len(listSessions); i++ {
		if listSessions[i].Token == tokenTarget {
			listSessions = append(listSessions[:i], listSessions[i+1:]...)
		}
	}

	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	if session.Token == "" || session.Username == "" || session.Expiry.IsZero() {
        return errors.New("bad request")
	}

	records, err := u.db.Load("sessions")
	if err != nil {
		return err
	}

	var listSessions []model.Session
	err = json.Unmarshal(records, &listSessions)
	if err != nil {
		return err
	}
	
	listSessions = append(listSessions, session)
	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	if token == "" {
	    return model.Session{}, errors.New("bad request")
	}

	records, err := u.db.Load("sessions")
	if err != nil {
		return model.Session{}, err
	}

	var listSessions []model.Session
	err = json.Unmarshal(records, &listSessions)
	if err != nil {
		return model.Session{}, err
	}

	for i := 0; i < len(listSessions); i++ {
		if listSessions[i].Token == token {
			if time.Now().After(listSessions[i].Expiry) {
				listSessions = append(listSessions[:i], listSessions[i+1:]...)
				jsonData, err := json.Marshal(listSessions)
				if err != nil {
					return model.Session{}, err
				}
				err = u.db.Save("sessions", jsonData)
				if err != nil {
					return model.Session{}, err
				}
				return model.Session{}, errors.New("Token is Expired!")
			} else {
				return listSessions[i], nil
			}
		}
	}
	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
