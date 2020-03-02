package internal

import (
	"math"
)

func f133(ctx *Context, l0 float32, l1 float32) float32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l1
	s2i32 = int32(math.Float32bits(s2f32))
	s3i32 = 2147483647
	s2i32 = s2i32 & s3i32
	s3i32 = 2139095041
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s1f32 = l1
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l2 = s0i32
	s1f32 = l1
	s1i32 = int32(math.Float32bits(s1f32))
	l3 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s0i32 = l2
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l3
		s1i32 = -2147483648
		s0i32 = s0i32 & s1i32
		s1i32 = 1
		s0i32 = s0i32 | s1i32
		goto lbl2
	}
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = l4
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 0
	s3i32 = l2
	s4i32 = l3
	s3i32 = s3i32 ^ s4i32
	s4i32 = -1
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
lbl2:
	s0f32 = math.Float32frombits(uint32(s0i32))
	l1 = s0f32
lbl1:
	s0f32 = l1
	return s0f32
}
