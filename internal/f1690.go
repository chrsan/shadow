package internal

import (
	"unsafe"
)

func f1690(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+52)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
		s3i32 = l3
		s4i32 = l2
		s5i32 = l2
		s6i32 = 24
		s5i32 = s5i32 + s6i32
		f215(ctx, s0i32, s1i32, s2f32, s3i32, s4i32, s5i32)
		s0i32 = l2
		s1i32 = 1
		ctx.Mem[int(s0i32+52)] = uint8(s1i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+53)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
		s3i32 = l3
		s4i32 = l2
		s5i32 = 16
		s4i32 = s4i32 + s5i32
		s5i32 = l2
		s6i32 = 32
		s5i32 = s5i32 + s6i32
		f215(ctx, s0i32, s1i32, s2f32, s3i32, s4i32, s5i32)
		s0i32 = l2
		s1i32 = 1
		ctx.Mem[int(s0i32+53)] = uint8(s1i32)
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = 0
	s0i32 = f310(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s3i32 = l3
		s4i32 = 8
		s3i32 = s3i32 | s4i32
		s4i32 = l3
		s5i32 = 0
		f215(ctx, s0i32, s1i32, s2f32, s3i32, s4i32, s5i32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = l3
		s3i32 = l2
		s0i32 = f399(ctx, s0i32, s1i32, s2i32, s3i32)
		l4 = s0i32
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
