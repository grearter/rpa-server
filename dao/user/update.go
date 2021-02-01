package user

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func Update(id, nick, mail, phone string, role *api.Role) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	up := bson.M{
		"nick":  nick,
		"mail":  mail,
		"phone": phone,
		"role":  role,
	}

	if err = c.UpdateId(id, up); err != nil {
		return
	}

	return
}

func UpdateAuth(id, auth string) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	up := bson.M{
		"auth": auth,
	}

	if err = c.UpdateId(id, up); err != nil {
		return
	}

	return
}
