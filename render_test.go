package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	gherkin "github.com/cucumber/gherkin-go"
	"github.com/stretchr/testify/require"
)

var (
	testDataDir = "testdata"
	update      = flag.Bool("update", false, "update .golden files")
)

func TestAll(t *testing.T) {

	tests := []struct {
		name        string
		gherkinFile string
		mdFile      string
	}{
		{
			name:        "Minimal feature",
			gherkinFile: "minimal.feature",
			mdFile:      "minimal.md",
		},
		{
			name:        "Datatables",
			gherkinFile: "datatables.feature",
			mdFile:      "datatables.md",
		},
		{
			name:        "Descriptions",
			gherkinFile: "descriptions.feature",
			mdFile:      "descriptions.md",
		},
		{
			name:        "Doc strings",
			gherkinFile: "docstrings.feature",
			mdFile:      "docstrings.md",
		},
		{
			name:        "Escaped pipes",
			gherkinFile: "escaped_pipes.feature",
			mdFile:      "escaped_pipes.md",
		},
		{
			name:        "Multiple tokens",
			gherkinFile: "example_token_multiple.feature",
			mdFile:      "example_token_multiple.md",
		},
		{
			name:        "Tokens everywhere",
			gherkinFile: "example_tokens_everywhere.feature",
			mdFile:      "example_tokens_everywhere.md",
		},
		{
			name:        "Incomplete background 1",
			gherkinFile: "incomplete_background.feature",
			mdFile:      "incomplete_background.md",
		},
		{
			name:        "Incomplete background 2",
			gherkinFile: "incomplete_background2.feature",
			mdFile:      "incomplete_background2.md",
		},
		{
			name:        "Incomplete feature 1",
			gherkinFile: "incomplete_feature1.feature",
			mdFile:      "incomplete_feature1.md",
		},
		{
			name:        "Incomplete feature 2",
			gherkinFile: "incomplete_feature2.feature",
			mdFile:      "incomplete_feature2.md",
		},
		{
			name:        "Incomplete feature 3",
			gherkinFile: "incomplete_feature3.feature",
			mdFile:      "incomplete_feature3.md",
		},
		{
			name:        "Incomplete scenario",
			gherkinFile: "incomplete_scenario.feature",
			mdFile:      "incomplete_scenario.md",
		},
		{
			name:        "Minimal example",
			gherkinFile: "minimal_example.feature",
			mdFile:      "minimal_example.md",
		},
		{
			name:        "Scenario outline minimal example",
			gherkinFile: "scenario_outline_minimal.feature",
			mdFile:      "scenario_outline_minimal.md",
		},
		{
			name:        "Scenario outline with docstring example",
			gherkinFile: "scenario_outline_with_docstring.feature",
			mdFile:      "scenario_outline_with_docstring.md",
		},
		{
			name:        "Scenario outlines with tags example",
			gherkinFile: "scenario_outlines_with_tags.feature",
			mdFile:      "scenario_outlines_with_tags.md",
		},
		{
			name:        "Several examples example",
			gherkinFile: "several_examples.feature",
			mdFile:      "several_examples.md",
		},
		{
			name:        "Tagged feature with scenario outline example",
			gherkinFile: "tagged_feature_with_scenario_outline.feature",
			mdFile:      "tagged_feature_with_scenario_outline.md",
		},
		{
			name:        "Tags example",
			gherkinFile: "tags.feature",
			mdFile:      "tags.md",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tr *testing.T) {
			gherkinFilePath := filepath.Join(testDataDir, test.gherkinFile)
			mdFilePath := filepath.Join(testDataDir, test.mdFile)

			inputGherkin, err := ioutil.ReadFile(gherkinFilePath)
			require.NoErrorf(t, err, "Error opening gherkin file %s", gherkinFilePath)

			parsedGherkin, err := gherkin.ParseGherkinDocument(strings.NewReader(string(inputGherkin)))
			require.NoErrorf(t, err, "Failed to parse gherkin document.")

			outputMarkdown := newRenderer().Render(parsedGherkin)

			if *update {
				ioutil.WriteFile(mdFilePath, []byte(outputMarkdown), 0644)
			}

			expectedMarkdown, err := ioutil.ReadFile(mdFilePath)
			require.NoErrorf(t, err, "Error opening markdown file %s", mdFilePath)

			assert.Equal(t, string(expectedMarkdown), outputMarkdown)

		})
	}

}
