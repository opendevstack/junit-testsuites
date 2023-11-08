package combine

import (
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCombine(t *testing.T) {
	tempDir := mkdirTempDir(t)
	defer rmTempDir(tempDir)
	b := new(bytes.Buffer)
	if err := CombineTestsuites(
		b,
		Options{
			FilesGlob:  "../testdata/fixtures/TEST-*.xml",
			Name:       "combined",
			KeepStdout: false,
			KeepStderr: false,
		}); err != nil {
		t.Fatal(err)
	}

	compareContent(t,
		golden(t, "combined.xml"),
		b,
	)
}

func golden(t *testing.T, path string) io.Reader {
	want, err := os.Open(filepath.Join("../testdata/golden", path))
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return want
}

func mkdirTempDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	return tempDir
}

func rmTempDir(tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		log.Printf("Failed to remove %s\n", tempDir)
	}
}

func compareContent(t *testing.T, wantFile, gotFile io.Reader) {
	t.Helper()
	got, err := io.ReadAll(gotFile)
	if err != nil {
		t.Error(err)
		return
	}
	want, err := io.ReadAll(wantFile)
	if err != nil {
		t.Error(err)
		return
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
