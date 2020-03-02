package internal

import (
	"unsafe"
)

func f1223(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
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
		l5 = s0i32
		s1i32 = l1
		s2i32 = l3
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s2i32 = 24
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		l7 = s1i32
		s2i32 = l5
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l5 = s2i32
		s3i32 = 11
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 * s2i32
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s2i32 = 5
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		s2i32 = 5
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 63488
		s1i32 = s1i32 & s2i32
		s2i32 = l4
		s3i32 = 16
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = l5
		s4i32 = 31
		s3i32 = s3i32 & s4i32
		s4i32 = l7
		s3i32 = s3i32 * s4i32
		s4i32 = 16
		s3i32 = s3i32 + s4i32
		l6 = s3i32
		s4i32 = 5
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = l6
		s3i32 = s3i32 + s4i32
		s4i32 = 5
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 | s2i32
		s2i32 = l4
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = l5
		s4i32 = 5
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 63
		s3i32 = s3i32 & s4i32
		s4i32 = l7
		s3i32 = s3i32 * s4i32
		s4i32 = 32
		s3i32 = s3i32 + s4i32
		l4 = s3i32
		s4i32 = 6
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = l4
		s3i32 = s3i32 + s4i32
		s4i32 = 6
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 65504
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
