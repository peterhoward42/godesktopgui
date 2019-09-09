package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/peterhoward42/godesktopgui/generate"
)

// htmlTemplate generates the html we wish to serve when we call its
// ExecuteTemplate method.
var htmlTemplate *template.Template

// A demo web server application, serving static content from a single endpoint.
// With the point of interest being that the html, css and javascript have
// been compiled-in to the executable.
func main() {

	// Prepare the html template that will be combined with a data model to
	// serve html pages.

	htmlTemplate = parseTemplate()

	// The html we serve has href links to css and .js files - the URLs of which
	// start with /files, so we route all /files requests to the standard
	// library http.FileServer. The FileServer requires that we provide
	// an http.FileSystem. And that is how the compiled-in files present
	// themselves. See the generate package for how this gets created.

	http.Handle("/files/", http.FileServer(generate.CompiledFileSystem))

	// The GUI home page has its own dedicated handler.
	http.HandleFunc("/thegui", guiHandler)

	fmt.Printf(
		"To see the GUI, visit this URL with your Web Browser:\n\n %s\n\n",
		"http://localhost:47066/thegui")

	// Spin-up the standard library's http server on a hard-coded port.
	http.ListenAndServe(":47066", nil)

}

// Provides a parsed html template, having first extracted the file
// representation of its text from a compiled resource.
func parseTemplate() *template.Template {
	htmlFilename := "files/templates/maingui.html"
	file, err := generate.CompiledFileSystem.Open(htmlFilename)
	if err != nil {
		log.Fatalf("Failed to open <%s>: %v", htmlFilename, err)
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read contents of html file: %v", err)
	}
	parsed_template, err := template.New("gui").Parse(string(contents))
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}
	return parsed_template
}

// A data structure for the model part of the example GUI's model-view pattern.
type GuiDataModel struct {
	Title       string
	Unwatch     int
	Star        int
	Fork        int
	Commits     int
	Branch      int
	Release     int
	Contributor int
	RowsInTable []TableRow
}

// A sub-model to the GuiDataModel that models a single row in the table
// displayed in the GUI.
type TableRow struct {
	File    string
	Comment string
	Ago     string
	Icon    string
}

// Sends the html required to render the GUI into the provided http
// response writer.
func guiHandler(w http.ResponseWriter, r *http.Request) {
	// Generate the html by plugging in data from the gui data model into the
	// prepared html template.
	err := htmlTemplate.ExecuteTemplate(w, "gui", gui_data())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Provides an illustrative, hard-coded instance of a GuiDataModel.
func gui_data() *GuiDataModel {
	gui_data := &GuiDataModel{
		Title:       "Golang Standalone GUI Example",
		Unwatch:     3,
		Star:        0,
		Fork:        2,
		Commits:     31,
		Release:     1,
		Contributor: 1,
		RowsInTable: []TableRow{},
	}
	gui_data.RowsInTable = append(gui_data.RowsInTable,
		TableRow{"do_this.go", "Initial commit", "1 month ago", "file"})
	gui_data.RowsInTable = append(gui_data.RowsInTable,
		TableRow{"do_that.go", "Initial commit", "1 month ago", "file"})
	gui_data.RowsInTable = append(gui_data.RowsInTable,
		TableRow{"index.go", "Initial commit", "1 month ago", "file"})
	gui_data.RowsInTable = append(gui_data.RowsInTable,
		TableRow{"resources", "Initial commit", "2 months ago", "folder-open"})
	gui_data.RowsInTable = append(gui_data.RowsInTable,
		TableRow{"docs", "Initial commit", "2 months ago", "folder-open"})
	return gui_data
}
