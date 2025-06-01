package matchers

import "battlemap/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Antique struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
	Service      string `json:"service"`
}

func NewAntique() Antique {

	var antique Antique

	antique.Version = "any"
	antique.Architecture = "any"
	antique.Manager = "any"
	antique.Service = "any"

	return antique

}

func ToAntique(value string) Antique {

	var antique Antique

	antique.Version = "any"
	antique.Architecture = "any"
	antique.Manager = "any"
	antique.Service = "any"

	antique.Parse(value)

	return antique

}

func (antique *Antique) IsIdentical(value Antique) bool {

	var result bool = false

	if antique.Name == value.Name &&
		antique.Version == value.Version &&
		antique.Architecture == value.Architecture &&
		antique.Manager == value.Manager &&
		antique.Service == value.Service {
		result = true
	}

	return result

}

func (antique *Antique) IsValid() bool {

	var result bool = false

	if antique.Name != "" {
		result = true
	}

	return result

}

func (antique *Antique) Matches(name string, version string, manager string, service string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(version, " ") {
		version = strings.TrimSpace(version[strings.Index(version, " ")+1:])
	}

	var matches_name bool = false
	var matches_version bool = false
	var matches_manager bool = false
	var matches_service bool = false

	if antique.Name == name {
		matches_name = true
	} else if antique.Name == "any" {
		matches_name = true
	}

	if antique.Version == "any" {

		matches_version = true

	} else if strings.HasPrefix(antique.Version, "<= ") {

		antique_version := types.ToVersion(antique.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(antique_version) {
			matches_version = true
		} else if other_version.IsBefore(antique_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(antique.Version, "< ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsBefore(antique_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(antique.Version, ">= ") {

		antique_version := types.ToVersion(antique.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(antique_version) {
			matches_version = true
		} else if other_version.IsAfter(antique_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(antique.Version, "> ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsAfter(antique_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(antique.Version, "= ") {

		antique_version := types.ToVersion(antique.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(antique_version) {
			matches_version = true
		}

	} else {

		antique_version := types.ToVersion(antique.Version)
		other_version := types.ToVersion(version)

		if other_version.IsSame(antique_version) {
			matches_version = true
		}

	}

	if antique.Manager == manager {
		matches_manager = true
	} else if antique.Manager == "any" {
		matches_manager = true
	}

	if antique.Service == service {
		matches_service = true
	} else if antique.Service == "any" {
		matches_service = true
	}

	return matches_name && matches_version && matches_manager && matches_service

}

func (antique *Antique) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	antique.Name = name
	antique.Version = version

	if architecture != "" {
		antique.Architecture = architecture
	}

}

func (antique *Antique) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		antique.Architecture = architecture.String()
	}

}

func (antique *Antique) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		antique.Manager = manager.String()
	}

}

func (antique *Antique) SetName(value string) {
	antique.Name = strings.TrimSpace(value)
}

func (antique *Antique) SetService(value string) {

	if value == "all" || value == "any" || value == "*" {
		antique.Service = "any"
	} else if value != "" {
		antique.Service = value
	}

}

func (antique *Antique) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		antique.Version = "any"
	} else if value != "" {
		antique.Version = value
	}

}

func (antique *Antique) Hash() string {

	var hash string

	if antique.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			antique.Name,
			antique.Version,
			antique.Architecture,
			antique.Manager,
			antique.Service,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
