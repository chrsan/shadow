package internal

import (
	"unsafe"
)

func f1627(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l1
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl2
	case 1:
		goto lbl1
	default:
		goto lbl3
	}
lbl3:
	s0i32 = 1073
	s1i32 = 1074
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	return s0i32
lbl2:
	s0i32 = 1075
	s1i32 = 1076
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	return s0i32
lbl1:
	s0i32 = 1077
	s1i32 = 1078
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	return s0i32
}
