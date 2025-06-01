package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Drive struct {
	Name       string `json:"name"`
	Mountpoint string `json:"mountpoint"`
	Type       string `json:"type"`
}

func NewDrive() Drive {

	var drive Drive

	drive.Mountpoint = "any"
	drive.Type = "any"

	return drive

}

func ToDrive(value string) Drive {

	var drive Drive

	drive.Type = "any"

	if strings.HasPrefix(value, "/") {
		drive.Name = "any"
		drive.SetMountpoint(value)
	} else if value != "" {
		drive.SetName(value)
		drive.Mountpoint = "any"
	}

	return drive

}

func (drive *Drive) IsIdentical(value Drive) bool {

	var result bool = false

	if drive.Name == value.Name &&
		drive.Mountpoint == value.Mountpoint &&
		drive.Type == value.Type {
		result = true
	}

	return result

}

func (drive *Drive) IsValid() bool {

	var result bool = false

	if drive.Name != "" {
		result = true
	}

	return result

}

func (drive *Drive) Matches(name string, mountpoint string, typ string) bool {

	var matches_name bool = false
	var matches_mountpoint bool = false
	var matches_type bool = false

	if drive.Name == name {
		matches_name = true
	} else if drive.Name == "any" {
		matches_name = true
	}

	if drive.Mountpoint == mountpoint {
		matches_mountpoint = true
	} else if drive.Mountpoint == "any" {
		matches_mountpoint = true
	}

	if drive.Type == typ {
		matches_type = true
	} else if drive.Type == "any" {
		matches_type = true
	}

	return matches_name && matches_mountpoint && matches_type

}

func (drive *Drive) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Name = "any"
	} else if value != "" {
		drive.Name = strings.TrimSpace(value)
	}

}

func (drive *Drive) SetMountpoint(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Mountpoint = "any"
	} else if value != "" {
		drive.Mountpoint = strings.TrimSpace(value)
	}

}

func (drive *Drive) SetType(value string) {

	if value == "all" || value == "any" || value == "*" {
		drive.Type = "any"
	} else if value != "" {
		drive.Type = strings.TrimSpace(value)
	}

}

func (drive *Drive) Hash() string {

	var hash string

	if drive.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			drive.Name,
			drive.Mountpoint,
			drive.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
