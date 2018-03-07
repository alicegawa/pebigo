package pebigo

// #cgo LDFLAGS: -lpebblesdb
// #include "pebblesdb/c.h"
import "C"

type Env struct {
	Env *C.leveldb_env_t
}

func NewDefaultEnv() *Env {
	return &Env{C.leveldb_create_default_env()}
}

func (env *Env) Close() {
	C.leveldb_env_destroy(env.Env)
}
