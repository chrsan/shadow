package internal

import (
	"unsafe"
)

func f1377(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 4201
		s1i32 = l0
		s2i32 = l1
		s0i32 = f379(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = 4201
	s1i32 = l0
	s2i32 = l1
	s0i32 = f379(ctx, s0i32, s1i32, s2i32)
lbl0:
	l1 = s0i32
	s1i32 = l0
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		return s0i32
	}
	s0i32 = l1
	return s0i32
}
