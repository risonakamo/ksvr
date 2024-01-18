// db v2
package ksvr

import (
	ksvr "ksvr/lib"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type DbV2 struct {
    // sentence table. stores sentences
    // key: `sentence:{hash}`
    // hash is sha256 hash of sentence.
    // val: the whole sentence as string
    sentence map[string]string
}

// create db v2 instance. give path. path should probably have _v2 at the end
// to denote v2
func CreateDbV2(dbpath string) *leveldb.DB {
    var db *leveldb.DB
    var err error
    db,err=leveldb.OpenFile(dbpath,nil)

    if err!=nil {
        log.Fatalln(err)
    }

    return db
}

// add sentence to db
func AddSentenceToDbV2(db *leveldb.DB,sentence string) {
    var key []byte=createSentenceKey(sentence)

    var e error=db.Put(key,[]byte(sentence),nil)

    if e!=nil {
        log.Fatalln("db put fail",e)
    }
}

// check if given sentence is in db
func SentenceInDbV2(db *leveldb.DB,sentence string) bool {
    var key []byte=createSentenceKey(sentence)

    var has bool
    has,_=db.Has(key,nil)

    return has
}

// convert sentence into key to find in sentence table
func createSentenceKey(sentence string) []byte {
    var sentenceHash []byte=ksvr.HashSentence(sentence)

    var key []byte=[]byte("sentence:")
    key=append(key,sentenceHash...)
    return key
}