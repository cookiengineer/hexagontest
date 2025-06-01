package matchers

import "battlemap/types"
import "encoding/hex"
import "strings"

type Provider struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Subnet  string `json:"subnet"`
}

func NewProvider() Provider {

	var provider Provider

	provider.Name = "any"
	provider.Country = "any"
	provider.Subnet = "any"

	return provider

}

func ToProvider(value string) Provider {

	var provider Provider

	provider.Name = "any"
	provider.Country = "any"
	provider.Subnet = "any"

	provider.SetName(value)

	return provider

}

func (provider *Provider) IsIdentical(value Provider) bool {

	var result bool = false

	if provider.Name == value.Name &&
		provider.Country == value.Country &&
		provider.Subnet == value.Subnet {
		result = true
	}

	return result

}

func (provider *Provider) IsValid() bool {

	var result bool = false

	if provider.Name != "any" || provider.Country != "any" || provider.Subnet != "any" {
		result = true
	}

	return result

}

func (provider *Provider) Matches(name string, country string, subnet string) bool {

	var matches_name bool = false
	var matches_country bool = false
	var matches_subnet bool = false

	if provider.Name == name {
		matches_name = true
	} else if provider.Name == "any" {
		matches_name = true
	}

	if provider.Country == country {
		matches_country = true
	} else if provider.Country == "any" {
		matches_country = true
	}

	if provider.Subnet != "any" && subnet != "any" && containsSubnet(provider.Subnet, subnet) {
		matches_subnet = true
	} else if provider.Subnet == "any" {
		matches_subnet = true
	}

	return matches_name && matches_country && matches_subnet

}

func (provider *Provider) SetName(value string) {
	provider.Name = strings.TrimSpace(value)
}

func (provider *Provider) SetCountry(value string) {

	if value == "all" || value == "any" || value == "*" {
		provider.Country = "any"
	} else if value != "" {
		provider.Country = value
	}

}

func (provider *Provider) SetSubnet(value string) {

	if value == "all" || value == "any" || value == "*" {

		provider.Subnet = "any"

	} else if strings.Contains(value, "/") {

		address, prefix := toSubnet(value)

		if address != "" && prefix != 0 {
			provider.Subnet = value
		}

	}

}

func (provider *Provider) Hash() string {

	var hash string

	if provider.Name != "any" {

		hash = provider.Name

	} else {

		if provider.Country != "any" {
			hash += "any-" + provider.Country
		} else {
			hash += "any-any"
		}

		if provider.Subnet != "any" {

			address, prefix := toSubnet(provider.Subnet)

			if types.IsIPv6(address) {

				ipv6 := types.ParseIPv6(address)

				if ipv6 != nil {
					bytes := ipv6.Bytes(prefix)
					hash += "-" + hex.EncodeToString(bytes)
				}

			} else if types.IsIPv4(address) {

				ipv4 := types.ParseIPv4(address)

				if ipv4 != nil {
					bytes := ipv4.Bytes(prefix)
					hash += "-" + hex.EncodeToString(bytes)
				}

			} else {

				hash += "-any"

			}

		} else {
			hash += "-any"
		}

	}

	return hash

}
