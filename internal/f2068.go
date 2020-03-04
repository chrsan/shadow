package internal

import (
	"unsafe"
)

func f2068(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s4i32 int32
	_ = s4i32
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
		l2 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l4
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		s2i32 = 24
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l6 = s2i32
		s3i32 = 24
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 + s2i32
		s2i32 = 23
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = -16777216
		s1i32 = s1i32 & s2i32
		s2i32 = l5
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = l6
		s4i32 = 255
		s3i32 = s3i32 & s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 1
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 | s2i32
		s2i32 = l5
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = l6
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 255
		s3i32 = s3i32 & s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 7
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 65280
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		s2i32 = l5
		s3i32 = 16
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = l6
		s4i32 = 16
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 255
		s3i32 = s3i32 & s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 15
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 16711680
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
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
