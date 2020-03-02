package internal

import (
	"unsafe"
)

func f1734(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	l1 = s0i32
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		f146(ctx, s0i32)
	}
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	return s0i32
}
