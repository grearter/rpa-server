package agent

import (
	"github.com/globalsign/mgo/bson"
	"github.com/grearter/rpa-server/api"
)

func Get(id bson.ObjectId) (agentAPI *api.Agent, err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.FindId(id).One(&agentAPI); err != nil {
		return
	}

	return
}
