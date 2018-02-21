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

func (it *Iterator) Destroy() {
	C.leveldb_iter_destroy(it.Iter)
	it.Iter = nil
}

func (it *Iterator) Valid() bool {
	return ucharToBool(C.leveldb_iter_valid(it.Iter))
}

func (it *Iterator) SeekToFirst() {
	C.leveldb_iter_seek_to_first(it.Iter)
}

func (it *Iterator) SeekToLast() {
	C.leveldb_iter_seek_to_last(it.Iter)
}

func (it *Iterator) Seek(key []byte) {
	k := (*C.char)(unsafe.Pointer(&key[0]))
	lenk := len(key)
	C.leveldb_iter_seek(it.Iter, k, C.size_t(lenk))
}

func (it *Iterator) Next() {
	C.leveldb_iter_next(it.Iter)
}

func (it *Iterator) Prev() {
	C.leveldb_iter_prev(it.Iter)
}

func (it *Iterator) Key() []byte {
	var klen C.size_t
	kdata := C.leveldb_iter_key(it.Iter, &klen)
	if kdata == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(kdata), C.int(klen))
}

func (it *Iterator) Value() []byte {
	var vlen C.size_t
	vdata := C.leveldb_iter_value(it.Iter, &vlen)
	if vdata == nil {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(vdata), C.int(vlen))
}

func (it *Iterator) GetError() error {
	var errStr *C.char
	C.leveldb_iter_get_error(it.Iter, &errStr)
	if errStr != nil {
		gs := C.GoString(errStr)
		C.leveldb_free(unsafe.Pointer(errStr))
		return IteratorError(gs)
	}
	return nil
}


