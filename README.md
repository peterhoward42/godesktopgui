Go has neither a native GUI, nor mature bindings to Qt or another similarly
sophisticated library. So this program explores a way for Go to produce a
locally running GUI app using an HTML5 web-app architecture, in which the
content delivery and the dedicated server are compiled together into a single
deployable executable. It comprises less than 100 lines of Go code. It
additionally, compiles the html, css and template files required into the
executable, so the executable has no runtime dependencies apart from a browser
to display it. The auxilliary files are converting into compilable Go source
code using the github.com/jteeuwen/go-bindata Go package. The example GUI is a
loose copy of the Github GUI, and its controls, layout and style are all
implented with the Bootstrap CSS library. Go's native html templating is used.
This repository includes a screenshot of the GUI and a pre-compiled binary (for
Windows). Nb. The individual file download links on Github's web interface
seems to be truncating or corrupting the image and the executable files if you
download them that way, but they are ok if you use the button to download the
zip file - or clone the repository.
