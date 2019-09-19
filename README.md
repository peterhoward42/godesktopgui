## Make a standalone desktop GUI application with Go

This code shows how to make a GUI Go program that will run locally standalone.
It is essentially a self-contained web server, dedicated to serving this GUI, 
using an HTML5 approach, and packaged as a single stand-alone executable with
no external dependencies. It's only around 100 lines of Go code.

It aims to be a how-to, and educational example. There's a summary explanation
below, but there's also a fuller, explanation in 
(these slides)[https://docs.google.com/presentation/d/1drkVWDZambK5NprhaBVqVuXOPw9rlo6mxrc0kzdsntg/edit?usp=sharing].

The slides cover:

- Using the Go `html.Template package` - as a Model/View pattern to 
  generate the HTML.
- The `github.com/shurcooL/vfsgen package` - that generates Go code that
  incapsulates a set of files, so they can be compiled-in the the app.
- The above step as a use case for `go generate`.
- Using `http.FileServer` to serve static files.
- The brilliant simplicity and utility of the `http.FileSystem` interface.
- The spooky `http.Di`r type.
- The `github.com/pkg/browser` package - for bring up a tab in the user's
  browser programmatically that points to a URL of your choice.
- How packages like (Bootstrap[https://getbootstrap.com/docs/3.3/] provide
  a relatively quick and easy way for non-Front-End specialists to compose 
  a decent looking GUI with rich contemporary controls.

## Here's What the GUI Looks Like

    (its very roughly modelled on the GitHub repo view)

![GUI screenshot](docs/screenshot.png?raw=true "Some title abc xxx")

Note that we use Bootstrap 3, which comes bundled with an icon library. If you
want to upgrade to Bootstrap 4 (which does not), you can easily use an
external icon library like https://useiconic.com/open . Iconic has explict
support for Bootstrap. You then include the font and css files it depends on
in the same way as the other files are included.

### Get it, Build it, Try it out

	go get github.com/peterhoward42/godesktopgui
    cd godesktopgui
    make run

