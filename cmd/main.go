package main

import (
	"fmt"
	"time"

	"github.com/mallvielfrass/docx"
)

// func main() {
// 	// Read from docx file
// 	r, err := docx.ReadDocxFile("./TestDocumentTag.docx")
// 	// Or read from memory
// 	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	docx1 := r.Editable()
// 	// Replace like https://golang.org/pkg/strings/#Replace
// 	//	docx1.Replace("old_1_1", "new_1_1", -1)
// 	header, body, footer := docx1.ParseNode()
// 	bodyParts := docx1.BodyParse(body)
// 	// newSimpleDimpleBlock := docx1.CreateNewBlock("Simple Dimple Block")
// 	// bodyParts = docx1.AddBlockAtTheEnd(newSimpleDimpleBlock, bodyParts)
// 	// bodyParts = docx1.AddBlockAtTheEnd(newSimpleDimpleBlock, bodyParts)
// 	// ns := docx1.EditTextInStringBlock(newSimpleDimpleBlock, "NEw Simple test")
// 	// bodyParts = docx1.AddBlockAtTheBeginning(ns, bodyParts)
// 	// // id, err := docx1.GetFirstElementContain("Simple Dimple Block", bodyParts)
// 	// // if err == nil {
// 	// // 	bodyParts[id] = docx1.EditTextInStringBlock(bodyParts[id], "EditedBlock")
// 	// // }
// 	// //or
// 	// docx1.ReplaceTextInBlock("Simple Dimple Block", "EditedBlock", bodyParts)
// 	//fmt.Println("bodyParts", bodyParts[1])
// 	_ = docx1.ParseBlockToStruct(bodyParts[1])
// 	BlockStr := docx1.CreateStructedBlock("LALALA")
// 	BlockStr.Head.Bold = true
// 	BlockStr.Head.FontSize = 48
// 	BlockStr.Head.Color = docx.Cyan

// 	bodyParts = docx1.AddBlockAtTheEnd(docx1.BlockToString(BlockStr), bodyParts)
// 	BlockStr.Body = "NEw String"
// 	BlockStr.Head.Color = docx.Lime
// 	BlockStr.Head.FontSize = 64
// 	BlockStr.Head.Strike = true
// 	bodyParts = docx1.AddBlockAtTheEnd(docx1.BlockToString(BlockStr), bodyParts)
// 	bodyFull := docx1.BodyGlue(bodyParts)

// 	// docx1.ReplaceWithTag("{div id=1}{/div}", "{div}Test{/div}")

// 	// docx1.Replace("old_1_2", "new_1_2", -1)
// 	// docx1.ReplaceHeader("out with the old", "in with the new")
// 	// docx1.ReplaceFooter("Change This Footer", "new footer")
// 	docx1.GlueNodes(header, bodyFull, footer)

// 	docx1.WriteToFile("./new_result_1.docx")

// 	// docx2 := r.Editable()
// 	// docx2.Replace("old_2_1", "new_2_1", -1)
// 	// docx2.Replace("old_2_2", "new_2_2", -1)
// 	// docx2.WriteToFile("./new_result_2.docx")

// 	// Or write to ioWriter
// 	// docx2.Write(ioWriter io.Writer)

// 	r.Close()
// }
func main() {
	for i := 0; i < 100; i++ {
		//fmt.Println("start")
		r, err := docx.ReadDocxFile("TestDocumentTag2.docx")
		if err != nil {
			fmt.Println(err)
			return
		}
		docx1 := r.Editable()
		doc, err := docx1.Parser()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = doc.ReplaceTagString("{Name}", "KVA")
		if err != nil {
			fmt.Println(err)
			return
		}
		//	fmt.Printf("doc:[%s]\n", docx1.content)
		//	fmt.Println(doc.WP[0].Body)
		//	exp := `<w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="00000A"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Name: KVA</w:t>`
		//assert.Equal(t, exp, doc.WP[0].Body[len(doc.WP[0].Body)-1].Body)
		//expDoc := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		//<w:document xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" mc:Ignorable="w14 wp14"><w:body><w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="00000A"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Name: KVA</w:t></w:r></w:p><w:p><w:pPr><w:pStyle w:val="Normal"/><w:rPr><w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs="" w:asciiTheme="minorHAnsi" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/><w:color w:val="00000A"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr></w:pPr><w:r><w:rPr><w:rFonts w:eastAsia="Calibri" w:cs="" w:cstheme="minorBidi" w:eastAsiaTheme="minorHAnsi"/><w:color w:val="00000A"/><w:kern w:val="0"/><w:sz w:val="24"/><w:szCs w:val="24"/><w:lang w:val="en-US" w:eastAsia="en-US" w:bidi="ar-SA"/></w:rPr><w:t>Log: {Log} Date {Date}</w:t></w:r></w:p><w:sectPr><w:headerReference w:type="default" r:id="rId2"/><w:footerReference w:type="default" r:id="rId3"/><w:type w:val="nextPage"/><w:pgSz w:w="12240" w:h="15840"/><w:pgMar w:left="1440" w:right="1440" w:header="720" w:top="1440" w:footer="720" w:bottom="1440" w:gutter="0"/><w:pgNumType w:fmt="decimal"/><w:formProt w:val="false"/><w:textDirection w:val="lrTb"/><w:docGrid w:type="default" w:linePitch="360" w:charSpace="0"/></w:sectPr></w:body></w:document>`

		err = docx1.Compile(fmt.Sprintf("./doc%d.docx", i), doc)
		if err != nil {
			fmt.Println(err)
			return
		}

		//assert.Equal(t, expDoc, docx1.content)
		err = doc.ReplaceTagString("{Log}", "Test")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = doc.ReplaceTagString("{Date}", "Daaaaata")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = docx1.Compile(fmt.Sprintf("./docX%d.docx", i), doc)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(5 * time.Second)
	}
}
