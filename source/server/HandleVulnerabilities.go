package server

import "encoding/json"
import "fmt"
import "net/http"

func HandleVulnerabilities(cache *Cache) {

	http.HandleFunc("/api/vulnerabilities", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			identifiers := make([]string, 0)

			for id, _ := range cache.Vulnerabilities {
				identifiers = append(identifiers, id)
			}

			payload, err := json.MarshalIndent(identifiers, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/vulnerabilities: ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/vulnerabilities: error")

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

	http.HandleFunc("/api/vulnerabilities/{name}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			name          := request.PathValue("name")
			vulnerability := cache.GetVulnerability(name)

			if vulnerability != nil {

				payload, _ := json.MarshalIndent(vulnerability, "", "\t")

				fmt.Println("> GET /api/vulnerabilities/" + name + ": ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/vulnerabilities/" + name + ": error")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte("{}"))

			}

		} else {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("[]"))

		}

	})

}
