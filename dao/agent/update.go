package agent

import "github.com/globalsign/mgo/bson"

func Update(id, hostname, ip string) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	up := bson.M{
		"$set": bson.M{
			"hostname": hostname,
			"ip":       ip,
		},
	}

	if err = c.UpdateId(id, up); err != nil {
		return
	}

	return
}
