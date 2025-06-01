package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Country struct {
	Name       string `json:"name"`
	Continent  string `json:"continent"`
	Allegiance string `json:"allegiance"`
}

func NewCountry() Country {

	var country Country

	country.Continent = "any"
	country.Allegiance = "any"

	return country

}

func ToCountry(value string) Country {

	var country Country

	country.Name = "any"
	country.Continent = "any"
	country.Allegiance = "any"

	country.SetName(value)

	return country

}

func (country *Country) IsIdentical(value Country) bool {

	var result bool = false

	if country.Name == value.Name &&
		country.Continent == value.Continent &&
		country.Allegiance == value.Allegiance {
		result = true
	}

	return result

}

func (country *Country) IsValid() bool {

	var result bool = false

	if country.Name != "" {
		result = true
	}

	return result

}

func (country *Country) Matches(name string, continent string, allegiance string) bool {

	var matches_name bool = false
	var matches_continent bool = false
	var matches_allegiance bool = false

	if country.Name == name {
		matches_name = true
	} else if country.Name == "any" {
		matches_name = true
	}

	if country.Continent == continent {
		matches_continent = true
	} else if country.Continent == "any" {
		matches_continent = true
	}

	if country.Allegiance == allegiance {
		matches_allegiance = true
	} else if country.Allegiance == "any" {
		matches_allegiance = true
	}

	return matches_name && matches_continent && matches_allegiance

}

func (country *Country) SetAllegiance(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Allegiance = "any"
	} else if value != "" {
		country.Allegiance = value
	}

}

func (country *Country) SetName(value string) {
	country.Name = strings.TrimSpace(value)
}

func (country *Country) SetContinent(value string) {

	if value == "all" || value == "any" || value == "*" {
		country.Continent = "any"
	} else if value != "" {
		country.Continent = value
	}

}

func (country *Country) Hash() string {

	var hash string

	if country.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			country.Name,
			country.Continent,
			country.Allegiance,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
