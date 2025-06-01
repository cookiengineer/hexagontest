package matchers

import "battlemap/types"
import "strings"

type Distribution struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Vendor  string `json:"vendor"`
}

func NewDistribution() Distribution {

	var distribution Distribution

	distribution.Version = "any"
	distribution.Vendor = "any"

	return distribution

}

func (distribution *Distribution) IsIdentical(value Distribution) bool {

	var result bool = false

	if distribution.Name == value.Name &&
		distribution.Version == value.Version &&
		distribution.Vendor == value.Vendor {
		result = true
	}

	return result

}

func (distribution *Distribution) IsValid() bool {

	var result bool = false

	if distribution.Name != "" {
		result = true
	}

	return result

}

func (distribution *Distribution) Matches(name string, version string, vendor string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(version, " ") {
		version = strings.TrimSpace(version[strings.Index(version, " ")+1:])
	}

	var matches_name bool = false
	var matches_version bool = false
	var matches_vendor bool = false

	if distribution.Name == name {
		matches_name = true
	} else if distribution.Name == "any" {
		matches_name = true
	} else if strings.HasSuffix(distribution.Name, "-*") {

		prefix := distribution.Name[0:len(distribution.Name)-2]

		if strings.HasPrefix(name, prefix) {
			matches_name = true
		}

	}

	if distribution.Version == "any" {

		matches_version = true

	} else if strings.HasPrefix(distribution.Version, "<= ") {

		distro_version := types.ToVersion(distribution.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(distro_version) {
			matches_version = true
		} else if other_version.IsBefore(distro_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(distribution.Version, "< ") {

		distro_version := types.ToVersion(distribution.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsBefore(distro_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(distribution.Version, ">= ") {

		distro_version := types.ToVersion(distribution.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(distro_version) {
			matches_version = true
		} else if other_version.IsAfter(distro_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(distribution.Version, "> ") {

		distro_version := types.ToVersion(distribution.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsAfter(distro_version) {
			matches_version = true
		}

	} else {

		distro_version := types.ToVersion(distribution.Version)
		other_version := types.ToVersion(version)

		if other_version.IsSame(distro_version) {
			matches_version = true
		}

	}

	if distribution.Vendor == vendor {
		matches_vendor = true
	} else if distribution.Vendor == "any" {
		matches_vendor = true
	}

	return matches_name && matches_version && matches_vendor

}

func (distribution *Distribution) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		distribution.Name = "any"
	} else if value != "" {
		distribution.Name = strings.TrimSpace(value)
	}

}

func (distribution *Distribution) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		distribution.Version = "any"
	} else if value != "" {
		distribution.Version = value
	}

}

func (distribution *Distribution) SetVendor(value string) {
	distribution.Vendor = strings.TrimSpace(value)
}

func (distribution *Distribution) Hash() string {

	var hash string

	if distribution.Name != "" {

		hash = distribution.Name

		if distribution.Version != "any" {
			hash += "-" + distribution.Version
		}

		if distribution.Vendor != "any" {
			hash += "-" + distribution.Vendor
		}

	}

	return hash

}
