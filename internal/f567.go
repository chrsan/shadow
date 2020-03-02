package internal

import (
	"unsafe"
)

func f567(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 float64
	_ = l2
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1f64
	s0f64 = s0f64 - s1f64
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l2 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	l3 = s2f64
	s1f64 = s1f64 - s2f64
	s2f64 = 3
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l3
	s2f64 = l4
	s3f64 = l2
	s2f64 = s2f64 - s3f64
	s3f64 = l2
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 + s2f64
	l3 = s1f64
	s2f64 = l3
	s1f64 = s1f64 + s2f64
	s2f64 = l2
	s3f64 = l4
	s2f64 = s2f64 - s3f64
	s3i32 = l1
	s0i32 = f566(ctx, s0f64, s1f64, s2f64, s3i32)
	return s0i32
}
