package main

import (
	"github.com/docopt/docopt-go"
)

var usage = `Gherkin to Markdown converter

Usage:
	gherkin2markdown <file> [--ignoretags=<tags>]
	gherkin2markdown <srcdir> <destdir> [--ignoretags=<tags>]

Arguments:
	<file>  The file to be converted
	<srcdir>  The source directory to be converted
	<destdir>  The destination directory to output converted files to

Options:
	-h, --help          Show this help.
	--ignoretags=<tags>  Ignore scenarios or features with matching tag`

type arguments struct {
	File       string `docopt:"<file>"`
	SrcDir     string `docopt:"<srcdir>"`
	DestDir    string `docopt:"<destdir>"`
	IgnoreTags string `docopt:"--ignoretags"`
}

func getArguments(ss []string) arguments {
	args := arguments{}
	parseArguments(usage, ss, &args)
	return args
}

func parseArguments(u string, ss []string, args interface{}) {

	opts, err := docopt.ParseArgs(u, ss, "0.1.0")
	if err != nil {
		panic(err)
	}

	if err := opts.Bind(args); err != nil {
		panic(err)
	}
}
