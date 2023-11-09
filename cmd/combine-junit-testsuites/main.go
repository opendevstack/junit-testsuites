// Combine multiple JUnit XML files into a single "testsuites" file.
//
// Usage:
//
//	go run github.com/opendevstack/junit-testsuites \
//		"build/test-results/test/*.xml" > combined.xml
//
// See https://github.com/windyroad/JUnit-Schema.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/opendevstack/junit-testsuites/combine"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nExample:\n  combine-junit-testsuites \"build/test-results/test/*.xml\" > combined.xml\n")
	}

	opts := combine.Options{
		FilesGlob:  "*.xml",
		Name:       "combined",
		KeepStdout: false,
		KeepStderr: false,
	}
	flag.StringVar(&opts.FilesGlob, "files", opts.FilesGlob, "Glob pattern of JUnit XML files. Preferrably specified as an argument.")
	flag.StringVar(&opts.Name, "name", opts.Name, "Name of combined testsuites")
	flag.BoolVar(&opts.KeepStdout, "keep-stdout", opts.KeepStdout, "Whether to keep STDOUT of tests")
	flag.BoolVar(&opts.KeepStderr, "keep-stderr", opts.KeepStderr, "Whether to keep STDERR of tests")

	outFlag := flag.String("out", "", "Output filename. When unset, output is written to STDOUT")
	flag.Parse()

	// If non-flag arg is given, treat it as the file glob.
	if flag.Arg(0) != "" {
		opts.FilesGlob = flag.Arg(0)
	}

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
