package pebigo

// #cgo LDFLAGS: -lpebblesdb
// #include "pebblesdb/c.h"
import "C"

// CompressionOpt is a value for Options.SetCompression.
type CompressionOpt int

// Known compression arguments for Options.SetCompression.
const (
	NoCompression     = CompressionOpt(0)
	SnappyCompression = CompressionOpt(1)
)

// Options represent all of the available options when opening a database with
// Open. Options should be created with NewOptions.
//
// It is usually with to call SetCache with a cache object. Otherwise, all
// data will be read off disk.
//
// To prevent memory leaks, Close must be called on an Options when the
// program no longer needs it.
type Options struct {
	Opt *C.leveldb_options_t
}

// WriteOptions represent all of the available options when writing from a
// database.
//
// To prevent memory leaks, Close must called on a WriteOptions when the
// program no longer needs it.
type WriteOptions struct {
	Opt *C.leveldb_writeoptions_t
}

// ReadOptions represent all of the available options when reading from a
// database.
//
// To prevent memory leaks, Close must called on a ReadOptions when the
// program no longer needs it.
type ReadOptions struct {
	Opt *C.leveldb_readoptions_t
}
