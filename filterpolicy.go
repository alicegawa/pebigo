package pebigo

// #cgo LDFLAGS: -lpebblesdb
// #include "pebblesdb/c.h"
import "C"

type FilterPolicy struct {
	Policy *C.leveldb_filterpolicy_t
}

func NewBloomFileter(bitsPerKey int) *FileterPolicy {
	return &FilterPolicy{C.leveldb_filterpolicy_create_bloom(C.int(bitsPerKey))}
}

func (fp *FilterPolicy) Destroy() {
	C.leveldb_filterpolicy_destroy(fp.Policy)
}
