package user

import "github.com/grearter/rpa-server/api"

func Add(u *api.User) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.Insert(u); err != nil {
		return
	}

	return
}
