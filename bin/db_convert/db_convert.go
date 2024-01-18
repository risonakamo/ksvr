package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
    var cwd string
    cwd,_=os.Getwd()
    var dbPath string=filepath.Join(cwd,"../ksvr/db")

	var db *leveldb.DB
	db,_=leveldb.OpenFile(dbPath,nil)

    var iter iterator.Iterator=db.NewIterator(
        util.BytesPrefix([]byte("sentence:")),
        nil,
    )

    for iter.Next() {
        // var sentenceInfo ksvr.SentenceInfo
        // json.Unmarshal(iter.Value(),&sentenceInfo)

        // spew.Dump(sentenceInfo)
        fmt.Println(string(iter.Value()))
    }
}