package internal

import (
	"unsafe"
)

func f1546(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 float32
	_ = l16
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s10i32 int32
	_ = s10i32
	var s11i32 int32
	_ = s11i32
	var s12i32 int32
	_ = s12i32
	var s13i32 int32
	_ = s13i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l15 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s5i32 = l5
	s6i32 = l6
	s7i32 = l7
	s8i32 = l8
	s9i32 = l9
	s10i32 = l10
	s11i32 = l11
	s12i32 = l12
	s13i32 = l13
	s0i32 = f383(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32, s13i32)
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l16 = s0f32
		s1f32 = 1
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l15
			s1f32 = l16
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l2 = s0i32
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l2
				s1f32 = l16
				*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
				s0i32 = l6
				s1i32 = l2
				s2i32 = 4
				s1i32 = s1i32 + s2i32
				l14 = s1i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
				goto lbl2
			}
			s0i32 = l6
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			s1i32 = l15
			s2i32 = 12
			s1i32 = s1i32 + s2i32
			f39(ctx, s0i32, s1i32)
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l14 = s0i32
		lbl2:
			s0i32 = l1
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l14
			s3i32 = l6
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
			s2i32 = s2i32 - s3i32
			s3i32 = -4
			s2i32 = s2i32 + s3i32
			s0i32 = f47(ctx, s0i32, s1i32, s2i32)
			l2 = s0i32
			s0i32 = l10
			s1i32 = l1
			s2i32 = l10
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s3i32 = l2
			s1i32 = f31(ctx, s1i32, s2i32, s3i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l11
			s1i32 = l1
			s2i32 = l11
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s3i32 = l2
			s1i32 = f31(ctx, s1i32, s2i32, s3i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l12
			s1i32 = l1
			s2i32 = l12
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s3i32 = l2
			s1i32 = f31(ctx, s1i32, s2i32, s3i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l13
			s1i32 = l1
			s2i32 = l13
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s3i32 = l2
			s1i32 = f31(ctx, s1i32, s2i32, s3i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		s1i32 = l1
		s2i32 = l5
		s3i32 = l6
		s4i32 = l7
		s5i32 = l10
		s6i32 = l11
		s7i32 = l12
		s8i32 = l13
		s0i32 = f1313(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32, s8i32)
		l14 = s0i32
	}
	s0i32 = l15
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l14
	return s0i32
}
