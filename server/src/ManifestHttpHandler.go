package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func pathWithCurrentDirectoryForFile(currentDirectory string, file string) string {
	// most likely I've just reinvented bycicle here!
	var builder strings.Builder
	builder.WriteString(currentDirectory)
	builder.WriteString(fmt.Sprintf("/%v", file))
	return builder.String()
}

func hasManifestJSONFile() bool {
	currentDirectory, getWdError := os.Getwd()
	if getWdError != nil {
		panic(getWdError)
	}
	manifestFile := pathWithCurrentDirectoryForFile(currentDirectory, "manifest.json")
	_, error := os.Stat(manifestFile)
	return !os.IsNotExist(error)
}

func manifest(writer http.ResponseWriter, request *http.Request) {
	if !strings.Contains(request.URL.Path, "loadManifest") {
		panic("Handler attached to wrong request!")
	}

	manifestData := loadManifest()

	error := ioutil.WriteFile("manifest.json", manifestData, 0666)
	if error != nil {
		panic(error)
	}
}
