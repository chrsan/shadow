package internal

import (
	"unsafe"
)

func f1252(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
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
	var s5i32 int32
	_ = s5i32
	s0i32 = l3
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l8 = s0i32
		s0i32 = l2
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l9 = s0i32
		s0i32 = l2
		s1i32 = 16
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l10 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l6
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = l2
		s1i32 = l7
		s2i32 = l1
		s3i32 = l6
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l5 = s2i32
		if s2i32 == 0 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl2
		}
		s1i32 = l4
		s2i32 = l5
		s3i32 = 65535
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			goto lbl2
		}
		s1i32 = l5
		s2i32 = 15
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l5
		s3i32 = 11
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 + s2i32
		s2i32 = l10
		s3i32 = l7
		s4i32 = 255
		s3i32 = s3i32 & s4i32
		l2 = s3i32
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 5
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l2
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 31
		s2i32 = s2i32 & s3i32
		l2 = s2i32
		s3i32 = 4
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = l2
		s2i32 = s2i32 + s3i32
		s3i32 = l8
		s4i32 = l7
		s5i32 = 16
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		l2 = s4i32
		s5i32 = 255
		s4i32 = s4i32 & s5i32
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s3i32 = 5
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = l2
		s2i32 = s2i32 + s3i32
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		s2i32 = l5
		s3i32 = 6
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 31
		s2i32 = s2i32 & s3i32
		l2 = s2i32
		s3i32 = 4
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = l2
		s2i32 = s2i32 + s3i32
		s3i32 = l9
		s4i32 = l7
		s5i32 = 8
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 255
		s4i32 = s4i32 & s5i32
		l2 = s4i32
		s3i32 = s3i32 - s4i32
		s2i32 = s2i32 * s3i32
		s3i32 = 5
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = l2
		s2i32 = s2i32 + s3i32
		s3i32 = 8
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		s2i32 = -16777216
		s1i32 = s1i32 | s2i32
	lbl2:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
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
