package internal

import (
	"math"
)

func f134(ctx *Context, l0 float64, l1 float64) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0f64 = l1
	s0f64 = math.Abs(s0f64)
	l4 = s0f64
	s0f64 = l0
	s0f64 = math.Abs(s0f64)
	l5 = s0f64
	s1f64 = 3.4028234663852886e+38
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f64 = l4
	s1f64 = 3.4028234663852886e+38
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	s1f64 = l1
	s1f32 = float32(s1f64)
	s1i32 = int32(math.Float32bits(s1f32))
	l2 = s1i32
	s2i32 = 2147483647
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
	s1i32 = l2
	s2i32 = l2
	s3i32 = 0
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s1i32 = 0
	s2f64 = l0
	s2f32 = float32(s2f64)
	s2i32 = int32(math.Float32bits(s2f32))
	l2 = s2i32
	s3i32 = 2147483647
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = l2
	s3i32 = l2
	s4i32 = 0
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l2 = s1i32
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l2
	s2i32 = l3
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	if s1i32 < s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	return s0i32
lbl0:
	s0f64 = l0
	s1f64 = l1
	s0f64 = s0f64 - s1f64
	s0f64 = math.Abs(s0f64)
	s1f64 = l4
	s2f64 = l5
	s3f64 = l5
	s4f64 = l4
	if s3f64 < s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	s0f64 = s0f64 / s1f64
	s1f64 = 1.9073486328125e-06
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	return s0i32
}
