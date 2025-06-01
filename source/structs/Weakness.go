package structs

import "strings"

func isWeaknessImpact(value string) bool {

	var result bool = false

	impacts := []string{
		"Alter Execution Logic",
		"Bypass Protection Mechanism",
		"DoS: Amplification",
		"DoS: Crash, Exit, or Restart",
		"DoS: Instability",
		"DoS: Resource Consumption (CPU)",
		"DoS: Resource Consumption (Memory)",
		"DoS: Resource Consumption (Other)",
		"Execute Unauthorized Code or Commands",
		"Gain Privileges or Assume Identity",
		"Hide Activities",
		"Modify Application Data",
		"Modify Files or Directories",
		"Modify Memory",
		"Quality Degradation",
		"Read Application Data",
		"Read Files or Directories",
		"Read Memory",
		"Reduce Maintainability",
		"Reduce Performance",
		"Reduce Reliability",
		"Unexpected State",
		"Varies by Context",
	}

	for i := 0; i < len(impacts); i++ {

		if impacts[i] == value {
			result = true
			break
		}

	}

	return result

}

func isWeaknessScope(value string) bool {

	var result bool = false

	scopes := []string{
		"Access Control",
		"Accountability",
		"Authentication",
		"Authorization",
		"Availability",
		"Confidentiality",
		"Integrity",
		"Non-Repudiation",
	}

	for s := 0; s < len(scopes); s++ {

		if scopes[s] == value {
			result = true
			break
		}

	}

	return result

}

type Weakness struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Impacts     []string `json:"impacts"`
	References  []string `json:"references"`
	Scopes      []string `json:"scopes"`
}

func NewWeakness() Weakness {

	var weakness Weakness

	weakness.Name = ""
	weakness.Description = ""
	weakness.Impacts = make([]string, 0)
	weakness.References = make([]string, 0)
	weakness.Scopes = make([]string, 0)

	return weakness

}

func ToWeakness(value string) Weakness {

	var weakness Weakness

	weakness.Name = ""
	weakness.Description = ""
	weakness.Impacts = make([]string, 0)
	weakness.References = make([]string, 0)
	weakness.Scopes = make([]string, 0)

	if strings.HasPrefix(value, "CWE-") {
		weakness.SetName(value)
	}

	return weakness

}

func (weakness *Weakness) IsValid() bool {

	var result bool = false

	if weakness.Name != "" {

		if len(weakness.Impacts) > 0 || len(weakness.Scopes) > 0 {
			result = true
		}

	}

	return result

}

func (weakness *Weakness) SetName(value string) {

	if strings.HasPrefix(value, "CWE-") {
		weakness.Name = strings.TrimSpace(value)
	}

}

func (weakness *Weakness) SetDescription(value string) {
	weakness.Description = strings.TrimSpace(value)
}

func (weakness *Weakness) AddImpact(value string) {

	if isWeaknessImpact(value) {

		var found bool = false

		for i := 0; i < len(weakness.Impacts); i++ {

			if weakness.Impacts[i] == value {
				found = true
				break
			}

		}

		if found == false {
			weakness.Impacts = append(weakness.Impacts, value)
		}

	}

}

func (weakness *Weakness) RemoveImpact(value string) {

	var index int = -1

	for i := 0; i < len(weakness.Impacts); i++ {

		if weakness.Impacts[i] == value {
			index = i
			break
		}

	}

	if index != -1 {
		weakness.Impacts = append(weakness.Impacts[:index], weakness.Impacts[index+1:]...)
	}

}

func (weakness *Weakness) SetImpacts(values []string) {

	var filtered []string

	for v := 0; v < len(values); v++ {

		var value = values[v]

		if isWeaknessImpact(value) {
			filtered = append(filtered, value)
		}

	}

	weakness.Impacts = filtered

}

func (weakness *Weakness) AddReference(value string) {

	if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {

		var found bool = false

		for r := 0; r < len(weakness.References); r++ {

			if weakness.References[r] == value {
				found = true
				break
			}

		}

		if found == false {
			weakness.References = append(weakness.References, value)
		}

	}

}

func (weakness *Weakness) RemoveReference(value string) {

	var index int = -1

	for r := 0; r < len(weakness.References); r++ {

		if weakness.References[r] == value {
			index = r
			break
		}

	}

	if index != -1 {
		weakness.References = append(weakness.References[:index], weakness.References[index+1:]...)
	}

}

func (weakness *Weakness) SetReferences(values []string) {

	var filtered []string

	for v := 0; v < len(values); v++ {

		var value = values[v]

		if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") {
			filtered = append(filtered, value)
		}

	}

	weakness.References = filtered

}

func (weakness *Weakness) AddScope(value string) {

	if isWeaknessScope(value) {

		var found bool = false

		for s := 0; s < len(weakness.Scopes); s++ {

			if weakness.Scopes[s] == value {
				found = true
				break
			}

		}

		if found == false {
			weakness.Scopes = append(weakness.Scopes, value)
		}

	}

}

func (weakness *Weakness) RemoveScope(value string) {

	var index int = -1

	for s := 0; s < len(weakness.Scopes); s++ {

		if weakness.Scopes[s] == value {
			index = s
			break
		}

	}

	if index != -1 {
		weakness.Scopes = append(weakness.Scopes[:index], weakness.Scopes[index+1:]...)
	}

}

func (weakness *Weakness) SetScopes(values []string) {

	var filtered []string

	for v := 0; v < len(values); v++ {

		var value = values[v]

		if isWeaknessScope(value) {
			filtered = append(filtered, value)
		}

	}

	weakness.Scopes = filtered

}
