package internal

import (
	"unsafe"
)

func f694(ctx *Context, l0 float32, l1 float32, l2 float32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0f32 = l0
	s1f32 = l1
	s0f32 = s0f32 - s1f32
	l0 = s0f32
	s0f32 = -s0f32
	s1f32 = l0
	s2f32 = l0
	s3f32 = 0
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l5 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l6 = s0f32
	s1f32 = l0
	s2f32 = l1
	s1f32 = s1f32 - s2f32
	s2f32 = l2
	s1f32 = s1f32 + s2f32
	l0 = s1f32
	s1f32 = -s1f32
	s2f32 = l0
	s3i32 = l5
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l0 = s1f32
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l6
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l0
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l6
	s1f32 = l0
	s0f32 = s0f32 / s1f32
	l0 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l0
	s2f32 = l0
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1f32 = l0
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = 1
	l4 = s0i32
lbl0:
	s0i32 = l4
	return s0i32
}
