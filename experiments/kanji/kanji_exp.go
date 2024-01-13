package main

import (
	"fmt"
	"strings"

	"github.com/KEINOS/go-joyokanjis/kanjis"
)

func main() {
	var expStrings []string=[]string{
		"本当のこと、知ってるから。誤魔化しても、分かってるから。",
		"そんな一生懸命なイブだから、わたしは放っておけなくて、好きになれたんだと思う。",
		"　だからわたしも、イブのために何かしてあげなきゃって、一生懸命でいてあげなきゃって……そう思えるのよ、イブ……。",
		"　大体、イブってば……何してるのかしら……。",
		"伊吹「数日前から、例のパーティーの準備が進んでいまして……」",
		"真衣「でもさ、今の話を総合すると……！」",
		"真優「先輩なら、それでも特攻しそうだけど」",
		"　……や、やばいです。",
		"　傍にいてほしいの、離れたくないのっ！",
		"アリス「イブは、ホント、ばかすぎなのよ……ふふ……」",
		"伊吹「あうぅあぁ～……ドキドキしましたぁ～……」",
		"伊吹「お嬢様の貞操は、わたしのものですっ！！！」		",
	}

	fmt.Println(kanjis.IsJoyoKanji('漢'))

	res:=extractKanjis(expStrings[0])
	fmt.Println(runeArrayToStrArray(res))

	res=extractKanjis(expStrings[1])
	fmt.Println(runeArrayToStrArray(res))

	res=extractKanjis(expStrings[2])
	fmt.Println(runeArrayToStrArray(res))

	fmt.Println(strings.Trim(expStrings[2],"\u3000 "))
}

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

// convert array of runes into array of strings, for printing purposes
func runeArrayToStrArray(runes []rune) []string {
	var res []string

	for i := range runes {
		res=append(res,string(runes[i]))
	}

	return res
}