package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/uw-labs/gherkin2markdown"
)

func main() {
	if err := command(os.Args[1:], os.Stdout); err != nil {
		if _, err = fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

func command(ss []string, w io.Writer) error {
	args := getArguments(ss)

	ignoreTags := extractTags(args.IgnoreTags)

	if args.File == "" {
		return g2md.ConvertFiles(args.SrcDir, args.DestDir, ignoreTags...)
	}

	data, err := g2md.ConvertFileToString(args.File, ignoreTags...)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, data)
	return err
}

func extractTags(s string) []string {
	tags := strings.Split(s, ",")
	tagsTrimmed := make([]string, len(tags))
	for _, tag := range tags {
		tagsTrimmed = append(tagsTrimmed, strings.TrimSpace(tag))
	}
	return tagsTrimmed
}
