package matchers

import "battlemap/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Update struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	Manager      string `json:"manager"`
}

func NewUpdate() Update {

	var update Update

	update.Version = "any"
	update.Architecture = "any"
	update.Manager = "any"

	return update

}

func ToUpdate(value string) Update {

	var update Update

	update.Version = "any"
	update.Architecture = "any"
	update.Manager = "any"

	update.Parse(value)

	return update

}

func (update *Update) IsIdentical(value Update) bool {

	var result bool = false

	if update.Name == value.Name &&
		update.Version == value.Version &&
		update.Architecture == value.Architecture &&
		update.Manager == value.Manager {
		result = true
	}

	return result

}

func (update *Update) IsValid() bool {

	var result bool = false

	if update.Name != "" {
		result = true
	}

	return result

}

func (update *Update) Matches(name string, version string, manager string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(version, " ") {
		version = strings.TrimSpace(version[strings.Index(version, " ")+1:])
	}

	var matches_name bool = false
	var matches_version bool = false
	var matches_manager bool = false

	if update.Name == name {
		matches_name = true
	} else if update.Name == "any" {
		matches_name = true
	}

	if update.Version == "any" {

		matches_version = true

	} else if strings.HasPrefix(update.Version, "<= ") {

		update_version := types.ToVersion(update.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(update_version) {
			matches_version = true
		} else if other_version.IsBefore(update_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(update.Version, "< ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsBefore(update_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(update.Version, ">= ") {

		update_version := types.ToVersion(update.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(update_version) {
			matches_version = true
		} else if other_version.IsAfter(update_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(update.Version, "> ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsAfter(update_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(update.Version, "= ") {

		update_version := types.ToVersion(update.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(update_version) {
			matches_version = true
		}

	} else {

		update_version := types.ToVersion(update.Version)
		other_version := types.ToVersion(version)

		if other_version.IsSame(update_version) {
			matches_version = true
		}

	}

	if update.Manager == manager {
		matches_manager = true
	} else if update.Manager == "any" {
		matches_manager = true
	}

	return matches_name && matches_version && matches_manager

}

func (update *Update) Parse(value string) {

	name, version, architecture := parseVersionCondition(value)

	update.Name = name
	update.Version = version

	if architecture != "" {
		update.Architecture = architecture
	}

}

func (update *Update) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		update.Architecture = architecture.String()
	}

}

func (update *Update) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		update.Manager = manager.String()
	}

}

func (update *Update) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		update.Name = "any"
	} else if value != "" {
		update.Name = strings.TrimSpace(value)
	}

}

func (update *Update) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" {
		update.Version = "any"
	} else if value != "" {
		update.Version = value
	}

}

func (update *Update) Hash() string {

	var hash string

	if update.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			update.Name,
			update.Version,
			update.Architecture,
			update.Manager,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
