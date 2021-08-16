package docx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditBlockWithNewLine(t *testing.T) {
	r, err := ReadDocxFile("./testingXML/TestDocumentTag2.docx")
	if err != nil {
		t.Error(err)
		return
	}
	docx1 := r.Editable()
	doc, err := docx1.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	err = doc.EditBlockWithNewLine("{Name}", "K\nA")
	if err != nil {
		t.Error(err)
		return
	}
	err = docx1.Compile("./testingDOC/doc236.docx", doc)
	if err != nil {
		t.Error(err)
		return
	}
}
func TestCreateMarkedStringListWithRealDoc(t *testing.T) {
	r, err := ReadDocxFile("./testingXML/TestDocumentTag2.docx")
	if err != nil {
		t.Error(err)
		return
	}
	docx1 := r.Editable()
	doc, err := docx1.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	list := doc.CreateMarkedStringList(MarkerParams{}, "ls", "ls", "test", "45aw5", "fff")
	for _, item := range list {
		doc.AppendWPBlockInToEnd(item)
	}

	err = docx1.Compile("./testingDOC/RL2.docx", doc)
	if err != nil {
		t.Error(err)
		return
	}
}
func TestCreateMarkedStringList(t *testing.T) {
	type args struct {
		mp     MarkerParams
		letter []string
	}
	tests := []struct {
		name string
		args args
		want []WP
	}{
		{
			name: "WithOut Params",
			args: args{
				mp: MarkerParams{},
				letter: []string{
					"раз",
					"два",
					"три",
				},
			},
			want: []WP{
				{
					Tag: "w:p",
					Body: []WPTokens{
						{
							Tag:    "w:pPr",
							Body:   `<w:pStyle w:val="Normal"/><w:numPr><w:ilvl w:val="0"/><w:numId w:val="1"/></w:numPr><w:bidi w:val="0"/><w:jc w:val="left"/><w:rPr></w:rPr>`,
							Attr:   ``,
							Status: Open,
						},
						{
							Tag:    "w:r",
							Body:   `<w:rPr></w:rPr><w:t>раз</w:t>`,
							Attr:   ``,
							Status: Open,
						},
					},
				},
				{
					Tag: "w:p",
					Body: []WPTokens{
						{
							Tag:    "w:pPr",
							Body:   `<w:pStyle w:val="Normal"/><w:numPr><w:ilvl w:val="0"/><w:numId w:val="1"/></w:numPr><w:bidi w:val="0"/><w:jc w:val="left"/><w:rPr></w:rPr>`,
							Attr:   ``,
							Status: Open,
						},
						{
							Tag:    "w:r",
							Body:   `<w:rPr></w:rPr><w:t>два</w:t>`,
							Attr:   ``,
							Status: Open,
						},
					},
				},
				{
					Tag: "w:p",
					Body: []WPTokens{
						{
							Tag:    "w:pPr",
							Body:   `<w:pStyle w:val="Normal"/><w:numPr><w:ilvl w:val="0"/><w:numId w:val="1"/></w:numPr><w:bidi w:val="0"/><w:jc w:val="left"/><w:rPr></w:rPr>`,
							Attr:   ``,
							Status: Open,
						},
						{
							Tag:    "w:r",
							Body:   `<w:rPr></w:rPr><w:t>три</w:t>`,
							Attr:   ``,
							Status: Open,
						},
					},
				},
			},
		},
	}
	r, err := ReadDocxFile("./testingXML/TestDocumentWithMarker.docx")
	if err != nil {
		t.Error(err)
		return
	}
	docx1 := r.Editable()
	doc, err := docx1.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doc.CreateMarkedStringList(tt.args.mp, tt.args.letter...)
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestAtomicWPTokensToString(t *testing.T) {
	type args struct {
		token WPTokens
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TokenToString",
			args: args{
				token: WPTokens{
					Tag:    "w:pPr",
					Body:   `<w:pStyle w:val="Normal"/><w:numPr><w:ilvl w:val="0"/><w:numId w:val="1"/></w:numPr><w:bidi w:val="0"/><w:jc w:val="left"/><w:rPr></w:rPr>`,
					Attr:   ``,
					Status: Open,
				},
			},
			want: `<w:pPr><w:pStyle w:val="Normal"/><w:numPr><w:ilvl w:val="0"/><w:numId w:val="1"/></w:numPr><w:bidi w:val="0"/><w:jc w:val="left"/><w:rPr></w:rPr></w:pPr>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AtomicWPTokensToString(tt.args.token); got != tt.want {
				t.Errorf("AtomicWPTokensToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
