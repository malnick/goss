package outputs

import (
	"io"
	"time"

	"github.com/aelsabbahy/goss/resource"
)

type Silent struct{}

func (r Silent) SetReportURL(stringified string) error { return nil }

func (r Silent) Output(w io.Writer, results <-chan []resource.TestResult, startTime time.Time) (exitCode int) {
	var failed int
	for resultGroup := range results {
		for _, testResult := range resultGroup {
			switch testResult.Result {
			case resource.FAIL:
				failed++
			}
		}
	}

	if failed > 0 {
		return 1
	}
	return 0
}

func init() {
	RegisterOutputer("silent", &Silent{})
}
