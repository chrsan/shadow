package internal

import (
	"unsafe"
)

func f1551(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s1i32 = 26392
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 36
	s0i32 = s0i32 + s1i32
	s0i32 = f41(ctx, s0i32)
	s0i32 = l0
	s0i32 = f388(ctx, s0i32)
	s0i32 = l0
	return s0i32
}
