package robottask

import "github.com/grearter/rpa-server/api"

func Add(robotTaskAPI *api.RobotTask) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.Insert(robotTaskAPI); err != nil {
		return
	}

	return
}
