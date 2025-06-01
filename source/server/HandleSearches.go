package server

import "encoding/json"
import "fmt"
import "net/http"

func HandleSearches(cache *Cache) {

	http.HandleFunc("/api/vulnerabilities/by-system/{name}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			name   := request.PathValue("name")
			system := cache.GetSystem(name)

			vulnerabilities := make([]string, 0)

			for name, vulnerability := range cache.Vulnerabilities {

				for _, pkg := range system.Packages {

					if vulnerability.MatchesPackage(pkg) {
						fmt.Println(name)
						vulnerabilities = append(vulnerabilities, name)
					}

				}

			}

			payload, err := json.MarshalIndent(vulnerabilities, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/vulnerabilities/by-system/" + name + ": ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/vulnerabilities/by-system/" + name + ": error")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte("[]"))

			}

		} else {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("[]"))

		}

	})

}
