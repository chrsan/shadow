package internal

import (
	"unsafe"
)

func f224(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s0i32 = f81(ctx, s0i32, s1i32)
	s0i32 = l0
	return s0i32
}
