package types

import "strconv"
import "strings"

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewMaintainer() Maintainer {

	var maintainer Maintainer

	return maintainer

}

func ParseMaintainer(value string) *Maintainer {

	var result *Maintainer = nil
	var maintainer Maintainer

	if value != "" && maintainer.Parse(value) == true {
		result = &maintainer
	}

	return result

}

func ToMaintainer(value string) Maintainer {

	var maintainer Maintainer

	if value != "" {
		maintainer.Parse(value)
	}

	return maintainer

}

func (maintainer *Maintainer) IsIdentical(value Maintainer) bool {

	var result bool = false

	if maintainer.Name == value.Name {
		result = true
	}

	return result

}

func (maintainer *Maintainer) IsValid() bool {

	if maintainer.Name != "" {

		if maintainer.Email != "" && strings.Contains(maintainer.Email, "@") {
			return true
		}

	}

	return false

}

func (maintainer *Maintainer) Parse(value string) bool {

	var result bool = false

	if strings.Contains(value, " <") && strings.Contains(value, "@") && strings.HasSuffix(value, ">") {

		tmp := strings.Split(value[:len(value)-1], " <")
		name := strings.TrimSpace(tmp[0])
		email := strings.TrimSpace(tmp[1])

		if name != "" {
			maintainer.Name = name
			result = true
		}

		if email != "" {
			maintainer.Email = email
		}

	} else if strings.HasPrefix(value, "<") && strings.Contains(value, "@") && strings.HasSuffix(value, ">") {

		email := strings.TrimSpace(value[1 : len(value)-2])

		if email != "" {
			maintainer.Email = email
			result = true
		}

	} else if strings.Contains(value, "@") && !strings.Contains(value, " ") {

		email := strings.TrimSpace(value)

		if email != "" {
			maintainer.Email = email
			result = true
		}

	} else {

		name := strings.TrimSpace(value)

		if name != "" {
			maintainer.Name = name
			result = true
		}

	}

	return result

}

func (maintainer Maintainer) MarshalJSON() ([]byte, error) {

	if maintainer.Name != "" && maintainer.Email != "" {
		return []byte(strconv.Quote(maintainer.Name + " <" + maintainer.Email + ">")), nil
	} else if maintainer.Name != "" {
		return []byte(strconv.Quote(maintainer.Name)), nil
	} else if maintainer.Email != "" {
		return []byte(strconv.Quote(maintainer.Email)), nil
	} else {
		return []byte(strconv.Quote("")), nil
	}

}

func (maintainer *Maintainer) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	check := ToMaintainer(unquoted)

	if check.Name != "" {
		maintainer.Name = check.Name
	}

	if check.Email != "" {
		maintainer.Email = check.Email
	}

	return nil

}
