package internal

import (
	"unsafe"
)

func f1660(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = 4
	s0i32 = f17(ctx, s0i32)
	l0 = s0i32
	s1i32 = 25744
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	return s0i32
}
