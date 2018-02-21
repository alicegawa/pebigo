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

func NewOptions() *Options {
	opt := C.leveldb_options_create()
	return &Option{opt}
}

func NewReadOptions() *ReadOptions {
	opt := C.leveldb_readoptions_create()
	return &ReadOptions{opt}
}

func NewWriteOptions() *WriteOptions {
	opt := C.leveldb_writeoptions_create()
	return &WriteOptions{opt}
}

func (o *Options) Close() {
	C.leveldb_options_destroy(o.Opt)
}

func (o *Options) SetComparator(cmp *C.leveldb_comparaotr_t) {
	C.leveldb_options_set_comparator(o.Opt, cmp)
}

func (o *Options) SetFilterPolicy(fp *FilterPolicy) {
	var policy *C.leveldb_filter_policy_t
	if fp != nil {
		policy = fp.Policy
	}
	C.leveldb_options_set_fileter_policy(o.Opt, fp)
}

func (o *Options) SetCreateIfMissing(b bool) {
	C.leveldb_options_set_create_if_missing(o.Opt, boolToUchar(b))
}

func (o *Options) SetErrorIfExists(b bool) {
	C.leveldb_options_set_error_if_exists(o.Opt, booToUchar(b))
}

func (o *Options) SetParanoidChecks(b bool) {
	C.leveldb_options_set_paranoid_checks(o.Opt, boolToUchar(b))
}

func (o *Options) SetEnv(env *Env) {
	C.leveldb_options_set_env(o.Opt, env.Env)
}

func (o *Options) SetInfoLog(log *C.leveldb_logger_t) {
	C.leveldb_options_set_info_log(o.Opt, log)
}

func (o *Options) SetWriteBufferSize(s int) {
	C.leveldb_options_set_write_buffer_size(o.Opt, C.size_t(s))
}

func (o *Options) SetMaxOpenFiles(n int) {
	C.leveldb_options_set_max_open_files(o.Opt, C.int(n))
}

func (o *Options) SetCache(cache *Cache) {
	C.leveldb_options_set_cache(o.Opt, cache.Cache)
}

func (o *Options) SetBlockSize(s int) {
	C.leveldb_options_set_block_size(o.Opt, C.size_t(s))
}

func (o *Options) SetBlockRestartInterval(n int) {
	C.leveldb_options_set_block_restart_interval(o.Opt, C.int(n))
}

func (o *Options) SetCompression(n int) {
	C.leveldb_options_set_compression(o.Opt, C.int(n))
}

func (ro *ReadOptions) Close() {
	C.leveldb_readoptions_destroy(ro.Opt)
}

func (ro *ReadOptions) SetVerifyChecksums(b bool) {
	C.leveldb_readoptions_set_verify_checksums(ro.Opt, boolToUchar(b))
}

func (ro *ReadOptions) SetFillCache(b bool) {
	C.leveldb_readoptions_set_fill_cache(ro.Opt, boolToUchar(b))
}

func (ro *ReadOptions) SetSnapshot(snap *Snapshot) {
	if snap != nil {
		C.leveldb_readoptions_set_snapshot(ro.Opt, snap.snap)
	}
}

func (wo *WriteOptions) Close() {
	C.leveldb_writeoptions_destroy(wo.Opt)
}

func (wo *WriteOptions) SetSync(b bool) {
	C.leveldb_writeoptions_set_sync(wo.Opt, boolToUchar(b))
}
