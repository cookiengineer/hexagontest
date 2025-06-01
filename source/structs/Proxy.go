package structs

import "battlemap/types"
import "math/rand"

type Proxy struct {
	Domain    string         `json:"domain"`
	Addresses []string       `json:"addresses"`
	Port      uint16         `json:"port"`
	Protocol  types.Protocol `json:"protocol"`
}

func NewProxy() Proxy {

	var proxy Proxy

	proxy.Domain = ""
	proxy.Port = 0
	proxy.Protocol = types.ProtocolANY
	proxy.Addresses = make([]string, 0)

	return proxy

}

func (proxy *Proxy) IsValid() bool {

	var result bool = false

	if proxy.Domain != "" && proxy.Port != 0 && proxy.Protocol != types.ProtocolANY {

		if len(proxy.Addresses) > 0 {
			result = true
		}

	}

	return result

}

func (proxy *Proxy) AddAddress(value string) {

	valid := false

	if types.IsIPv6(value) {

		ipv6 := types.ParseIPv6(value)

		if ipv6 != nil {
			value = ipv6.String()
			valid = true
		}

	} else if types.IsIPv4(value) {

		ipv4 := types.ParseIPv4(value)

		if ipv4 != nil {
			value = ipv4.String()
			valid = true
		}

	}

	if valid == true {

		found := false

		for a := 0; a < len(proxy.Addresses); a++ {

			if proxy.Addresses[a] == value {
				found = true
				break
			}

		}

		if found == false {
			proxy.Addresses = append(proxy.Addresses, value)
		}

	}

}

func (proxy *Proxy) RemoveAddress(value string) {

	var index int = -1

	for a := 0; a < len(proxy.Addresses); a++ {

		if proxy.Addresses[a] == value {
			index = a
			break
		}

	}

	if index != -1 {
		proxy.Addresses = append(proxy.Addresses[:index], proxy.Addresses[index+1:]...)
	}

}

func (proxy *Proxy) SetAddresses(value []string) {

	filtered := make([]string, 0)

	for v := 0; v < len(value); v++ {

		if types.IsIPv6(value[v]) {

			ipv6 := types.ParseIPv6(value[v])

			if ipv6 != nil {
				filtered = append(filtered, ipv6.String())
			}

		} else if types.IsIPv4(value[v]) {

			ipv4 := types.ParseIPv4(value[v])

			if ipv4 != nil {
				filtered = append(filtered, ipv4.String())
			}

		}

	}

	proxy.Addresses = filtered

}

func (proxy *Proxy) SetDomain(value string) {

	if types.IsDomain(value) {

		domain := types.ParseDomain(value)

		if domain != nil {
			proxy.Domain = domain.String()
		}

	}

}

func (proxy *Proxy) SetPort(value uint16) {

	if value > 0 && value < 65535 {
		proxy.Port = value
	}

}

func (proxy *Proxy) SetProtocol(value types.Protocol) {
	proxy.Protocol = value
}

func (proxy *Proxy) RandomizeAddress() string {

	var result string

	if types.SupportsIPv4() {

		filtered := make([]string, 0)

		for a := 0; a < len(proxy.Addresses); a++ {

			address := proxy.Addresses[a]

			if types.IsIPv4(address) {
				filtered = append(filtered, address)
			}

		}

		if len(filtered) > 0 {
			result = filtered[rand.Intn(len(filtered))]
		}

	} else if types.SupportsIPv6() {

		filtered := make([]string, 0)

		for a := 0; a < len(proxy.Addresses); a++ {

			address := proxy.Addresses[a]

			if types.IsIPv6(address) {
				filtered = append(filtered, address)
			}

		}

		if len(filtered) > 0 {
			result = filtered[rand.Intn(len(filtered))]
		}

	}

	return result

}

