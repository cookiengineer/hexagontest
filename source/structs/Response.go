package structs

import "battlemap/matchers"
import "battlemap/types"
import "time"

type Response struct {
	Type            string                   `json:"type"`
	Datetime        types.Datetime           `json:"datetime"`
	Connections     []matchers.Connection    `json:"connections"`
	Distributions   []matchers.Distribution  `json:"distributions"`
	Networks        []matchers.Network       `json:"networks"`
	Packages        []matchers.Package       `json:"packages"`
	Programs        []matchers.Program       `json:"programs"`
	Users           []matchers.User          `json:"users"`
	Vulnerabilities []matchers.Vulnerability `json:"vulnerabilities"`
}

func NewResponse(typ string) Response {

	var response Response

	response.Type = "none"
	response.SetDatetime(time.Now().Format(time.RFC3339))
	response.Connections = make([]matchers.Connection, 0)
	response.Distributions = make([]matchers.Distribution, 0)
	response.Networks = make([]matchers.Network, 0)
	response.Packages = make([]matchers.Package, 0)
	response.Programs = make([]matchers.Program, 0)
	response.Users = make([]matchers.User, 0)
	response.Vulnerabilities = make([]matchers.Vulnerability, 0)

	response.SetType(typ)

	return response

}

func (response *Response) IsValid() bool {

	var result bool = false

	if response.Type == "Update" {

		if len(response.Packages) > 0 || len(response.Programs) > 0 {
			result = true
		}

	} else if response.Type == "Forbid" {

		if len(response.Connections) > 0 || len(response.Networks) > 0 || len(response.Programs) > 0 || len(response.Users) > 0 {
			result = true
		}

	} else if response.Type == "Permit" {

		if len(response.Connections) > 0 || len(response.Networks) > 0 || len(response.Programs) > 0 || len(response.Users) > 0 {
			result = true
		}

	} else if response.Type == "Recon" {

		if len(response.Connections) > 0 || len(response.Networks) > 0 {
			result = true
		}

	} else if response.Type == "Intel" {

		if len(response.Connections) > 0 || len(response.Networks) > 0 {
			result = true
		}

	} else if response.Type == "Conquer" {

		if len(response.Distributions) > 0 && len(response.Vulnerabilities) > 0 {

			if len(response.Connections) > 0 || len(response.Networks) > 0 {
				result = true
			}

		}

	} else if response.Type == "Persist" {

		if len(response.Distributions) > 0 && len(response.Vulnerabilities) > 0 {

			if len(response.Connections) > 0 || len(response.Networks) > 0 {
				result = true
			}

		}

	} else if response.Type == "Exfil" {

		if len(response.Connections) > 0 || len(response.Networks) > 0 {

			if len(response.Programs) > 0 || len(response.Users) > 0 {
				result = true
			}

		}

	} else if response.Type == "Destroy" {

		if len(response.Packages) > 0 || len(response.Programs) > 0 || len(response.Users) > 0 {
			result = true
		}

	}

	return result

}

