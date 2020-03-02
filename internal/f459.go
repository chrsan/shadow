package internal

import (
	"unsafe"
)

func f459(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		f669(ctx, s0i32)
	}
}
