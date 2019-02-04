package models

type Profile struct {
	UserName string `xorm:"pk not null"`
	Name     string
	Company  string
	Position string
	Password string
}

func ReadProfile() []Profile {
	var profile = make([]Profile, 0)
	err := engine.Find(&profile)
	printError(err)

	return profile
}

func CreateProfile(username, name, company, position, password string) error {
	var profile = Profile{
		UserName: username,
		Name:     name,
		Company:  company,
		Position: position,
		Password: password,
	}
	_, err := engine.Insert(profile)

	return err
}

func MatchUsernamePass(username, pass string) (bool, error) {
	var profile = Profile{
		UserName: username,
		Password: pass,
	}
	found, err := engine.Get(&profile)

	return found, err
}

func printError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
