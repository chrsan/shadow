package internal

import (
	"math"
)

func f250(ctx *Context, l0 float64, l1 int32) float64 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	s0i32 = l1
	s1i32 = 1024
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l0
		s1f64 = 8.98846567431158e+307
		s0f64 = s0f64 * s1f64
		l0 = s0f64
		s0i32 = l1
		s1i32 = 2047
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = -1023
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			goto lbl0
		}
		s0f64 = l0
		s1f64 = 8.98846567431158e+307
		s0f64 = s0f64 * s1f64
		l0 = s0f64
		s0i32 = l1
		s1i32 = 3069
		s2i32 = l1
		s3i32 = 3069
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
		s1i32 = -2046
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl0
	}
	s0i32 = l1
	s1i32 = -1023
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f64 = l0
	s1f64 = 2.2250738585072014e-308
	s0f64 = s0f64 * s1f64
	l0 = s0f64
	s0i32 = l1
	s1i32 = -2045
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1022
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl0
	}
	s0f64 = l0
	s1f64 = 2.2250738585072014e-308
	s0f64 = s0f64 * s1f64
	l0 = s0f64
	s0i32 = l1
	s1i32 = -3066
	s2i32 = l1
	s3i32 = -3066
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
	s1i32 = 2044
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl0:
	s0f64 = l0
	s1i32 = l1
	s2i32 = 1023
	s1i32 = s1i32 + s2i32
	s1i64 = int64(uint32(s1i32))
	s2i64 = 52
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	s1f64 = math.Float64frombits(uint64(s1i64))
	s0f64 = s0f64 * s1f64
	return s0f64
}
