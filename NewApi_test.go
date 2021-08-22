package docx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	r, err := ReadDocxFile("./testingXML/TestDocumentTag.docx")
	if err != nil {
		t.Error(err)
		return
	}
	docx1 := r.Editable()

	f := Document{
		[]WP{
			{
				Tag: "w:p",
				Body: []WPTokens{
					{
						Tag:    "w:pPr",
						Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
						Status: Open,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t>This is a</w:t>`,
						Status: Open,
					},
					//<w:bookmarkStart w:id="0" w:name="_GoBack"/><w:bookmarkEnd w:id="0"/><w:r><w:rPr></w:rPr><w:t xml:space="preserve"> word document.</w:t>
					{
						Tag:    "w:bookmarkStart",
						Body:   "",
						Attr:   `w:id="0" w:name="_GoBack"`,
						Status: Self,
					},
					{
						Tag:    "w:bookmarkEnd",
						Body:   "",
						Attr:   `w:id="0"`,
						Status: Self,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t xml:space="preserve"> word document.</w:t>`,
						Status: Open,
					},
				},
			},
		},
		SectPr{
			Tag:  "w:sectPr",
			Body: `<w:headerReference w:type="default" r:id="rId2"/><w:footerReference w:type="default" r:id="rId3"/><w:type w:val="nextPage"/><w:pgSz w:w="12240" w:h="15840"/><w:pgMar w:left="1440" w:right="1440" w:header="720" w:top="1440" w:footer="720" w:bottom="1440" w:gutter="0"/><w:pgNumType w:fmt="decimal"/><w:formProt w:val="false"/><w:textDirection w:val="lrTb"/><w:docGrid w:type="default" w:linePitch="360" w:charSpace="0"/>`,
		},
	}
	res, err := docx1.Parser()
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, f, res)
}
func TestAddNewBlock(t *testing.T) {
	r, err := ReadDocxFile("./testingXML/TestDocumentTag.docx")
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
	expected := Document{
		[]WP{
			{
				Tag: "w:p",
				Body: []WPTokens{
					{
						Tag:    "w:pPr",
						Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
						Status: Open,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t>This is a</w:t>`,
						Status: Open,
					},
					{
						Tag:    "w:bookmarkStart",
						Body:   "",
						Attr:   `w:id="0" w:name="_GoBack"`,
						Status: Self,
					},
					{
						Tag:    "w:bookmarkEnd",
						Body:   "",
						Attr:   `w:id="0"`,
						Status: Self,
					},
					{
						Tag:    "w:r",
						Body:   `<w:rPr></w:rPr><w:t xml:space="preserve"> word document.</w:t>`,
						Status: Open,
					},
				},
			},
			{
				Tag: "w:p",
				Body: []WPTokens{
					{
						Tag:    "w:pPr",
						Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
						Status: Open,
					},
					{
						Tag:    "w:r",
						Body:   "<w:rPr></w:rPr><w:t>Simple</w:t>",
						Status: Open,
					},
				},
			},
		},
		SectPr{
			Tag:  "w:sectPr",
			Body: `<w:headerReference w:type="default" r:id="rId2"/><w:footerReference w:type="default" r:id="rId3"/><w:type w:val="nextPage"/><w:pgSz w:w="12240" w:h="15840"/><w:pgMar w:left="1440" w:right="1440" w:header="720" w:top="1440" w:footer="720" w:bottom="1440" w:gutter="0"/><w:pgNumType w:fmt="decimal"/><w:formProt w:val="false"/><w:textDirection w:val="lrTb"/><w:docGrid w:type="default" w:linePitch="360" w:charSpace="0"/>`,
		},
	}
	doc.AddNewBlock("Simple")
	//fmt.Println(doc)
	assert.Equal(t, expected, doc)
}
func TestTemplateBlock(t *testing.T) {
	exp := WP{
		Tag: "w:p",
		Body: []WPTokens{
			{
				Tag:    "w:pPr",
				Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
				Status: Open,
			},
			{
				Tag:    "w:r",
				Body:   fmt.Sprintf("<w:rPr></w:rPr><w:t>%s</w:t>", "test"),
				Status: Open,
			},
		},
	}
	res := templateBlock("test")
	assert.Equal(t, exp, res)
}
func TestGetBlockByIDName(t *testing.T) {
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
	id, err := doc.GetBlockIDByTagDeprecated("{Name}")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 0, id)
}
func TestGetBlockByIDNameBroken(t *testing.T) {
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
	_, err = doc.GetBlockIDByTagDeprecated("{BrokenName}")
	if err == nil {
		t.Error(err, "expected error")
		return
	}
	//assert.Equal(t, 0, id)
}
func TestEditBlockByID(t *testing.T) {
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
	id, err := doc.GetBlockIDByTagDeprecated("{Name}")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 0, id)

}
func TestReplaceTagString(t *testing.T) {
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
	err = doc.ReplaceTagString("{Name}", "KVA")
	if err != nil {
		t.Error(err)
		return
	}
	//	fmt.Println(doc.WP[0].Body)
	exp := `<w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="000000"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:shd w:fill="FFFFFF" w:val="clear"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Name: KVA</w:t>`
	assert.Equal(t, exp, doc.WP[0].Body[len(doc.WP[0].Body)-1].Body)
}

func TestBodyToString(t *testing.T) {
	r, err := ReadDocxFile("./testingXML/TestDocumentTag.docx")
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
	expected := `<w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>This is a</w:t></w:r><w:bookmarkStart w:id="0" w:name="_GoBack"/><w:bookmarkEnd w:id="0"/><w:r><w:rPr></w:rPr><w:t xml:space="preserve"> word document.</w:t></w:r></w:p><w:sectPr><w:headerReference w:type="default" r:id="rId2"/><w:footerReference w:type="default" r:id="rId3"/><w:type w:val="nextPage"/><w:pgSz w:w="12240" w:h="15840"/><w:pgMar w:left="1440" w:right="1440" w:header="720" w:top="1440" w:footer="720" w:bottom="1440" w:gutter="0"/><w:pgNumType w:fmt="decimal"/><w:formProt w:val="false"/><w:textDirection w:val="lrTb"/><w:docGrid w:type="default" w:linePitch="360" w:charSpace="0"/></w:sectPr>`
	res := doc.BodyToString()
	assert.Equal(t, expected, res)
}
func TestWpTokenToString(t *testing.T) {
	templ := WP{
		Tag: "w:p",
		Body: []WPTokens{
			{
				Tag:    "w:pPr",
				Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
				Status: Open,
			},
			{
				Tag:    "w:r",
				Body:   fmt.Sprintf("<w:rPr></w:rPr><w:t>%s</w:t>", "test"),
				Status: Open,
			},
		},
	}
	expected := `<w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>test</w:t></w:r></w:p>`
	res := wpTokenToString(templ)
	assert.Equal(t, expected, res)
	//t.Error(t, expected)
}
func TestCompile(t *testing.T) {
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
	err = doc.ReplaceTagString("{Name}", "KVA")
	if err != nil {
		t.Error(err)
		return
	}
	//	fmt.Printf("doc:[%s]\n", docx1.content)
	//	fmt.Println(doc.WP[0].Body)
	//	exp := `<w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="00000A"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Name: KVA</w:t>`
	//assert.Equal(t, exp, doc.WP[0].Body[len(doc.WP[0].Body)-1].Body)
	expDoc := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" mc:Ignorable="w14 wp14"><w:body><w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:color w:val="000000"/></w:rPr></w:pPr><w:r><w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="000000"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:shd w:fill="FFFFFF" w:val="clear"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Name: KVA</w:t></w:r></w:p><w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:color w:val="000000"/></w:rPr></w:pPr><w:r><w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="000000"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Log: {Log} Date {Date}</w:t></w:r></w:p><w:sectPr><w:headerReference w:type="default" r:id="rId2"/><w:footerReference w:type="default" r:id="rId3"/><w:type w:val="nextPage"/><w:pgSz w:w="12240" w:h="15840"/><w:pgMar w:left="1440" w:right="1440" w:header="720" w:top="1440" w:footer="720" w:bottom="1440" w:gutter="0"/><w:pgNumType w:fmt="decimal"/><w:formProt w:val="false"/><w:textDirection w:val="lrTb"/><w:docGrid w:type="default" w:linePitch="360" w:charSpace="0"/></w:sectPr></w:body></w:document>`
	err = docx1.Compile("./testingDOC/doc.docx", doc)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, expDoc, docx1.content)
	err = doc.ReplaceTagString("{Log}", "Test")
	if err != nil {
		t.Error(err)
		return
	}
	err = doc.ReplaceTagString("{Date}", "Daaaaata")
	if err != nil {
		t.Error(err)
		return
	}
	err = docx1.Compile("./testingDOC/doc2.docx", doc)
	if err != nil {
		t.Error(err)
		return
	}
}

type Exp struct {
	Origin   string
	Expected string
}

func TestScreening(t *testing.T) {
	equal := []Exp{
		{
			Origin:   "Э<div>lala&</div>",
			Expected: "Э&lt;div&gt;lala&amp;&lt;/div&gt;",
		},
	}
	for _, item := range equal {
		assert.Equal(t, item.Expected, Screening(item.Origin))
	}
}
