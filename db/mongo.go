package db

import (
	"github.com/thearyanahmed/kloudlabllc/config"

	mgo "gopkg.in/mgo.v2"
)

var instance *mgo.Session

var err error

// GetInstance return copy of db session
func GetInstance(c *config.Configuration) *mgo.Session {

	if instance == nil {
		instance, err = mgo.Dial(c.DataBaseConnectionURL)
		if err != nil {
			panic(err)
		}
	}

	return instance.Copy()
}
