package user

import (
	"github.com/globalsign/mgo"
	"github.com/grearter/rpa-server/conf"
)

const (
	dbName         = "rpa"
	collectionName = "user"
)

func getC() *mgo.Collection {
	return conf.MongoSession.Copy().DB(dbName).C(collectionName)
}
