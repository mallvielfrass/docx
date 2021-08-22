package algo

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkContainSimple(b *testing.B) {
	d := struct {
		Substr  string
		Pattern string
	}{
		Substr:  "KKKKKKKKKKKKKDWEDADEA45w4ww4-9,w409aumpaw5vnwam5vwam5w45myawp49v5ymw4a95owvmamw94amw4p9vw4maw94paw49amw4p9aw49pawmaw49mwa49o4haw4mo9awh49pmwa84haw94othawo48t,wmt8oaw4hrtaw8ortma,w4o8m7hwao84t7mhaw8o4thawo84tahw4t,ow4th7wao84thmawm8owharo8wash{Code}eaaaaaa",
		Pattern: "{Code}",
	}
	for i := 0; i < b.N; i++ {
		res := strings.Contains(d.Substr, d.Pattern)
		if !(res) {
			fmt.Println("error")
		}
	}

}
func BenchmarkContainKMP(b *testing.B) {
	d := struct {
		Substr  string
		Pattern string
	}{
		Substr:  "KKKKKKKKKKKKKDWEDADEA45w4ww4-9,w409aumpaw5vnwam5vwam5w45myawp49v5ymw4a95owvmamw94amw4p9vw4maw94paw49amw4p9aw49pawmaw49mwa49o4haw4mo9awh49pmwa84haw94othawo48t,wmt8oaw4hrtaw8ortma,w4o8m7hwao84t7mhaw8o4thawo84tahw4t,ow4th7wao84thmawm8owharo8wash{Code}eaaaaaa",
		Pattern: "{Code}",
	}
	for i := 0; i < b.N; i++ {
		_, err := KMPSearch(d.Substr, d.Pattern)
		if err != nil {
			fmt.Println("error")
		}
	}

}
