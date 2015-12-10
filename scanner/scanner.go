package scanner

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"github.com/composer22/domain-scanner/logger"
)

const (
	maxJobs = 1000000 // The jobq maximum number of jobs to hold. We need something for non blocking.
)

var (
	maxScannerSleep = 100 * time.Millisecond // How long should the scanner sleep before checking for results.
)

// Scanner is a manager of scanning jobs and evaluates the results of the workers.
type Scanner struct {
	FilePath   string         // File to read domains from.
	MaxWorkers int            // Maximum number of workers to process jobs.
	StartTime  time.Time      // When the scanner started runnning.
	EndTime    time.Time      // When the scanner ended.
	mu         sync.Mutex     // For locking access.
	wg         sync.WaitGroup // Synchronize close() of job channel.
	log        *logger.Logger // Logger for writing errors.
	jobq       chan *scanJob  // Channel to send jobs.
	doneCh     chan *scanJob  // Channel to receive done jobs.
}

// New is a factory function that creates a new Scanner instance.
func New(filePath string, maxWorkers int) *Scanner {
	return &Scanner{
		FilePath:   filePath,
		MaxWorkers: maxWorkers,
		log:        logger.New(logger.UseDefault, false),
		jobq:       make(chan *scanJob, maxJobs),
		doneCh:     make(chan *scanJob, maxJobs),
	}
}

// PrintVersionAndExit prints the version of the scanner then exits.
func PrintVersionAndExit() {
	fmt.Printf("domain-scanner version %s\n", version)
	os.Exit(0)
}

// Run starts the scanner and manages the jobs.
func (s *Scanner) Run() {
	// Trap all signals to quit.
	s.handleSignals()

	s.mu.Lock()

	// Spin up the workers
	for i := 0; i < s.MaxWorkers; i++ {
		s.wg.Add(1)
		go scanWorker(s.jobq, s.doneCh, &s.wg)
	}

	s.StartTime = time.Now()
	s.mu.Unlock()

	// Read the domains from the file and create scan jobs.
	file, err := os.Open(s.FilePath)
	if err != nil {
		s.log.Errorf("Cannot open file %s", s.FilePath)
		return
	}
	defer file.Close()
	numJobs := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p, _ := url.Parse(fmt.Sprintf("http://%s", scanner.Text()))
		s.jobq <- scanJobNew(p)
		numJobs++
	}

	if err := scanner.Err(); err != nil {
		s.log.Errorf(err.Error())
		return
	}

	fmt.Println("requested\trendered\tcode\terror")

	// Loop looking for results.
	for numJobs > 0 {
		select {
		case j, ok := <-s.doneCh:
			if !ok {
				return
			}
			s.evaluate(j)
			numJobs--
		default:
			runtime.Gosched()
		}
	}
}

// Stop performs close out procedures.
func (s *Scanner) Stop() {
	s.EndTime = time.Now()
	close(s.jobq)
	s.wg.Wait()
	close(s.doneCh)
}

// handleSignals responds to operating system interrupts such as application kills.
func (s *Scanner) handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			s.Stop()
			os.Exit(0)
		}
	}()
}

// evaluate examines the result of the job and prints out the result.
func (s *Scanner) evaluate(j *scanJob) {
	if j.Response != nil {
		fmt.Printf("%s\t%s\t%d\t\n", j.URL.Host, j.Response.Request.URL, j.Response.StatusCode)
	} else {
		fmt.Printf("%s\t*error\t0\t%s\n", j.URL.Host, j.Error)
	}
}
