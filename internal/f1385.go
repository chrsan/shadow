package internal

import (
	"unsafe"
)

func f1385(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 184
	s2i32 = 0
	f16(ctx, s0i32, s1i32, s2i32)
	s0i32 = 1
	return s0i32
}
