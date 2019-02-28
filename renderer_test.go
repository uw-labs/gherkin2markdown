package g2md

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
		excludeTags string
		gherkinFile string
		mdFile      string
	}{
		{
			name:        "Minimal feature",
			excludeTags: "",
			gherkinFile: "minimal.feature",
			mdFile:      "minimal.md",
		},
		{
			name:        "Datatables",
			excludeTags: "",
			gherkinFile: "datatables.feature",
			mdFile:      "datatables.md",
		},
		{
			name:        "Descriptions",
			excludeTags: "",
			gherkinFile: "descriptions.feature",
			mdFile:      "descriptions.md",
		},
		{
			name:        "Doc strings",
			excludeTags: "",
			gherkinFile: "docstrings.feature",
			mdFile:      "docstrings.md",
		},
		{
			name:        "Escaped pipes",
			excludeTags: "",
			gherkinFile: "escaped_pipes.feature",
			mdFile:      "escaped_pipes.md",
		},
		{
			name:        "Multiple tokens",
			excludeTags: "",
			gherkinFile: "example_token_multiple.feature",
			mdFile:      "example_token_multiple.md",
		},
		{
			name:        "Tokens everywhere",
			excludeTags: "",
			gherkinFile: "example_tokens_everywhere.feature",
			mdFile:      "example_tokens_everywhere.md",
		},
		{
			name:        "Incomplete background 1",
			excludeTags: "",
			gherkinFile: "incomplete_background.feature",
			mdFile:      "incomplete_background.md",
		},
		{
			name:        "Incomplete background 2",
			excludeTags: "",
			gherkinFile: "incomplete_background2.feature",
			mdFile:      "incomplete_background2.md",
		},
		{
			name:        "Incomplete feature 1",
			excludeTags: "",
			gherkinFile: "incomplete_feature1.feature",
			mdFile:      "incomplete_feature1.md",
		},
		{
			name:        "Incomplete feature 2",
			excludeTags: "",
			gherkinFile: "incomplete_feature2.feature",
			mdFile:      "incomplete_feature2.md",
		},
		{
			name:        "Incomplete feature 3",
			excludeTags: "",
			gherkinFile: "incomplete_feature3.feature",
			mdFile:      "incomplete_feature3.md",
		},
		{
			name:        "Incomplete scenario",
			excludeTags: "",
			gherkinFile: "incomplete_scenario.feature",
			mdFile:      "incomplete_scenario.md",
		},
		{
			name:        "Minimal example",
			excludeTags: "",
			gherkinFile: "minimal_example.feature",
			mdFile:      "minimal_example.md",
		},
		{
			name:        "Scenario outline minimal example",
			excludeTags: "",
			gherkinFile: "scenario_outline_minimal.feature",
			mdFile:      "scenario_outline_minimal.md",
		},
		{
			name:        "Scenario outline with docstring example",
			excludeTags: "",
			gherkinFile: "scenario_outline_with_docstring.feature",
			mdFile:      "scenario_outline_with_docstring.md",
		},
		{
			name:        "Scenario outlines with tags example",
			excludeTags: "",
			gherkinFile: "scenario_outlines_with_tags.feature",
			mdFile:      "scenario_outlines_with_tags.md",
		},
		{
			name:        "Several examples example",
			excludeTags: "",
			gherkinFile: "several_examples.feature",
			mdFile:      "several_examples.md",
		},
		{
			name:        "Tagged feature with scenario outline example",
			excludeTags: "",
			gherkinFile: "tagged_feature_with_scenario_outline.feature",
			mdFile:      "tagged_feature_with_scenario_outline.md",
		},
		{
			name:        "Tags example",
			excludeTags: "",
			gherkinFile: "tags.feature",
			mdFile:      "tags.md",
		},
		{
			name:        "Exclude scenario with tags example",
			excludeTags: "@exclusion_tag1",
			gherkinFile: "scenarios_tagged_to_exclude.feature",
			mdFile:      "scenarios_tagged_to_exclude.md",
		},
		{
			name:        "Exclude feature with tag example",
			excludeTags: "@exclusion_tag",
			gherkinFile: "feature_excluded_with_tag.feature",
			mdFile:      "feature_excluded_with_tag.md",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tr *testing.T) {
			gherkinFilePath := filepath.Join(testDataDir, test.gherkinFile)
			mdFilePath := filepath.Join(testDataDir, test.mdFile)

			inputGherkin, err := ioutil.ReadFile(gherkinFilePath)
			require.NoErrorf(tr, err, "Error opening gherkin file %s", gherkinFilePath)

			parsedGherkin, err := gherkin.ParseGherkinDocument(strings.NewReader(string(inputGherkin)))
			require.NoErrorf(tr, err, "Failed to parse gherkin document.")

			outputMarkdown := newRenderer([]string{test.excludeTags}).Render(parsedGherkin)

			if *update {
				assert.NoError(tr, ioutil.WriteFile(mdFilePath, []byte(outputMarkdown), 0644))
			}

			expectedMarkdown, err := ioutil.ReadFile(mdFilePath)
			require.NoErrorf(tr, err, "Error opening markdown file %s", mdFilePath)

			assert.Equal(tr, string(expectedMarkdown), outputMarkdown)
		})
	}

}
