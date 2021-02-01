package robottask

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
	"time"
)

func AppendLog(id bson.ObjectId, robotMsg *api.RobotMessage) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	up := bson.M{
		"$push": bson.M{
			"messages": robotMsg,
		},
		"$set": bson.M{
			"ut": time.Now().UnixNano(),
		},
	}

	if err = c.UpdateId(id, up); err != nil {
		return
	}

	return
}
