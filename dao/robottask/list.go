package robottask

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func List(robotID bson.ObjectId, offset, limit int) (total int, taskAPIs []*api.RobotTask, err error) {
	c := getC()
	defer c.Database.Session.Close()

	q := bson.M{"robotId": robotID}
	sel := bson.M{"messages": false}

	total, err = c.Find(q).Count()
	if err != nil {
		return
	}

	if total == 0 {
		return
	}

	if err = c.Find(q).Sort("-ct").Select(sel).Skip(offset).Limit(limit).All(&taskAPIs); err != nil {
		return
	}

	return
}
