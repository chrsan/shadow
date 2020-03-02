package internal

import (
	"math"
	"unsafe"
)

func f160(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 3456
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3448)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3440)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3432)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3424)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+3416)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l6 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 3416
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = l3
		f69(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l6 = s0i32
		s0i32 = l5
		s1i32 = 3416
		s0i32 = s0i32 + s1i32
	} else {
		s0i32 = l6
	}
	l7 = s0i32
	s0i32 = l2
	s1i32 = l6
	s2i32 = l5
	s3i32 = 3408
	s2i32 = s2i32 + s3i32
	s0i32 = f1132(ctx, s0i32, s1i32, s2i32)
	l6 = s0i32
	s1i32 = 3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s3i32 = l7
		f505(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1i32 = l5
	s2i32 = 3392
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = l1
	s4i32 = l3
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = 2
	f55(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3392)]))
	l8 = s0f32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3400)]))
	l9 = s1f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l9
		l10 = s0f32
		s0f32 = l8
		l9 = s0f32
		goto lbl3
	}
	s0i32 = l5
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3400)])) = s1f32
	s0i32 = l5
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3392)])) = s1f32
	s0f32 = l8
	l10 = s0f32
lbl3:
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3396)]))
	l12 = s0f32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3404)]))
	l8 = s1f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l8
		l11 = s0f32
		s0f32 = l12
		l8 = s0f32
		goto lbl5
	}
	s0i32 = l5
	s1f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3404)])) = s1f32
	s0i32 = l5
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3396)])) = s1f32
	s0f32 = l12
	l11 = s0f32
lbl5:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l12 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l11
		s1f32 = 1
		s0f32 = s0f32 + s1f32
		l11 = s0f32
		s0f32 = l10
		s1f32 = 1
		s0f32 = s0f32 + s1f32
		l10 = s0f32
		s0f32 = l8
		s1f32 = -1
		s0f32 = s0f32 + s1f32
		l8 = s0f32
		s0f32 = l9
		s1f32 = -1
		s0f32 = s0f32 + s1f32
		l9 = s0f32
		goto lbl7
	}
	s0f32 = l11
	s1i32 = l6
	s2i32 = 2
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3408)]))
		l12 = s1f32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3412)]))
		goto lbl9
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l3 = s1i32
	s1i32 = l5
	s2f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3380)])) = s2f32
	s1i32 = l5
	s2f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3376)])) = s2f32
	s1i32 = l3
	s2i32 = l5
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 3376
	s3i32 = s3i32 + s4i32
	s4i32 = 1
	f238(ctx, s1i32, s2i32, s3i32, s4i32)
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f32 = float32(math.Abs(float64(s1f32)))
	l12 = s1f32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1f32 = float32(math.Abs(float64(s1f32)))
lbl9:
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l13 = s1f32
	s0f32 = s0f32 + s1f32
	l11 = s0f32
	s0f32 = l10
	s1f32 = l12
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l12 = s1f32
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s0f32 = l8
	s1f32 = l13
	s0f32 = s0f32 - s1f32
	l8 = s0f32
	s0f32 = l9
	s1f32 = l12
	s0f32 = s0f32 - s1f32
	l9 = s0f32
lbl7:
	s0f32 = l11
	s1f32 = 8.5070587e+37
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l10
	s1f32 = 8.5070587e+37
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l8
	s1f32 = -8.5070587e+37
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l9
	s1f32 = -8.5070587e+37
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 1
	s1f32 = l9
	s1f32 = -s1f32
	s2f32 = l9
	s3f32 = l9
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = 32767
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl11
	}
	s0i32 = 1
	s1f32 = l8
	s1f32 = -s1f32
	s2f32 = l8
	s3f32 = l8
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = 32767
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl11
	}
	s0i32 = 1
	s1f32 = l10
	s1f32 = -s1f32
	s2f32 = l10
	s3f32 = l10
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	s2f32 = 32767
	if s1f32 <= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl11
	}
	s0f32 = l11
	s0f32 = -s0f32
	s1f32 = l11
	s2f32 = l11
	s3f32 = 0
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
	s1f32 = 32767
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
lbl11:
	l3 = s0i32
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l7
	f505(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl0
lbl12:
	s0i32 = l5
	s1f32 = l11
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l11 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l11
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l11 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l11
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l11 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l11
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl13
	}
	s1i32 = -2147483648
lbl13:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3388)])) = uint32(s1i32)
	s0i32 = l5
	s1f32 = l10
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l10 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l10
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l10 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l10
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
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
		goto lbl15
	}
	s1i32 = -2147483648
lbl15:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3384)])) = uint32(s1i32)
	s0i32 = l5
	s1f32 = l8
	s1f32 = float32(math.Floor(float64(s1f32)))
	l8 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l8
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l8 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l8
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l8 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l8
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl17
	}
	s1i32 = -2147483648
lbl17:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3380)])) = uint32(s1i32)
	s0i32 = l5
	s1f32 = l9
	s1f32 = float32(math.Floor(float64(s1f32)))
	l9 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l9
	s4f32 = 2.1474835e+09
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l9 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l9
	s4f32 = -2.1474835e+09
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l9 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l9
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl19
	}
	s1i32 = -2147483648
lbl19:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3376)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	l1 = s1i32
	s2i32 = l1
	s3i32 = 20
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s3i32 = int32(ctx.Mem[int(s3i32+40)])
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l5
	s3i32 = 3376
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 3344
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	s2i32 = 3332
	s3i32 = 3332
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l3 = s0i32
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l7
	s3i32 = l2
	s0i32 = f182(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l6
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l0 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl24
	case 1:
		goto lbl23
	default:
		goto lbl22
	}
lbl24:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 3392
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l1
		f1792(ctx, s0i32, s1i32, s2i32)
		goto lbl21
	}
	s0i32 = l5
	s1i32 = 3392
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l1
	f219(ctx, s0i32, s1i32, s2i32)
	goto lbl21
lbl23:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 3392
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 3408
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = l1
		f1791(ctx, s0i32, s1i32, s2i32, s3i32)
		goto lbl21
	}
	s0i32 = l5
	s1i32 = 3392
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 3408
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s3i32 = l1
	f1773(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl21
lbl22:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 3392
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l1
		f1794(ctx, s0i32, s1i32, s2i32)
		goto lbl21
	}
	s0i32 = l5
	s1i32 = 3392
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l1
	f1782(ctx, s0i32, s1i32, s2i32)
lbl21:
	s0i32 = l3
	f43(ctx, s0i32)
lbl0:
	s0i32 = l5
	s1i32 = 3456
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
