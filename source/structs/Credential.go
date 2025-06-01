package structs

import "bytes"
import "strings"

type Credential struct {
	Name      string   `json:"name"`
	Keys      [][]byte `json:"keys"`
	Passwords []string `json:"passwords"`
	Type      string   `json:"type"`
}

func NewCredential(name string, typ string) Credential {

	var credential Credential

	credential.SetName(name)
	credential.SetType(typ)
	credential.Keys = make([][]byte, 0)
	credential.Passwords = make([]string, 0)

	return credential

}

func (credential *Credential) IsValid() bool {

	var result bool = false

	if credential.Name != "" && credential.Type != "" {

		if len(credential.Keys) > 0 || len(credential.Passwords) > 0 {
			result = true
		}

	}

	return result

}

func (credential *Credential) SetName(value string) {
	credential.Name = strings.TrimSpace(value)
}

func (credential *Credential) AddKey(value []byte) {

	var found bool = false

	for k := 0; k < len(credential.Keys); k++ {

		if bytes.Equal(credential.Keys[k], value) {
			found = true
			break
		}

	}

	if found == false {
		credential.Keys = append(credential.Keys, value)
	}

}

func (credential *Credential) RemoveKey(value []byte) {

	var index int = -1

	for k := 0; k < len(credential.Keys); k++ {

		if bytes.Equal(credential.Keys[k], value) {
			index = k
			break
		}

	}

	if index != -1 {
		credential.Keys = append(credential.Keys[:index], credential.Keys[index+1:]...)
	}

}

func (credential *Credential) SetKeys(values [][]byte) {

	var filtered [][]byte

	for v := 0; v < len(values); v++ {
		filtered = append(filtered, values[v])
	}

	credential.Keys = filtered

}

func (credential *Credential) AddPassword(value string) {

	value = strings.TrimSpace(value)

	var found bool = false

	for p := 0; p < len(credential.Passwords); p++ {

		if credential.Passwords[p] == value {
			found = true
			break
		}

	}

	if found == false {
		credential.Passwords = append(credential.Passwords, value)
	}

}

func (credential *Credential) RemovePassword(value string) {

	value = strings.TrimSpace(value)

	var index int = -1

	for p := 0; p < len(credential.Passwords); p++ {

		if credential.Passwords[p] == value {
			index = p
			break
		}

	}

	if index != -1 {
		credential.Passwords = append(credential.Passwords[:index], credential.Passwords[index+1:]...)
	}

}

func (credential *Credential) SetPasswords(values []string) {

	var filtered []string

	for v := 0; v < len(values); v++ {
		filtered = append(filtered, strings.TrimSpace(values[v]))
	}

	credential.Passwords = filtered

}

func (credential *Credential) SetType(value string) {

	if value == "user" {
		credential.Type = "user"
	} else if value == "program" {
		credential.Type = "program"
	}

}
