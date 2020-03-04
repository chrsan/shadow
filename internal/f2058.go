package internal

import (
	"unsafe"
)

func f2058(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l6
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l4 = s1i32
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l4
		s1i32 = s1i32 | s2i32
		s2i32 = 132184095
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l4 = s2i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l4
		s2i32 = s2i32 | s3i32
		s3i32 = 132184095
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l4 = s2i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l4
		s2i32 = s2i32 | s3i32
		s3i32 = 132184095
		s2i32 = s2i32 & s3i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s2i32 = 18
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 2016
		s1i32 = s1i32 & s2i32
		s2i32 = l4
		s3i32 = 2
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 63519
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l3
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
