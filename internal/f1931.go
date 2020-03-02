package internal

import (
	"math"
	"unsafe"
)

func f1931(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 float32
	_ = l10
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
	var s4f32 float32
	_ = s4f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l7 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1i32 = f24(ctx, s1i32)
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l7
	s1i32 = 12
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s0i32 = f38(ctx, s0i32)
		l7 = s0i32
		s1i32 = l1
		s2i32 = 0
		f139(ctx, s0i32, s1i32, s2i32)
		s0i32 = l7
		s1i32 = l7
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 4
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l7
		s2i32 = l2
		s3i32 = l3
		s4i32 = l4
		s5i32 = l5
		f435(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l7
		f34(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l2
	s1i32 = l6
	s2i32 = l1
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0.125
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s1f32 = l10
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	s1f32 = 0.25
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 0.125
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s1f32 = l10
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	s1f32 = 0.25
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = 0.125
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s1f32 = l10
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	s1f32 = 0.25
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = 0.125
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	s1f32 = l10
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	s1f32 = 0.25
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = -1
	s0i32 = s0i32 ^ s1i32
	s1i32 = l5
	s0i32 = s0i32 & s1i32
	l5 = s0i32
lbl3:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l5
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l6
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
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
		goto lbl7
	}
	s1i32 = -2147483648
lbl7:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
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
		goto lbl9
	}
	s1i32 = -2147483648
lbl9:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
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
		goto lbl11
	}
	s1i32 = -2147483648
lbl11:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l6
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
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
		goto lbl13
	}
	s1i32 = -2147483648
lbl13:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l1
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
		s1i32 = l1
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		s0i64 = s0i64 - s1i64
		l8 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l1
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])))
		s1i32 = l1
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
		s0i64 = s0i64 - s1i64
		l9 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i64 = l8
		s1i64 = l9
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l6
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l6
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		goto lbl5
	}
	s0i32 = l4
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s1i32 = l0
	s2i32 = l6
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s0i32 = f129(ctx, s0i32)
	goto lbl4
lbl6:
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l0
		f245(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+40)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		ctx.Mem[int(s0i32+41)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l1
		s1i32 = f92(ctx, s1i32)
		ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	}
	s0i32 = l4
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l8 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l9 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i64 = l8
	s1i64 = l9
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l6
	s1i32 = l7
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l6
	s1i32 = l2
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l6
	s1i32 = l1
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l6
	s1i32 = l3
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l6
	s1i32 = l6
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s0i32 = f135(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
lbl17:
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = l4
	s3i32 = l5
	f1184(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl4
lbl5:
	s0i32 = l0
	s1i32 = l0
	s2i32 = l6
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	s3i32 = l4
	f268(ctx, s0i32, s1i32, s2i32, s3i32)
lbl4:
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+40)])
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		s3i32 = -1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		ctx.Mem[int(s1i32+41)] = uint8(s2i32)
		goto lbl20
	}
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	l1 = s2i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	ctx.Mem[int(s1i32+41)] = uint8(s2i32)
	s1i32 = l1
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl19
	}
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i32 = f92(ctx, s1i32)
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l1
		s1i32 = f81(ctx, s1i32, s2i32)
		s1i32 = l1
		f143(ctx, s1i32)
		s1i32 = l0
		s2i32 = 1
		ctx.Mem[int(s1i32+40)] = uint8(s2i32)
		goto lbl20
	}
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+40)])
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl19
	}
lbl20:
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	goto lbl18
lbl19:
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	s1i32 = f92(ctx, s1i32)
lbl18:
	ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
lbl1:
	s0i32 = l6
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
