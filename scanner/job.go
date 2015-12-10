package scanner

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

// scanJob is a transport packet that represents URL that needs processing.
type scanJob struct {
	URL       *url.URL       `json:"url"`       // The URL to scan.
	StartTime time.Time      `json:"startTime"` // The start time of the scan.
	EndTime   time.Time      `json:"endTime"`   // The end time of the scan.
	Response  *http.Response `json:"response"`  // Response returned from the scan.
	Body      io.ReadCloser  `json:"body"`      // Body returned from the scan.
	Error     string         `json:"error"`     // Error returned from the request.
}

// scanJobNew is a factory for creating a new job instance.
func scanJobNew(u *url.URL) *scanJob {
	return &scanJob{
		URL: u,
	}
}

// String is an implentation of the Stringer interface so the structure is returned as a
// string to fmt.Print() etc.
func (s *scanJob) String() string {
	j, _ := json.Marshal(s)
	return string(j)
}
