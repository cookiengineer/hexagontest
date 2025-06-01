package matchers

import "battlemap/types"
import utils_strings "battlemap/utils/strings"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

func parseCPEValue(value string) string {

	value = utils_strings.ToASCII(strings.TrimSpace(value))

	if value == "*" {
		value = ""
	} else if value == "-" {
		value = ""
	}

	return value

}

type Product struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Product string `json:"product"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

func NewProduct() Product {

	var product Product

	product.Name = "any:any"
	product.Product = "any"
	product.Version = "any"
	product.Vendor = "any"
	product.Type = "any"

	return product

}

func ToProduct(value string) Product {

	var product Product

	if strings.HasPrefix(value, "cpe:") {

		product.Version = "any"
		product.Vendor = "any"
		product.Type = "any"

		value = strings.ReplaceAll(value[4:], "\\:", "%COLON%")
		tmp := strings.Split(strings.TrimSpace(value), ":")

		for t := 0; t < len(tmp); t++ {
			tmp[t] = strings.ReplaceAll(tmp[t], "%COLON%", ":")
		}

		if len(tmp) >= 4 && tmp[0] == "2.3" {

			cpe_part := parseCPEValue(tmp[1])

			if cpe_part == "/a" || cpe_part == "a" {
				product.SetType("software")
			} else if cpe_part == "/h" || cpe_part == "h" {
				product.SetType("hardware")
			} else if cpe_part == "/o" || cpe_part == "o" {
				product.SetType("system")
			}

			cpe_vendor := parseCPEValue(tmp[2])

			if cpe_vendor != "" {
				product.SetVendor(cpe_vendor)
			}

			cpe_product := parseCPEValue(tmp[3])

			if cpe_product != "" {
				product.SetProduct(cpe_product)
			}

			if len(tmp) >= 6 {

				cpe_version := parseCPEValue(tmp[4])
				cpe_update := parseCPEValue(tmp[5])

				cpe_version = stripArchitecture(cpe_version)
				cpe_update = stripArchitecture(cpe_update)

				if cpe_version != "" && cpe_update != "" {
					product.SetVersion(cpe_version + "-" + cpe_update)
				} else if cpe_version != "" {
					product.SetVersion(cpe_version)
				} else if cpe_update != "" {
					product.SetVersion(cpe_update)
				}

			} else if len(tmp) == 5 {

				cpe_version := parseCPEValue(tmp[4])
				cpe_version = stripArchitecture(cpe_version)

				if cpe_version != "" {
					product.SetVersion(cpe_version)
				}

			}

		} else if len(tmp) >= 3 {

			cpe_part := parseCPEValue(tmp[0])

			if cpe_part == "/a" {
				product.SetType("software")
			} else if cpe_part == "/h" {
				product.SetType("hardware")
			} else if cpe_part == "/o" {
				product.SetType("system")
			}

			cpe_vendor := parseCPEValue(tmp[1])

			if cpe_vendor != "" {
				product.SetVendor(cpe_vendor)
			}

			cpe_product := parseCPEValue(tmp[2])

			if cpe_product != "" {
				product.SetProduct(cpe_product)
			}

			if len(tmp) >= 5 {

				cpe_version := parseCPEValue(tmp[3])
				cpe_update := parseCPEValue(tmp[4])

				cpe_version = stripArchitecture(cpe_version)
				cpe_update = stripArchitecture(cpe_update)

				if cpe_version != "" && cpe_update != "" {
					product.SetVersion(cpe_version + "-" + cpe_update)
				} else if cpe_version != "" {
					product.SetVersion(cpe_version)
				} else if cpe_update != "" {
					product.SetVersion(cpe_update)
				}

			} else if len(tmp) == 4 {

				cpe_version := parseCPEValue(tmp[3])
				cpe_version = stripArchitecture(cpe_version)

				if cpe_version != "" {
					product.SetVersion(cpe_version)
				}

			}

		}

	}

	return product

}

func (product *Product) IsIdentical(value Product) bool {

	var result bool = false

	if product.Name == value.Name &&
		product.Vendor == value.Vendor &&
		product.Product == value.Product &&
		product.Version == value.Version &&
		product.Type == value.Type {
		result = true
	}

	return result

}

func (product *Product) IsValid() bool {

	var result bool = false

	if product.Name != "" && product.Name != "any:any" {
		result = true
	}

	return result

}

func (product *Product) Matches(name string, version string, vendor string, typ string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(version, " ") {
		version = strings.TrimSpace(version[strings.Index(version, " ")+1:])
	}

	var matches_product bool = false
	var matches_version bool = false
	var matches_vendor bool = false
	var matches_type bool = false

	if product.Product == name {
		matches_product = true
	} else if product.Product == "any" {
		matches_product = true
	}

	if product.Version == "any" {

		matches_version = true

	} else if strings.HasPrefix(product.Version, "<= ") {

		product_version := types.ToVersion(product.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(product_version) {
			matches_version = true
		} else if other_version.IsBefore(product_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(product.Version, "< ") {

		product_version := types.ToVersion(product.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsBefore(product_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(product.Version, ">= ") {

		product_version := types.ToVersion(product.Version[3:])
		other_version := types.ToVersion(version)

		if other_version.IsSame(product_version) {
			matches_version = true
		} else if other_version.IsAfter(product_version) {
			matches_version = true
		}

	} else if strings.HasPrefix(product.Version, "> ") {

		product_version := types.ToVersion(product.Version[2:])
		other_version := types.ToVersion(version)

		if other_version.IsAfter(product_version) {
			matches_version = true
		}

	} else {

		product_version := types.ToVersion(product.Version)
		other_version := types.ToVersion(version)

		if other_version.IsSame(product_version) {
			matches_version = true
		}

	}

	if product.Vendor == vendor {
		matches_vendor = true
	} else if product.Vendor == "any" {
		matches_vendor = true
	}

	if product.Type == typ {
		matches_type = true
	} else if product.Type == "any" {
		matches_type = true
	}

	return matches_product && matches_version && matches_vendor && matches_type

}

func (product *Product) SetProduct(value string) {

	if value == "any" || value == "*" || value == "-" {
		product.Product = "any"
		product.Name = product.Vendor + ":" + product.Product
	} else if value != "" {
		product.Product = utils_strings.ToASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product
	}

}

func (product *Product) SetType(value string) {

	if value == "any" || value == "*" || value == "-" {
		product.Type = "any"
	} else if value == "hardware" {
		product.Type = "hardware"
	} else if value == "software" {
		product.Type = "software"
	} else if value == "system" {
		product.Type = "system"
	}

}

func (product *Product) SetVendor(value string) {

	if value == "any" || value == "*" || value == "-" {
		product.Vendor = "any"
		product.Name = product.Vendor + ":" + product.Product
	} else if value != "" {
		product.Vendor = utils_strings.ToASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product
	}

}

func (product *Product) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" || value == "-" {
		product.Version = "any"
	} else if value != "" {
		product.Version = value
	}

}

func (product *Product) Hash() string {

	var hash string

	if product.Name != "any:any" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			product.Name,
			product.Vendor,
			product.Product,
			product.Version,
			product.Type,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
