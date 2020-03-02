package internal

import (
	"math"
)

func f20(ctx *Context, l0 int32, l1 float32) int32 {
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s5f32 float32
	_ = s5f32
	s0i32 = l0
	s1i32 = 14
	s2i32 = -1
	s3i32 = -1
	s4i32 = -1
	s5f32 = l1
	s5i32 = int32(math.Float32bits(s5f32))
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	return s0i32
}
