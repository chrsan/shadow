package internal

import (
	"unsafe"
)

func f970(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
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
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
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
	s0i32 = 3
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l4 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = f24(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
	lbl3:
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l8 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l10 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l11 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l12 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l13 = s0f32
		s0i32 = l1
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1f32
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
		s1f32 = s1f32 * s2f32
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l6 = s2f32
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		l7 = s2f32
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+32)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l1
		s1f32 = l5
		s2f32 = l12
		s1f32 = s1f32 * s2f32
		s2f32 = l6
		s3f32 = l13
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l7
		s3f32 = l11
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l5
		s2f32 = l9
		s1f32 = s1f32 * s2f32
		s2f32 = l6
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l7
		s3f32 = l8
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l1
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		if s0i32 != 0 {
			goto lbl3
		}
		goto lbl1
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l1
	s1i32 = l2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = l2
	s2i32 = 36
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
lbl1:
}
