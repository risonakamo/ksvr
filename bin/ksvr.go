package main

import (
	"fmt"
	ksvr "ksvr/lib"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/atotto/clipboard"
	"github.com/patrickmn/go-cache"
)

const sentenceProcessorsCount int=5

func main() {
    // --- config vars ---
    var cwd string
    cwd,_=os.Getwd()
    var dbPath string=filepath.Join(cwd,"db")


    // --- main vars ---
    var sentenceCache *cache.Cache=ksvr.NewSentenceCache()
    var db *ksvr.KsvrDb=ksvr.NewKsvrDb(dbPath)

    var sentences chan string=make(chan string,10)
    var wg sync.WaitGroup

    // spawn clipboard worker
    clipboardWorker(
        5,
        sentences,
        &wg,
    )

    // spawn sentence processor workers
    for i:=0;i<sentenceProcessorsCount;i++ {
        wg.Add(1)
        go sentenceProcessorWorker(
            sentenceCache,
            db,

            sentences,

            &wg,
        )
    }

    fmt.Println("hm")

    wg.Wait()
}

// constantly pulls from clipboard at an interval and pushes into
// sentence processing channel. this function does not block; it spawns
// another func that never ends.
//
// adds to waitgroup, but never finishes
func clipboardWorker(
    watchInterval float32, // seconds

    sentences chan<- string,

    wg *sync.WaitGroup,
) {
    var ticker *time.Ticker=time.NewTicker(time.Duration(5)*time.Second)

    var lastSentence string=""

    wg.Add(1)
    go func() {
        for {
            <-ticker.C
            fmt.Println("hello")

            var sentence string
            var e error
            sentence,e=clipboard.ReadAll()

            if e!=nil {
                fmt.Println("error while doing clipboard read")
                continue
            }

            // last sentence check. don't handle if it's the same as the thing we
            // just handled
            if sentence==lastSentence {
                fmt.Println("skipping, same as last")
                continue
            }

            lastSentence=sentence

            fmt.Println("preparing to process:",sentence)
            sentences<-sentence
        }
    }()
}

// processes sentences. recvs sentences from the sentence channel.
// runs sentence through checks to see if it should be added to db.
// if not, then does nothing. if valid for db, adds to the db.
func sentenceProcessorWorker(
    sentenceCache *cache.Cache,
    db *ksvr.KsvrDb,

    sentences <-chan string,

    wg *sync.WaitGroup,
) {
    var sentence string
    for sentence = range sentences {
        // mem cache check. don't handle anything that was already seen in the
        // last few minutes
        if ksvr.SentenceInCache(sentenceCache,sentence) {
            fmt.Println("in mem cache")
            continue
        }

        ksvr.AddSentenceToCache(sentenceCache,sentence)



        // db check. don't add anything to the db that's already in there
        var inDb bool=db.SentenceInDb(sentence)

        if inDb {
            fmt.Println("already in db")
            continue
        }


        // kanji check. if the sentence has no kanjis, dont add it
        var kanjis []rune=ksvr.ExtractKanjis(sentence)

        if len(kanjis)==0 {
            fmt.Println("no kanjis")
            continue
        }



        // if got to this point, ready to add it to db. add it to db.
        db.AddSentence(sentence)

        fmt.Println("added:",sentence)
    }

    wg.Done()
}