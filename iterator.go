package pebigo

// #cgo LDFLAGS: -lpebblesdb
// #include "pebblesdb/c.h"
import "C"

import (
	"unsafe"
)

type IteratorError String

func (e IteratorError) Error() string {
	return string(e)
}

// Iterator is a read-only iterator through a LevelDB database. It provides a
// way to seek to specific keys and iterate through the keyspace from that
// point, as well as access the values of those keys.
//
// Care must be taken when using an Iterator. If the method Valid returns
// false, calls to Key, Value, Next, and Prev will result in panics. However,
// Seek, SeekToFirst, SeekToLast, GetError, Valid, and Close will still be
// safe to call.
//
// GetError will only return an error in the event of a LevelDB error. It will
// return a nil on iterators that are simply invalid. Given that behavior,
// GetError is not a replacement for a Valid.
//
// A typical use looks like:
//
// 	db := levigo.Open(...)
//
// 	it := db.NewIterator(readOpts)
// 	defer it.Close()
// 	for it.Seek(mykey); it.Valid(); it.Next() {
// 		useKeyAndValue(it.Key(), it.Value())
// 	}
// 	if err := it.GetError() {
// 		...
// 	}
//
// To prevent memory leaks, an Iterator must have Close called on it when it
// is no longer needed by the program.
type Iterator struct {
	Iter *C.leveldb_iterator_t
}