func (response *Response) MatchesAntique(antique Antique) bool {

	var result bool = false

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Matches(antique.Name, antique.Version.String(), antique.Manager.String(), "any") {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesConnection(connection types.Connection) bool {

	var result bool = false

	for c := 0; c < len(response.Connections); c++ {

		if response.Connections[c].Matches(connection.Source.Host, connection.Source.Port, connection.Protocol.String(), connection.Type) {
			result = true
			break
		} else if response.Connections[c].Matches(connection.Target.Host, connection.Target.Port, connection.Protocol.String(), connection.Type) {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesDistribution(distribution Distribution) bool {

	var result bool = false

	for d := 0; d < len(response.Distributions); d++ {

		if response.Distributions[d].Matches(distribution.Name, distribution.Version, distribution.Vendor) {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesNetwork(network Network) bool {

	var result bool = false

	for n := 0; n < len(response.Networks); n++ {

		for s := 0; s < len(network.Subnets); s++ {

			if response.Networks[n].Matches(network.Subnets[s].Name, network.Subnets[s].String()) {
				result = true
				break
			}

		}

		if result == true {
			break
		}

	}

	return result

}

func (response *Response) MatchesPackage(pkg Package) bool {

	var result bool = false

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Matches(pkg.Name, pkg.Version.String(), pkg.Manager.String(), pkg.Vendor) {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesProgram(program Program) bool {

	var result bool = false

	for p := 0; p < len(response.Programs); p++ {

		if response.Programs[p].Matches(program.Name, program.Command) {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesUpdate(update Update) bool {

	var result bool = false

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].Matches(update.Name, update.Version.String(), update.Manager.String(), "any") {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesUser(user types.User) bool {

	var result bool = false

	for u := 0; u < len(response.Users); u++ {

		if response.Users[u].Matches(user.Name, user.Password, user.Type) {
			result = true
			break
		}

	}

	return result

}

func (response *Response) MatchesVulnerability(vulnerability Vulnerability) bool {

	var result bool = false

	for v := 0; v < len(response.Vulnerabilities); v++ {

		if response.Vulnerabilities[v].Matches(vulnerability.Name) {
			result = true
			break
		}

	}

	return result

}

func (response *Response) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		response.Datetime = datetime
	}

}

func (response *Response) SetType(value string) {

	if value == "Update" {
		response.Type = value
	} else if value == "Forbid" {
		response.Type = value
	} else if value == "Permit" {
		response.Type = value
	} else if value == "Recon" {
		response.Type = value
	} else if value == "Intel" {
		response.Type = value
	} else if value == "Conquer" {
		response.Type = value
	} else if value == "Persist" {
		response.Type = value
	} else if value == "Exfil" {
		response.Type = value
	} else if value == "Destroy" {
		response.Type = value
	}

}

func (response *Response) AddConnection(value matchers.Connection) {

	if value.IsValid() {

		var found bool = false

		for c := 0; c < len(response.Connections); c++ {

			if response.Connections[c].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Connections = append(response.Connections, value)
		}

	}

}

func (response *Response) RemoveConnection(value matchers.Connection) {

	var index int = -1

	for c := 0; c < len(response.Connections); c++ {

		if response.Connections[c].IsIdentical(value) {
			index = c
			break
		}

	}

	if index != -1 {
		response.Connections = append(response.Connections[:index], response.Connections[index+1:]...)
	}

}

func (response *Response) SetConnections(value []matchers.Connection) {

	var filtered []matchers.Connection

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Connections = filtered

}

func (response *Response) AddDistribution(value matchers.Distribution) {

	if value.IsValid() {

		var found bool = false

		for d := 0; d < len(response.Distributions); d++ {

			if response.Distributions[d].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Distributions = append(response.Distributions, value)
		}

	}

}

func (response *Response) RemoveDistribution(value matchers.Distribution) {

	var index int = -1

	for d := 0; d < len(response.Distributions); d++ {

		if response.Distributions[d].IsIdentical(value) {
			index = d
			break
		}

	}

	if index != -1 {
		response.Distributions = append(response.Distributions[:index], response.Distributions[index+1:]...)
	}

}

func (response *Response) SetDistributions(value []matchers.Distribution) {

	var filtered []matchers.Distribution

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Distributions = filtered

}

func (response *Response) AddNetwork(value matchers.Network) {

	if value.IsValid() {

		var found bool = false

		for n := 0; n < len(response.Networks); n++ {

			if response.Networks[n].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Networks = append(response.Networks, value)
		}

	}

}

func (response *Response) RemoveNetwork(value matchers.Network) {

	var index int = -1

	for n := 0; n < len(response.Networks); n++ {

		if response.Networks[n].IsIdentical(value) {
			index = n
			break
		}

	}

	if index != -1 {
		response.Networks = append(response.Networks[:index], response.Networks[index+1:]...)
	}

}

func (response *Response) SetNetworks(value []matchers.Network) {

	var filtered []matchers.Network

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Networks = filtered

}

func (response *Response) AddPackage(value matchers.Package) {

	if value.IsValid() {

		var found bool = false

		for p := 0; p < len(response.Packages); p++ {

			if response.Packages[p].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Packages = append(response.Packages, value)
		}

	}

}

func (response *Response) RemovePackage(value matchers.Package) {

	var index int = -1

	for p := 0; p < len(response.Packages); p++ {

		if response.Packages[p].IsIdentical(value) {
			index = p
			break
		}

	}

	if index != -1 {
		response.Packages = append(response.Packages[:index], response.Packages[index+1:]...)
	}

}

func (response *Response) SetPackages(value []matchers.Package) {

	var filtered []matchers.Package

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Packages = filtered

}

func (response *Response) AddProgram(value matchers.Program) {

	if value.IsValid() {

		var found bool = false

		for p := 0; p < len(response.Programs); p++ {

			if response.Programs[p].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Programs = append(response.Programs, value)
		}

	}

}

func (response *Response) RemoveProgram(value matchers.Program) {

	var index int = -1

	for p := 0; p < len(response.Programs); p++ {

		if response.Programs[p].IsIdentical(value) {
			index = p
			break
		}

	}

	if index != -1 {
		response.Programs = append(response.Programs[:index], response.Programs[index+1:]...)
	}

}

func (response *Response) SetPrograms(value []matchers.Program) {

	var filtered []matchers.Program

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Programs = filtered

}

func (response *Response) AddUser(value matchers.User) {

	if value.IsValid() {

		var found bool = false

		for u := 0; u < len(response.Users); u++ {

			if response.Users[u].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Users = append(response.Users, value)
		}

	}

}

func (response *Response) RemoveUser(value matchers.User) {

	var index int = -1

	for u := 0; u < len(response.Users); u++ {

		if response.Users[u].IsIdentical(value) {
			index = u
			break
		}

	}

	if index != -1 {
		response.Users = append(response.Users[:index], response.Users[index+1:]...)
	}

}

func (response *Response) SetUsers(value []matchers.User) {

	var filtered []matchers.User

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Users = filtered

}

func (response *Response) AddVulnerability(value matchers.Vulnerability) {

	if value.IsValid() {

		var found bool = false

		for v := 0; v < len(response.Vulnerabilities); v++ {

			if response.Vulnerabilities[v].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			response.Vulnerabilities = append(response.Vulnerabilities, value)
		}

	}

}

func (response *Response) RemoveVulnerability(value matchers.Vulnerability) {

	var index int = -1

	for v := 0; v < len(response.Vulnerabilities); v++ {

		if response.Vulnerabilities[v].IsIdentical(value) {
			index = v
			break
		}

	}

	if index != -1 {
		response.Vulnerabilities = append(response.Vulnerabilities[:index], response.Vulnerabilities[index+1:]...)
	}

}

func (response *Response) SetVulnerabilities(value []matchers.Vulnerability) {

	var filtered []matchers.Vulnerability

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	response.Vulnerabilities = filtered

}
