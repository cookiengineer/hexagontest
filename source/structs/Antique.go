package structs

import "battlemap/types"
import "net/url"
import "strings"

type Antique struct {
	Name         string             `json:"name"`
	Version      types.Version      `json:"version"`
	Architecture types.Architecture `json:"architecture"`
	Manager      types.Manager      `json:"manager"`
	Service      string             `json:"service"`
	URL          string             `json:"url"`
}

func NewAntique(manager string, service string) Antique {

	var antique Antique

	antique.SetArchitecture("any")
	antique.SetManager(manager)
	antique.SetService(service)

	return antique

}

func (antique *Antique) IsValid() bool {

	if antique.Name != "" {

		var result bool = true

		if antique.Version.IsValid() == false {
			result = false
		}

		if antique.Architecture.IsValid() == false {
			result = false
		}

		if antique.Manager.IsValid() == false {
			result = false
		}

		if antique.Service == "" {
			result = false
		}

		return result

	}

	return false

}

func (antique *Antique) SetArchitecture(value string) {

	architecture := types.ParseArchitecture(value)

	if architecture != nil {
		antique.Architecture = *architecture
	}

}

func (antique *Antique) SetManager(value string) {

	manager := types.ParseManager(value)

	if manager != nil {
		antique.Manager = *manager
	}

}

func (antique *Antique) SetName(value string) {
	antique.Name = strings.TrimSpace(value)
}

func (antique *Antique) SetService(value string) {
	antique.Service = strings.TrimSpace(value)
}

func (antique *Antique) SetURL(value string) {

	tmp, err := url.ParseRequestURI(value)

	if err == nil {

		if tmp.Scheme == "https" || tmp.Scheme == "http" {
			antique.URL = tmp.String()
		}

	}

}

func (antique *Antique) SetVersion(value string) {

	version := types.ToVersion(value)

	if version.IsValid() {
		antique.Version = version
	}

}
