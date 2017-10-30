package outputs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/aelsabbahy/goss/resource"
	"github.com/fatih/color"
)

type Http struct{}

func (r Http) Output(w io.Writer, results <-chan []resource.TestResult, startTime time.Time) (exitCode int) {
	color.NoColor = true

	out, failed := makeMap(results, startTime)

	j, _ := json.MarshalIndent(out, "", "    ")
	fmt.Fprintln(w, string(j))
	if err := postReport(j, r.Report); err != nil {
		fmt.Errorf("errors sending report: %s", err.Error())
	}

	if failed > 0 {
		return 1
	}

	return 0

}

func postReport(json []byte, u *url.URL) error {
	resp, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	fmt.Printf("status code from report URL %s: %s\n", resp.Request.URL.String(), resp.Status)

	return nil
}
