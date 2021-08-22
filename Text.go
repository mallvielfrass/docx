package docx

import (
	"fmt"
	"strings"

	"github.com/mallvielfrass/docx/algo"
)

func (d *Document) ReplaceTextByTag(pattern string) error {
	id, err := d.GetBlockIDByTag(pattern)
	if err != nil {
		return err
	}
	_ = id
	return fmt.Errorf("replace failed")
}
func (d *Document) GetBlockIDByTag(tag string) (int, error) {
	for i, WPItem := range d.WP { //итерация всех параграфов
		var body string
		for _, token := range WPItem.Body {
			if token.Tag == "w:r" {
				//	fmt.Printf("i %d z %d\n", i, z)
				res, err := wpParser(token.Body)
				if err != nil {
					return 0, err
				}
				for _, wtTag := range res {
					if wtTag.Tag == "w:t" {
						body += wtTag.Body
					}
				}
			}
		}
		//поиск, есть ли в параграфе паттерн
		if strings.Contains(body, tag) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("tag not found")
}

//RebuildBlocks (pattern string, source []string) (expectedArray []string, blockIDWithTag int, err)
func RebuildBlocks(pattern string, source []string) ([]string, int, error) {
	start, stop, shift, err := algo.FindMatchInArray(source, pattern)
	if err != nil {
		return []string{}, 0, err
	}
	if start == stop { //если паттерн находится только в одном элементе
		return source, start, nil
	}
	//______________________________
	//fmt.Printf("start: %d| stop: %d| shift: %d\n", start, stop, shift)
	//variables block_______
	var newSource []string
	var firstElem string
	patternLen := len(pattern)
	patternLenCounter := 0
	//closed[variables block]
	for i, item := range source {

		if start <= i && i <= stop {
			//блок если в item находится часть pattern
			if stop == i {
				//работа с последним блоком, в котором есть паттерн
				//	fmt.Printf("patternLen:[%d]| patternLenCounter: [%d]\n", patternLen, patternLenCounter)
				shift := patternLen - patternLenCounter
				firstElem += item[:shift]
				newSource = append(newSource, item[shift:])
				continue
			}
			//fmt.Printf("start[%d] item[%v]\n", i, item)

			if i == start {
				//работа с первым блоком, в котором есть паттерн
				if shift == 0 {
					//	fmt.Printf("shift : [%s|%d]\n", item, len(item))
					firstElem += item
					patternLenCounter += len(item)
					newSource = append(newSource, "")
				} else {
					//fmt.Printf("shift before: [%s|%d], after: [%s|%d]\n", item[:start], len(item[:start]), item[start:], len(item[start:]))
					firstElem += item[start:]
					patternLenCounter += len(item[start:])
					newSource = append(newSource, item[:start])
				}
				continue
				//closed[работа с первым блоком, в котором есть паттерн]
			}
			//____работа с средними блоками, в которых есть паттерн
			firstElem += item
			patternLenCounter += len(item)
			newSource = append(newSource, "")
			//____closed[работа с средними блоками, в которых есть паттерн]
			//closed[блок если в item находится часть pattern]
		} else {
			newSource = append(newSource, item) //сохранение остальных беспаттерновых блоков без изменений
		}
	}
	newSource[start] += firstElem
	//	fmt.Printf("newSource: [%+v] patternLen: [%d], plCounter: [%d]\n", newSource, patternLen, patternLenCounter)
	return newSource, start, nil
}
func ExtractWPToArrayTextString(wp WP) ([]string, error) {
	return []string{}, fmt.Errorf("extract failed")
}
