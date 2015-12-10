package scanner

import (
	"fmt"
	"os"
)

const usageStr = `
Description: A simple domain scanner in golang to validate hostname reachable.

Usage: domain-scanner [options...]

Server options:
    -f, -filepath FILEPATH          FILEPATH to list of domains to scan.
    -X, -procs MAX                  MAX processor cores to use from the
	                                 machine (default 1).
    -W, -workers MAX                MAX running workers allowed (default: 4).

Common options:
    -h, -help                       Show this message.
    -V, -version                    Show version.

Example:

    # Scan input.txt; 1 processor; 2 min max; 10 worker go routines.

    ./domain-scanner -f "/foo/bar/input.txt" -X 1 -W 10
`

// end help text

// PrintUsageAndExit is used to print out command line options.
func PrintUsageAndExit() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
