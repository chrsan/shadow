package internal

import (
	"unsafe"
)

func f109(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f11(ctx, s0i32)
		s0i32 = l0
		f12(ctx, s0i32)
	}
}
