package internal

import (
	"unsafe"
)

func f1160(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l0
	s1i32 = 14
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 3136
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = 1
	l3 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 3136
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl0:
	s0i32 = l3
	return s0i32
}