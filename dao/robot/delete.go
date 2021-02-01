package robot

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func Delete(id bson.ObjectId) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	q := bson.M{
		"_id":    id,
		"status": api.RobotStatusStopped,
	}

	if err = c.Remove(q); err != nil {
		if err == mgo.ErrNotFound {
			err = errors.New("robot运行中")
		}
		return
	}

	return
}
