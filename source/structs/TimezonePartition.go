package structs

import "sort"
import "strings"

type TimezonePartition struct {
	Offset     string   `json:"offset"`
	Population uint     `json:"population"`
	Weight     float64  `json:"weight"`
	Names      []string `json:"names"`
}

func NewTimezonePartition(offset string, weight float64) TimezonePartition {

	var tzpartition TimezonePartition

	tzpartition.SetOffset(offset)
	tzpartition.SetWeight(weight)

	return tzpartition

}

func (tzpartition *TimezonePartition) IsValid() bool {

	var result bool = false

	if tzpartition.Offset != "" && tzpartition.Weight != 0.0 {
		result = true
	}

	return result

}

func (tzpartition *TimezonePartition) AddName(value string) {

	value = strings.TrimSpace(value)

	if value != "" {

		var found bool = false

		for n := 0; n < len(tzpartition.Names); n++ {

			if tzpartition.Names[n] == value {
				found = true
				break
			}

		}

		if found == false {
			tzpartition.Names = append(tzpartition.Names, value)
			sort.Strings(tzpartition.Names)
		}

	}

}

func (tzpartition *TimezonePartition) RemoveName(value string) {

	var index int = -1

	for n := 0; n < len(tzpartition.Names); n++ {

		if tzpartition.Names[n] == value {
			index = n
			break
		}

	}

	if index != -1 {
		tzpartition.Names = append(tzpartition.Names[:index], tzpartition.Names[index+1:]...)
	}

}

func (tzpartition *TimezonePartition) SetNames(values []string) {
	tzpartition.Names = values
}

func (tzpartition *TimezonePartition) SetOffset(value string) {

	if strings.HasPrefix(value, "+") || strings.HasPrefix(value, "-") {

		if len(value) == 6 && string(value[3]) == ":" {
			tzpartition.Offset = value
		}

	}

}

func (tzpartition *TimezonePartition) SetPopulation(value uint) {
	tzpartition.Population = value
}

func (tzpartition *TimezonePartition) SetWeight(value float64) {
	tzpartition.Weight = value
}
