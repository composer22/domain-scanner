# domain-scanner
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Current Release](https://img.shields.io/badge/release-v0.1.1--alpha-brightgreen.svg)](https://github.com/composer22/domain-scanner/releases/tag/v0.0.1)

A scanner to follow domain names and resolve redirects/rendering.

## Description

With a list of domains in a file, this application will run concurrent requests to the domains
to retrun the resolution of the domain and the status.

## Usage

```
Usage: domain-scanner [options...]

Server options:
    -f, -filepath FILEPATH          FILEPATH to list of domains to scan.
    -X, -procs MAX                  MAX processor cores to use from the
	                                 machine (default 1).
    -W, -workers MAX                MAX running workers allowed (default: 4).

Common options:
    -h, --help                       Show this message.
    -V, --version                    Show version.

Example:

    # Scan input.txt; 1 processor; 2 min max; 10 worker go routines.

    ./domain-scanner -f "/foo/bar/input.txt" -X 1 -W 10

```
## Building

This code currently requires version 1.42 or higher of Golang.

Information on Golang installation, including pre-built binaries, is available at
<http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `domain-scanner` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

## License

(The MIT License)

Copyright (c) 2015 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
