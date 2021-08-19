package docx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestReplaceText(t *testing.T) {
// 	diff := []struct {
// 		Original string
// 		Expected []WP
// 	}{
// 		{},
// 	}
// 	for _, diffItem := range diff {
// 		res, err := GetTextFromXML(diffItem.Original)
// 		if err != nil {
// 			t.Error(err)
// 			return
// 		}
// 		//	fmt.Printf("res: %s\n", res)
// 		assert.Equal(t, diffItem.Expected, res)
// 	}
// }
func TestGetBlockIDByTagNew(t *testing.T) {
	diff := []struct {
		TestName string
		Tag      string
		//EditWord    string
		SourceDoc   Document
		ExpectedID  int
		ExpectedDoc Document
	}{
		{
			TestName: "First",
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
								Body:   `<w:rPr></w:rPr><w:t>e}</w:t>`,
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
			idWP, _, _, _, err := tt.SourceDoc.GetBlockIDByTagNew(tt.Tag)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.ExpectedID, idWP)
			//doc := tt.SourceDoc.WP[id]
			// got := doc.CreateMarkedStringList(tt.args.mp, tt.args.letter...)
			//	assert.Equal(t, tt.ExpectedDoc, tt.SourceDoc)

		})
	}
}
func TestReplaceTextByIDWP(t *testing.T) {
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
			err := tt.SourceDoc.ReplaceTextByIDWP(tt.Tag, tt.EditWord)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.ExpectedDoc, tt.SourceDoc)
			//doc := tt.SourceDoc.WP[id]
			// got := doc.CreateMarkedStringList(tt.args.mp, tt.args.letter...)
			//	assert.Equal(t, tt.ExpectedDoc, tt.SourceDoc)

		})
	}
}
