package types

import "strings"

type User struct {
	ID       uint16  `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Folder   string  `json:"folder"`
	Groups   []Group `json:"groups"`
	Shell    string  `json:"shell"`
	Type     string  `json:"type"`
}

func NewUser() User {

	var user User

	user.Groups = make([]Group, 0)

	return user

}

func ToUser(name string, id uint16) User {

	var user User

	if id == 65535 {
		user.ID = id
	} else {
		user.SetID(id)
	}

	user.SetName(name)
	user.Groups = make([]Group, 0)
	user.SetType("user")

	return user

}

func (user *User) IsValid() bool {

	if user.Name != "" {

		var result bool = true

		if user.ID == 0 && user.Name != "root" {
			result = false
		}

		return result

	}

	return false

}

func (user *User) SetID(value uint16) {

	// User "nobody" is last user id
	if value >= 0 && value <= 65534 {
		user.ID = value
	}

}

func (user *User) SetName(value string) {
	user.Name = strings.TrimSpace(value)
}

func (user *User) SetPassword(value string) {
	user.Password = value
}

func (user *User) SetFolder(value string) {

	if strings.HasPrefix(value, "/") {
		user.Folder = value
	}

}

func (user *User) AddGroup(value Group) {

	if value.IsValid() {

		var found bool = false

		for g := 0; g < len(user.Groups); g++ {

			if user.Groups[g].ID == value.ID {
				found = true
				break
			}

		}

		if found == false {
			user.Groups = append(user.Groups, value)
		}

	}

}

func (user *User) RemoveGroup(value Group) {

	var index int = -1

	for g := 0; g < len(user.Groups); g++ {

		if user.Groups[g].ID == value.ID {
			index = g
			break
		}

	}

	if index != -1 {
		user.Groups = append(user.Groups[:index], user.Groups[index+1:]...)
	}

}

func (user *User) SetGroups(value []Group) {

	var filtered []Group

	for v := 0; v < len(value); v++ {

		var group = value[v]

		if group.IsValid() {
			filtered = append(filtered, group)
		}

	}

	user.Groups = filtered

}

func (user *User) SetShell(value string) {

	if strings.HasPrefix(value, "/") {
		user.Shell = strings.TrimSpace(value)
	}

}

func (user *User) SetType(value string) {

	if value == "user" {
		user.Type = "user"
	} else if value == "program" {
		user.Type = "program"
	}

}
