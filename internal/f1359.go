package internal

import (
	"math"
)

func f1359(ctx *Context, l0 float32) float32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
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
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l2 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s1i32 = 1065353216
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1065353216
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = 3.1415925
			s1f32 = 0
			s2i32 = l2
			s3i32 = 0
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			return s0f32
		}
		s0f32 = 0
		s1f32 = l0
		s2f32 = l0
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 / s1f32
		return s0f32
	}
	s0i32 = l1
	s1i32 = 1056964607
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 1.5707963
		s1i32 = l1
		s2i32 = 847249409
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl2
		}
		s0f32 = 7.5497894e-08
		s1f32 = l0
		s2f32 = l0
		s1f32 = s1f32 * s2f32
		l3 = s1f32
		s2f32 = l3
		s3f32 = l3
		s4f32 = -0.008656363
		s3f32 = s3f32 * s4f32
		s4f32 = -0.042743422
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 * s3f32
		s3f32 = 0.16666587
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = l3
		s3f32 = -0.70662963
		s2f32 = s2f32 * s3f32
		s3f32 = 1
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 / s2f32
		s2f32 = l0
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 - s1f32
		s1f32 = l0
		s0f32 = s0f32 - s1f32
		s1f32 = 1.5707963
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0i32 = l2
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 1.5707963
		s1f32 = l0
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l0 = s1f32
		s1f32 = float32(math.Sqrt(float64(s1f32)))
		l3 = s1f32
		s2f32 = l3
		s3f32 = l0
		s4f32 = l0
		s5f32 = l0
		s6f32 = -0.008656363
		s5f32 = s5f32 * s6f32
		s6f32 = -0.042743422
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 * s5f32
		s5f32 = 0.16666587
		s4f32 = s4f32 + s5f32
		s3f32 = s3f32 * s4f32
		s4f32 = l0
		s5f32 = -0.70662963
		s4f32 = s4f32 * s5f32
		s5f32 = 1
		s4f32 = s4f32 + s5f32
		s3f32 = s3f32 / s4f32
		s2f32 = s2f32 * s3f32
		s3f32 = -7.5497894e-08
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 - s1f32
		l0 = s0f32
		s1f32 = l0
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0f32 = 1
	s1f32 = l0
	s0f32 = s0f32 - s1f32
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	l0 = s0f32
	s1f32 = l0
	s2f32 = l0
	s3f32 = -0.008656363
	s2f32 = s2f32 * s3f32
	s3f32 = -0.042743422
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = 0.16666587
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l0
	s2f32 = -0.70662963
	s1f32 = s1f32 * s2f32
	s2f32 = 1
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 / s1f32
	s1f32 = l0
	s1f32 = float32(math.Sqrt(float64(s1f32)))
	l3 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l0
	s2f32 = l3
	s2i32 = int32(math.Float32bits(s2f32))
	s3i32 = -4096
	s2i32 = s2i32 & s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l0 = s2f32
	s3f32 = l0
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = l3
	s3f32 = l0
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 / s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l0
	s0f32 = s0f32 + s1f32
	l0 = s0f32
	s1f32 = l0
	s0f32 = s0f32 + s1f32
lbl2:
	return s0f32
}
