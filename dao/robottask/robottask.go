package robottask

import (
	"github.com/globalsign/mgo"
	"github.com/grearter/rpa-server/conf"
)

const (
	dbName         = "rpa"
	collectionName = "robot_task"
)

func getC() *mgo.Collection {
	return conf.MongoSession.Copy().DB(dbName).C(collectionName)
}
