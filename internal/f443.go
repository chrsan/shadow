package internal

import (
	"math"
	"unsafe"
)

func f443(ctx *Context, l0 int32, l1 float32, l2 int32) float32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float64
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
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
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l4 = s0i32
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1i32 = f67(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i64
		s0i32 = l3
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = l5
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l6 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l7 = s0f32
		s0i32 = l3
		s1f32 = l1
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = s1f32
		s0i32 = l3
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1f32 = l6
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1f32 = l7
		s2f32 = l1
		s1f32 = s1f32 * s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = s1f32
		s0i32 = l0
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l5 = s0i64
		s0i32 = l3
		s1i32 = 1065353216
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = l5
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l3
		s2i32 = l3
		s3i32 = 48
		s2i32 = s2i32 + s3i32
		f970(ctx, s0i32, s1i32, s2i32)
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s0f64 = float64(s0f32)
		l8 = s0f64
		s1f64 = l8
		s0f64 = s0f64 * s1f64
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f64 = float64(s1f32)
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
		s2f64 = float64(s2f32)
		s1f64 = s1f64 * s2f64
		s0f64 = s0f64 / s1f64
		s0f64 = math.Sqrt(s0f64)
		s0f32 = float32(s0f64)
		l1 = s0f32
	}
	s0i32 = l3
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0f32 = l1
	return s0f32
}
