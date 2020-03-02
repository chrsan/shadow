package internal

import (
	"math/bits"
	"unsafe"
)

func f926(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
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
	s0i32 = l1
	s1i32 = l1
	s2i32 = 4
	if uint32(s1i32) < uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
	} else {
		s1i32 = l1
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		l5 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1i32 = l0
		l3 = s1i32
	lbl1:
		s1i32 = l2
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l2 = s2i32
		s3i32 = -862048943
		s2i32 = s2i32 * s3i32
		s3i32 = 17
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = l2
		s4i32 = 380141568
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 | s3i32
		s3i32 = 461845907
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 ^ s2i32
		s2i32 = 13
		s1i32 = int32(bits.RotateLeft32(uint32(s1i32), int(s2i32)))
		s2i32 = 5
		s1i32 = s1i32 * s2i32
		s2i32 = -430675100
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s1i32 = l3
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1i32 = l1
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 3
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl1
		}
		s1i32 = l0
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s1i32 = l4
		s2i32 = l5
		s1i32 = s1i32 - s2i32
	}
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = 2
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = 0
		l3 = s1i32
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		switch s1i32 {
		case 0:
			goto lbl4
		case 1:
			goto lbl5
		default:
			goto lbl3
		}
	lbl5:
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+2)])
		s2i32 = 16
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l3 = s1i32
	lbl4:
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+1)])
		s2i32 = 8
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l3
		s1i32 = s1i32 | s2i32
		l3 = s1i32
	lbl3:
		s1i32 = l3
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		s1i32 = s1i32 ^ s2i32
		l0 = s1i32
		s2i32 = -862048943
		s1i32 = s1i32 * s2i32
		s2i32 = 17
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = l0
		s3i32 = 380141568
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 | s2i32
		s2i32 = 461845907
		s1i32 = s1i32 * s2i32
		s2i32 = l2
		s1i32 = s1i32 ^ s2i32
	} else {
		s1i32 = l2
	}
	s0i32 = s0i32 ^ s1i32
	l0 = s0i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l0
	s0i32 = s0i32 ^ s1i32
	s1i32 = -2048144789
	s0i32 = s0i32 * s1i32
	l0 = s0i32
	s1i32 = 13
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l0
	s0i32 = s0i32 ^ s1i32
	s1i32 = -1028477387
	s0i32 = s0i32 * s1i32
	l0 = s0i32
	s1i32 = 16
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = l0
	s0i32 = s0i32 ^ s1i32
	return s0i32
}
