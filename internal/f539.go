package internal

import (
	"unsafe"
)

func f539(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
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
	s0i32 = l0
	l1 = s0i32
	s0i32 = l0
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl3:
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 != 0 {
		goto lbl3
	}
	goto lbl0
lbl1:
lbl4:
	s0i32 = l1
	l2 = s0i32
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s1i32 = -1
	s0i32 = s0i32 ^ s1i32
	s1i32 = l3
	s2i32 = -16843009
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 & s1i32
	s1i32 = -2139062144
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l1 = s0i32
		goto lbl0
	}
lbl6:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	l3 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	l2 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl6
	}
lbl0:
	s0i32 = l1
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	return s0i32
}
