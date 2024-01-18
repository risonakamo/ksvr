// ksvr kv database implementation
// see db-design.md for specs

package ksvr

import (
	"encoding/json"
	ksvr "ksvr/lib"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// db wrapper
type KsvrDb struct {
    db *leveldb.DB
}

type DbV1 struct {
	// sentence table.
	// key: `sentence:{hash}`
	// hash is sha256 hash of sentence.
	sentence map[string]SentenceInfo
}

// item in sentence table
type SentenceInfo struct {
	sentence string
	kanjis   []rune
}

// create ksvr db
func NewKsvrDb(dbpath string) *KsvrDb {
    var db *leveldb.DB
    var err error
    db,err=leveldb.OpenFile(dbpath,nil)

    if err!=nil {
        log.Fatalln(err)
    }

    return &KsvrDb {
        db:db,
    }
}

// add sentence to sentence table
func (self *KsvrDb) AddSentence(sentence string) {
    var key []byte=createSentenceKey(sentence)

    var sentenceItem SentenceInfo=SentenceInfo{
        sentence:sentence,
        kanjis:ksvr.ExtractKanjis(sentence),
    }

    var jsondata []byte
    var e error
    jsondata,e=json.Marshal(sentenceItem)

    if e!=nil {
        log.Fatalln("json marshal fail",e)
    }

    e=self.db.Put(key,jsondata,nil)

    if e!=nil {
        log.Fatalln("db put fail",e)
    }
}

// check if given sentence is in db
func (self *KsvrDb) SentenceInDb(sentence string) bool {
    var key []byte=createSentenceKey(sentence)

    var has bool
    has,_=self.db.Has(key,nil)

    return has
}