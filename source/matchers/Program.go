package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Program struct {
	Name      string   `json:"name"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

func NewProgram() Program {

	var program Program

	program.Command = "any"
	program.Arguments = make([]string, 0)

	return program

}

func ToProgram(value string) Program {

	var program Program

	program.Arguments = make([]string, 0)

	if strings.HasPrefix(value, "/") {
		program.Name = "any"
		program.SetCommand(value)
	} else {
		program.SetName(value)
		program.Command = "any"
	}

	return program

}

func (program *Program) IsIdentical(value Program) bool {

	var result bool = false

	if program.Name == value.Name && program.Command == value.Command {
		result = true
	}

	return result

}

func (program *Program) IsValid() bool {

	var result bool = false

	if program.Name != "" {
		result = true
	}

	return result

}

func (program *Program) Matches(name string, command string) bool {

	var matches_name bool = false
	var matches_command bool = false

	if program.Name == name {
		matches_name = true
	} else if program.Name == "any" {
		matches_name = true
	}

	if program.Command == command {
		matches_command = true
	} else if program.Command == "any" {
		matches_command = true
	}

	return matches_name && matches_command

}

func (program *Program) MatchesArguments(arguments []string) bool {

	var matches_arguments bool = false

	if len(program.Arguments) == len(arguments) {

		matches_arguments = true

		for a := 0; a < len(program.Arguments); a++ {

			if program.Arguments[a] != arguments[a] {
				matches_arguments = false
				break
			}

		}

	}

	return matches_arguments

}

func (program *Program) SetArguments(value []string) {
	program.Arguments = value
}

func (program *Program) SetCommand(value string) {

	if value == "all" || value == "any" || value == "*" {
		program.Command = "any"
	} else if value != "" {
		program.Command = strings.TrimSpace(value)
	}

}

func (program *Program) SetName(value string) {

	if value == "all" || value == "any" || value == "*" {
		program.Name = "any"
	} else if strings.Contains(value, "/") {
		program.Name = strings.TrimSpace(value[0:strings.LastIndex(value, "/")])
	} else if value != "" {
		program.Name = strings.TrimSpace(value)
	}

}

func (program *Program) Hash() string {

	var hash string

	if program.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			program.Name,
			program.Command,
			strings.Join(program.Arguments, " "),
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
