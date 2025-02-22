package docx

import (
	"fmt"
	"strings"
)

type MarkerParams struct {
}

func returnTemplateMarkerWP() WP {
	return WP{
		Tag: "w:p",
		Body: []WPTokens{
			{
				Tag: "w:pPr",
				Body: AtomicWPTokensToString(WPTokens{
					Tag:    "w:pStyle",
					Attr:   ` w:val="Normal"`,
					Body:   "",
					Status: Self,
				}) +
					AtomicWPTokensToString(WPTokens{
						Tag:    "w:numPr",
						Attr:   ``,
						Body:   `<w:ilvl w:val="0"/><w:numId w:val="1"/>`,
						Status: Open,
					}) +
					AtomicWPTokensToString(WPTokens{
						Tag:    "w:bidi",
						Attr:   ` w:val="0"`,
						Body:   "",
						Status: Self,
					}) +
					AtomicWPTokensToString(WPTokens{
						Tag:    "w:jc",
						Attr:   ` w:val="left"`,
						Body:   "",
						Status: Self,
					}) +
					AtomicWPTokensToString(WPTokens{
						Tag:    "w:rPr",
						Attr:   ``,
						Body:   "",
						Status: Open,
					}),
				Attr:   ``,
				Status: Open,
			},
			{
				Tag:    "w:r",
				Body:   "<w:rPr></w:rPr><w:t>%s</w:t>",
				Attr:   ``,
				Status: Open,
			},
		},
	}
}
func (d *Document) CreateMarkedStringList(mp MarkerParams, letter ...string) []WP {
	//	fmc.Printfln("letter: %v", letter)
	// templateMarkerWP := WP{
	// 	Tag: "w:p",
	// 	Body: []WPTokens{
	// 		{
	// 			Tag: "w:pPr",
	// 			Body: AtomicWPTokensToString(WPTokens{
	// 				Tag:    "w:pStyle",
	// 				Attr:   ` w:val="Normal"`,
	// 				Body:   "",
	// 				Status: Self,
	// 			}) +
	// 				AtomicWPTokensToString(WPTokens{
	// 					Tag:    "w:numPr",
	// 					Attr:   ``,
	// 					Body:   `<w:ilvl w:val="0"/><w:numId w:val="1"/>`,
	// 					Status: Open,
	// 				}) +
	// 				AtomicWPTokensToString(WPTokens{
	// 					Tag:    "w:bidi",
	// 					Attr:   ` w:val="0"`,
	// 					Body:   "",
	// 					Status: Self,
	// 				}) +
	// 				AtomicWPTokensToString(WPTokens{
	// 					Tag:    "w:jc",
	// 					Attr:   ` w:val="left"`,
	// 					Body:   "",
	// 					Status: Self,
	// 				}) +
	// 				AtomicWPTokensToString(WPTokens{
	// 					Tag:    "w:rPr",
	// 					Attr:   ``,
	// 					Body:   "",
	// 					Status: Open,
	// 				}),
	// 			Attr:   ``,
	// 			Status: Open,
	// 		},
	// 		{
	// 			Tag:    "w:r",
	// 			Body:   "<w:rPr></w:rPr><w:t>%s</w:t>",
	// 			Attr:   ``,
	// 			Status: Open,
	// 		},
	// 	},
	// } //wpTokenToString
	if mp != (MarkerParams{}) {
		//With Params
		return []WP{returnTemplateMarkerWP()}
	}
	var wpArray []WP
	for _, item := range letter {
		//fmc.Printfln("ranger: %s", item)
		if strings.Contains(item, "\n") {
			arr := strings.Split(item, "\n")
			for _, arrItem := range arr {
				tempBlock := returnTemplateMarkerWP()
				tempBlock.Body[1].Body = fmt.Sprintf(tempBlock.Body[1].Body, Screening(arrItem))
				wpArray = append(wpArray, tempBlock)
			}
		} else {
			//tempBlock := templateMarkerWP
			tempBlock := returnTemplateMarkerWP()

			//	fmc.Printfln("TemplateStruct: %s", returnTemplateMarkerWP().Body[1].Body)
			//	fmc.Printfln("%d) Item: %s", i, Screening(item))
			tempBlock.Body[1].Body = fmt.Sprintf(tempBlock.Body[1].Body, Screening(item))
			//	fmc.Printfln("tempBlock.Body[i].Body : %s", tempBlock.Body[1].Body)
			wpArray = append(wpArray, tempBlock)
		}
	}
	return wpArray
}
func AtomicWPTokensToString(token WPTokens) string {
	var attr string
	var body string
	if 0 < len(token.Attr) && string(token.Attr[0]) != " " {
		attr = " " + token.Attr
	} else {
		attr = token.Attr
	}
	if token.Status == Open {
		body += "<" + token.Tag + attr + ">" + token.Body + "</" + token.Tag + ">"
	}
	if token.Status == Self {
		body += "<" + token.Tag + attr + "/>"
	}
	return body
}
func GetTextFromXML(src string) (string, error) {
	res, err := wpParser(src)
	if err != nil {
		return "", err
	}
	var text string
	for _, item := range res {
		//		fmt.Printf("res[]: %s\n", item.Tag)
		if item.Tag == "w:r" {
			res, err := wpParser(item.Body)
			if err != nil {
				return "", err
			}
			for _, wtTag := range res {
				if wtTag.Tag == "w:t" {
					//fmt.Printf("w:t: %s\n", wtTag.Body)
					text += wtTag.Body
				}
			}
		}
	}
	return text, nil
}
