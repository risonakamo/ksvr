// functions implementing ksvr last sentence mem cache

package ksvr

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// create ksvr sentence cache
//
// cache interface: map[string]bool
// key: the full sentence text
// val: always true bool. not used.
//
// cache is meant to store around the last minute of sentences
func NewSentenceCache() *cache.Cache {
    return cache.New(
        1*time.Minute,
        5*time.Minute,
    )
}

// add sentence to the cache
func AddSentence(c *cache.Cache,sentence string) {
    c.Add(
        sentence,
        true,
        cache.DefaultExpiration,
    )
}

// check if the sentence is in the cache
func SentenceInCache(c *cache.Cache,sentence string) bool {
    var found bool
    _,found=c.Get(sentence)
    return found
}