package user

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func List() (userAPIs []*api.User, err error) {
	c := getC()
	defer c.Database.Session.Close()

	sel := bson.M{
		"_id":   true,
		"nick":  true,
		"mail":  true,
		"phone": true,
		"role":  true,
	}

	if err = c.Find(nil).Sort("_id").Select(sel).All(&userAPIs); err != nil {
		return
	}

	return
}
