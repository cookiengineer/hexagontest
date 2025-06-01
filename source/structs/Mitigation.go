package structs

import utils_strings "battlemap/utils/strings"
import "strings"

type Mitigation struct {
	Name     string   `json:"name"`
	Incident Incident `json:"incident"`
	Response Response `json:"response"`
}

func NewMitigation(incident Incident, response Response) Mitigation {

	var mitigation Mitigation

	mitigation.SetIncident(incident)
	mitigation.SetResponse(response)
	mitigation.Name = mitigation.Hash()

	return mitigation

}

func (mitigation *Mitigation) IsValid() bool {

	var result bool = false

	if mitigation.Incident.IsValid() && mitigation.Response.IsValid() {
		result = true
	}

	return result

}

func (mitigation *Mitigation) SetIncident(value Incident) {

	if value.IsValid() {
		mitigation.Incident = value
	}

}

func (mitigation *Mitigation) SetResponse(value Response) {

	if value.IsValid() {
		mitigation.Response = value
		mitigation.Name = mitigation.Hash()
	}

}

func (mitigation *Mitigation) Hash() string {

	var hash string

	response := mitigation.Response

	if response.Type == "Update" {

		if len(response.Packages) > 0 {

			hashes := make([]string, 0)

			for p := 0; p < len(response.Packages); p++ {
				hashes = append(hashes, response.Packages[p].Hash())
			}

			hash = response.Type + "-Packages-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for p := 0; p < len(response.Programs); p++ {
				hashes = append(hashes, response.Programs[p].Hash())
			}

			hash = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Forbid" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for p := 0; p < len(response.Programs); p++ {
				hashes = append(hashes, response.Programs[p].Hash())
			}

			hash = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Users) > 0 {

			hashes := make([]string, 0)

			for u := 0; u < len(response.Users); u++ {
				hashes = append(hashes, response.Users[u].Hash())
			}

			hash = response.Type + "-Users-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Permit" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for p := 0; p < len(response.Programs); p++ {
				hashes = append(hashes, response.Programs[p].Hash())
			}

			hash = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Users) > 0 {

			hashes := make([]string, 0)

			for u := 0; u < len(response.Users); u++ {
				hashes = append(hashes, response.Users[u].Hash())
			}

			hash = response.Type + "-Users-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Recon" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Intel" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Conquer" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Persist" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Exfil" {

		if len(response.Connections) > 0 {

			hashes := make([]string, 0)

			for c := 0; c < len(response.Connections); c++ {
				hashes = append(hashes, response.Connections[c].Hash())
			}

			hash = response.Type + "-Connections-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Networks) > 0 {

			hashes := make([]string, 0)

			for n := 0; n < len(response.Networks); n++ {
				hashes = append(hashes, response.Networks[n].Hash())
			}

			hash = response.Type + "-Networks-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	} else if response.Type == "Destroy" {

		if len(response.Packages) > 0 {

			hashes := make([]string, 0)

			for p := 0; p < len(response.Packages); p++ {
				hashes = append(hashes, response.Packages[p].Hash())
			}

			hash = response.Type + "-Packages-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Programs) > 0 {

			hashes := make([]string, 0)

			for p := 0; p < len(response.Programs); p++ {
				hashes = append(hashes, response.Programs[p].Hash())
			}

			hash = response.Type + "-Programs-" + strings.Join(utils_strings.Unique(hashes), ",")

		} else if len(response.Users) > 0 {

			hashes := make([]string, 0)

			for u := 0; u < len(response.Users); u++ {
				hashes = append(hashes, response.Users[u].Hash())
			}

			hash = response.Type + "-Users-" + strings.Join(utils_strings.Unique(hashes), ",")

		}

	}

	return hash

}
