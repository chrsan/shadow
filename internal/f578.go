package internal

import (
	"unsafe"
)

func f578(ctx *Context, l0 int32, l1 int32, l2 float32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
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
	var s6i32 int32
	_ = s6i32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = l5
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l4 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1f32 = l2
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1i32 = l4
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	f39(ctx, s0i32, s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l4 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s2i32 = s2i32 - s3i32
	s3i32 = -4
	s2i32 = s2i32 + s3i32
	s0i32 = f47(ctx, s0i32, s1i32, s2i32)
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = l5
	s1f32 = 1
	s2f32 = l2
	s1f32 = s1f32 / s2f32
	l2 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l4 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1f32 = l2
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1i32 = l4
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl2
	}
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	f39(ctx, s0i32, s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l4 = s0i32
lbl2:
	s0i32 = l7
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l4
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s2i32 = s2i32 - s3i32
	s3i32 = -4
	s2i32 = s2i32 + s3i32
	s0i32 = f47(ctx, s0i32, s1i32, s2i32)
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s1i32 = l1
	s2i32 = l0
	s3i32 = l0
	s4i32 = l0
	s5i32 = l1
	s6i32 = l3
	s4i32 = f31(ctx, s4i32, s5i32, s6i32)
	s3i32 = f213(ctx, s3i32, s4i32)
	s4i32 = l6
	s2i32 = f31(ctx, s2i32, s3i32, s4i32)
	s0i32 = f36(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
