package internal

import (
	"math"
)

func f1417(ctx *Context, l0 float64) float64 {
	var l1 float64
	_ = l1
	var l2 float64
	_ = l2
	var l3 float64
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
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
	var s5f64 float64
	_ = s5f64
	s0f64 = l0
	s0f64 = math.Abs(s0f64)
	l2 = s0f64
	s1f64 = 1.6940658945086007e-21
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = 0
	} else {
		s0f64 = l2
		s1f64 = l2
		s2f64 = l2
		s2i64 = int64(math.Float64bits(s2f64))
		s3i64 = 32
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s2i32 = int32(s2i64)
		s3i32 = 3
		s2i32 = i32DivU(s2i32, s3i32)
		s3i32 = 715094163
		s2i32 = s2i32 + s3i32
		s2i64 = int64(uint32(s2i32))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s2f64 = math.Float64frombits(uint64(s2i64))
		l1 = s2f64
		s3f64 = l1
		s2f64 = s2f64 * s3f64
		s3f64 = l1
		s2f64 = s2f64 * s3f64
		l3 = s2f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 + s1f64
		s1f64 = l1
		s0f64 = s0f64 * s1f64
		s1f64 = l2
		s2f64 = l3
		s3f64 = l3
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 / s1f64
		l1 = s0f64
		s1f64 = l2
		s2f64 = l2
		s3f64 = l1
		s4f64 = l1
		s5f64 = l1
		s4f64 = s4f64 * s5f64
		s3f64 = s3f64 * s4f64
		l1 = s3f64
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 * s1f64
		s1f64 = l2
		s2f64 = l1
		s3f64 = l1
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 / s1f64
		l1 = s0f64
		s1f64 = l2
		s2f64 = l2
		s3f64 = l1
		s4f64 = l1
		s5f64 = l1
		s4f64 = s4f64 * s5f64
		s3f64 = s3f64 * s4f64
		l1 = s3f64
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 * s1f64
		s1f64 = l2
		s2f64 = l1
		s3f64 = l1
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 / s1f64
		l2 = s0f64
		s0f64 = -s0f64
		s1f64 = l2
		s2f64 = l0
		s3f64 = 0
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f64 = s0f64
		} else {
			s0f64 = s1f64
		}
	}
	return s0f64
}
