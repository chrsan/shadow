package internal

import (
	"unsafe"
)

func f1166(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 1303
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 3094
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 3052
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 2916
	s1i32 = l0
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
