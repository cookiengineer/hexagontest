package matchers

import "battlemap/types"
import "strings"

type Timezone struct {
	Name   string `json:"name"`
	Offset string `json:"offset"`
}

func NewTimezone() Timezone {

	var timezone Timezone

	timezone.Name = "any"
	timezone.Offset = "any"

	return timezone

}

func ToTimezone(value string) Timezone {

	var timezone Timezone

	timezone.Name = "any"
	timezone.Offset = "any"

	timezone.Parse(value)

	return timezone

}

func (timezone *Timezone) IsIdentical(value Timezone) bool {

	var result bool = false

	if timezone.Name == value.Name && timezone.Offset == value.Offset {
		result = true
	}

	return result

}

func (timezone *Timezone) IsValid() bool {

	var result bool = false

	if timezone.Name != "" {
		result = true
	}

	return result

}

func (timezone *Timezone) Matches(name string, offset string) bool {

	// Compatibility with "<operator> <offset>" syntax
	if strings.Contains(offset, " ") {
		offset = strings.TrimSpace(offset[strings.Index(offset, " ")+1:])
	}

	var matches_name bool = false
	var matches_offset bool = false

	if timezone.Name == name {
		matches_name = true
	} else if timezone.Name == "any" {
		matches_name = true
	} else if strings.HasPrefix(timezone.Name, "*/") {

		if strings.HasSuffix(name, timezone.Name[strings.Index(timezone.Name, "*/")+1:]) {
			matches_name = true
		}

	} else if strings.HasSuffix(timezone.Name, "/*") {

		if strings.HasPrefix(name, timezone.Name[0:strings.Index(timezone.Name, "/*")+1]) {
			matches_name = true
		}

	}

	if timezone.Offset == "any" {
		matches_offset = true
	} else if strings.HasPrefix(timezone.Offset, "<= ") {

		timezone_offset := types.NewTime()
		timezone_offset.Offset(timezone.Offset[3:])

		other_offset := types.NewTime()
		other_offset.Offset(offset)

		if other_offset.IsSame(timezone_offset) {
			matches_offset = true
		} else if other_offset.IsBefore(timezone_offset) {
			matches_offset = true
		}

	} else if strings.HasPrefix(timezone.Offset, "< ") {

		timezone_offset := types.NewTime()
		timezone_offset.Offset(timezone.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(offset)

		if other_offset.IsBefore(timezone_offset) {
			matches_offset = true
		}

	} else if strings.HasPrefix(timezone.Offset, ">= ") {

		timezone_offset := types.NewTime()
		timezone_offset.Offset(timezone.Offset[3:])

		other_offset := types.NewTime()
		other_offset.Offset(offset)

		if other_offset.IsSame(timezone_offset) {
			matches_offset = true
		} else if other_offset.IsAfter(timezone_offset) {
			matches_offset = true
		}

	} else if strings.HasPrefix(timezone.Offset, "> ") {

		timezone_offset := types.NewTime()
		timezone_offset.Offset(timezone.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(offset)

		if other_offset.IsAfter(timezone_offset) {
			matches_offset = true
		}

	} else if strings.HasPrefix(timezone.Offset, "= ") {

		timezone_offset := types.NewTime()
		timezone_offset.Offset(timezone.Offset[2:])

		other_offset := types.NewTime()
		other_offset.Offset(offset)

		if other_offset.IsSame(timezone_offset) {
			matches_offset = true
		}

	} else {

		timezone_offset := types.NewTime()
		timezone_offset.Offset(timezone.Offset)

		other_offset := types.NewTime()
		other_offset.Offset(offset)

		if other_offset.IsSame(timezone_offset) {
			matches_offset = true
		}

	}

	return matches_name && matches_offset

}

func (timezone *Timezone) Parse(value string) {

	name, offset := parseOffsetCondition(value)

	timezone.Name = name
	timezone.Offset = offset

}

func (timezone *Timezone) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		timezone.Name = "any"
	} else if value != "" {
		timezone.Name = value
	}

}

func (timezone *Timezone) SetOffset(value string) {

	if value == "all" || value == "any" || value == "*" {
		timezone.Name = "any"
	} else if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		if len(value) == 6 && string(value[3]) == ":" {
			timezone.Offset = value
		}

	}

}

func (timezone *Timezone) Hash() string {

	var hash string

	if timezone.Name != "" {

		hash = strings.Join([]string{
			timezone.Name,
			timezone.Offset,
		}, "-")

	}

	return hash

}
