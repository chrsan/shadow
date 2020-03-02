package internal

import (
	"unsafe"
)

func f1684(ctx *Context, l0 int32, l1 int32) float32 {
	var l2 float32
	_ = l2
	var s1i32 int32
	_ = s1i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0f32 = 0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	s2f32 = -1
	s3i32 = l1
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l2 = s1f32
	s2f32 = 0
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = 1
	s1f32 = l2
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l2
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	s1f32 = 1.4142135
	s2f32 = 1
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	l2 = s3f32
	s4f32 = 1
	s5f32 = l2
	s6f32 = 1
	if s5f32 > s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+44)]))
	l0 = s4i32
	s5i32 = 48
	s4i32 = s4i32 & s5i32
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l2 = s2f32
	s3f32 = l2
	s4f32 = 1.4142135
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = l2
	s3i32 = l0
	s4i32 = 12
	s3i32 = s3i32 & s4i32
	s4i32 = 8
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s0f32 = s0f32 * s1f32
lbl0:
	return s0f32
}
