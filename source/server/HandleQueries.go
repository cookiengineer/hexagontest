package server

import "encoding/json"
import "fmt"
import "net/http"

func HandleQueries(cache *Cache) {

	http.HandleFunc("/api/query/vulnerabilities-by-system/{name}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			name   := request.PathValue("name")
			system := cache.GetSystem(name)

			vulnerabilities := cache.QueryVulnerabilitiesByDistribution(system.Distribution.Name)
			payload, err := json.MarshalIndent(vulnerabilities, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/query/vulnerabilities-by-system/" + name + ": ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/query/vulnerabilities-by-system/" + name + ": error")

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

	http.HandleFunc("/api/query/vulnerabilities-by-distribution/{query}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			query           := request.PathValue("query")
			vulnerabilities := cache.QueryVulnerabilitiesByDistribution(query)
			payload, err    := json.MarshalIndent(vulnerabilities, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/query/vulnerabilities-by-distribution/" + query + ": ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/query/vulnerabilities-by-distribution/" + query + ": error")

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

	http.HandleFunc("/api/query/vulnerabilities-by-distribution-and-package/{query_distribution}/{query_package}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			query_distribution := request.PathValue("query_distribution")
			query_package      := request.PathValue("query_package")
			vulnerabilities    := cache.QueryVulnerabilitiesByDistributionAndPackage(query_distribution, query_package)
			payload, err       := json.MarshalIndent(vulnerabilities, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/query/vulnerabilities-by-distribution-and-package/" + query_distribution + "/" + query_package + ": ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/query/vulnerabilities-by-distribution-and-package/" + query_distribution + "/" + query_package + ": error")

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
