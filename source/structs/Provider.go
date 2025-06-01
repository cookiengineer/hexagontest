package structs

import "strings"

type Provider struct {
	Name    string   `json:"name"`
	Address []string `json:"address"`
	Country string   `json:"country"`
	Subnets []Subnet `json:"subnets"`
}

func NewProvider(name string) Provider {

	var provider Provider

	provider.SetName(name)

	provider.Address = make([]string, 0)
	provider.Subnets = make([]Subnet, 0)

	return provider

}

func (provider *Provider) IsValid() bool {

	var result bool = false

	if provider.Name != "" && provider.Country != "" && len(provider.Subnets) > 0 {
		result = true
	}

	return result

}

func (provider *Provider) SetAddress(value []string) {

	var filtered []string

	for v := 0; v < len(value); v++ {

		line := strings.TrimSpace(value[v])

		if line != "" {
			filtered = append(filtered, line)
		}

	}

	provider.Address = filtered

}

func (provider *Provider) SetCountry(value string) {

	if len(value) == 2 {
		provider.Country = value
	}

}

func (provider *Provider) SetName(value string) {
	provider.Name = strings.TrimSpace(value)
}

func (provider *Provider) AddSubnet(value Subnet) {

	if value.IsValid() {

		var found bool = false

		for s := 0; s < len(provider.Subnets); s++ {

			if provider.Subnets[s].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			provider.Subnets = append(provider.Subnets, value)
		}

	}

}

func (provider *Provider) RemoveSubnet(value Subnet) {

	var index int = -1

	for s := 0; s < len(provider.Subnets); s++ {

		if provider.Subnets[s].IsIdentical(value) {
			index = s
			break
		}

	}

	if index != -1 {
		provider.Subnets = append(provider.Subnets[:index], provider.Subnets[index+1:]...)
	}

}

func (provider *Provider) SetSubnets(value []Subnet) {

	var filtered []Subnet

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	provider.Subnets = filtered

}
