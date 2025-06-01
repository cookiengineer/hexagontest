package matchers

import "strings"

type Weakness struct {
	Name   string `json:"name"`
	Impact string `json:"impact"`
	Scope  string `json:"scope"`
}

func NewWeakness() Weakness {

	var weakness Weakness

	weakness.Name = "any"
	weakness.Impact = "any"
	weakness.Scope = "any"

	return weakness

}

func ToWeakness(value string) Weakness {

	var weakness Weakness

	weakness.Name = "any"
	weakness.Impact = "any"
	weakness.Scope = "any"

	if strings.HasPrefix(value, "CWE-") {
		weakness.SetName(value)
	}

	return weakness

}

func (weakness *Weakness) IsIdentical(value Weakness) bool {

	var result bool = false

	if weakness.Name == value.Name &&
		weakness.Impact == value.Impact &&
		weakness.Scope == value.Scope {
		result = true
	}

	return result

}

func (weakness *Weakness) IsValid() bool {

	var result bool = false

	if weakness.Name != "" {
		result = true
	}

	return result

}

func (weakness *Weakness) Matches(name string, impacts []string, scopes []string) bool {

	var matches_name bool = false
	var matches_impact bool = false
	var matches_scope bool = false

	if weakness.Name == name {
		matches_name = true
	} else if weakness.Name == "any" {
		matches_name = true
	}

	if weakness.Impact == "any" {
		matches_impact = true
	} else {

		for i := 0; i < len(impacts); i++ {

			if weakness.Impact == impacts[i] {
				matches_impact = true
				break
			}

		}

	}

	if weakness.Scope == "any" {
		matches_scope = true
	} else {

		for s := 0; s < len(scopes); s++ {

			if weakness.Scope == scopes[s] {
				matches_scope = true
				break
			}

		}

	}

	return matches_name && matches_impact && matches_scope

}

func (weakness *Weakness) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		weakness.Name = "any"
	} else if strings.HasPrefix(value, "CWE-") {
		weakness.Name = value
	}

}

func (weakness *Weakness) SetImpact(value string) {

	if value == "all" || value == "any" || value == "*" {
		weakness.Impact = "any"
	} else if value == "Alter Execution Logic" {
		weakness.Impact = "Alter Execution Logic"
	} else if value == "Bypass Protection Mechanism" {
		weakness.Impact = "Bypass Protection Mechanism"
	} else if value == "DoS: Amplification" {
		weakness.Impact = "DoS: Amplification"
	} else if value == "DoS: Crash, Exit, or Restart" {
		weakness.Impact = "DoS: Crash, Exit, or Restart"
	} else if value == "DoS: Instability" {
		weakness.Impact = "DoS: Instability"
	} else if value == "DoS: Resource Consumption (CPU)" {
		weakness.Impact = "DoS: Resource Consumption (CPU)"
	} else if value == "DoS: Resource Consumption (Memory)" {
		weakness.Impact = "DoS: Resource Consumption (Memory)"
	} else if value == "DoS: Resource Consumption (Other)" {
		weakness.Impact = "DoS: Resource Consumption (Other)"
	} else if value == "Execute Unauthorized Code or Commands" {
		weakness.Impact = "Execute Unauthorized Code or Commands"
	} else if value == "Gain Privileges or Assume Identity" {
		weakness.Impact = "Gain Privileges or Assume Identity"
	} else if value == "Hide Activities" {
		weakness.Impact = "Hide Activities"
	} else if value == "Modify Application Data" {
		weakness.Impact = "Modify Application Data"
	} else if value == "Modify Files or Directories" {
		weakness.Impact = "Modify Files or Directories"
	} else if value == "Modify Memory" {
		weakness.Impact = "Modify Memory"
	} else if value == "Quality Degradation" {
		weakness.Impact = "Quality Degradation"
	} else if value == "Read Application Data" {
		weakness.Impact = "Read Application Data"
	} else if value == "Read Files or Directories" {
		weakness.Impact = "Read Files or Directories"
	} else if value == "Read Memory" {
		weakness.Impact = "Read Memory"
	} else if value == "Reduce Maintainability" {
		weakness.Impact = "Reduce Maintainability"
	} else if value == "Reduce Performance" {
		weakness.Impact = "Reduce Performance"
	} else if value == "Reduce Reliability" {
		weakness.Impact = "Reduce Reliability"
	} else if value == "Unexpected State" {
		weakness.Impact = "Unexpected State"
	} else if value == "Varies by Context" {
		weakness.Impact = "Varies by Context"
	}

}

func (weakness *Weakness) SetScope(value string) {

	if value == "all" || value == "any" || value == "*" {
		weakness.Scope = "any"
	} else if value == "Access Control" {
		weakness.Scope = "Access Control"
	} else if value == "Accountability" {
		weakness.Scope = "Accountability"
	} else if value == "Authentication" {
		weakness.Scope = "Authentication"
	} else if value == "Authorization" {
		weakness.Scope = "Authorization"
	} else if value == "Availability" {
		weakness.Scope = "Availability"
	} else if value == "Confidentiality" {
		weakness.Scope = "Confidentiality"
	} else if value == "Integrity" {
		weakness.Scope = "Integrity"
	} else if value == "Non-Repudiation" {
		weakness.Scope = "Non-Repudiation"
	}

}
