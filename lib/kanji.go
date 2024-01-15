// functions for doing kanji related tasks

package ksvr

import "github.com/KEINOS/go-joyokanjis/kanjis"

// return list of all unique kanjis extracted from a string
func extractKanjis(str string) []rune {
	var res []rune

	for _,char := range str {
		if kanjis.IsJoyoKanji(char) {
			res=append(res,char)
		}
	}

	return res
}