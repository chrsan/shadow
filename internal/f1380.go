package internal

import (
	"unsafe"
)

func f1380(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	s0i32 = f5(ctx, s0i32)
	return s0i32
}
