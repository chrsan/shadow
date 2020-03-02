package internal

import (
	"math"
)

func f15(ctx *Context, l0 float32, l1 float32) float32 {
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
	if s0i32 != 0 {
		s0f32 = l0
		s1f32 = l0
		s2f32 = l1
		s1f32 = float32(math.Max(float64(s1f32), float64(s2f32)))
		s2f32 = l1
		s2i32 = int32(math.Float32bits(s2f32))
		s3i32 = 2147483647
		s2i32 = s2i32 & s3i32
		s3i32 = 2139095040
		if uint32(s2i32) > uint32(s3i32) {
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
	s0f32 = l1
	return s0f32
}
