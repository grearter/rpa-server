package agent

import "github.com/grearter/rpa-server/api"

func Add(agentAPI *api.Agent) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.Insert(c); err != nil {
		return
	}

	return
}
