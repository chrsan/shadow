package internal

import (
	"unsafe"
)

func f1621(ctx *Context, l0 int32, l1 float32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s1f32 float32
	_ = s1f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l3 = s0i32
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 36
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		f66(ctx, s0i32, s1i32)
	}
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	s1f32 = l1
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
}
