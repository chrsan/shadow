package internal

import (
	"unsafe"
)

func f965(ctx *Context, l0 int32, l1 float32, l2 float32, l3 int32) {
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
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
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0f32
	s0i32 = l3
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3f32 = l1
	s2f32 = s2f32 * s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = 1
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+32)]))
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
	s5f32 = l1
	s4f32 = s4f32 * s5f32
	s5i32 = l0
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+28)]))
	s6f32 = l2
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 + s4f32
	l4 = s3f32
	s2f32 = s2f32 / s3f32
	s3f32 = l4
	s4f32 = l4
	s5f32 = 0
	if s4f32 != s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l4 = s2f32
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l3
	s1f32 = l5
	s2f32 = l6
	s3f32 = l1
	s2f32 = s2f32 * s3f32
	s3f32 = l7
	s4f32 = l2
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l4
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
}
