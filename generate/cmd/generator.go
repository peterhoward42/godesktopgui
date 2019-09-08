//go:generate go run generator.go

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	rootDir := "../../resources"
	var fs http.FileSystem = http.Dir(rootDir)

    // Output the generated file in a directory different from this one,
    // so that it declaring itself as belonging to a non-main package, does 
    // not clash with this one being obliged to (being a cmd).

	options := vfsgen.Options{
		Filename:     "../generated.go",
		PackageName:  "generate",
		VariableName: "CompiledFileSystem",
	}

	err := vfsgen.Generate(fs, options)
	if err != nil {
		log.Fatalf("Failed to generate Go code to provide compiled file system: %v", err)
	}
}
