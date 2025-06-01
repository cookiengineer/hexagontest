package types

type Manager string

const (

	ManagerANY Manager = "any"

	// Supported Distributions
	ManagerAPK    Manager = "apk"
	ManagerAPT    Manager = "apt"
	ManagerDNF    Manager = "dnf"
	ManagerPACMAN Manager = "pacman"
	ManagerRPM    Manager = "rpm"
	ManagerTDNF   Manager = "tdnf"
	ManagerZYPPER Manager = "zypper"

	// TODO: Unsupported Distributions
	ManagerPKG    Manager = "pkg"    // FreeBSD
	ManagerPKGSRC Manager = "pkgsrc" // NetBSD, OpenBSD
	ManagerMSI    Manager = "msi"    // Microsoft Installer

	// Programming Languages
	ManagerCARGO      Manager = "cargo"      // Rust
	ManagerCHOCOLATEY Manager = "chocolatey" // C#
	ManagerCOCOAPODS  Manager = "cocoapods"  // Cocoa
	ManagerCOMPOSER   Manager = "composer"   // PHP
	ManagerCONAN      Manager = "conan"      // C++
	ManagerCONDA      Manager = "conda"      // Python
	ManagerCRAN       Manager = "cran"       // R
	ManagerGEM        Manager = "gem"        // Ruby
	ManagerGO         Manager = "go"         // Go
	ManagerGRADLE     Manager = "gradle"     // Java
	ManagerHACKAGE    Manager = "hackage"    // Haskell
	ManagerHEX        Manager = "hex"        // Erlang
	ManagerMAVEN      Manager = "maven"      // Java
	ManagerNPM        Manager = "npm"        // node.js
	ManagerNUGET      Manager = "nuget"      // C#
	ManagerPEAR       Manager = "pear"       // PHP
	ManagerPHAR       Manager = "phar"       // PHP
	ManagerPIP        Manager = "pip"        // Python

)

func IsManager(value string) bool {

	var result bool = false

	// Distributions
	if value == string(ManagerAPK) {
		result = true
	} else if value == string(ManagerAPT) {
		result = true
	} else if value == string(ManagerDNF) {
		result = true
	} else if value == string(ManagerPACMAN) {
		result = true
	} else if value == string(ManagerRPM) {
		result = true
	} else if value == string(ManagerTDNF) {
		result = true
	} else if value == string(ManagerZYPPER) {
		result = true
	} else if value == string(ManagerPKG) {
		result = true
	} else if value == string(ManagerPKGSRC) {
		result = true
	} else if value == string(ManagerMSI) {
		result = true
	}

	// Programming Languages
	if value == string(ManagerCARGO) {
		result = true
	} else if value == string(ManagerCHOCOLATEY) {
		result = true
	} else if value == string(ManagerCOCOAPODS) {
		result = true
	} else if value == string(ManagerCOMPOSER) {
		result = true
	} else if value == string(ManagerCONAN) {
		result = true
	} else if value == string(ManagerCONDA) {
		result = true
	} else if value == string(ManagerCRAN) {
		result = true
	} else if value == string(ManagerGEM) {
		result = true
	} else if value == string(ManagerGO) {
		result = true
	} else if value == string(ManagerGRADLE) {
		result = true
	} else if value == string(ManagerHACKAGE) {
		result = true
	} else if value == string(ManagerHEX) {
		result = true
	} else if value == string(ManagerMAVEN) {
		result = true
	} else if value == string(ManagerNPM) {
		result = true
	} else if value == string(ManagerNUGET) {
		result = true
	} else if value == string(ManagerPEAR) {
		result = true
	} else if value == string(ManagerPHAR) {
		result = true
	} else if value == string(ManagerPIP) {
		result = true
	}

	return result

}

func ParseManager(value string) *Manager {

	var result *Manager = nil

	if IsManager(value) {
		manager := Manager(value)
		result = &manager
	}

	return result

}

func (manager Manager) String() string {
	return string(manager)
}

func (manager *Manager) IsValid() bool {

	var result bool = false

	if IsManager(manager.String()) {
		result = true
	}

	return result

}
