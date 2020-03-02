package internal

import (
	"unsafe"
)

func f1054(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s1i32 = 20
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl3
	case 2:
		goto lbl4
	case 3:
		goto lbl4
	case 4:
		goto lbl3
	case 5:
		goto lbl4
	case 6:
		goto lbl4
	case 7:
		goto lbl4
	case 8:
		goto lbl3
	case 9:
		goto lbl3
	case 10:
		goto lbl3
	case 11:
		goto lbl4
	case 12:
		goto lbl4
	case 13:
		goto lbl4
	case 14:
		goto lbl3
	case 15:
		goto lbl5
	case 16:
		goto lbl3
	case 17:
		goto lbl5
	case 18:
		goto lbl3
	case 19:
		goto lbl4
	default:
		goto lbl6
	}
lbl6:
	goto lbl1
lbl5:
	s0i32 = l1
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 0
	l0 = s0i32
	s0i32 = 2
	l3 = s0i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl2
	case 2:
		goto lbl1
	default:
		goto lbl0
	}
lbl4:
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 0
	return s0i32
lbl3:
	s0i32 = 1
	l3 = s0i32
	goto lbl1
lbl2:
	s0i32 = l1
	l3 = s0i32
lbl1:
	s0i32 = 1
	l0 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl0:
	s0i32 = l0
	return s0i32
}
