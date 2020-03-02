package internal

import (
	"unsafe"
)

func f1680(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
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
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l6 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l7 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l5
	s0f32 = s0f32 + s1f32
	s1f32 = l6
	s0f32 = s0f32 - s1f32
	l8 = s0f32
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l0
		s1f32 = l8
		s2f32 = l7
		f681(ctx, s0i32, s1f32, s2f32)
		s0i32 = l0
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 - s2f32
		s2f32 = l6
		s1f32 = s1f32 - s2f32
		s2f32 = l5
		s3i32 = l1
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s4i32 = l2
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 + s3f32
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		return
	}
	s0i32 = l0
	s1f32 = l8
	s2f32 = l7
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	s2f32 = l6
	s1f32 = s1f32 - s2f32
	s2f32 = l5
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = l2
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 + s3f32
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
}
