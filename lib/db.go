// ksvr kv database implementation
// see db-design.md for specs

package ksvr

import (
	"encoding/json"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// db wrapper
type KsvrDb struct {
    db *leveldb.DB
}

// item in sentence table
type SentenceInfo struct {
    sentence string
    kanjis []rune
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
        kanjis:extractKanjis(sentence),
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

// convert sentence into key to find in sentence table
func createSentenceKey(sentence string) []byte {
    var sentenceHash []byte=HashSentence(sentence)

    var key []byte=[]byte("sentence:")
    key=append(key,sentenceHash...)
    return key
}