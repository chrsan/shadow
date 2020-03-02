package internal

import (
	"unsafe"
)

func f571(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 float64
	_ = l5
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
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = l1
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2f64 = 3
	s1f64 = s1f64 * s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l3
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2f64 = 3
	s1f64 = s1f64 * s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l4
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l1
	s1i32 = l1
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f64 = l5
	s3i32 = l3
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2f64 = s2f64 - s3f64
	s3i32 = l2
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 - s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l2
	s1i32 = l2
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3f64 = 3
	s2f64 = s2f64 * s3f64
	s3i32 = l3
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l5 = s3f64
	s4f64 = l5
	s3f64 = s3f64 + s4f64
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l3
	s1i32 = l3
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3f64 = 3
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
}
