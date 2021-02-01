package robottask

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func Get(id bson.ObjectId) (taskAPI *api.RobotTask, err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.FindId(id).One(&taskAPI); err != nil {
		return
	}

	return
}

func GetLast(robotID bson.ObjectId) (taskAPI *api.RobotTask, err error) {
	c := getC()
	defer c.Database.Session.Close()

	q := bson.M{"robotId": robotID}

	if err = c.Find(q).Sort("-ct").One(&taskAPI); err != nil {
		return
	}

	return
}
