package structs

import utils_strings "battlemap/utils/strings"
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

func stripArchitecture(value string) string {

	if strings.HasPrefix(value, "x86_64_") {
		value = strings.TrimSpace(value[7:])
	} else if strings.HasPrefix(value, "x86_") {
		value = strings.TrimSpace(value[4:])
	} else if strings.HasSuffix(value, "_x86_64") {
		value = strings.TrimSpace(value[0 : len(value)-7])
	} else if strings.HasSuffix(value, "_x86") {
		value = strings.TrimSpace(value[0 : len(value)-4])
	} else if strings.HasSuffix(value, "_32-bit") {
		value = strings.TrimSpace(value[0 : len(value)-7])
	} else if strings.HasSuffix(value, "_64-bit") {
		value = strings.TrimSpace(value[0 : len(value)-7])
	} else if value == "x86_64" {
		value = ""
	} else if value == "x86" {
		value = ""
	} else if value == "32-bit" {
		value = ""
	} else if value == "64-bit" {
		value = ""
	} else if value == "*" {
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
	State   string `json:"state"`
}

func NewProduct(typ string) Product {

	var product Product

	product.Name = ""
	product.Product = ""
	product.Version = ""
	product.Vendor = ""
	product.State = "invalid"
	product.SetType(typ)

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
		product.Version == value.Version &&
		product.Vendor == value.Vendor &&
		product.Type == value.Type {
		result = true
	}

	return result

}

func (product *Product) IsValid() bool {

	if product.Name != "" && product.State != "invalid" {

		if product.Vendor != "" && product.Type != "" {
			return true
		}

	}

	return false

}

func (product *Product) SetProduct(value string) {

	if value == "*" || value == "-" {
		product.Product = ""
		product.Name = ""
	} else if value != "" {
		product.Product = utils_strings.ToASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product
	}

}

func (product *Product) SetState(value string) {

	if value == "edited" {
		product.State = "edited"
	} else if value == "published" {
		product.State = "published"
	} else if value == "invalid" {
		product.State = "invalid"
	}

}

func (product *Product) SetType(value string) {

	if value == "hardware" {
		product.Type = "hardware"
	} else if value == "software" {
		product.Type = "software"
	} else if value == "system" {
		product.Type = "system"
	}

}

func (product *Product) SetVendor(value string) {

	if value == "*" || value == "-" {
		product.Vendor = ""
		product.Name = ""
	} else if value != "" {
		product.Vendor = utils_strings.ToASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product
	}

}

func (product *Product) SetVersion(value string) {

	if value == "any" || value == "*" || value == "-" {
		product.Version = "any"
	} else if value != "" {
		product.Version = value
	}

}
