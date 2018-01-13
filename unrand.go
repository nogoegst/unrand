// unrand.go - replace crypto/rand.Reader with math/rand one.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to unrand, using the creative
// commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package unrand

import (
	"crypto/rand"
	mrand "math/rand"
)

var reader = rand.Reader

func Replace(seed int64) {
	rand.Reader = mrand.New(mrand.NewSource(seed))
}

func Recover() {
	rand.Reader = reader
}
