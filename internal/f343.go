package internal

import (
	"unsafe"
)

func f343(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l0
	s0i32 = f1944(ctx, s0i32)
	return s0i32
}
