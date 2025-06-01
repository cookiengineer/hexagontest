package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Incident struct {
	Type    string `json:"type"`
	Keyword string `json:"keyword"`
}

func NewIncident() Incident {

	var incident Incident

	incident.Keyword = "any"

	return incident

}

func ToIncident(value string) Incident {

	var incident Incident

	incident.Keyword = "any"

	if strings.Contains(value, ":") {
		incident.SetType(value[0:strings.Index(value, ":")])
		incident.SetKeyword(value[strings.Index(value, ":"):])
	} else {
		incident.SetType(value)
	}

	return incident

}

func (incident *Incident) IsIdentical(value Incident) bool {

	var result bool = false

	if incident.Type == value.Type &&
		incident.Keyword == value.Keyword {
		result = true
	}

	return result

}

func (incident *Incident) IsValid() bool {

	var result bool = false

	if incident.Type != "" {
		result = true
	}

	return result

}

func (incident *Incident) MatchesType(typ string) bool {

	var matches_type bool = false

	if incident.Type == typ {
		matches_type = true
	} else if incident.Type == "any" {
		matches_type = true
	}

	return matches_type

}

func (incident *Incident) MatchesKeyword(value string) bool {

	var matches_keyword bool = false

	if incident.Keyword == value {
		matches_keyword = true
	} else if incident.Keyword == "any" {
		matches_keyword = true
	}

	return matches_keyword

}

func (incident *Incident) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		incident.Type = "any"
	} else if value != "" {
		incident.Type = strings.TrimSpace(value)
	}

}

func (incident *Incident) SetKeyword(value string) {

	if value == "all" || value == "any" || value == "*" {
		incident.Keyword = "any"
	} else if value != "" {
		incident.Keyword = strings.TrimSpace(value)
	}

}

func (incident *Incident) Hash() string {

	var hash string

	if incident.Type != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			incident.Type,
			incident.Keyword,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
