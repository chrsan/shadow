package internal

import (
	"unsafe"
)

func f1915(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l8 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l7
	s0i32 = s0i32 * s1i32
	s1i32 = l1
	s2i32 = l8
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 3472
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
lbl0:
	l1 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i64 = l5
		s0i32 = int32(s0i64)
		l2 = s0i32
	lbl2:
		s0i32 = l1
		s1i32 = l2
		s2i32 = l3
		s0i32 = f27(ctx, s0i32, s1i32, s2i32)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l4
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
		s0i32 = l4
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l6
		if s0i32 != 0 {
			goto lbl2
		}
	}
}
