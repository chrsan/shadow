package internal

import (
	"unsafe"
)

func f1951(ctx *Context, l0 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s1i32 = 23760
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	s1i32 = 1
	f91(ctx, s0i32, s1i32)
	s0i32 = l0
	f12(ctx, s0i32)
}
