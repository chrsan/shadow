package internal

import (
	"unsafe"
)

func f65(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l0
		s1i32 = f354(ctx, s1i32)
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = 1
lbl0:
	return s0i32
}
