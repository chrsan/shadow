package internal

import (
	"math"
	"unsafe"
)

func f687(ctx *Context, l0 int32, l1 int32) {
	var l2 int64
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float64
	_ = l17
	var l18 float64
	_ = l18
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l4 = s0f32
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	l9 = s0f32
	s0f32 = 1
	s1f32 = l4
	s2f32 = 1
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 / s1f32
	l3 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l10 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l11 = s2f32
	s3f32 = l4
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
	l14 = s4f32
	s3f32 = s3f32 * s4f32
	l6 = s3f32
	s4f32 = l6
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	l5 = s0f32
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1f32 = l3
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l12 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l13 = s3f32
	s4f32 = l4
	s5i32 = l0
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
	l15 = s5f32
	s4f32 = s4f32 * s5f32
	l7 = s4f32
	s5f32 = l7
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l8 = s1f32
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s1f32 = l16
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = 1
		s1f32 = l4
		s1f64 = float64(s1f32)
		l17 = s1f64
		s2f64 = 1
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 / s1f64
		s1f64 = 0.5
		s0f64 = s0f64 * s1f64
		l18 = s0f64
		s1f64 = l17
		s2f64 = l17
		s1f64 = s1f64 + s2f64
		l17 = s1f64
		s2f32 = l15
		s2f64 = float64(s2f32)
		s1f64 = s1f64 * s2f64
		s2f32 = l13
		s2f64 = float64(s2f32)
		s1f64 = s1f64 + s2f64
		s2f32 = l12
		s2f64 = float64(s2f32)
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 * s1f64
		s0f32 = float32(s0f64)
		l8 = s0f32
		s0f64 = l18
		s1f64 = l17
		s2f32 = l14
		s2f64 = float64(s2f32)
		s1f64 = s1f64 * s2f64
		s2f32 = l11
		s2f64 = float64(s2f32)
		s1f64 = s1f64 + s2f64
		s2f32 = l10
		s2f64 = float64(s2f32)
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 * s1f64
		s0f32 = float32(s0f64)
		l5 = s0f32
	}
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i64
	s0i32 = l1
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l1
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l1
	s1i64 = l2
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l1
	s1f32 = l5
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l1
	s1f32 = l3
	s2f32 = l11
	s3f32 = l6
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l3
	s3f32 = l13
	s4f32 = l7
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 * s3f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l3
	s2f32 = l6
	s3f32 = l10
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	s1i64 = int64(uint32(s1i32))
	s2f32 = l3
	s3f32 = l7
	s4f32 = l12
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 * s3f32
	s2i32 = int32(math.Float32bits(s2f32))
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i64
	s0i32 = l1
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0i32 = l1
	s1i64 = l2
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = l1
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
}
