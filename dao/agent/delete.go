package agent

import "github.com/globalsign/mgo"

func Delete(id string) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.RemoveId(id); err != nil {
		if err == mgo.ErrNotFound {
			err = nil
		}
		return
	}

	return
}
