package internal

import (
	"unsafe"
)

func f441(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	var s6i32 int32
	_ = s6i32
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
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		f2114(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l3
		s2i32 = l2
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		l1 = s2i32
		f441(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = l3
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		f441(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	s4i32 = l1
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
	s5i32 = l1
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+24)]))
	s6i32 = l1
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+28)]))
	f175(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32, s6f32)
lbl0:
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
