package internal

import (
	"math"
)

func f497(ctx *Context, l0 int32) float32 {
	var l1 int32
	_ = l1
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
	s0i32 = l0
	s1i32 = 1023
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = -2147483648
	s0i32 = s0i32 & s1i32
	s1i32 = l1
	s2i32 = 1056964608
	s1i32 = s1i32 | s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s2f32 = -0.5
	s1f32 = s1f32 + s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = l0
	s3i32 = 10
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 31
	s2i32 = s2i32 & s3i32
	l0 = s2i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl0
	}
	s1i32 = l1
	s2i32 = 13
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l1 = s1i32
	s1i32 = l1
	s2i32 = 2139095040
	s1i32 = s1i32 | s2i32
	s2i32 = l0
	s3i32 = 31
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl0
	}
	s1i32 = l0
	s2i32 = 23
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l1
	s1i32 = s1i32 | s2i32
	s2i32 = 939524096
	s1i32 = s1i32 + s2i32
lbl0:
	s0i32 = s0i32 | s1i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	return s0f32
}
