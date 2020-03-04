package internal

import (
	"unsafe"
)

func f1093(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l0 = s0i32
	lbl1:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 1936876899
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 8
			s0i32 = s0i32 + s1i32
			return s0i32
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = 0
	return s0i32
}
