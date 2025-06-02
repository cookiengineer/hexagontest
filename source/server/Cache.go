package server

import "battlemap/structs"
import "encoding/json"
import "fmt"
import "io/fs"
import "os"
import "strconv"
import "strings"

type Cache struct {
	Folder          string                                    `json:"folder"`
	Systems         map[string]*structs.System                `json:"systems"`
	Vulnerabilities map[string]*structs.Vulnerability         `json:"vulnerabilities"`
	lookup          map[string]map[string]map[string][]string `json:"lookup"`
}

func NewCache(folder string) *Cache {

	var cache Cache

	cache.Folder = strings.TrimSpace(folder)
	cache.Systems = make(map[string]*structs.System)
	cache.Vulnerabilities = make(map[string]*structs.Vulnerability)

	// Cache Lookup structure: distribution.Name > distribution.Version > package.Name > []structs.Vulnerability.Name
	cache.lookup = make(map[string]map[string]map[string][]string)

	fmt.Println("Cache: \"" + cache.Folder + "\"")

	return &cache

}

func (cache *Cache) Init() bool {

	var result bool

	count_systems := 0
	count_vulnerabilities := 0

	if cache.Folder != "" {

		fsys := os.DirFS(cache.Folder)

		files1, err1 := fs.ReadDir(fsys, "systems")

		if err1 == nil {

			for _, file := range files1 {

				name := file.Name()

				if strings.HasSuffix(name, ".json") {

					buffer, err11 := fs.ReadFile(fsys, "systems/" + file.Name())

					if err11 == nil {

						var system structs.System

						err12 := json.Unmarshal(buffer, &system)

						if err12 == nil {
							cache.SetSystem(system)
							count_systems++
						}

					}

				}

			}

		}

		files2, err2 := fs.ReadDir(fsys, "vulnerabilities")

		if err2 == nil {

			for _, file := range files2 {

				name := file.Name()

				if strings.HasSuffix(name, ".json") {

					buffer, err21 := fs.ReadFile(fsys, "vulnerabilities/" + file.Name())

					if err21 == nil {

						var vulnerability structs.Vulnerability

						err22 := json.Unmarshal(buffer, &vulnerability)

						if err22 == nil {
							cache.SetVulnerability(vulnerability)
							count_vulnerabilities++
						}

					}

				}

			}

		}

	}

	fmt.Println("Cache: " + strconv.Itoa(count_systems) + " Systems")
	fmt.Println("Cache: " + strconv.Itoa(count_vulnerabilities) + " Vulnerabilities")

	for distribution_name, _ := range cache.lookup {

		for distribution_version, _ := range cache.lookup[distribution_name] {
			fmt.Println("Cache: " + distribution_name + "-" + distribution_version)
		}

	}

	return result

}

func (cache *Cache) GetSystem(name string) *structs.System {

	var result *structs.System

	tmp, ok := cache.Systems[name]

	if ok == true {
		result = tmp
	}

	return result

}

func (cache *Cache) SetSystem(system structs.System) {

	if system.Name != "" {
		cache.Systems[system.Name] = &system
	}

}

func (cache *Cache) GetVulnerability(name string) *structs.Vulnerability {

	var result *structs.Vulnerability

	tmp, ok := cache.Vulnerabilities[name]

	if ok == true {
		result = tmp
	}

	return result

}

