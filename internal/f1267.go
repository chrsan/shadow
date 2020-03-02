package internal

import (
	"unsafe"
)

func f1267(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		l1 = s2i32
		s3i32 = l2
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l0 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = 255
		s2i32 = l3
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l4
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
		s0i32 = l4
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l2
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
