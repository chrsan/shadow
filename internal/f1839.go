package internal

import (
	"unsafe"
)

func f1839(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+536)]))
	l2 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l3 = s0i32
	s1i32 = 68
	s0i32 = s0i32 | s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+540)]))
	s2i32 = l2
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 532
		s0i32 = s0i32 + s1i32
		s1i32 = 68
		s2i32 = 4
		f18(ctx, s0i32, s1i32, s2i32)
		s0i32 = 0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+536)]))
		l2 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l3 = s0i32
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = l3
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = 68
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+536)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 0
	s2i32 = 68
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
	l3 = s0i32
	s1i32 = l1
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+564)]))
	s0i32 = f640(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = f318(ctx, s0i32)
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s0i32 = f217(ctx, s0i32)
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
}
