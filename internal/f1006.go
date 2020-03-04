package internal

import (
	"math"
	"unsafe"
)

func f1006(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 float32
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
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0f32
	s0i32 = l1
	s1i32 = l2
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+16)])
	if s2i32 != 0 {
		s2i32 = l3
		s3f32 = l5
		s2f32 = f180(ctx, s2i32, s3f32)
	} else {
		s2f32 = l5
	}
	s3f32 = 128
	s2f32 = float32(math.Min(float64(s2f32), float64(s3f32)))
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l4
	s0i32 = f1026(ctx, s0i32, s1i32, s2f32, s3i32, s4i32)
	return s0i32
}
