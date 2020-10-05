package service

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession() (*Session, error) {
	mongoHost := os.Getenv("MONGODB_HOST")
	mongoPort := os.Getenv("MONGODB_PORT")
	connectionString := mongoHost + ":" + mongoPort
	log.Println("Mongo URL:", connectionString)

	session, err := mgo.Dial(connectionString)
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
