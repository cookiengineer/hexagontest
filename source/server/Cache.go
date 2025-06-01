package server

import "battlemap/structs"

type Cache struct {
	Systems         map[string]*structs.System        `json:"systems"`
	Vulnerabilities map[string]*structs.Vulnerability `json:"vulnerabilities"`
}

func NewCache() *Cache {

	var cache Cache

	cache.Systems = make(map[string]*structs.System)
	cache.Vulnerabilities = make(map[string]*structs.Vulnerability)

	return &cache

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

func (cache *Cache) SetVulnerability(vulnerability structs.Vulnerability) {

	if vulnerability.Name != "" {
		cache.Vulnerabilities[vulnerability.Name] = &vulnerability
	}

}
