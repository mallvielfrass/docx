package docx

import (
	"fmt"

	"github.com/mallvielfrass/ooxml"
)

func (d *Docx) Parser() (Document, error) {
	_, body, _ := d.ParseNode()
	//	fmt.Println(body)
	nodes, err := ooxml.GetParentNodes(body)
	if err != nil {
		fmt.Println(err)
		return Document{}, err
	}
	var doc Document
	for _, item := range nodes {
		//	fmc.Printfln("#gbt%d) #ybt%s", i+1, item.Name)
		switch item.Name {
		case "w:p":
			doc.WP = append(doc.WP, WP{
				Tag:  item.Name,
				Body: item.Body,
			})
		case "w:sectPr":
			doc.SectPr = SectPr{
				Tag:  item.Name,
				Body: item.Body,
			}
		}

	}
	return doc, nil
}
func templateBlock() {}
func (d *Document) AddNewBlock(s string) {
	d.WP = append(d.WP, WP{
		Tag:  "w:p",
		Body: fmt.Sprintf(`<w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>%s</w:t></w:r>`, s),
	})
}
