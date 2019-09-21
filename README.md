## Make a standalone desktop GUI application with Go

This code shows how to make a GUI Go program that will run locally standalone.

It is essentially a self-contained web server, dedicated to serving this GUI, 
using an HTML5 approach, and packaged as a single stand-alone executable with
no external dependencies. It's only around 100 lines of Go code.

There's a walk-through of the key technical aspects [here](https://docs.google.com/presentation/d/1drkVWDZambK5NprhaBVqVuXOPw9rlo6mxrc0kzdsntg/edit?usp=sharing),  
which includes:

- Using the Go `html.Template package` - as a Model/View pattern to 
  generate the HTML.
- The `github.com/shurcooL/vfsgen package` - that generates Go code that
  incapsulates a set of files, so they can be compiled-in to the app.
- Using `go generate` to run the code generation step.
- Using `http.FileServer` to serve static files.
- The brilliant simplicity and utility of the `http.FileSystem` interface.
- The thought provoking `http.Dir` type.
- The `github.com/pkg/browser` package - for programmatically bringing up a 
  tab in the user's browser that points to a URL of your choice.
- How packages like [Bootstrap](https://getbootstrap.com/docs/3.3) provide
  a relatively quick and easy way for non-Front-End specialists to compose 
  a decent looking GUI with rich contemporary controls.

## Here's What the GUI Looks Like

-------

![GUI screenshot](docs/screenshot.png?raw=true "Some title abc xxx")

-------

Note that we use Bootstrap 3, which comes bundled with an icon library. If you
want to upgrade to Bootstrap 4 (which does not), you can easily use an
external icon library like https://useiconic.com/open . Iconic has explict
support for Bootstrap. You then include the font and css files it depends on
in the same way as the other files are included.

### Get it, Build it, Try it out

	go get github.com/peterhoward42/godesktopgui
    cd godesktopgui
    make

This should bring up a tab in your default web browser showing the GUI.
