package internal

// module: wasi_snapshot_preview1, field: fd_seek
func f4(ctx *Context, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	return ctx.f.F4(ctx, l0, l1, l2, l3)
}
