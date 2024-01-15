package main

import (
	ksvr "ksvr/lib"
	"os"
	"path/filepath"

	"github.com/patrickmn/go-cache"
)

func main() {
    // --- config vars ---
    var cwd string
    cwd,_=os.Getwd()
    var dbPath string=filepath.Join(cwd,"db")



    // --- test vars ---
    var sentence string="本当のこと、知ってるから。誤魔化しても、分かってるから。"


    // --- main vars ---
    var lastSentence string=""
    var sentenceCache *cache.Cache=ksvr.NewSentenceCache()


    // last sentence check
    if sentence==lastSentence {
        return
    }

    lastSentence=sentence


    // mem cache check
    if ksvr.SentenceInCache(sentenceCache,sentence) {
        return
    }

    ksvr.AddSentence(sentenceCache,sentence)



    // db check
    var sentenceHash []byte=ksvr.HashSentence(sentence)

}