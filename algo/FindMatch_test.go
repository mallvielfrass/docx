package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMatch(t *testing.T) {
	diff := []struct {
		TestName                           string
		Source                             string
		ShiftPatternInSource               int
		FirstIndexPatternCharMatchToSource int
		LastIndexPatternCharMatchToSource  int
		//LastShiftOfPattern   int

		Pattern string
	}{
		{
			TestName:                           "Complimentary",
			Source:                             "{Name}",
			Pattern:                            "{Name}",
			ShiftPatternInSource:               0,
			FirstIndexPatternCharMatchToSource: 0,
			LastIndexPatternCharMatchToSource:  5,
		},
		{
			TestName:                           "Shift",
			Source:                             "Zaza{Name}3",
			Pattern:                            "{Name}",
			ShiftPatternInSource:               4,
			FirstIndexPatternCharMatchToSource: 0,
			LastIndexPatternCharMatchToSource:  5,
		},
		{
			TestName:                           "FirstPartOfSource",
			Source:                             "Zaza{Nam",
			Pattern:                            "{Name}",
			ShiftPatternInSource:               4,
			FirstIndexPatternCharMatchToSource: 0,
			LastIndexPatternCharMatchToSource:  3,
		},
		{
			TestName:                           "SecondPartOfSource",
			Source:                             "e}",
			Pattern:                            "{Name}",
			ShiftPatternInSource:               0,
			FirstIndexPatternCharMatchToSource: 4,
			LastIndexPatternCharMatchToSource:  5,
		},
	}
	for _, tt := range diff {
		t.Run(tt.TestName, func(t *testing.T) {
			ShiftPatternInSource, FirstIndexPatternCharMatchToSource, LastIndexPatternCharMatchToSource, err := FindMatch(tt.Source, tt.Pattern)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.ShiftPatternInSource, ShiftPatternInSource)
			assert.Equal(t, tt.FirstIndexPatternCharMatchToSource, FirstIndexPatternCharMatchToSource)
			assert.Equal(t, tt.LastIndexPatternCharMatchToSource, LastIndexPatternCharMatchToSource)

		})
	}
}

// func TestFindMatchBroken(t *testing.T) {
// 	diff := []struct {
// 		TestName                           string
// 		Source                             string
// 		ShiftPatternInSource               int
// 		FirstIndexPatternCharMatchToSource int
// 		LastIndexPatternCharMatchToSource  int
// 		//LastShiftOfPattern   int

// 		Pattern string
// 	}{
// 		{
// 			TestName:                           "BrokenTestFirst",
// 			Source:                             "{Name}",
// 			Pattern:                            "Z",
// 			ShiftPatternInSource:               0,
// 			FirstIndexPatternCharMatchToSource: 0,
// 			LastIndexPatternCharMatchToSource:  0,
// 		},
// 	}
// 	for _, tt := range diff {
// 		t.Run(tt.TestName, func(t *testing.T) {
// 			_, _, _, err := FindMatch(tt.Source, tt.Pattern)
// 			if err == nil {
// 				t.Error("expected error")
// 			}
// 		})
// 	}
// }
func TestContainCharWithShift(t *testing.T) {
	diff := []struct {
		TestName string
		Substr   string
		Pattern  string
		ExpBool  bool
		ExpID    int
	}{
		{
			TestName: "FirstElement",
			Substr:   "qwerty",
			Pattern:  "q",
			ExpBool:  true,
			ExpID:    0,
		},
		{
			TestName: "ThirdElement",
			Substr:   "qwerty",
			Pattern:  "e",
			ExpBool:  true,
			ExpID:    2,
		},
		{
			TestName: "BrokenTest",
			Substr:   "qwerty",
			Pattern:  "G",
			ExpBool:  false,
			ExpID:    0,
		},
	}
	for _, tt := range diff {
		t.Run(tt.TestName, func(t *testing.T) {
			id, bl := ContainStringWithRuneCharShift([]rune(tt.Substr), []rune(tt.Pattern)[0])
			assert.Equal(t, tt.ExpBool, bl)
			assert.Equal(t, tt.ExpID, id)
		})
	}
}

func TestFindMatchInArray(t *testing.T) {
	type args struct {
		arr     []string
		pattern string
	}
	tests := []struct {
		name                                string
		args                                args
		StartIndexElementWithPartOfPattern  int
		StopIndexElementWithPartOfPattern   int
		ShiftBeginningPatternInFirstElement int
	}{
		{name: "First",
			args: args{
				arr:     []string{"lalaЭЭd{c", "h", "eck}f", "rdad"},
				pattern: "{check}",
			},
			StartIndexElementWithPartOfPattern:  0,
			StopIndexElementWithPartOfPattern:   2,
			ShiftBeginningPatternInFirstElement: 9,
		},
		{name: "Second",
			args: args{
				arr:     []string{"lalaЭЭ", " ыывлоц3о2424", "!№*№?*", "{check}rdad"},
				pattern: "{check}",
			},
			StartIndexElementWithPartOfPattern:  3,
			StopIndexElementWithPartOfPattern:   3,
			ShiftBeginningPatternInFirstElement: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartIndexElementWithPartOfPattern, StopIndexElementWithPartOfPattern, ShiftBeginningPatternInFirstElement, err := FindMatchInArray(tt.args.arr, tt.args.pattern)
			if err != nil {
				t.Error(err)
			}
			if StartIndexElementWithPartOfPattern != tt.StartIndexElementWithPartOfPattern {
				t.Errorf("FindMatchInArray().StartIndex = %v, want %v", StartIndexElementWithPartOfPattern, tt.StartIndexElementWithPartOfPattern)
			}
			if StopIndexElementWithPartOfPattern != tt.StopIndexElementWithPartOfPattern {
				t.Errorf("FindMatchInArray().StopIndex = %v, want %v", StopIndexElementWithPartOfPattern, tt.StopIndexElementWithPartOfPattern)
			}
			if ShiftBeginningPatternInFirstElement != tt.ShiftBeginningPatternInFirstElement {
				t.Errorf("FindMatchInArray().ShiftPattern = %v, want %v", ShiftBeginningPatternInFirstElement, tt.ShiftBeginningPatternInFirstElement)
			}
		})
	}
}
