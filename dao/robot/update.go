package robot

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func UpdateStatus(id bson.ObjectId, status api.RobotStatus) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.UpdateId(id, bson.M{"$set": bson.M{"status": status}}); err != nil {
		return
	}

	return
}
