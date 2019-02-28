package g2md

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/cucumber/gherkin-go"
)

const featureFileExtension = ".feature"

// Convert reads data from the provided reader, converts it to markdown and returns the result as a string.
func Convert(r io.Reader, ignoreTags ...string) (string, error) {
	d, err := gherkin.ParseGherkinDocument(r)

	if err != nil {
		return "", err
	}

	return newRenderer(ignoreTags).Render(d), nil
}

// ConvertFileToString loads a file, converts it to markdown and returns the result as a string.
func ConvertFileToString(fileName string, ignoreTags ...string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return Convert(f, ignoreTags...)
}

// ConvertFile loads a source file, converts it to markdown and writes it to the given destination.
func ConvertFile(sourceName, destName string, ignoreTags ...string) error {
	data, err := ConvertFileToString(sourceName, ignoreTags...)
	if err != nil {
		return err
	}

	out, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, data)
	if err != nil {
		_ = out.Close()
		return err
	}

	return out.Close()
}

// ConvertFiles reads all gherkin files in the source directory, converts them to markdown
// and writes the result to the target directory.
func ConvertFiles(sourceDir, destDir string, ignoreTags ...string) error {
	var sources []string

	err := filepath.Walk(sourceDir, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && filepath.Ext(p) == featureFileExtension {
			sources = append(sources, p)
		}
		return nil
	})

	if err != nil {
		return err
	}

	var eg errgroup.Group

	for _, source := range sources {
		eg.Go(newConvertJob(source, sourceDir, destDir, ignoreTags))
	}

	return eg.Wait()
}

func newConvertJob(source, sourceDir, destDir string, ignoreTags []string) func() error {
	return func() error {
		dest, err := filepath.Rel(sourceDir, source)
		if err != nil {
			return err
		}

		dest = strings.TrimSuffix(filepath.Join(destDir, dest), featureFileExtension) + ".md"
		if err = os.MkdirAll(filepath.Dir(dest), 0700); err != nil {
			return err
		}

		if err = ConvertFile(source, dest, ignoreTags...); err != nil {
			return err
		}
		return nil
	}
}
