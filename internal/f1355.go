package internal

import (
	"math"
)

func f1355(ctx *Context, l0 float32, l1 int32) float32 {
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
	s0i32 = l1
	s1i32 = 128
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s1f32 = 1.7014118e+38
		s0f32 = s0f32 * s1f32
		l0 = s0f32
		s0i32 = l1
		s1i32 = 255
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = -127
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			goto lbl0
		}
		s0f32 = l0
		s1f32 = 1.7014118e+38
		s0f32 = s0f32 * s1f32
		l0 = s0f32
		s0i32 = l1
		s1i32 = 381
		s2i32 = l1
		s3i32 = 381
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
		s1i32 = -254
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl0
	}
	s0i32 = l1
	s1i32 = -127
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l0
	s1f32 = 1.1754944e-38
	s0f32 = s0f32 * s1f32
	l0 = s0f32
	s0i32 = l1
	s1i32 = -253
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 126
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl0
	}
	s0f32 = l0
	s1f32 = 1.1754944e-38
	s0f32 = s0f32 * s1f32
	l0 = s0f32
	s0i32 = l1
	s1i32 = -378
	s2i32 = l1
	s3i32 = -378
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
	s1i32 = 252
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl0:
	s0f32 = l0
	s1i32 = l1
	s2i32 = 23
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 1065353216
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	s0f32 = s0f32 * s1f32
	return s0f32
}
