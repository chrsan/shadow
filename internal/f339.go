package internal

import (
	"unsafe"
)

func f339(ctx *Context, l0 int32, l1 int32, l2 float32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	var s6f32 float32
	_ = s6f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l7 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l4 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0f32
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l3 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l6 = s2f32
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s2f32 = l8
	s1f32 = s1f32 + s2f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s3f32 = l3
	s4f32 = l3
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 - s3f32
	s3f32 = l6
	s2f32 = s2f32 + s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s1f32 = s1f32 * s2f32
	s2f32 = l6
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1f32 = l5
	s2f32 = l4
	s3f32 = l5
	s2f32 = s2f32 - s3f32
	l3 = s2f32
	s3f32 = l3
	s2f32 = s2f32 + s3f32
	s3f32 = l5
	s4f32 = l7
	s5f32 = l4
	s6f32 = l4
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l2
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
}
