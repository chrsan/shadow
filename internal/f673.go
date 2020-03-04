package internal

import (
	"unsafe"
)

func f673(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	var s5i32 int32
	_ = s5i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1i32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s1i32 = 3
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 & s1i32
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 0
	s2i32 = l1
	s3i32 = l2
	s2i32 = s2i32 + s3i32
	l3 = s2i32
	s3i32 = l1
	s4i32 = 1
	s3i32 = s3i32 + s4i32
	l4 = s3i32
	s4i32 = l3
	s5i32 = l4
	if uint32(s4i32) > uint32(s5i32) {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l1
	s2i32 = s2i32 - s3i32
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
}
