package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	var store=cache.New(
        1*time.Minute,
        5*time.Minute,
    )

    store.Set("thing",true,cache.DefaultExpiration)

    var res0 interface{}
    var res bool
    var found bool
    res0,found=store.Get("thing")
    res=res0.(bool)

    res2,found:=store.Get("huh")

    if found {
        fmt.Println(res)
    }

    fmt.Println(res2)
}