// domain-scanner is a simple validation routine that calls a list of domains and returns status checks
package main

import (
	"flag"
	"runtime"

	"github.com/composer22/domain-scanner/scanner"
)

// main is the main entry point for the application.
func main() {
	var inputFilePath string
	var procs int
	var maxWorkers int
	var showVersion bool

	flag.StringVar(&inputFilePath, "f", "input.txt", "List of domains to scan.")
	flag.StringVar(&inputFilePath, "filepath", "input.txt", "List of domains to scan.")
	flag.IntVar(&procs, "X", scanner.DefaultMaxProcs, "Maximum processor cores to use.")
	flag.IntVar(&procs, "procs", scanner.DefaultMaxProcs, "Maximum processor cores to use.")
	flag.IntVar(&maxWorkers, "W", scanner.DefaultMaxWorkers, "Maximum Job Workers.")
	flag.IntVar(&maxWorkers, "workers", scanner.DefaultMaxWorkers, "Maximum Job Workers.")
	flag.BoolVar(&showVersion, "V", false, "Show version")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Usage = scanner.PrintUsageAndExit
	flag.Parse()

	// Version flag request?
	if showVersion {
		scanner.PrintVersionAndExit()
	}

	runtime.GOMAXPROCS(procs)
	s := scanner.New(inputFilePath, maxWorkers)
	s.Run()
}
