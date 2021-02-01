package agent

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func List(keyword string, offset, limit int) (total int, agentAPI []*api.Agent, err error) {
	c := getC()
	defer c.Database.Session.Close()

	q := bson.M{
		"$or": []bson.M{
			{"hostname": bson.M{"$regex": keyword}},
			{"ip": bson.M{"$regex": keyword}},
		},
	}

	total, err = c.Find(q).Count()
	if err != nil {
		return
	}

	if total == 0 {
		return
	}

	if err = c.Find(q).Sort("-ct").Skip(offset).Limit(limit).All(&agentAPI); err != nil {
		return
	}

	return
}
