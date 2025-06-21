package idgen

import (
	"crypto/sha1"
	"encoding/hex"
	"iter"
)

const IDLength = 8

func Generate(title string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for {
			h := sha1.Sum([]byte(title))
			str := hex.EncodeToString(h[:])
			str = str[:IDLength]
			if !yield(str) {
				return
			}

			title += "_"
		}
	}
}
