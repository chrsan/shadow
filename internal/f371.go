package internal

import (
	"math"
)

func f371(ctx *Context, l0 float32) float32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
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
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l1 = s0i32
	s1i32 = 8388608
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l1
	s3i32 = -1
	if s2i32 > s3i32 {
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
		s0i32 = l1
		s1i32 = 2147483647
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = -1
			s1f32 = l0
			s2f32 = l0
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 / s1f32
			return s0f32
		}
		s0i32 = l1
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l0
			s1f32 = l0
			s0f32 = s0f32 - s1f32
			s1f32 = 0
			s0f32 = s0f32 / s1f32
			return s0f32
		}
		s0f32 = l0
		s1f32 = 3.3554432e+07
		s0f32 = s0f32 * s1f32
		s0i32 = int32(math.Float32bits(s0f32))
		l1 = s0i32
		s0i32 = -152
		l2 = s0i32
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 2139095039
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = -127
	l2 = s0i32
	s0f32 = 0
	l0 = s0f32
	s0i32 = l1
	s1i32 = 1065353216
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l1
	s1i32 = 4913933
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 8388607
	s0i32 = s0i32 & s1i32
	s1i32 = 1060439283
	s0i32 = s0i32 + s1i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	s1f32 = -1
	s0f32 = s0f32 + s1f32
	l0 = s0f32
	s1f32 = l0
	s2f32 = l0
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	l3 = s1f32
	s0f32 = s0f32 - s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = -4096
	s0i32 = s0i32 & s1i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l4 = s0f32
	s1f32 = 1.4428711
	s0f32 = s0f32 * s1f32
	s1f32 = l0
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	s2f32 = l3
	s1f32 = s1f32 - s2f32
	s2f32 = l0
	s3f32 = l0
	s4f32 = 2
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	l0 = s2f32
	s3f32 = l3
	s4f32 = l0
	s5f32 = l0
	s4f32 = s4f32 * s5f32
	l0 = s4f32
	s5f32 = l0
	s6f32 = l0
	s5f32 = s5f32 * s6f32
	l0 = s5f32
	s6f32 = 0.28498787
	s5f32 = s5f32 * s6f32
	s6f32 = 0.6666666
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = l0
	s6f32 = l0
	s7f32 = 0.24279079
	s6f32 = s6f32 * s7f32
	s7f32 = 0.40000972
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l0 = s1f32
	s2f32 = 1.4428711
	s1f32 = s1f32 * s2f32
	s2f32 = l0
	s3f32 = l4
	s2f32 = s2f32 + s3f32
	s3f32 = -0.00017605285
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	s1i32 = l2
	s2i32 = l1
	s3i32 = 23
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 + s2i32
	s1f32 = float32(s1i32)
	s0f32 = s0f32 + s1f32
	l0 = s0f32
lbl0:
	s0f32 = l0
	return s0f32
}
