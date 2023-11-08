package combine

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jstemmer/go-junit-report/v2/junit"
)

type Options struct {
	FilesGlob  string
	Name       string
	KeepStdout bool
	KeepStderr bool
}

// CombineTestsuites combines the Junit XML files given through opts
// and writes the result to w.
func CombineTestsuites(w io.Writer, opts Options) error {
	matches, err := filepath.Glob(opts.FilesGlob)
	if err != nil {
		return fmt.Errorf("invalid JUnit glob pattern: %s", err)
	}
	if len(matches) == 0 {
		return errors.New("JUnit glob pattern matched no files")
	}

	suites := &junit.Testsuites{Name: opts.Name}

	for _, m := range matches {
		xmlFile, err := os.Open(m)
		if err != nil {
			return fmt.Errorf("open %s: %s", m, err)
		}
		defer xmlFile.Close()

		dec := xml.NewDecoder(xmlFile)
		var ts junit.Testsuite
		err = dec.Decode(&ts)
		if err != nil {
			return fmt.Errorf("decode %s: %s", m, err)
		}

		if !opts.KeepStdout {
			ts.SystemOut.Data = ""
		}
		if !opts.KeepStderr {
			ts.SystemErr.Data = ""
		}

		suites.AddSuite(ts)
	}

	return suites.WriteXML(w)
}
