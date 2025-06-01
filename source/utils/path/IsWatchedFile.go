package path

import "strings"

func IsWatchedFile(value string) bool {

	var result bool = false

	folder := ""
	file := ""

	if strings.Contains(value, "/") {
		folder = value[0:strings.LastIndex(value, "/")]
		file = value[strings.LastIndex(value, "/")+1:]
	} else {
		file = value
	}

	watched_extensions := []string{
		"bin", "sh", "so", // Linux
		"app", "dylib",    // MacOS
		"exe", "dll",      // Windows
	}

	watched_folders := []string{
		"/bin/",
		"/lib/",
		"/lib32/",
		"/lib64/",
		"/run",
		"/sbin/",
		"/tmp",
		"/usr/bin/",
		"/usr/lib/",
		"/usr/lib32/",
		"/usr/lib64/",
		"/usr/sbin/",
		"/usr/local/bin/",
		"/usr/local/lib/",
		"/usr/local/lib32/",
		"/usr/local/lib64/",
		"/usr/local/sbin/",
	}

	if folder != "" {

		for w1 := 0; w1 < len(watched_folders); w1++ {

			prefix := watched_folders[w1]

			if strings.HasPrefix(folder, prefix) {

				if strings.Contains(file, ".") {

					for w2 := 0; w2 < len(watched_extensions); w2++ {

						suffix := "." + watched_extensions[w2]

						if strings.HasSuffix(file, suffix) {
							result = true
							break
						}

					}

				}

			}

			if result == true {
				break
			}

		}

	}

	if result == false && file != "" {

		if strings.Contains(file, ".so.") {

			// Special case: symbolic links to shared object files
			result = true

		} else if strings.Contains(file, ".") {

			for w2 := 0; w2 < len(watched_extensions); w2++ {

				suffix := watched_extensions[w2]

				if strings.HasSuffix(file, "." + suffix) {
					result = true
					break
				}

			}

		}

	}

	return result

}
