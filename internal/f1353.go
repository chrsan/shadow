package internal

import (
	"unsafe"
)

func f1353(ctx *Context) {
	var l0 int32
	_ = l0
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = 15476
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s1i32 = 19135
	s2i32 = 0
	s0i32 = f375(ctx, s0i32, s1i32, s2i32)
	s0i32 = 10
	s1i32 = l0
	f545(ctx, s0i32, s1i32)
	f63(ctx)
	panic("unreachable executed")
}
