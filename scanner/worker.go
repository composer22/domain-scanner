package scanner

import (
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var (
	workerMaxSleep = 10 * time.Millisecond // How long should a worker sleep between jobq peeks.
)

// scanWorker is used as a go routine wrapper to handle URL scan jobs.
func scanWorker(jobq chan *scanJob, doneCh chan *scanJob, wg *sync.WaitGroup) {
	defer wg.Done()
	cl := &http.Client{
		Timeout: 30 * time.Second,
	}

	for {
		select {
		case j, ok := <-jobq:
			if !ok {
				return
			}
			j.StartTime = time.Now()
			resp, err := cl.Get(j.URL.String())
			if err != nil {
				j.Error = err.Error()
			} else {
				io.Copy(ioutil.Discard, resp.Body)
				resp.Body.Close()
				j.Response = resp
			}
			j.EndTime = time.Now()
			doneCh <- j
		default:
		}
		runtime.Gosched()
		//		time.Sleep(workerMaxSleep) // Sleep before peeking again.
	}
}
