// ksvr kv database implementation
// see db-design.md for specs

package ksvr

import (
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
func (self *KsvrDb) AddSentence(
    sentenceHash string,
    sentence SentenceInfo,
) {

}