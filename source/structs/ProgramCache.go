package structs

type ProgramCache struct {
	Map      map[uint]*Program `json:"map"`
	Names    map[string][]uint `json:"names"`
	Commands map[string][]uint `json:"commands"`
}

func NewProgramCache() ProgramCache {

	var cache ProgramCache

	cache.Map = make(map[uint]*Program)
	cache.Names = make(map[string][]uint, 0)
	cache.Commands = make(map[string][]uint, 0)

	return cache

}

func (cache *ProgramCache) AddProgram(value Program) {

	if value.PID != 0 && value.Name != "" && value.Command != "" {

		_, ok := cache.Map[value.PID]

		if ok == false {

			cache.Map[value.PID] = &value

			_, ok1 := cache.Commands[value.Command]

			if ok1 == true {
				cache.Commands[value.Command] = append(cache.Commands[value.Command], value.PID)
			} else {
				cache.Commands[value.Command] = []uint{value.PID}
			}

			_, ok2 := cache.Names[value.Name]

			if ok2 == true {
				cache.Names[value.Name] = append(cache.Names[value.Name], value.PID)
			} else {
				cache.Names[value.Name] = []uint{value.PID}
			}

		}

	}

}

func (cache *ProgramCache) RemoveProgram(value Program) {

	if value.PID != 0 {

		_, ok := cache.Map[value.PID]

		if ok == true {
			cache.Remove(value.PID)
		}

	}

}

func (cache *ProgramCache) QueryCommand(value string) []*Program {

	var result []*Program

	pids, ok1 := cache.Commands[value]

	if ok1 == true {

		for p := 0; p < len(pids); p++ {

			pid := pids[p]

			program, ok2 := cache.Map[pid]

			if ok2 == true {
				result = append(result, program)
			}

		}

	}

	return result

}

func (cache *ProgramCache) QueryPID(value uint) *Program {

	var pointer *Program = nil

	if value != 0 {

		program, ok := cache.Map[value]

		if ok == true {
			pointer = program
		}

	}

	return pointer

}

func (cache *ProgramCache) QueryName(value string) []*Program {

	var result []*Program

	pids, ok1 := cache.Names[value]

	if ok1 == true {

		for p := 0; p < len(pids); p++ {

			pid := pids[p]

			program, ok2 := cache.Map[pid]

			if ok2 == true {
				result = append(result, program)
			}

		}

	}

	return result

}

func (cache *ProgramCache) Remove(value uint) {

	if value != 0 {

		program, ok := cache.Map[value]

		if ok == true {

			pids_by_command, ok1 := cache.Commands[program.Command]

			if ok1 == true {

				var filtered []uint

				for p := 0; p < len(pids_by_command); p++ {

					if pids_by_command[p] != value {
						filtered = append(filtered, pids_by_command[p])
					}

				}

				cache.Commands[program.Command] = filtered

			}

			pids_by_name, ok2 := cache.Names[program.Name]

			if ok2 == true {

				var filtered []uint

				for p := 0; p < len(pids_by_command); p++ {

					if pids_by_name[p] != value {
						filtered = append(filtered, pids_by_name[p])
					}

				}

				cache.Names[program.Name] = filtered

			}

			delete(cache.Map, value)

		}

	}

}
