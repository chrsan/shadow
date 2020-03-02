package internal

type ImportedFuncs interface {
	// module: env, field: round
	F0(ctx *Context, p0 float64) float64
	// module: env, field: sem_init
	F1(ctx *Context, p0 int32, p1 int32, p2 int32) int32
	// module: wasi_snapshot_preview1, field: environ_sizes_get
	F2(ctx *Context, p0 int32, p1 int32) int32
	// module: wasi_snapshot_preview1, field: fd_write
	F3(ctx *Context, p0 int32, p1 int32, p2 int32, p3 int32) int32
	// module: wasi_snapshot_preview1, field: fd_seek
	F4(ctx *Context, p0 int32, p1 int64, p2 int32, p3 int32) int32
	// module: wasi_snapshot_preview1, field: fd_close
	F5(ctx *Context, p0 int32) int32
	// module: wasi_snapshot_preview1, field: proc_exit
	F6(ctx *Context, p0 int32)
	// module: env, field: sem_wait
	F7(ctx *Context, p0 int32) int32
	// module: env, field: sem_post
	F8(ctx *Context, p0 int32) int32
	// module: env, field: emscripten_notify_memory_growth
	F9(ctx *Context, p0 int32)
	// module: wasi_snapshot_preview1, field: environ_get
	F10(ctx *Context, p0 int32, p1 int32) int32
	// module: env, field: sem_destroy
	F11(ctx *Context, p0 int32) int32
}
