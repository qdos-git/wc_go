
This is a basic Python wrapper around a Go binary, to experiment with
this idea, generally.

For example, it could be useful to compile a binary for multiple
architectures, operating systems, and then call them as needed.

All arguments are passed to the binary when `wc.py` is called, and it
is to be called just like a traditional `wc` implementation.

Careful configuration has made it callable from any directory, once
the `.whl` Wheel` package is installed, by utilising the `site`
standard library.