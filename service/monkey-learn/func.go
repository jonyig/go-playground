package monkey_learn

import other_folder "go-playground/service/monkey-learn/other-folder"

func getAdmin() string {
	return "admin"
}

func getUsers() []string {
	admin := getAdmin()
	return []string{admin}
}

func getOtherUsers() []string {
	admin := other_folder.GetAdmin1()
	return []string{admin}
}
