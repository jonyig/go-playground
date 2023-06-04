package monkey_learn

func getAdmin() string {
	return "admin"
}

func getUsers() []string {
	admin := getAdmin()
	return []string{admin}
}
