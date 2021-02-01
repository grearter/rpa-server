package user

import "github.com/globalsign/mgo/bson"

func Auth(id, auth string) (ok bool, err error) {
	c := getC()
	defer c.Database.Session.Close()

	count, err := c.Find(bson.M{"_id": id, "auth": auth}).Count()
	if err != nil {
		return
	}

	ok = count > 0
	return
}
