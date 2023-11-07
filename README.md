
# junit-testsuites

This repository provides a simple utility tool (`combine-junit-testsuites`) to combine multiple JUnit XML files, each representing one testsuite, into a single JUnit XML file representing multiple testsuites.

The main functionality is also available as a Go package.

## Usage (CLI)

```
go install github.com/opendevstack/junit-testsuites/cmd/combine-junit-testsuites@latest

combine-junit-testsuites \
    -junit-glob='build/test-results/test/*.xml' \
    -name=combined > combined.xml
```

## Usage (Go)

```
import (
	"log"
	"os"

	"github.com/opendevstack/junit-testsuites/combine"
)

opts := combine.Options{
    JUnitGlob:  "*.xml",
    Name:       "combined",
    KeepStdout: false,
    KeepStderr: false,
}
if err := combine.CombineTestsuites(os.Stdout, opts); err != nil {
    log.Fatal(err)
}
```

## Background

See https://github.com/windyroad/JUnit-Schema for more information on the JUnit XML schema.
