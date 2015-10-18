package main

import (
    "fmt"
	"html/template"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/peterhoward42/gostandalonegui/resources"
)

// This program demonstrates the potential of Go to create a standalone,
// desktop, application with a sophisticated GUI. It is implemented as a
// bundled-up html5 web application and does require a Browser to render it.
//
// It compiles everything required into a standard Go statically linked and
// autonomuous executable. The executable includes a local http server, the web
// app, and more unusually, all the resource files like CSS and templates. This
// means that there is nothing to deploy apart from the executable file, and
// furthermore, no disk access is performed at run time. 
//
// The Boost CSS library is used without any additions or alteration to create
// the GUI. The auxilliary files are converted into resources that can be
// compiled into the program using the github.com/jteeuwen/go-bindata Go
// package.  
func main() {

	// Unpack the compiled file resources into an in-memory virtual file system.
	virtual_fs := &assetfs.AssetFS{
		Asset: resources.Asset, AssetDir: resources.AssetDir, Prefix: ""}

	// Prepare an html template that will be combined with a data model to
    // serve html pages.
	gui_html_template = extract_and_parse_html_template()

	// Route incoming web page requests for static URLs (like css files) to 
    // the standard library's file server.
	http.Handle("/static/", http.FileServer(virtual_fs))

    // Route incoming web page requests for the GUI home page to the dedicated
    // handler.
	http.HandleFunc("/thegui", gui_home_page_handler)

    fmt.Printf(
        "To see the GUI, visit this URL with your Web Browser:\n\n %s\n\n",
        "http://localhost:47066/thegui")

	// Spin-up the standard library's http server on a hard-coded port.
	http.ListenAndServe(":47066", nil)

}

// Provides a parsed html template, having first extracted the file
// representation of its text from a compiled resource.
func extract_and_parse_html_template() *template.Template {
	// Expose errors by permitting panic response.
	bytes, _ := resources.Asset("templates/maingui.html")
	parsed_template, _ := template.New("gui").Parse(string(bytes))
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
func gui_home_page_handler(w http.ResponseWriter, r *http.Request) {
	// Generate the html by plugging in data from the gui data model into the
	// prepared html template.
	err := gui_html_template.ExecuteTemplate(w, "gui", gui_data())
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

// Makes the the GUI template available at module-scope.
var gui_html_template *template.Template = nil
