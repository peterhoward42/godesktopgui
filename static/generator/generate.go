//go:generate go run generate.go

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	rootDir := "../files"
	var fs http.FileSystem = http.Dir(rootDir)

	options := vfsgen.Options{
		Filename:     "../generated.go",
		PackageName:  "static",
		VariableName: "Assets",
	}

	err := vfsgen.Generate(fs, options)
	if err != nil {
		log.Fatalf("Failed to generate Go code for assets: %v", err)
	}
}
