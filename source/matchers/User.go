package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func NewUser() User {

	var user User

	user.Password = "any"
	user.Type = "any"

	return user

}

func ToUser(value string) User {

	var user User

	user.Password = "any"
	user.Type = "any"

	user.SetName(value)

	return user

}

func (user *User) IsIdentical(value User) bool {

	var result bool = false

	if user.Name == value.Name &&
		user.Password == value.Password &&
		user.Type == value.Type {
		result = true
	}

	return result

}

func (user *User) IsValid() bool {

	var result bool = false

	if user.Name != "" && user.Password != "" {
		result = true
	}

	return result

}

func (user *User) Matches(name string, password string, typ string) bool {

	var matches_name bool = false
	var matches_password bool = false
	var matches_type bool = false

	if user.Name == name {
		matches_name = true
	} else if user.Name == "any" {
		matches_name = true
	}

	if user.Password == password {
		matches_password = true
	} else if user.Password == "any" {
		matches_password = true
	}

	if user.Type == typ {
		matches_type = true
	} else if user.Type == "any" {
		matches_type = true
	}

	return matches_name && matches_password && matches_type

}

func (user *User) SetName(value string) {
	user.Name = strings.TrimSpace(value)
}

func (user *User) SetPassword(value string) {

	if value == "all" || value == "any" || value == "*" {
		user.Password = "any"
	} else if value != "" {
		user.Password = value
	}

}

func (user *User) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		user.Type = "any"
	} else if value == "user" {
		user.Type = "user"
	} else if value == "program" {
		user.Type = "program"
	}

}

func (user *User) Hash() string {

	var hash string

	if user.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			user.Name,
			user.Password,
			user.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
