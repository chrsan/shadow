package internal

import (
	"unsafe"
)

func f1812(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s1i32 = 24364
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	f128(ctx, s0i32)
	s0i32 = l0
	s0i32 = f35(ctx, s0i32)
	s0i32 = l0
	return s0i32
}
