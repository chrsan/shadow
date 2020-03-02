package internal

import (
	"unsafe"
)

func f935(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s4i32 int32
	_ = s4i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s1i32 = 16777216
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l3
			s2i32 = -16777217
			if uint32(s1i32) <= uint32(s2i32) {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1i32 = l3
				s2i32 = 256
				s3i32 = l3
				s4i32 = 24
				s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
				s2i32 = s2i32 - s3i32
				l3 = s2i32
				s3i32 = l0
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
				l4 = s3i32
				s4i32 = 16711935
				s3i32 = s3i32 & s4i32
				s2i32 = s2i32 * s3i32
				s3i32 = 8
				s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
				s3i32 = 16711935
				s2i32 = s2i32 & s3i32
				s3i32 = l4
				s4i32 = 8
				s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
				s4i32 = 16711935
				s3i32 = s3i32 & s4i32
				s4i32 = l3
				s3i32 = s3i32 * s4i32
				s4i32 = -16711936
				s3i32 = s3i32 & s4i32
				s2i32 = s2i32 | s3i32
				s1i32 = s1i32 + s2i32
			} else {
				s1i32 = l3
			}
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		}
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l3 = s0i32
		s0i32 = l2
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
