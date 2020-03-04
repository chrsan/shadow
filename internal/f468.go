package internal

import (
	"math"
	"unsafe"
)

func f468(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var s4i32 int32
	_ = s4i32
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
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+12)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+28)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+28)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+44)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+44)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+60)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+60)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+76)])
	l6 = s0i32
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l6
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+76)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+92)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+92)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 4
	f237(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+12)])
	s2i32 = 253
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	f237(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+28)])
	s2i32 = 253
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+28)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	f281(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+44)])
	s2i32 = 253
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+44)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	f281(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = 4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+60)])
	s2i32 = 253
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+60)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(s1i32)
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl0
	}
	s1i32 = -2147483648
lbl0:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l7 = s0i32
	s0i32 = l5
	s1i32 = l1
	s1f32 = float32(s1i32)
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl2
	}
	s1i32 = -2147483648
lbl2:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l7
	s1f32 = float32(s1i32)
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl4
	}
	s1i32 = -2147483648
lbl4:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l5 = s0i32
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f32 = float32(s1i32)
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = l5
	s1i32 = l2
	s1f32 = float32(s1i32)
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl8
	}
	s1i32 = -2147483648
lbl8:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l7
	s1f32 = float32(s1i32)
	l10 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l10
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl10
	}
	s1i32 = -2147483648
lbl10:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l6 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l7 = s2i32
	s2f32 = float32(s2i32)
	l12 = s2f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l6
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l8 = s3i32
	s2i32 = s2i32 - s3i32
	s2f32 = float32(s2i32)
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l6
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l5 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l9 = s2i32
	s2f32 = float32(s2i32)
	l13 = s2f32
	s1f32 = s1f32 + s2f32
	l10 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l5
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l2
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	l3 = s3i32
	s2i32 = s2i32 - s3i32
	s2f32 = float32(s2i32)
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l5
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l10
	} else {
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l10 = s1f32
		s2i32 = l6
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3f32 = l10
		s2f32 = s2f32 - s3f32
		s3f32 = l12
		s2f32 = s2f32 * s3f32
		s3i32 = l1
		s4i32 = l7
		s3i32 = s3i32 + s4i32
		s4i32 = l8
		s3i32 = s3i32 - s4i32
		s3f32 = float32(s3i32)
		s2f32 = s2f32 / s3f32
		s1f32 = s1f32 + s2f32
		l10 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l6
		s1f32 = l10
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l11 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	}
	s1f32 = l11
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l10 = s1f32
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3f32 = l10
		s2f32 = s2f32 - s3f32
		s3f32 = l13
		s2f32 = s2f32 * s3f32
		s3i32 = l2
		s4i32 = l9
		s3i32 = s3i32 + s4i32
		s4i32 = l3
		s3i32 = s3i32 - s4i32
		s3f32 = float32(s3i32)
		s2f32 = s2f32 / s3f32
		s1f32 = s1f32 + s2f32
		l10 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l5
		s1f32 = l10
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	}
	s0i32 = l0
	s1i64 = 38654705673
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l0
	return s0i32
}
