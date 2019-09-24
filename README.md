# Make a standalone desktop GUI application with Go

This code shows how to make a GUI Go program that will run locally standalone.

It is essentially a self-contained web server, dedicated to serving this GUI, 
using an HTML5 approach, and packaged as a single stand-alone executable with
no external dependencies. It's only around 100 lines of Go code.


## Here's What the GUI Looks Like

-------

![GUI screenshot](docs/screenshot.png?raw=true "Some title abc xxx")

-------

## Get it, Build it, Try it out

	go get github.com/peterhoward42/godesktopgui
    cd godesktopgui
    make

## Compiling-in the Asset Files

>All the code described in this section is in `cmd/main.go`.

The App is just a main() that starts a web server to serve the HTML at the
`/thegui` endpoint, and then launches a browser tab pointing at the server.

The dynamic aspect of the GUI comes from the HTML being generated on-the-fly,
using Go's `html.Template` package - executing against a `GuiData` structure 
instance.

>i.e. Model/View pattern.

The HTML template is the first example of a file we want to be compiled-in.
The original file lives in the source tree at 
`resources/files/templates/maingui.html`.

We should digress first into Go's `http.FileSystem` interface. Which is very 
simple, defining just a single method:

	Open(name string) (File, error)

If you look at the `parseTemplateFromVirtualFileSystem()` function, you'll find
code that uses this `Open` method on an `http.FileSystem` to read in the 
template file.:

	generate.CompiledFileSystem.Open()

So it seems there's a package somewhee called `generate`, which is exporting
an `http.FileSystem` attribute called `CompiledFileSystem`.

## Generating the Source Code Required

The code that defines the `generate` package is auto-generated using the
`github.com/shurcooL/vfsgen` package; which is capabable of sucking
up files from a real directory tree and expressing their contents as compilable
Go source files

You can see the code for the generation command in
`generate/cmd/generator.go`.

To avoid having to run it manually to build the program, it's built in to
the Makefile dependency graph with the `generate` target.

You'll see that the `generate` make target does not run the command explicitly
but instead does

	go generate

This is a standard Go capability. It looks for comments structured like this 
one in `generate.cmd.generator.go`

	//go:generate go run generator.go

And does what they ask it to.

## CSS and Javascript Files

The HTML we serve makes reference to CSS files and to Javascript files (from the
Bootstrap library).  The files we need also live in that `resources/files`
directory tree, and thus also get incorporated into the compiled-in virtual file
system.

The HTML refers to these files using URLs like:

	/files/css/xxx etc.

So we configure the web server to satisfy any requests for URLs in the style
of `/files/xxx` by serving the file called `xxx` from the compiled-in file 
system:

	http.Handle("/files/", http.FileServer(generate.CompiledFileSystem))

During development, when you are iterating on the HTML, it's handy to 
replace this with a `http.FileServer`that uses the original files instead.

	http.Handle("/files/", http.FileServer(http.Dir(<path>)))

	Where <path> is the directory that contains `/files`.


# Bootstrap

Note that we use Bootstrap 3, which comes bundled with an icon library. If you
want to upgrade to Bootstrap 4 (which does not), you can easily use an
external icon library like https://useiconic.com/open . Iconic has explict
support for Bootstrap. You then include the font and css files it depends on,
in the same way as the other files are included.

