package internal

import (
	"unsafe"
)

func f554(ctx *Context, l0 int32, l1 int32, l2 float32, l3 int32) int32 {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
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
	var s1i64 int64
	_ = s1i64
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 180
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2f32 = l2
	s3i32 = l5
	s4i32 = 12
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 8
	s4i32 = s4i32 + s5i32
	s5i32 = l5
	s6i32 = 4
	s5i32 = s5i32 + s6i32
	s0i32 = f1407(ctx, s0i32, s1i32, s2f32, s3i32, s4i32, s5i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		goto lbl0
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 52
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l0
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l0
		s1i32 = 28
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)]))
		l2 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)]))
		l12 = s0f32
	lbl3:
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)]))
		l16 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)]))
		l13 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l14 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l15 = s0f32
		s0i32 = l10
		s0i32 = f45(ctx, s0i32)
		l4 = s0i32
		s1f32 = l13
		s2f32 = l2
		s3f32 = l14
		s2f32 = s2f32 * s3f32
		s3f32 = l12
		s4f32 = l15
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		l13 = s2f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l4
		s1f32 = l16
		s2f32 = l12
		s3f32 = l14
		s2f32 = s2f32 * s3f32
		s3f32 = l2
		s4f32 = l15
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 - s3f32
		l12 = s2f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l9
		s0i32 = f88(ctx, s0i32)
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)]))
		l11 = s0i32
		s0i32 = l0
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])))
		l4 = s0i32
		s0i32 = l8
		s1i32 = 3
		s0i32 = f89(ctx, s0i32, s1i32)
		l7 = s0i32
		s1i32 = l4
		s2i32 = -2
		s1i32 = s1i32 + s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
		s0i32 = l7
		s1i32 = l4
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
		s0i32 = l7
		s1i32 = l11
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0f32 = l13
		l2 = s0f32
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)]))
	l2 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l13 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)]))
	l12 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l14 = s0f32
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	s0i32 = f45(ctx, s0i32)
	l3 = s0i32
	s1f32 = l2
	s2f32 = l14
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l3
	s1f32 = l12
	s2f32 = l13
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s0i32 = f88(ctx, s0i32)
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])))
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)]))
	l6 = s0i32
	s0i32 = l0
	s1i32 = 52
	s0i32 = s0i32 + s1i32
	s1i32 = 3
	s0i32 = f89(ctx, s0i32, s1i32)
	l4 = s0i32
	s1i32 = l3
	s2i32 = -2
	s1i32 = s1i32 + s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
	s0i32 = l4
	s1i32 = l3
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
	s0i32 = l4
	s1i32 = l6
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	return s0i32
}
