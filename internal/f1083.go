package internal

import (
	"unsafe"
)

func f1083(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	f243(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+52)])
	s2i32 = 253
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+52)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = 0
	f242(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+76)])
	s2i32 = 253
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+76)] = uint8(s1i32)
}
