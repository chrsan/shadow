package internal

import (
	"unsafe"
)

func f1832(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l2
	s1i32 = 56
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 76695845
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+536)]))
		l2 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l3 = s0i32
		s1i32 = l1
		s2i32 = 56
		s1i32 = s1i32 * s2i32
		l1 = s1i32
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
			s1i32 = l1
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
		l0 = s1i32
		s2i32 = l1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+536)])) = uint32(s1i32)
		s0i32 = l0
		return s0i32
	}
	f63(ctx)
	panic("unreachable executed")
}
