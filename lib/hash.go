// functions related to sentence hashing

package ksvr

import (
	"crypto/sha256"
	"hash"
)

// hash the sentence into bytes
func HashSentence(sentence string) []byte {
    var hasher hash.Hash=sha256.New()

    hasher.Write([]byte(sentence))
    return hasher.Sum(nil)
}