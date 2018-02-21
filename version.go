package pebigo

// #cgo LDFLAGS: -lpebblesdb
// #include "pebblesdb/c.h"
import "C"

func GetLevelDBMajorVersion() int {
	return int(C.leveldb_major_version())
}

func GetLevelDBMinorVersion() int {
	return int(C.leveldb_minor_version())
}
