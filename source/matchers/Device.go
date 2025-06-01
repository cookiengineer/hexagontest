package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Device struct {
	Name   string `json:"name"`
	Bus    string `json:"bus"`
	System struct {
		Device string `json:"device"`
		Vendor string `json:"vendor"`
		Name   string `json:"name"`
	} `json:"system"`
	Subsystem struct {
		Device string `json:"device"`
		Vendor string `json:"vendor"`
		Name   string `json:"name"`
	} `json:"subsystem"`
}

func NewDevice() Device {

	var device Device

	device.Name = "any"
	device.Bus = "any"

	return device

}

func ToDevice(value string) Device {

	var device Device

	device.Name = "any"
	device.Bus = "any"

	device.SetName(value)

	return device

}

func (device *Device) IsIdentical(value Device) bool {

	var result bool = false

	if device.Name == value.Name &&
		device.Bus == value.Bus &&
		device.System.Device == value.System.Device &&
		device.System.Vendor == value.System.Vendor &&
		device.System.Name == value.System.Name &&
		device.Subsystem.Device == value.Subsystem.Device &&
		device.Subsystem.Vendor == value.Subsystem.Vendor &&
		device.Subsystem.Name == value.Subsystem.Name {
		result = true
	}

	return result

}

func (device *Device) IsValid() bool {

	var result bool = false

	if device.Bus != "" {

		if device.Name != "" {
			result = true
		} else {

			if device.System.Device != "" && device.System.Vendor != "" {
				result = true
			} else if device.Subsystem.Device != "" && device.Subsystem.Vendor != "" {
				result = true
			}

		}

	}

	return result

}

func (dev *Device) Matches(name string, bus string) bool {

	var matches_name bool = false
	var matches_bus bool = false

	if dev.Name == name {
		matches_name = true
	} else if dev.Name == "any" {
		matches_name = true
	}

	if dev.Bus == bus {
		matches_bus = true
	} else if dev.Bus == "any" {
		matches_bus = true
	}

	return matches_name && matches_bus

}

func (dev *Device) MatchesSystem(vendor string, device string, name string) bool {

	var matches_vendor bool = false
	var matches_device bool = false
	var matches_name bool = false

	if dev.System.Vendor == vendor {
		matches_vendor = true
	} else if vendor == "any" {
		matches_vendor = true
	}

	if dev.System.Device == device {
		matches_device = true
	} else if device == "any" {
		matches_device = true
	}

	if dev.System.Name == name {
		matches_name = true
	} else if name == "any" {
		matches_name = true
	}

	return matches_vendor && matches_device && matches_name

}

func (dev *Device) MatchesSubsystem(vendor string, device string, name string) bool {

	var matches_vendor bool = false
	var matches_device bool = false
	var matches_name bool = false

	if dev.Subsystem.Vendor == vendor {
		matches_vendor = true
	} else if vendor == "any" {
		matches_vendor = true
	}

	if dev.Subsystem.Device == device {
		matches_device = true
	} else if device == "any" {
		matches_device = true
	}

	if dev.Subsystem.Name == name {
		matches_name = true
	} else if name == "any" {
		matches_name = true
	}

	return matches_vendor && matches_device && matches_name

}

func (dev *Device) SetName(value string) {
	dev.Name = strings.TrimSpace(value)
}

func (dev *Device) SetBus(value string) {

	if value == "any" {
		dev.Bus = "any"
	} else if value == "hid" {
		dev.Bus = "hid"
	} else if value == "i2c" {
		dev.Bus = "i2c"
	} else if value == "pci" {
		dev.Bus = "pci"
	} else if value == "scsi" {
		dev.Bus = "scsi"
	} else if value == "usb" {
		dev.Bus = "usb"
	} else if value == "other" {
		dev.Bus = "other"
	}

}

func (dev *Device) SetSystem(vendor string, device string, name string) {

	dev.System.Vendor = strings.TrimSpace(vendor)
	dev.System.Device = strings.TrimSpace(device)
	dev.System.Name = strings.TrimSpace(name)

}

func (dev *Device) SetSubsystem(vendor string, device string, name string) {

	dev.Subsystem.Vendor = strings.TrimSpace(vendor)
	dev.Subsystem.Device = strings.TrimSpace(device)
	dev.Subsystem.Name = strings.TrimSpace(name)

}

func (device *Device) Hash() string {

	var hash string

	if device.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			device.Name,
			device.Bus,
			device.System.Device,
			device.System.Vendor,
			device.System.Name,
			device.Subsystem.Device,
			device.Subsystem.Vendor,
			device.Subsystem.Name,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
