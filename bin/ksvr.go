package main

import (
	"fmt"
	ksvr "ksvr/lib"
	"log"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/patrickmn/go-cache"
)

// --- test vars ---
// const sentence string="本当のこと、知ってるから。誤魔化しても、分かってるから。"

func main() {
    var sentence string
    var e error
    sentence,e=clipboard.ReadAll()

    if e!=nil {
        log.Fatalln("clipboard err",e)
    }

    fmt.Println(sentence)

    // --- config vars ---
    var cwd string
    cwd,_=os.Getwd()
    var dbPath string=filepath.Join(cwd,"db")


    // --- main vars ---
    var lastSentence string=""
    var sentenceCache *cache.Cache=ksvr.NewSentenceCache()
    var db *ksvr.KsvrDb=ksvr.NewKsvrDb(dbPath)


    // last sentence check. don't handle if it's the same as the thing we
    // just handled
    if sentence==lastSentence {
        return
    }

    lastSentence=sentence


    // mem cache check. don't handle anything that was already seen in the
    // last few minutes
    if ksvr.SentenceInCache(sentenceCache,sentence) {
        return
    }

    ksvr.AddSentence(sentenceCache,sentence)



    // db check. don't add anything to the db that's already in there
    var inDb bool=db.SentenceInDb(sentence)

    if inDb {
        return
    }


    // kanji check. if the sentence has no kanjis, dont add it
    var kanjis []rune=ksvr.ExtractKanjis(sentence)

    if len(kanjis)==0 {
        return
    }



    // if got to this point, ready to add it to db. add it to db.
    db.AddSentence(sentence)

    fmt.Println("added:",sentence)
}