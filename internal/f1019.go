package internal

import (
	"math"
	"unsafe"
)

func f1019(ctx *Context, l0 float32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l6
	s1f32 = l0
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l0 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f32
	s0i32 = l6
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l9 = s2f32
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l6
	s1f32 = l0
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0f32 = l9
	s0f32 = float32(math.Floor(float64(s0f32)))
	l0 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l0
	s3f32 = 2.1474835e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l0 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l0
	s3f32 = -2.1474835e+09
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l0 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	l7 = s0i32
	s0f32 = l8
	s0f32 = float32(math.Floor(float64(s0f32)))
	l0 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l0
	s3f32 = 2.1474835e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l0 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l0
	s3f32 = -2.1474835e+09
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l0 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl2
	}
	s0i32 = -2147483648
lbl2:
	l1 = s0i32
	s0f32 = 0
	l0 = s0f32
	s0i32 = l3
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l10 = s1f32
		s0f32 = s0f32 - s1f32
		l11 = s0f32
		s0f32 = l8
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l0 = s1f32
		s0f32 = s0f32 - s1f32
		l12 = s0f32
		s0f32 = l9
		s1f32 = l10
		s0f32 = s0f32 - s1f32
		l10 = s0f32
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1f32 = l0
		s0f32 = s0f32 - s1f32
		l0 = s0f32
	}
	s0i32 = l6
	s1f32 = l10
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
	s0i32 = l6
	s1f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
	s0i32 = l6
	s1f32 = l11
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
	s0i32 = l6
	s1f32 = l0
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l6
	s1f32 = l9
	s2i32 = l7
	s2f32 = float32(s2i32)
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = s1f32
	s0i32 = l6
	s1f32 = l8
	s2i32 = l1
	s2f32 = float32(s2i32)
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0i32 = l6
	s1i32 = 28024
	s2i64 = 0
	s3i32 = 40
	f222(ctx, s0i32, s1i32, s2i64, s3i32)
	s0i32 = 104
	s0i32 = f17(ctx, s0i32)
	l1 = s0i32
	s1i32 = 21620
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1
	f201(ctx, s0i32, s1i32)
	s0i32 = l1
	f424(ctx, s0i32)
	s0i32 = l6
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
