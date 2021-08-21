package algo

import (
	"testing"
)

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
