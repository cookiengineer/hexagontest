package structs

import "battlemap/matchers"
import "battlemap/types"
import "time"

type Incident struct {
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

func NewIncident(typ string) Incident {

	var incident Incident

	incident.Type = "none"
	incident.SetDatetime(time.Now().Format(time.RFC3339))
	incident.Connections = make([]matchers.Connection, 0)
	incident.Distributions = make([]matchers.Distribution, 0)
	incident.Networks = make([]matchers.Network, 0)
	incident.Packages = make([]matchers.Package, 0)
	incident.Programs = make([]matchers.Program, 0)
	incident.Users = make([]matchers.User, 0)
	incident.Vulnerabilities = make([]matchers.Vulnerability, 0)

	incident.SetType(typ)

	return incident

}

func (incident *Incident) IsValid() bool {

	var result bool = false

	if incident.Type == "Recon" {

		if len(incident.Connections) > 0 || len(incident.Networks) > 0 {

			if len(incident.Programs) > 0 || len(incident.Users) > 0 {
				result = true
			}

		}

	} else if incident.Type == "Intel" {

		if len(incident.Vulnerabilities) > 0 {

			if len(incident.Packages) > 0 || len(incident.Programs) > 0 {
				result = true
			}

		}

	} else if incident.Type == "Conquer" {

		if len(incident.Distributions) > 0 && len(incident.Vulnerabilities) > 0 {

			if len(incident.Packages) > 0 || len(incident.Programs) > 0 {
				result = true
			}

		}

	} else if incident.Type == "Persist" {

		if len(incident.Distributions) > 0 && len(incident.Vulnerabilities) > 0 {

			if len(incident.Packages) > 0 || len(incident.Programs) > 0 || len(incident.Users) > 0 {
				result = true
			}

		}

	} else if incident.Type == "Exfil" {

		if len(incident.Connections) > 0 || len(incident.Networks) > 0 {

			if len(incident.Programs) > 0 || len(incident.Users) > 0 {
				result = true
			}

		}

	} else if incident.Type == "Destroy" {

		if len(incident.Packages) > 0 || len(incident.Programs) > 0 || len(incident.Users) > 0 {
			result = true
		}

	}

	return result

}

func (incident *Incident) MatchesAntique(antique Antique) bool {

	var result bool = false

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(antique.Name, antique.Version.String(), antique.Manager.String(), "any") {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesConnection(connection types.Connection) bool {

	var result bool = false

	for c := 0; c < len(incident.Connections); c++ {

		if incident.Connections[c].Matches(connection.Source.Host, connection.Source.Port, connection.Protocol.String(), connection.Type) {
			result = true
			break
		} else if incident.Connections[c].Matches(connection.Target.Host, connection.Target.Port, connection.Protocol.String(), connection.Type) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesDistribution(distribution Distribution) bool {

	var result bool = false

	for d := 0; d < len(incident.Distributions); d++ {

		if incident.Distributions[d].Matches(distribution.Name, distribution.Version, distribution.Vendor) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesNetwork(network Network) bool {

	var result bool = false

	for n := 0; n < len(incident.Networks); n++ {

		for s := 0; s < len(network.Subnets); s++ {

			if incident.Networks[n].Matches(network.Subnets[s].Name, network.Subnets[s].String()) {
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

func (incident *Incident) MatchesPackage(pkg Package) bool {

	var result bool = false

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(pkg.Name, pkg.Version.String(), pkg.Manager.String(), pkg.Vendor) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesProgram(program Program) bool {

	var result bool = false

	for p := 0; p < len(incident.Programs); p++ {

		if incident.Programs[p].Matches(program.Name, program.Command) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesUpdate(update Update) bool {

	var result bool = false

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].Matches(update.Name, update.Version.String(), update.Manager.String(), "any") {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesUser(user types.User) bool {

	var result bool = false

	for u := 0; u < len(incident.Users); u++ {

		if incident.Users[u].Matches(user.Name, user.Password, user.Type) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) MatchesVulnerability(vulnerability Vulnerability) bool {

	var result bool = false

	for v := 0; v < len(incident.Vulnerabilities); v++ {

		if incident.Vulnerabilities[v].Matches(vulnerability.Name) {
			result = true
			break
		}

	}

	return result

}

func (incident *Incident) SetDatetime(value string) {

	datetime := types.ToDatetime(value)

	if datetime.IsValid() {
		incident.Datetime = datetime
	}

}

func (incident *Incident) SetType(value string) {

	if value == "Recon" {
		incident.Type = value
	} else if value == "Intel" {
		incident.Type = value
	} else if value == "Conquer" {
		incident.Type = value
	} else if value == "Persist" {
		incident.Type = value
	} else if value == "Exfil" {
		incident.Type = value
	} else if value == "Destroy" {
		incident.Type = value
	}

}

func (incident *Incident) AddConnection(value matchers.Connection) {

	if value.IsValid() {

		var found bool = false

		for c := 0; c < len(incident.Connections); c++ {

			if incident.Connections[c].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Connections = append(incident.Connections, value)
		}

	}

}

func (incident *Incident) RemoveConnection(value matchers.Connection) {

	var index int = -1

	for c := 0; c < len(incident.Connections); c++ {

		if incident.Connections[c].IsIdentical(value) {
			index = c
			break
		}

	}

	if index != -1 {
		incident.Connections = append(incident.Connections[:index], incident.Connections[index+1:]...)
	}

}

func (incident *Incident) SetConnections(value []matchers.Connection) {

	var filtered []matchers.Connection

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Connections = filtered

}

func (incident *Incident) AddDistribution(value matchers.Distribution) {

	if value.IsValid() {

		var found bool = false

		for d := 0; d < len(incident.Distributions); d++ {

			if incident.Distributions[d].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Distributions = append(incident.Distributions, value)
		}

	}

}

func (incident *Incident) RemoveDistribution(value matchers.Distribution) {

	var index int = -1

	for d := 0; d < len(incident.Distributions); d++ {

		if incident.Distributions[d].IsIdentical(value) {
			index = d
			break
		}

	}

	if index != -1 {
		incident.Distributions = append(incident.Distributions[:index], incident.Distributions[index+1:]...)
	}

}

func (incident *Incident) SetDistributions(value []matchers.Distribution) {

	var filtered []matchers.Distribution

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Distributions = filtered

}

func (incident *Incident) AddNetwork(value matchers.Network) {

	if value.IsValid() {

		var found bool = false

		for n := 0; n < len(incident.Networks); n++ {

			if incident.Networks[n].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Networks = append(incident.Networks, value)
		}

	}

}

func (incident *Incident) RemoveNetwork(value matchers.Network) {

	var index int = -1

	for n := 0; n < len(incident.Networks); n++ {

		if incident.Networks[n].IsIdentical(value) {
			index = n
			break
		}

	}

	if index != -1 {
		incident.Networks = append(incident.Networks[:index], incident.Networks[index+1:]...)
	}

}

func (incident *Incident) SetNetworks(value []matchers.Network) {

	var filtered []matchers.Network

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Networks = filtered

}

func (incident *Incident) AddPackage(value matchers.Package) {

	if value.IsValid() {

		var found bool = false

		for p := 0; p < len(incident.Packages); p++ {

			if incident.Packages[p].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Packages = append(incident.Packages, value)
		}

	}

}

func (incident *Incident) RemovePackage(value matchers.Package) {

	var index int = -1

	for p := 0; p < len(incident.Packages); p++ {

		if incident.Packages[p].IsIdentical(value) {
			index = p
			break
		}

	}

	if index != -1 {
		incident.Packages = append(incident.Packages[:index], incident.Packages[index+1:]...)
	}

}

func (incident *Incident) SetPackages(value []matchers.Package) {

	var filtered []matchers.Package

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Packages = filtered

}

func (incident *Incident) AddProgram(value matchers.Program) {

	if value.IsValid() {

		var found bool = false

		for p := 0; p < len(incident.Programs); p++ {

			if incident.Programs[p].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Programs = append(incident.Programs, value)
		}

	}

}

func (incident *Incident) RemoveProgram(value matchers.Program) {

	var index int = -1

	for p := 0; p < len(incident.Programs); p++ {

		if incident.Programs[p].IsIdentical(value) {
			index = p
			break
		}

	}

	if index != -1 {
		incident.Programs = append(incident.Programs[:index], incident.Programs[index+1:]...)
	}

}

func (incident *Incident) SetPrograms(value []matchers.Program) {

	var filtered []matchers.Program

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Programs = filtered

}

func (incident *Incident) AddUser(value matchers.User) {

	if value.IsValid() {

		var found bool = false

		for u := 0; u < len(incident.Users); u++ {

			if incident.Users[u].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Users = append(incident.Users, value)
		}

	}

}

func (incident *Incident) RemoveUser(value matchers.User) {

	var index int = -1

	for u := 0; u < len(incident.Users); u++ {

		if incident.Users[u].IsIdentical(value) {
			index = u
			break
		}

	}

	if index != -1 {
		incident.Users = append(incident.Users[:index], incident.Users[index+1:]...)
	}

}

func (incident *Incident) SetUsers(value []matchers.User) {

	var filtered []matchers.User

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Users = filtered

}

func (incident *Incident) AddVulnerability(value matchers.Vulnerability) {

	if value.IsValid() {

		var found bool = false

		for v := 0; v < len(incident.Vulnerabilities); v++ {

			if incident.Vulnerabilities[v].IsIdentical(value) {
				found = true
				break
			}

		}

		if found == false {
			incident.Vulnerabilities = append(incident.Vulnerabilities, value)
		}

	}

}

func (incident *Incident) RemoveVulnerability(value matchers.Vulnerability) {

	var index int = -1

	for v := 0; v < len(incident.Vulnerabilities); v++ {

		if incident.Vulnerabilities[v].IsIdentical(value) {
			index = v
			break
		}

	}

	if index != -1 {
		incident.Vulnerabilities = append(incident.Vulnerabilities[:index], incident.Vulnerabilities[index+1:]...)
	}

}

func (incident *Incident) SetVulnerabilities(value []matchers.Vulnerability) {

	var filtered []matchers.Vulnerability

	for v := 0; v < len(value); v++ {

		if value[v].IsValid() {
			filtered = append(filtered, value[v])
		}

	}

	incident.Vulnerabilities = filtered

}
