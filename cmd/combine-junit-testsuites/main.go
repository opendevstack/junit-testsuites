// Combine multiple JUnit XML files into a single "testsuites" file.
//
// Usage:
//
//		go run github.com/opendevstack/junit-testsuites \
//			-junit-glob='build/test-results/test/*.xml' \
//	        -name=combined > combined.xml
//
// See https://github.com/windyroad/JUnit-Schema.
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/opendevstack/junit-testsuites/combine"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func main() {
	opts := combine.Options{
		JUnitGlob:  "*.xml",
		Name:       "combined",
		KeepStdout: false,
		KeepStderr: false,
	}
	flag.StringVar(&opts.JUnitGlob, "junit-glob", opts.JUnitGlob, "Glob pattern of JUnit XML files")
	flag.StringVar(&opts.Name, "name", opts.Name, "Name of combined testsuites")
	flag.BoolVar(&opts.KeepStdout, "keep-stdout", opts.KeepStdout, "Whether to keep STDOUT of tests")
	flag.BoolVar(&opts.KeepStderr, "keep-stderr", opts.KeepStderr, "Whether to keep STDERR of tests")

	outFlag := flag.String("out", "", "Output filename. When unset, output is written to STDOUT")
	flag.Parse()

	var w io.Writer
	if *outFlag == "" {
		w = os.Stdout
	} else {
		outFile, err := os.Create(*outFlag)
		if err != nil {
			log.Fatalf("create %s: %s", *outFlag, err)
		}
		w = outFile
	}
	if err := combine.CombineTestsuites(w, opts); err != nil {
		log.Fatal(err)
	}
}

func safeFilename(str string) string {
	return strings.Trim(nonAlphanumericRegex.ReplaceAllString(str, "-"), "-")
}
