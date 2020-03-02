package internal

import (
	"unsafe"
)

func f542(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
	s1i32 = -48
	s0i32 = s0i32 + s1i32
	s1i32 = 10
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s0i32 = l0
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l2
		s2i32 = 10
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		s1i32 = 10
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l2
	return s0i32
}
