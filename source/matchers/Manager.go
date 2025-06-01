package matchers

import "strings"

type Manager struct {
	Name string `json:"name"`
}

func NewManager() Manager {

	var manager Manager

	return manager

}

func ToManager(value string) Manager {

	var manager Manager

	manager.Name = "any"

	manager.SetName(value)

	return manager

}

func (manager *Manager) IsIdentical(value Manager) bool {

	var result bool = false

	if manager.Name == value.Name {
		result = true
	}

	return result

}

func (manager *Manager) IsValid() bool {

	var result bool = false

	if manager.Name != "" {
		result = true
	}

	return result

}

func (manager *Manager) Matches(name string) bool {

	var matches_name bool = false

	if manager.Name == name {
		matches_name = true
	} else if manager.Name == "any" {
		matches_name = true
	}

	return matches_name

}

func (manager *Manager) SetName(value string) {
	manager.Name = strings.TrimSpace(value)
}

func (manager *Manager) Hash() string {

	var hash string

	if manager.Name != "" {
		hash = manager.Name
	}

	return hash

}
