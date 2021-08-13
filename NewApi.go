package docx

import (
	"fmt"
	"strings"

	"github.com/mallvielfrass/ooxml"
)

const Open = 4
const Self = 5

func wpParser(s string) ([]WPTokens, error) {
	nodes, err := ooxml.GetParentNodes(s)
	if err != nil {
		fmt.Println(err)
		return []WPTokens{}, err
	}
	_ = nodes
	var d []WPTokens
	for _, item := range nodes {
		//	fmc.Printfln("\t#rbt%d) #ybt%s", i+1, item.Name)
		status := Self
		switch item.TagStatus {
		case ooxml.TagComplimentary:
			status = Open
		case ooxml.TagSelfClosed:
			status = Self
		}
		d = append(d, WPTokens{
			Tag:    item.Name,
			Body:   item.Body,
			Attr:   item.Args,
			Status: status,
		})
	}
	return d, nil
}
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
			body, err := wpParser(item.Body)
			if err != nil {
				fmt.Println(err)
				return Document{}, err
			}
			doc.WP = append(doc.WP, WP{
				Tag:  item.Name,
				Body: body,
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
func templateBlock(s string) WP {
	r := WP{

		Tag: "w:p",
		Body: []WPTokens{
			{
				Tag:    "w:pPr",
				Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
				Status: Open,
			},
			{
				Tag:    "w:r",
				Body:   fmt.Sprintf("<w:rPr></w:rPr><w:t>%s</w:t>", s),
				Status: Open,
			},
		},
	}
	return r
}
func (d *Document) AddNewBlock(s string) {
	d.WP = append(d.WP, WP{
		Tag: "w:p",
		//Body: fmt.Sprintf(`<w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr></w:rPr><w:t>%s</w:t></w:r>`, s),
		Body: []WPTokens{
			{
				Tag:    "w:pPr",
				Body:   `<w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr>`,
				Status: Open,
			},
			{
				Tag:    "w:r",
				Body:   fmt.Sprintf("<w:rPr></w:rPr><w:t>%s</w:t>", s),
				Status: Open,
			},
		},
	})
}
func (d *Document) GetBlockIDByTag(s string) (int, error) {
	for i, item := range d.WP {
		for _, itemD := range item.Body {
			// if strings.Contains(itemD.Body s) {

			// 		}
			if itemD.Tag == "w:r" {
				if strings.Contains(itemD.Body, s) {
					return i, nil
				}
			}
		}

	}
	return 0, fmt.Errorf("block with tag not found")
}
func (d *Document) EditBlockByID(id int) {

}
func (d *Document) ReplaceTagString(oldTag, newString string) error {
	id, err := d.GetBlockIDByTag(oldTag)
	if err != nil {
		//t.Error(err)
		return err
	}
	for i, itemD := range d.WP[id].Body {
		// if strings.Contains(itemD.Body s) {

		// 		}
		if itemD.Tag == "w:r" {
			if strings.Contains(itemD.Body, oldTag) {
				d.WP[id].Body[i].Body = strings.Replace(itemD.Body, oldTag, newString, -1)
				return nil
			}
		}
	}
	return fmt.Errorf("tag not found")
}
func (d *Document) GetBlockByID(id int) WP {
	return d.WP[id]
}

func (d *Docx) Compile(path string, doc Document) error {

	head, _, footer := d.ParseNode()
	//	fmt.Printf("head:[%s]\n", head)
	//	fmt.Printf("footer :[%s]\n", footer)
	body := "<w:body>" + doc.BodyToString() + "</w:body>"
	xml := head + body + footer
	d.content = xml
	err := d.WriteToFile(path)
	return err
}
func (d *Document) BodyToString() string {
	var body string
	//fmc.Printfln("#rbtItem: #gbt[%v]", d.WP)
	for _, item := range d.WP {
		//	fmc.Printfln("#bbtItemWP: #gbt[%v]", item)
		body += wpTokenToString(item)
	}
	body += "<" + d.SectPr.Tag + ">" + d.SectPr.Body + "</" + d.SectPr.Tag + ">"
	return body
}
func wpTokenToString(item WP) string {
	//item.Tag
	var body string
	for _, it := range item.Body {
		//fmc.Printf("#ybtit: #gbt[%v] status[%d]\n", it.Tag, it.Status)
		var attr string
		if 0 < len(it.Attr) && string(it.Attr[0]) != " " {
			attr = " " + it.Attr
		} else {
			attr = it.Attr
		}
		if it.Status == Open {
			body += "<" + it.Tag + attr + ">" + it.Body + "</" + it.Tag + ">"
		}
		if it.Status == Self {
			body += "<" + it.Tag + attr + "/>"
		}

	}
	return "<w:p>" + body + "</w:p>"
}
func Screening(s string) string {
	var retString string
	for _, item := range s {
		//fmc.Printfln("char: [%s]", string(item))
		it := string(item)
		switch it {
		case "<":
			retString += "&lt;"
		case ">":
			retString += "&gt;"
		case "&":
			retString += "&amp;"
		case "'":
			retString += "&apos;"
		default:
			retString += it
		}
	}
	return retString
}
