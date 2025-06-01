package main

import "battlemap/server"
import "embed"
import "errors"
import "fmt"
import "io/fs"
import "log"
import "net/http"
import "os"
import "strings"

//go:embed public/*
var EMBED_FS embed.FS

func main() {

	cwd, err0 := os.Getwd()

	if err0 == nil {

		if strings.HasSuffix(cwd, "/source") {
			// go run call
			cwd = cwd[0:len(cwd)-7]
		}

		cache := server.NewCache(cwd + "/data")
		fsys, _ := fs.Sub(EMBED_FS, "public")
		fsrv := http.FileServer(http.FS(fsys))

		cache.Init()

		http.Handle("/", fsrv)

		server.HandleSystems(cache)
		server.HandleVulnerabilities(cache)
		server.HandleSearches(cache)

		fmt.Println("Listening on http://localhost:3000")

		err := http.ListenAndServe(":3000", nil)

		if err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(errors.New("Inaccessible process working directory"))
	}

}
