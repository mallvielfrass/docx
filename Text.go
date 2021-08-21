package docx

import (
	"fmt"
	"strings"
)

//______________________________________________(wpID,wr1ID,wtID, wrXID)
func (d *Document) GetBlockIDByTagNew(s string) (int, int, int, int, error) {
	for i, WPItem := range d.WP {
		var text string
		//	d := 0
		for z, token := range WPItem.Body {
			if token.Tag == "w:r" {
				//	fmt.Printf("i %d z %d\n", i, z)
				res, err := wpParser(token.Body)
				if err != nil {
					return 0, 0, 0, 0, err
				}
				for n, wtTag := range res {
					if wtTag.Tag == "w:t" {
						//fmt.Printf("w:t: %s\n", wtTag.Body)
						if strings.Contains(wtTag.Body, s) {
							return i, z, z, n, err
						}
						text += wtTag.Body
					}
				}
			}
			//d = z
		}
		if strings.Contains(text, s) {
			return i, 0, len(WPItem.Body), 0, nil
		}
	}
	return 0, 0, 0, 0, fmt.Errorf("tag not found")
}
func (d *Document) ReplaceTextByIDWP(tag, dest string) error {
	wpID, wr1ID, wrXID, wtID, err := d.GetBlockIDByTagNew(tag)
	if err != nil {
		return err
	}
	if wr1ID == wrXID {
		wr := d.WP[wpID].Body[wr1ID]
		res, err := wpParser(wr.Body)
		if err != nil {
			return err
		}
		wt := res[wtID]
		//fmt.Println("wt: ", wt.Body)
		wt.Body = strings.Replace(wt.Body, tag, Screening(dest), 1)

		res[wtID] = wt
		var wrstrBody string
		for _, item := range res {
			wrstrBody += AtomicWPTokensToString(item)
		}
		d.WP[wpID].Body[wr1ID].Body = wrstrBody
		return nil
		//	wrstrBody :=
	}
	// for _, item := range d.WP[wpID].Body {

	// }
	//tgi := 0
	//tgn := 0
	wbody := d.WP[wpID].Body
	for i := 0; i < len(wbody); i++ {
		item := wbody[i]
		fmt.Println(item.Tag)
		if item.Tag == "w:r" {
			res, err := wpParser(item.Body)
			if err != nil {
				return err
			}
			for _, wtTag := range res {
				if wtTag.Tag == "w:t" {

					//fmt.Printf("w:t: %s\n", wtTag.Body)
					// if

					// if strings.Contains(wtTag.Body, tag[tgi:tgn]) {
					// 	// 	return   err
					// }
					// text += wtTag.Body
				}
			}
		}
	}
	return fmt.Errorf("replace error")
}
