package shadow

import (
	"math"

	"github.com/chrsan/shadow/internal"
)

type importedFuncs struct{}

// module: env, field: round
func (importedFuncs) F0(ctx *internal.Context, p0 float64) float64 {
	return math.Round(p0)
}

// module: env, field: sem_init
func (importedFuncs) F1(ctx *internal.Context, p0 int32, p1 int32, p2 int32) int32 {
	panic("not implemented")
}

// module: wasi_snapshot_preview1, field: environ_sizes_get
func (importedFuncs) F2(ctx *internal.Context, p0 int32, p1 int32) int32 {
	return 0
}

// module: wasi_snapshot_preview1, field: fd_write
func (importedFuncs) F3(ctx *internal.Context, p0 int32, p1 int32, p2 int32, p3 int32) int32 {
	panic("not implemented")
}

// module: wasi_snapshot_preview1, field: fd_seek
func (importedFuncs) F4(ctx *internal.Context, p0 int32, p1 int64, p2 int32, p3 int32) int32 {
	panic("not implemented")
}

// module: wasi_snapshot_preview1, field: fd_close
func (importedFuncs) F5(ctx *internal.Context, p0 int32) int32 {
	panic("not implemented")
}

// module: wasi_snapshot_preview1, field: proc_exit
func (importedFuncs) F6(ctx *internal.Context, p0 int32) {
	panic("not implemented")
}

// module: env, field: sem_wait
func (importedFuncs) F7(ctx *internal.Context, p0 int32) int32 {
	panic("not implemented")
}

// module: env, field: sem_post
func (importedFuncs) F8(ctx *internal.Context, p0 int32) int32 {
	panic("not implemented")
}

// module: env, field: emscripten_notify_memory_growth
func (importedFuncs) F9(ctx *internal.Context, p0 int32) {
	panic("not implemented")
}

// module: wasi_snapshot_preview1, field: environ_get
func (importedFuncs) F10(ctx *internal.Context, p0 int32, p1 int32) int32 {
	return 0
}

// module: env, field: sem_destroy
func (importedFuncs) F11(ctx *internal.Context, p0 int32) int32 {
	panic("not implemented")
}
