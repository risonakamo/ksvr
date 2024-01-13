package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
    var cwd string
    cwd,_=os.Getwd()
    var dbdir string=filepath.Join(cwd,"db")
    fmt.Println("db path",dbdir)

    var db *leveldb.DB
    var err error
    db,err=leveldb.OpenFile(dbdir,nil)

    if err!=nil {
        fmt.Println("db open error")
        log.Fatalln(err)
    }

    db.Put(
        []byte("huh"),
        []byte("thing123123"),
        nil,
    )

    val,err:=db.Get([]byte("huh2"),nil)

    if err!=nil {
        fmt.Println("db get error")
        log.Fatalln(err)
    }

    fmt.Println(string(val))

    db.Close()
}