package lcr

func dynamicPassword(password string, target int) string {
	newPass := password + password
	return newPass[target : len(password)+target]
}
