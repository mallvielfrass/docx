package main

import "github.com/mallvielfrass/docx"

func main() {
	// Read from docx file
	r, err := docx.ReadDocxFile("./TestDocumentTag.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	//	docx1.Replace("old_1_1", "new_1_1", -1)
	header, body, footer := docx1.ParseNode()
	bodyParts := docx1.BodyParse(body)
	// newSimpleDimpleBlock := docx1.CreateNewBlock("Simple Dimple Block")
	// bodyParts = docx1.AddBlockAtTheEnd(newSimpleDimpleBlock, bodyParts)
	// bodyParts = docx1.AddBlockAtTheEnd(newSimpleDimpleBlock, bodyParts)
	// ns := docx1.EditTextInStringBlock(newSimpleDimpleBlock, "NEw Simple test")
	// bodyParts = docx1.AddBlockAtTheBeginning(ns, bodyParts)
	// // id, err := docx1.GetFirstElementContain("Simple Dimple Block", bodyParts)
	// // if err == nil {
	// // 	bodyParts[id] = docx1.EditTextInStringBlock(bodyParts[id], "EditedBlock")
	// // }
	// //or
	// docx1.ReplaceTextInBlock("Simple Dimple Block", "EditedBlock", bodyParts)
	//fmt.Println("bodyParts", bodyParts[1])
	_ = docx1.ParseBlockToStruct(bodyParts[1])
	BlockStr := docx1.CreateStructedBlock("LALALA")
	BlockStr.Head.Bold = true
	BlockStr.Head.FontSize = 48
	BlockStr.Head.Color = docx.Cyan

	bodyParts = docx1.AddBlockAtTheEnd(docx1.BlockToString(BlockStr), bodyParts)
	BlockStr.Body = "NEw String"
	BlockStr.Head.Color = docx.Lime
	BlockStr.Head.FontSize = 64
	BlockStr.Head.Strike = true
	bodyParts = docx1.AddBlockAtTheEnd(docx1.BlockToString(BlockStr), bodyParts)
	bodyFull := docx1.BodyGlue(bodyParts)

	// docx1.ReplaceWithTag("{div id=1}{/div}", "{div}Test{/div}")

	// docx1.Replace("old_1_2", "new_1_2", -1)
	// docx1.ReplaceHeader("out with the old", "in with the new")
	// docx1.ReplaceFooter("Change This Footer", "new footer")
	docx1.GlueNodes(header, bodyFull, footer)

	docx1.WriteToFile("./new_result_1.docx")

	// docx2 := r.Editable()
	// docx2.Replace("old_2_1", "new_2_1", -1)
	// docx2.Replace("old_2_2", "new_2_2", -1)
	// docx2.WriteToFile("./new_result_2.docx")

	// Or write to ioWriter
	// docx2.Write(ioWriter io.Writer)

	r.Close()
}
