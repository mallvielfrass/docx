package docx

import (
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
				Tag:  "w:p",
				Body: `<w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>This is a</w:t></w:r><w:bookmarkStart w:id="0" w:name="_GoBack"/><w:bookmarkEnd w:id="0"/><w:r><w:rPr></w:rPr><w:t xml:space="preserve"> word document.</w:t></w:r>`,
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
				Tag:  "w:p",
				Body: `<w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>This is a</w:t></w:r><w:bookmarkStart w:id="0" w:name="_GoBack"/><w:bookmarkEnd w:id="0"/><w:r><w:rPr></w:rPr><w:t xml:space="preserve"> word document.</w:t></w:r>`,
			},
			{
				Tag:  "w:p",
				Body: `<w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>Simple</w:t></w:r>`,
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

}
