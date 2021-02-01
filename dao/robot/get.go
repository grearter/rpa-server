package robot

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func Get(id bson.ObjectId) (robotAPI *api.Robot, err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.FindId(id).One(&robotAPI); err != nil {
		return
	}

	return
}
