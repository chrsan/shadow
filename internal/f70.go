package internal

import (
	"math"
	"unsafe"
)

func f70(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
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
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
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
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4408)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l3 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l4 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l5 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l6 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l7 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l8 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l9 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		l10 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l11 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l3 = s0i32
		s0i32 = l2
		s1i32 = 128
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l0
		s2i32 = 4412
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 40
		s2i32 = s2i32 + s3i32
		s0i32 = f1494(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l14 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l15 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l18 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l19 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l2
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	l20 = s1f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+52)]))
	l21 = s2f32
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l17 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	s2f32 = l15
	s3f32 = l16
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l16 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	s2f32 = l15
	s3f32 = l19
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l2
	s1f32 = l20
	s2f32 = l21
	s3f32 = l18
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l15 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l2
	s1f32 = l15
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l2
	s1f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l2
	s1f32 = l17
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l2
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	s2f32 = l14
	s3f32 = l16
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l4 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l2
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l15
	s3f32 = l17
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l5 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = l2
	s1i32 = 36
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = l0
	s1i32 = 4420
	s0i32 = s0i32 + s1i32
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0i64
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4412)]))
	l13 = s0i64
	s0i32 = l2
	s1i32 = l2
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = l12
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l13
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = 1
	s1i32 = l1
	s1f32 = math.Float32frombits(uint32(s1i32))
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s2i32 = l3
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	s2i32 = l4
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	s2i32 = l0
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	l14 = s1f32
	s2f32 = l14
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	s0i32 = f135(ctx, s0i32, s1i32)
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
lbl0:
	l0 = s0i32
	s0i32 = l2
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
