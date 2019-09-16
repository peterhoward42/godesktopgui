## Make a standalone desktop GUI application with Go

Go has neither a native GUI, nor mature bindings to Qt or another similarly
sophisticated GUI library. So this program explores a way for Go to produce a
locally running GUI that looks like this:

![GUI screenshot](docs/screenshotimg.png?raw=true "Some title abc xxx")

The app uses an HTML5 web-app architecture, in which the
files required for content delivery, and the dedicated server, are
compiled together into a single deployable executable file.

It is less than 100 lines of Go code.

The auxilliary files are converting into compilable Go source code using
https://github.com/shurcooL/vfsgen . The example GUI is a loose copy of
[the Github GUI](https://github.com/peterhoward42/godesktopgui) .

The page controls, layout and style are all implented with the
Bootstrap 3 CSS library: https://getbootstrap.com/docs/3.3/ .

Note that Bootstrap 3 comes bundled with an icon library. If you want to
upgrade to Bootstrap 4 (which does not), you can easily use a different icon
library like https://useiconic.com/open . Iconic has explict support for
Bootstrap. You then include the font and css files it depends on in the same
way as the other files are included.

The GUI html is produced using Go's native html
templating: https://golang.org/pkg/html/template/ .

### Build it and try it out

	go get github.com/peterhoward42/godesktopgui
    cd godesktopgui
    make run

