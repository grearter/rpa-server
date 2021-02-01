package user

// Exist 用户是否存在
func Exist(id string) (ok bool, err error) {
	c := getC()
	defer c.Database.Session.Close()

	cnt, err := c.FindId(id).Count()
	if err != nil {
		return
	}

	ok = cnt > 0
	return
}
