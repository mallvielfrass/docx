package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMPSearch(t *testing.T) {
	diff := []struct {
		TestName string
		Substr   string
		Pattern  string
		Shift    int
	}{
		{
			TestName: "FirstElement",
			Substr:   "qwerty",
			Pattern:  "q",
			Shift:    0,
		},
		{
			TestName: "ThirdElement",
			Substr:   "qwerty",
			Pattern:  "e",
			Shift:    2,
		},
		{
			TestName: "ArrElements",
			Substr:   "qwerty",
			Pattern:  "ert",
			Shift:    2,
		},
		// {
		// 	TestName: "BrokenTest",
		// 	Substr:   "qwerty",
		// 	Pattern:  "G",
		// },
	}
	for _, tt := range diff {
		t.Run(tt.TestName, func(t *testing.T) {
			shift, err := KMPSearch(tt.Substr, tt.Pattern)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tt.Shift, shift)
			//assert.Equal(t, tt.ExpID, id)
		})
	}
}
func TestKMPSearchBroken(t *testing.T) {
	diff := []struct {
		TestName string
		Substr   string
		Pattern  string
		Shift    int
	}{
		{
			TestName: "T1",
			Substr:   "qwerty",
			Pattern:  "qqwert",
			Shift:    0,
		},
		{
			TestName: "T2",
			Substr:   "qwerty",
			Pattern:  "et",
			Shift:    0,
		},
		{
			TestName: "T3",
			Substr:   "qwerty",
			Pattern:  "Ert",
			Shift:    0,
		},
		{
			TestName: "T4",
			Substr:   "qwerty",
			Pattern:  "G",
			Shift:    0,
		},
	}
	for _, tt := range diff {
		t.Run(tt.TestName, func(t *testing.T) {
			shift, err := KMPSearch(tt.Substr, tt.Pattern)
			if err == nil {
				t.Error(err)
			}
			assert.Equal(t, tt.Shift, shift)
			//assert.Equal(t, tt.ExpID, id)
		})
	}
}
