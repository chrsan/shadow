package internal

import (
	"unsafe"
)

func f130(ctx *Context, l0 int32, l1 float32, l2 float32, l3 float32, l4 float32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		if s0i32 != 0 {
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
			s1i32 = l5
			s2i32 = -1
			s1i32 = s1i32 ^ s2i32
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l9 = s0f32
			s0i32 = l5
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l8 = s0f32
		}
		s0i32 = l6
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		l5 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		l5 = s0i32
		s1f32 = l9
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l5
		s1f32 = l8
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	}
	s0i32 = l6
	s1i32 = l0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 2
	s2f32 = 0
	s0i32 = f49(ctx, s0i32, s1i32, s2f32)
	l5 = s0i32
	s1f32 = l3
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l5
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l5
	s1f32 = l1
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l5
	s1f32 = l4
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l0
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l6
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
