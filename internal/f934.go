package internal

import (
	"unsafe"
)

func f934(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l0
		s1i32 = l4
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l3 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l3
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l3 = s1i32
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 16711680
		s1i32 = s1i32 & s2i32
		s2i32 = l3
		s3i32 = -16711936
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		s2i32 = l3
		s3i32 = 16
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
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