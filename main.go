package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if err := command(os.Args[1:], os.Stdout); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

func command(ss []string, w io.Writer) error {
	args := getArguments(ss)

	ignoreTags := extractTags(args.IgnoreTags)

	if args.File == "" {
		return convertFiles(args.SrcDir, args.DestDir, ignoreTags)
	}

	return convertFile(args.File, ignoreTags, w)
}

func extractTags(s string) []string {
	tags := strings.Split(s, ",")
	var tagsTrimmed []string
	for _, tag := range tags {
		tagsTrimmed = append(tagsTrimmed, strings.TrimSpace(tag))
	}
	return tagsTrimmed
}
