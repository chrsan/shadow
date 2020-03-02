package internal

import (
	"unsafe"
)

func f2112(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 float32
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
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l4 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l2 = s2f32
	s3f32 = l2
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	l8 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l3 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
	l5 = s3f32
	s2f32 = s2f32 - s3f32
	s3f32 = 3
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l7 = s2f32
	s1f32 = s1f32 - s2f32
	l9 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l7
	s2f32 = l5
	s3f32 = l3
	s4f32 = l3
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 + s2f32
	l5 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	s3f32 = l2
	s4f32 = l4
	s3f32 = s3f32 - s4f32
	s4f32 = 3
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l6
	s2f32 = s2f32 - s3f32
	l4 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l2
	s2f32 = l6
	s1f32 = s1f32 - s2f32
	l2 = s1f32
	s2f32 = l9
	s1f32 = s1f32 * s2f32
	s2f32 = l3
	s3f32 = l7
	s2f32 = s2f32 - s3f32
	l3 = s2f32
	s3f32 = l4
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = l2
	s3f32 = l5
	s2f32 = s2f32 * s3f32
	s3f32 = l3
	s4f32 = l8
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	s3i32 = l1
	s0i32 = f122(ctx, s0f32, s1f32, s2f32, s3i32)
	return s0i32
}
