package internal

import (
	"unsafe"
)

func f1223(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = 0
	l3 = s0i32
	s0i32 = l2
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l0
		s1i32 = l3
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l3
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l4 = s1i32
		s2i32 = 8
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 63488
		s1i32 = s1i32 & s2i32
		s2i32 = l4
		s3i32 = 5
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 2016
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		s2i32 = l4
		s3i32 = 19
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 31
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
