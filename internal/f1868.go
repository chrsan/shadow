package internal

import (
	"unsafe"
)

func f1868(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 1064
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		s0i32 = 0
	} else {
		s0i32 = 1
		s1i32 = 2
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+42)])
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
	}
	return s0i32
}
