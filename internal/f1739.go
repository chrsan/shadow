package internal

import (
	"unsafe"
)

func f1739(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l7 = s1i32
		s2i32 = l2
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+72)]))
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+68)]))
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l8 = s1i32
		s2i32 = l2
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
	lbl1:
		s0i32 = 0
		l1 = s0i32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+80)])
		if s0i32 != 0 {
		lbl4:
			s0i32 = l5
			s1i32 = l1
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l6 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
			s0i32 = l6
			s1i32 = 24
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l6 = s0i32
			s1i32 = 255
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l1
				s1i32 = l2
				s0i32 = s0i32 + s1i32
				l9 = s0i32
				s1i32 = l9
				s1i32 = int32(ctx.Mem[int(s1i32+0)])
				s2i32 = l6
				s3i32 = 255
				s2i32 = s2i32 ^ s3i32
				s1i32 = s1i32 * s2i32
				s2i32 = 257
				s1i32 = s1i32 * s2i32
				s2i32 = 127
				s1i32 = s1i32 + s2i32
				s2i32 = 16
				s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
				s2i32 = l6
				s1i32 = s1i32 + s2i32
				ctx.Mem[int(s0i32+0)] = uint8(s1i32)
				goto lbl5
			}
			s0i32 = l1
			s1i32 = l2
			s0i32 = s0i32 + s1i32
			s1i32 = 255
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		lbl5:
			s0i32 = l1
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i32 = l3
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			goto lbl2
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl7:
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = l1
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+3)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
	lbl2:
		s0i32 = l5
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l2
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l4
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