func (cache *Cache) QueryVulnerabilitiesByDistribution(query string) []string {

	distribution_name := ""
	distribution_version := ""

	if strings.Contains(query, "-") {
		distribution_name    = query[0:strings.Index(query, "-")]
		distribution_version = query[strings.Index(query, "-")+1:]
	} else {
		distribution_name = query
		distribution_version = "any"
	}

	fmt.Println("Cache Query: " + distribution_name + " " + distribution_version)

	result := make([]string, 0)

	if distribution_name != "" && distribution_name != "any" {

		_, ok1 := cache.lookup[distribution_name]

		if ok1 == true {

			found_map := make(map[string]bool)

			if distribution_version == "any" {

				_, ok2 := cache.lookup[distribution_name][distribution_version]

				if ok2 == true {

					for _, vulnerabilities := range cache.lookup[distribution_name][distribution_version] {

						for _, name := range vulnerabilities {
							found_map[name] = true
						}

					}

				} else {

					for distribution_version, _ := range cache.lookup[distribution_name] {

						for _, vulnerabilities := range cache.lookup[distribution_name][distribution_version] {

							for _, name := range vulnerabilities {
								found_map[name] = true
							}

						}

					}

				}

			} else if distribution_version != "" {

				_, ok2 := cache.lookup[distribution_name][distribution_version]

				if ok2 == true {

					for _, vulnerabilities := range cache.lookup[distribution_name][distribution_version] {

						for _, name := range vulnerabilities {
							found_map[name] = true
						}

					}

				}

			}

			if len(found_map) > 0 {

				for name, _ := range found_map {
					result = append(result, name)
				}

			}

		}

	}

	return result

}

func (cache *Cache) QueryVulnerabilitiesByDistributionAndPackage(query_distribution string, query_package string) []string {

	distribution_name := ""
	distribution_version := ""
	package_name := ""

	if strings.Contains(query_distribution, "-") {
		distribution_name    = query_distribution[0:strings.Index(query_distribution, "-")]
		distribution_version = query_distribution[strings.Index(query_distribution, "-")+1:]
	} else {
		distribution_name = query_distribution
		distribution_version = "any"
	}

	if query_package != "" {
		package_name = query_package
	} else {
		package_name = "any"
	}

	result := make([]string, 0)

	if distribution_name != "" && distribution_name != "any" {

		_, ok1 := cache.lookup[distribution_name]

		if ok1 == true {

			found_map := make(map[string]bool)

			if distribution_version == "any" {

				_, ok2 := cache.lookup[distribution_name]["any"]

				if ok2 == true {

					if package_name != "" && package_name != "any" {

						vulnerabilities, ok3 := cache.lookup[distribution_name]["any"][package_name]

						if ok3 == true {

							for _, name := range vulnerabilities {
								found_map[name] = true
							}

						}

					} else {
						result = cache.QueryVulnerabilitiesByDistribution(query_distribution)
					}

				} else {

					if package_name != "" && package_name != "any" {

						for distribution_version, _ := range cache.lookup[distribution_name] {

							vulnerabilities, ok3 := cache.lookup[distribution_name][distribution_version][package_name]

							if ok3 == true {

								for _, name := range vulnerabilities {
									found_map[name] = true
								}

							}

						}

					} else {
						result = cache.QueryVulnerabilitiesByDistribution(query_distribution)
					}

				}

			} else if distribution_version != "" {

				if package_name != "" && package_name != "any" {

					vulnerabilities, ok2 := cache.lookup[distribution_name][distribution_version][package_name]

					if ok2 == true {

						for _, name := range vulnerabilities {
							found_map[name] = true
						}

					}

				} else {
					result = cache.QueryVulnerabilitiesByDistribution(query_distribution)
				}

			}

			if len(found_map) > 0 {

				for name, _ := range found_map {
					result = append(result, name)
				}

			}

		}

	}

	return result

}

func (cache *Cache) SetVulnerability(vulnerability structs.Vulnerability) {

	if vulnerability.Name != "" {

		cache.Vulnerabilities[vulnerability.Name] = &vulnerability

		for _, distribution := range vulnerability.Distributions {

			_, ok1 := cache.lookup[distribution.Name]

			if ok1 == false {
				cache.lookup[distribution.Name] = make(map[string]map[string][]string)
			}

			_, ok2 := cache.lookup[distribution.Name][distribution.Version]

			if ok2 == false {
				cache.lookup[distribution.Name][distribution.Version] = make(map[string][]string)
			}

			for _, pkg := range vulnerability.Packages {

				_, ok3 := cache.lookup[distribution.Name][distribution.Version][pkg.Name]

				if ok3 == false {
					cache.lookup[distribution.Name][distribution.Version][pkg.Name] = make([]string, 0)
				}

				cache.lookup[distribution.Name][distribution.Version][pkg.Name] = append(cache.lookup[distribution.Name][distribution.Version][pkg.Name], vulnerability.Name)

			}

		}

	}

}
