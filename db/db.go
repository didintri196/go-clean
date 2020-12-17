package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

type DBConnection struct {
	Session *mgo.Session
}

func (conn *DBConnection) NewConnection() {
	//session, err := mgo.Dial("127.0.0.1:27017")
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  10 * time.Second,
		Database: "db_adms",
		// Username: "evenkita",
		// Password: "evenkita1234",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	conn.Session = session
}

func (conn *DBConnection) Close() {
	conn.Session.Close()
	return
}
