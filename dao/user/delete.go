package user

func Delete(id string) (err error) {
	c := getC()
	defer c.Database.Session.Close()

	if err = c.RemoveId(id); err != nil {
		return
	}

	return
}
