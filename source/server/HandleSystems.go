package server

import "encoding/json"
import "fmt"
import "net/http"

func HandleSystems(cache *Cache) {

	http.HandleFunc("/api/systems", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			identifiers := make([]string, 0)

			for id, _ := range cache.Systems {
				identifiers = append(identifiers, id)
			}

			payload, err := json.MarshalIndent(identifiers, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/systems: ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/systems: error")

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

	http.HandleFunc("/api/systems/{name}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			name := request.PathValue("name")
			system := cache.GetSystem(name)

			if system != nil {

				payload, _ := json.MarshalIndent(system, "", "\t")

				fmt.Println("> GET /api/systems/" + name + ": ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/systems/" + name + ": error")

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
