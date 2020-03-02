package internal

import (
	"unsafe"
)

func f1832(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+156)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 2
		ctx.Mem[int(s0i32+53)] = uint8(s1i32)
	}
}
