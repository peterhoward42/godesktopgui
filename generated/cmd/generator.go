// A command that reads a file system rooted at ../../resources, serializes
// what it finds into bytes, and then produces a Go source code file
// (../generated.go) that has the byte representations of those files
//  manifested as variables.
//
// The generated code is structured such that the exported variable
// CompiledFileSystem implements the http.FileSystem interface.
//
// The go:generate comment line below makes is possible to run the command
// easily in the right place with
//
//		cd here
//		go generate
//
//go:generate go run generator.go

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen" // Wherein the magic lies.
)

func main() {
	rootDir := "../../resources"

	// Rather neatly, not only does the generated source code implement
	// http.FileSystem, but the call below that generates the file,
	// **also** expects an http.FileSystem to specify its input.
	var fs http.FileSystem = http.Dir(rootDir)

	// Output the generated file in a directory different from this one,
	// so that it can declare itself as belonging to a non-main package, without
	// clashing with this file being (of necessity in the main package).

	options := vfsgen.Options{
		Filename:     "../generated.go",
		PackageName:  "generated", // What package declaration do we want in the generated file?
		VariableName: "CompiledFileSystem",
	}

	err := vfsgen.Generate(fs, options)
	if err != nil {
		log.Fatalf("Failed to generate Go code to provide compiled file system: %v", err)
	}
}
