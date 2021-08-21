package algo

import "fmt"

func ContainStringWithRuneCharShift(s []rune, substr rune) (int, bool) {
	//sb := []rune(substr)[0]
	for i, sChar := range s {
		if sChar == substr {
			return i, true
		}
	}
	return 0, false
}

//FindMatch return ShiftPatternInSource,
//FirstIndexPatternCharMatchToSource,
//LastIndexPatternCharMatchToSource,
//err
func FindMatch(substr, pattern string) (int, int, int, error) {
	//check full tag
	shift, err := KMPSearch(substr, pattern)
	if err == nil {
		return shift, 0, len([]rune(pattern)) - 1, nil
	}
	//_____
	// pShift := 0
	// patt := []rune(pattern)
	// for {
	// 	shift, err := KMPSearch(substr, string(patt[pShift]))
	// 	if err == nil {
	// 		return shift, pShift, pShift, nil
	// 	} else {
	// 		if pShift < len(patt) {
	// 			pShift += 1
	// 		} else {
	// 			return 0, 0, 0, fmt.Errorf("pattern not found")
	// 		}

	// 	}
	// }

	//____check first element pattern in substr
	// pShift := 0
	// pShiftT := 0
	// patt := []rune(pattern)
	// Lp := len(patt)
	// for {
	// 	searchPatternID, err := KMPSearch(substr, string(patt[pShift:pShiftT]))
	// 	if err != nil {
	// 		pShift += 1
	// 		pShiftT = pShift
	// 	} else {
	// 		pShiftT += 1
	// 	}

	// 	if Lp < pShift || Lp < pShiftT {
	// 		break
	// 	}
	// }
	return 0, 0, 0, fmt.Errorf("pattern not found")
}

//FindMatch(substr, pattern string)
//(
//ShiftPatternInSource int,
//FirstIndexPatternCharMatchToSource int,
//LastIndexPatternCharMatchToSource int,
//err error
//)
// func FindMatch(substr, pattern string) (int, int, int, error) {
// 	//____for
// 	ShiftPatternInSource := 0
// 	FirstIndexPatternCharMatchToSource := 0
// 	LastIndexPatternCharMatchToSource := 0
// 	//____return
// 	p := []rune(pattern)
// 	sub := []rune(substr)
// 	shiftSub := 0     //shift index of substr
// 	patternShift := 0 //shift index of pattern char
// 	firstPShift := 0  //
// 	lastShift := 0
// 	//pLen := len(p)
// 	CounterContainsPattternChar := 0
// 	CounterPattternChar := 0
// 	for {
// 		_, contain := ContainStringWithRuneCharShift(sub[shiftSub:], p[patternShift])
// 		if contain {
// 			patternShift += 1
// 			shiftSub += 1
// 		} else {
// 			shiftSub = 0
// 			patternShift = CounterPattternChar
// 			CounterPattternChar += 1
// 		}

// 	}

// 	// 	if contain {
// 	// 		CounterContainsPattternChar += 1
// 	// 		//	fmt.Printf("sub: %v | p: %v|contain : [%d]\n", sub[shiftSub:], p[patternShift], sh)
// 	// 		if patternShift == 0 {
// 	// 			firstPShift = shiftSub + sh
// 	// 		}
// 	// 		if patternShift+1 == pLen {
// 	// 			//	fmt.Printf("patternShift == pLen\n")
// 	// 			lastShift = shiftSub + sh
// 	// 			//return firstPShift, lastShift, nil
// 	// 			break
// 	// 		}
// 	// 		shiftSub += 1
// 	// 		patternShift += 1
// 	// 	} else {
// 	// 		shiftSub += 1
// 	// 		patternShift += 0
// 	// 	}
// 	// 	if shiftSub == len(sub) {
// 	// 		//		fmt.Printf("break\n")
// 	// 		break
// 	// 	}
// 	// }
// 	//temp fix
// 	_ = firstPShift
// 	_ = lastShift
// 	//temp fix closed
// 	if CounterContainsPattternChar == 0 {
// 		return 0, 0, 0, fmt.Errorf("pattern not found in string")
// 	}
// 	return ShiftPatternInSource, FirstIndexPatternCharMatchToSource, LastIndexPatternCharMatchToSource, nil
// }

//FindMatchInArray return StartIndexElementWithPartOfPattern, StopIndexElementWithPartOfPattern, ShiftBeginningPatternInFirstElement, err
func FindMatchInArray(arr []string, pattern string) (int, int, int, error) {

	strL := ""
	for _, item := range arr {
		strL += item
	}
	shift, err := KMPSearch(strL, pattern)
	if err != nil {
		return 0, 0, 0, err
	}

	globalSize := 0
	index := 0
	for i, item := range arr {
		globalSize += len(item)
		index = i
		if shift < globalSize {
			break
		}
	}
	LenOfFirstPart := len(arr[index]) - (globalSize - shift)
	// fmt.Printf("shift: %d|globalSize: %d\n", shift, globalSize)
	// fmt.Printf("itemIndex: %d\n", index)
	// fmt.Printf("itemArr: %s\n", arr[index])
	// fmt.Printf("LenOfFirstPart: %d\n", LenOfFirstPart)
	shiftSize := globalSize
	shiftIndex := index
	//	fmt.Printf("itemIndex: %s\n", arr)
	//	ShiftZ := shiftSize + len(pattern)
	//	fmt.Printf("ShiftZ: %d\n", ShiftZ)
	if globalSize < shift+len(pattern) {
		for _, item := range arr[index+1:] {

			shiftSize += len(item)
			shiftIndex += 1
			//	fmt.Printf("rangingItem: %s| ShiftZ %d|shiftSize %d\n", item, ShiftZ, shiftSize)
			if shift+len(pattern) < shiftSize+1 {
				break
			}
		}
	}

	// fmt.Printf("shift: %d\n", shift)
	// fmt.Printf("globalSize: %d\n", globalSize)
	// fmt.Printf("index: %d\n", index)
	// fmt.Printf("ShiftIndex: %d\n", shiftIndex)

	return index, shiftIndex, LenOfFirstPart, nil
}
