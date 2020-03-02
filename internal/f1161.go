package internal

import (
	"unsafe"
)

func f1161(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
lbl0:
	s0i32 = l1
	s1i32 = l3
	s2i32 = l1
	s3i32 = l3
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 255
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l7 = s0i32
			s0i32 = l3
			s1i32 = 0
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl1
			}
		lbl4:
			s0i32 = l4
			s1i32 = l7
			s2i32 = 1
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			l6 = s1i32
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = l6
			s1i32 = s1i32 + s2i32
			s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
			l6 = s1i32
			s2i32 = 11
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = l5
			s1i32 = s1i32 * s2i32
			s2i32 = 128
			s1i32 = s1i32 + s2i32
			l8 = s1i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = l8
			s1i32 = s1i32 + s2i32
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = 63488
			s1i32 = s1i32 & s2i32
			s2i32 = l6
			s3i32 = 31
			s2i32 = s2i32 & s3i32
			s3i32 = l5
			s2i32 = s2i32 * s3i32
			s3i32 = 128
			s2i32 = s2i32 + s3i32
			l8 = s2i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = l8
			s2i32 = s2i32 + s3i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s1i32 = s1i32 | s2i32
			s2i32 = l6
			s3i32 = 5
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = 63
			s2i32 = s2i32 & s3i32
			s3i32 = l5
			s2i32 = s2i32 * s3i32
			s3i32 = 128
			s2i32 = s2i32 + s3i32
			l6 = s2i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = l6
			s2i32 = s2i32 + s3i32
			s3i32 = 3
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = 8160
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 | s2i32
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			s0i32 = l7
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			s1i32 = l3
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			goto lbl1
		}
		s0i32 = l4
		s1i32 = l0
		s2i32 = l3
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 0
	s2i32 = l5
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
lbl1:
	s0i32 = l1
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l3
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l3 = s1i32
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l0
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+2)])
		l3 = s0i32
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		goto lbl0
	}
}
