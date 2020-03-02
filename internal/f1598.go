package internal

import (
	"unsafe"
)

func f1598(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
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
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l3
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0i32 = f1160(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = 6
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 51
		s1i32 = l0
		s2i32 = 255
		s1i32 = s1i32 & s2i32
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l0 = s0i32
	s1i32 = 7
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	l2 = s0i32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl3
	case 2:
		goto lbl2
	case 3:
		goto lbl2
	case 4:
		goto lbl2
	case 5:
		goto lbl4
	case 6:
		goto lbl5
	default:
		goto lbl0
	}
lbl5:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	goto lbl0
lbl4:
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	goto lbl0
lbl3:
	s0i32 = l1
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	goto lbl0
lbl2:
	s0i32 = 0
	l2 = s0i32
lbl0:
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l2
	return s0i32
}
