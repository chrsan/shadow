package internal

import (
	"unsafe"
)

func f500(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l3 = s0i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	s1i32 = 8190
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+24)])
		l4 = s0i32
		s1i32 = 5
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 1
			l1 = s0i32
			s0i32 = l4
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			switch s0i32 {
			case 0:
				goto lbl2
			case 1:
				goto lbl2
			case 2:
				goto lbl3
			case 3:
				goto lbl5
			case 4:
				goto lbl2
			default:
				goto lbl6
			}
		lbl6:
			s0i32 = l3
			s1i32 = 7
			s0i32 = s0i32 + s1i32
			s1i32 = 3
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			goto lbl1
		lbl5:
			s0i32 = 2
			l1 = s0i32
			goto lbl2
		}
		s0i32 = l2
		s1i32 = 60
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 3323
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 3282
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 3256
		s1i32 = l2
		f19(ctx, s0i32, s1i32)
		panic("unreachable executed")
	lbl3:
		s0i32 = 4
		l1 = s0i32
	lbl2:
		s0i32 = l1
		s1i32 = l3
		s0i32 = s0i32 * s1i32
	lbl1:
		l1 = s0i32
		s0i32 = l1
		s1i32 = l0
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+2)])))
		s0i32 = s0i32 * s1i32
		l0 = s0i32
		s1i32 = 3
		s0i32 = s0i32 * s1i32
		s1i32 = l0
		s2i32 = l4
		s3i32 = 2
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
	}
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l1
	return s0i32
}
