package internal

import (
	"unsafe"
)

func f945(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = 7
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl0
	case 1:
		goto lbl0
	case 2:
		goto lbl0
	case 3:
		goto lbl0
	case 4:
		goto lbl0
	case 5:
		goto lbl0
	default:
		goto lbl1
	}
lbl1:
	s0i32 = 1
	l1 = s0i32
lbl0:
	s0i32 = l1
	return s0i32
}
