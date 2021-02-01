package robot

import "github.com/grearter/rpa-server/api"

func Add(robotAPI *api.Robot) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.Insert(robotAPI); err != nil {
		return
	}

	return
}
