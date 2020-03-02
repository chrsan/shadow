package internal

import (
	"math"
)

func f62(ctx *Context, l0 float32, l1 float32) float32 {
	var l2 float32
	_ = l2
	var l3 float64
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0f32 float32
	_ = s0f32
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
	s0f32 = l0
	s1f32 = l0
	s0f32 = s0f32 * s1f32
	s1f32 = l1
	s2f32 = l1
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l2
		s0f32 = float32(math.Sqrt(float64(s0f32)))
		return s0f32
	}
	s0f32 = l0
	s0f64 = float64(s0f32)
	l3 = s0f64
	s1f64 = l3
	s0f64 = s0f64 * s1f64
	s1f32 = l1
	s1f64 = float64(s1f32)
	l3 = s1f64
	s2f64 = l3
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s0f64 = math.Sqrt(s0f64)
	s0f32 = float32(s0f64)
	return s0f32
}
