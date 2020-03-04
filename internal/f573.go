package internal

import (
	"unsafe"
)

func f573(ctx *Context, l0 int32, l1 int32, l2 float64) {
	var l3 float64
	_ = l3
	var l4 float64
	_ = l4
	var l5 float64
	_ = l5
	var l6 float64
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
	_ = l8
	var l9 float64
	_ = l9
	var l10 float64
	_ = l10
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
	s0f64 = l2
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		return
	}
	s0f64 = l2
	s1f64 = 1
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		return
	}
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0f64
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l6 = s0f64
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0f64
	s0i32 = l1
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f64
	s0i32 = l0
	s1f64 = 1
	s2f64 = l2
	s1f64 = s1f64 - s2f64
	l3 = s1f64
	s2f64 = l3
	s3f64 = l3
	s2f64 = s2f64 * s3f64
	l4 = s2f64
	s1f64 = s1f64 * s2f64
	l9 = s1f64
	s2i32 = l1
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f64 = s1f64 * s2f64
	s2f64 = l4
	s3f64 = 3
	s2f64 = s2f64 * s3f64
	s3f64 = l2
	s2f64 = s2f64 * s3f64
	l4 = s2f64
	s3i32 = l1
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l2
	s3f64 = l2
	s2f64 = s2f64 * s3f64
	l10 = s2f64
	s3f64 = l3
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 * s3f64
	l3 = s2f64
	s3i32 = l1
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l10
	s3f64 = l2
	s2f64 = s2f64 * s3f64
	l2 = s2f64
	s3i32 = l1
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+56)]))
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = l0
	s1f64 = l9
	s2f64 = l8
	s1f64 = s1f64 * s2f64
	s2f64 = l4
	s3f64 = l7
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l3
	s3f64 = l6
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	s2f64 = l2
	s3f64 = l5
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
}
