package g2md

import (
	"fmt"
	"regexp"
	"strings"

	messages "github.com/cucumber/cucumber-messages-go/v2"
	"github.com/willf/pad/utf8"
)

type renderer struct {
	*strings.Builder
	ignoreTags []string
}

func newRenderer(ignoreTags []string) renderer {
	return renderer{&strings.Builder{}, ignoreTags}
}

func (r renderer) Render(d *messages.GherkinDocument) string {
	r.renderFeature(d.Feature)

	return r.Builder.String()
}

func (r renderer) renderFeature(f *messages.Feature) {
	if f == nil {
		return
	}
	if r.matchesIgnoreTags(f.Tags) {
		return
	}

	r.writeLine("# " + f.Name)
	r.writeDescription(f.Description)

	for _, x := range f.Children {

		switch x := x.Value.(type) {
		case *messages.FeatureChild_Background:
			r.writeLine("")
			r.renderBackground(x.Background)
		case *messages.FeatureChild_Scenario:
			if !r.matchesIgnoreTags(x.Scenario.Tags) {
				r.writeLine("")
				r.renderScenario(x.Scenario)
			}
		default:
			panic("unreachable")
		}
	}
}

func (r renderer) matchesIgnoreTags(tags []*messages.Tag) bool {
	for _, tag := range tags {
		for _, ignoreTag := range r.ignoreTags {
			if tag.Name == ignoreTag {
				return true
			}
		}
	}
	return false
}

func (r renderer) renderBackground(b *messages.Background) {
	if len(strings.TrimSpace(b.Name)) == 0 {
		r.writeLine("## Background")
	} else {
		r.writeLine("## Background (" + b.Name + ")")
	}
	r.writeDescription(b.Description)
	r.renderSteps(b.Steps)
}

func (r renderer) renderScenario(s *messages.Scenario) {
	r.writeLine("## " + s.Name)
	r.writeDescription(s.Description)
	r.renderSteps(s.Steps)

	if len(s.Examples) != 0 {
		r.writeLine("")
		r.renderExamples(s.Examples)
	}
}

func (r renderer) renderSteps(ss []*messages.Step) {
	for i, s := range ss {
		r.writeLine("")
		r.renderStep(s, i == len(ss)-1)
	}
}

func (r renderer) renderDocString(d *messages.DocString) {
	if strings.Contains(d.Content, "`") {
		backTicksLength := maxRepeatingBackticks(d.Content)
		if backTicksLength >= 3 {
			r.writeLine(genRepeatingChars("`", backTicksLength+1) + d.ContentType)
			r.writeLine(d.Content)
			r.writeLine(genRepeatingChars("`", backTicksLength+1))
		}
		return
	}
	r.writeLine("```" + d.ContentType)
	r.writeLine(d.Content)
	r.writeLine("```")
}

func maxRepeatingBackticks(str string) int {
	var count, maxCount int
	for i := 0; i < len(str); i++ {
		if str[i] == '`' {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}

func genRepeatingChars(char string, length int) string {
	var genString string
	for i := 0; i < length; i++ {
		genString += char
	}
	return genString
}

func (r renderer) renderStep(s *messages.Step, last bool) {
	if last && s.Argument == nil && s.Text[len(s.Text)-1] != '.' {
		s.Text += "."
	}

	backtickParams := func(s string) string {
		re := regexp.MustCompile(`(?i)(<\s*[^>]*>(.*?))`)
		new := re.ReplaceAllString(s, "`$1`")
		return new
	}

	r.writeLine("_" + strings.TrimSpace(s.Keyword) + "_ " + backtickParams(s.Text))

	if s.Argument != nil {
		r.writeLine("")

		switch x := s.Argument.(type) {
		case *messages.Step_DocString:
			r.renderDocString(x.DocString)
		case *messages.Step_DataTable:
			r.renderDataTable(x.DataTable)
		default:
			panic(fmt.Sprintf("unreachable, type: %v", x))
		}
	}
}

func (r renderer) renderDataTable(dt *messages.DataTable) {

	headerRow := dt.GetRows()[0]

	ws := make([]int, len(headerRow.Cells))

	rows := dt.GetRows()
	rows = append(rows[:0], rows[1:]...)

	for _, r := range append([]*messages.TableRow{headerRow}, rows...) {
		for i, c := range r.Cells {
			if w := len(c.Value); w > ws[i] {
				ws[i] = w
			}
		}
	}

	r.renderCells(headerRow.Cells, ws)

	s := "|"

	for _, w := range ws {
		s += strings.Repeat("-", w+2) + "|"
	}

	r.writeLine(s)

	for _, t := range rows {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderExamples(es []*messages.Examples) {
	r.writeLine("### Examples")

	for _, e := range es {
		if e.Name != "" {
			r.writeLine("")
			r.writeLine("#### " + e.Name)
		}

		r.writeDescription(e.Description)

		r.writeLine("")
		r.renderTable(e.TableHeader, e.TableBody)
	}
}

func (r renderer) renderTable(h *messages.TableRow, rs []*messages.TableRow) {
	ws := make([]int, len(h.Cells))

	for _, r := range append([]*messages.TableRow{h}, rs...) {
		for i, c := range r.Cells {
			if w := len(c.Value); w > ws[i] {
				ws[i] = w
			}
		}
	}

	r.renderCells(h.Cells, ws)

	s := "|"

	for _, w := range ws {
		s += strings.Repeat("-", w+2) + "|"
	}

	r.writeLine(s)

	for _, t := range rs {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderCells(cs []*messages.TableCell, ws []int) {
	s := "|"

	for i, c := range cs {
		s += " " + utf8.Right(c.Value, ws[i], " ") + " |"
	}

	r.writeLine(s)
}

func (r renderer) writeDescription(s string) {
	if s != "" {
		r.writeLine("")
		r.writeLine(strings.TrimSpace(s))
	}
}

func (r renderer) writeLine(s string) {
	_, err := r.WriteString(s + "\n")

	if err != nil {
		panic(err)
	}
}
