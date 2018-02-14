package pebigo

// #cgo LDFLAGS: -lpebblesdb
// #include "pebblesdb/c.h"
import "C"

import (
	"unsafe"
)

// WriteBatch is a batching of Puts, and Deletes to be written atomically to a
// database. A WriteBatch is written when passed to DB.Write.
//
// To prevent memory leaks, call Close when the program no longer needs the
// WriteBatch object.
type WriteBatch struct {
	wbatch *C.leveldb_writebatch_t
}
