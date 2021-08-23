package docx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceTextByTag(t *testing.T) {
	diff := []struct {
		TestName  string
		Tag       string
		EditWord  string
		SourceDoc Document
		//SourceID    int
		ExpectedID  int
		ExpectedDoc Document
	}{
		{
			TestName: "WithSplit",
			Tag:      "{Name}",
			//EditWord: "Jack",
			ExpectedID: 1,
			EditWord:   "Jack",

			SourceDoc: Document{
				WP: []WP{
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
								Body:   `<w:rPr></w:rPr><w:t>{Nam</w:t>`,
								Attr:   ``,
								Status: Open,
							},
							{
								Tag:    "w:r",
								Body:   `<w:rPr></w:rPr><w:t>e}Lava</w:t>`,
								Attr:   ``,
								Status: Open,
							},
						},
					},
				},
				SectPr: SectPr{},
			},
			ExpectedDoc: Document{
				WP: []WP{
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
								Body:   `<w:rPr></w:rPr><w:t>{Name}</w:t>`,
								Attr:   ``,
								Status: Open,
							},
							{
								Tag:    "w:r",
								Body:   `<w:rPr></w:rPr><w:t>Lava</w:t>`,
								Attr:   ``,
								Status: Open,
							},
						},
					},
				},
				SectPr: SectPr{},
			},
		},
		{
			TestName: "WithOutSplit",
			Tag:      "{Name}",
			//EditWord: "Jack",
			ExpectedID: 1,
			EditWord:   "Jack",

			SourceDoc: Document{
				WP: []WP{
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
								Body:   `<w:rPr></w:rPr><w:t>{Name}</w:t>`,
								Attr:   ``,
								Status: Open,
							},
							{
								Tag:    "w:r",
								Body:   `<w:rPr></w:rPr><w:t>Lava</w:t>`,
								Attr:   ``,
								Status: Open,
							},
						},
					},
				},
				SectPr: SectPr{},
			},
			ExpectedDoc: Document{
				WP: []WP{
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
								Body:   `<w:rPr></w:rPr><w:t>Jack</w:t>`,
								Attr:   ``,
								Status: Open,
							},
							{
								Tag:    "w:r",
								Body:   `<w:rPr></w:rPr><w:t>Lava</w:t>`,
								Attr:   ``,
								Status: Open,
							},
						},
					},
				},
				SectPr: SectPr{},
			},
		},
	}
	for _, tt := range diff {

		t.Run(tt.TestName, func(t *testing.T) {
			err := tt.SourceDoc.ReplaceTextByTag(tt.Tag)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, tt.ExpectedDoc, tt.SourceDoc)
		})
	}
}
func TestRebuildBlocks(t *testing.T) {
	diff := []struct {
		TestName       string
		SourceArray    []string
		ExpectedArray  []string
		Pattern        string
		blockIDWithTag int
	}{
		{
			TestName: "WithSplitIndexZero",
			SourceArray: []string{
				"lala", "{te", "st", "}f", "rr",
			},
			ExpectedArray: []string{
				"lala", "{test}", "", "f", "rr",
			},
			Pattern:        "{test}",
			blockIDWithTag: 1,
		},
		{
			TestName: "WithSplitIndexNonZero",
			SourceArray: []string{
				"lala", "3{te", "st", "}f", "rr",
			},
			ExpectedArray: []string{
				"lala", "3{test}", "", "f", "rr",
			},
			Pattern:        "{test}",
			blockIDWithTag: 1,
		},
		{
			TestName: "WithOutSplit",
			SourceArray: []string{
				"lala", "{test}", "", "f", "rr",
			},
			ExpectedArray: []string{
				"lala", "{test}", "", "f", "rr",
			},
			Pattern:        "{test}",
			blockIDWithTag: 1,
		},
	}
	for _, tt := range diff {

		t.Run(tt.TestName, func(t *testing.T) {
			expectedArray, blockIDWithTag, err := RebuildBlocks(tt.Pattern, tt.SourceArray)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, tt.ExpectedArray, expectedArray)
			assert.Equal(t, tt.blockIDWithTag, blockIDWithTag)
		})
	}
}
func TestGetBlockIDByTag(t *testing.T) {
	diff := []struct {
		TestName   string
		Tag        string
		SourceDoc  Document
		ExpectedID int
	}{
		{
			TestName: "WithSplit",
			Tag:      "{Name}",
			//EditWord: "Jack",
			ExpectedID: 1,

			SourceDoc: Document{
				WP: []WP{
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
								Body:   `<w:rPr></w:rPr><w:t>{Nam</w:t>`,
								Attr:   ``,
								Status: Open,
							},
							{
								Tag:    "w:r",
								Body:   `<w:rPr></w:rPr><w:t>e}Lava</w:t>`,
								Attr:   ``,
								Status: Open,
							},
						},
					},
				},
				SectPr: SectPr{},
			},
		},
		{
			TestName: "WithOutSplit",
			Tag:      "{Name}",
			//EditWord: "Jack",
			ExpectedID: 1,

			SourceDoc: Document{
				WP: []WP{
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
								Body:   `<w:rPr></w:rPr><w:t>{Name}</w:t>`,
								Attr:   ``,
								Status: Open,
							},
							{
								Tag:    "w:r",
								Body:   `<w:rPr></w:rPr><w:t>Lava</w:t>`,
								Attr:   ``,
								Status: Open,
							},
						},
					},
				},
				SectPr: SectPr{},
			},
		},
	}
	for _, tt := range diff {
		t.Run(tt.TestName, func(t *testing.T) {
			id, err := tt.SourceDoc.GetBlockIDByTag(tt.Tag)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, tt.ExpectedID, id)
		})
	}
}
func TestExtractWPToArrayTextString(t *testing.T) {
	diff := []struct {
		TestName  string
		SourceDoc WP
		Expected  []string
	}{
		{
			TestName: "WithSplit",
			SourceDoc: WP{
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
						Body:   `<w:rPr></w:rPr><w:t>{Nam</w:t>`,
						Attr:   ``,
						Status: Open,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t>e}Lava</w:t>`,
						Attr:   ``,
						Status: Open,
					},
				},
			},
			Expected: []string{"{Nam", "e}Lava"},
		},
	}
	for _, tt := range diff {

		t.Run(tt.TestName, func(t *testing.T) {
			strArray, err := ExtractWPToArrayTextString(tt.SourceDoc)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, tt.Expected, strArray)
		})
	}
}
func TestBuildArrayTextStringToWP(t *testing.T) {
	diff := []struct {
		TestName      string
		SourceDoc     WP
		ExpectedDoc   WP
		SourceStrings []string
	}{
		{
			TestName:      "First",
			SourceStrings: []string{"{Nam", "e}Lava"},
			SourceDoc: WP{
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
						Body:   `<w:rPr></w:rPr><w:t></w:t>`,
						Attr:   ``,
						Status: Open,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t></w:t>`,
						Attr:   ``,
						Status: Open,
					},
				},
			},

			ExpectedDoc: WP{
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
						Body:   `<w:rPr></w:rPr><w:t>{Nam</w:t>`,
						Attr:   ``,
						Status: Open,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t>e}Lava</w:t>`,
						Attr:   ``,
						Status: Open,
					},
				},
			},
		},
	}
	for _, tt := range diff {

		t.Run(tt.TestName, func(t *testing.T) {
			strArray, err := BuildArrayTextStringToWP(tt.SourceDoc, tt.SourceStrings)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, tt.ExpectedDoc, strArray)
		})
	}
}
