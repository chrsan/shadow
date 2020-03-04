package internal

import (
	"math"
	"unsafe"
)

func f415(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 float32
	_ = l15
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
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = -1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l4
		s1i32 = 13844
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = 13836
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l7 = s0i32
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = float32(s1i32)
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = s1f32
		s0i32 = l4
		s1i32 = l7
		s1f32 = float32(s1i32)
		s2f32 = 1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = s1f32
		s0i32 = l4
		s1i32 = l6
		s1f32 = float32(s1i32)
		s2f32 = -1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
		s0i32 = l4
		s1i32 = l5
		s1f32 = float32(s1i32)
		s2f32 = -1
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = s1f32
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 13844
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 13836
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
lbl1:
	s0i32 = l1
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s0i32 = l4
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	l14 = s0i32
	s0i32 = 0
	l1 = s0i32
lbl3:
	s0i32 = l0
	s1i32 = l1
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s0i32 = f335(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = 96
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = 80
		s2i32 = s2i32 + s3i32
		s0i32 = f335(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	}
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	s1f32 = 64
	s0f32 = s0f32 * s1f32
	l15 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l15
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl6
	}
	s0i32 = -2147483648
lbl6:
	l5 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
	s1f32 = 64
	s0f32 = s0f32 * s1f32
	l15 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l15
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl8
	}
	s0i32 = -2147483648
lbl8:
	l6 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	s1f32 = 64
	s0f32 = s0f32 * s1f32
	l15 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l15
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl10
	}
	s0i32 = -2147483648
lbl10:
	l7 = s0i32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	s1f32 = 64
	s0f32 = s0f32 * s1f32
	l15 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l15
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl12
	}
	s0i32 = -2147483648
lbl12:
	l9 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l4
	s1i32 = l5
	s2i32 = l7
	s3i32 = l5
	s4i32 = l7
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 6
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	s2i32 = l9
	s3i32 = l6
	s4i32 = l9
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 6
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	s2i32 = l7
	s3i32 = l7
	s4i32 = l5
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 63
	s1i32 = s1i32 + s2i32
	s2i32 = 6
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l11 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	s2i32 = l9
	s3i32 = l9
	s4i32 = l6
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 63
	s1i32 = s1i32 + s2i32
	s2i32 = 6
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l12 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l12
	s1i32 = l10
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l11
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 | s1i32
	s0i64 = int64(s0i32)
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l4
	s1i32 = l2
	s2i32 = l4
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l8 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	l10 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	l11 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
	l12 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l8
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l11
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l10
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l12
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
lbl15:
	s0i32 = l4
	s1i32 = l2
	s2i32 = l4
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	s0i32 = f111(ctx, s0i32, s1i32, s2i32)
	l8 = s0i32
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 != 0 {
		goto lbl4
	}
lbl16:
	s0i32 = l9
	s1i32 = l7
	s2i32 = l6
	s3i32 = l5
	s4i32 = l14
	s5i32 = l3
	f315(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l8
	f110(ctx, s0i32)
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	goto lbl4
lbl14:
	s0i32 = l9
	s1i32 = l7
	s2i32 = l6
	s3i32 = l5
	s4i32 = 0
	s5i32 = l3
	f315(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
lbl4:
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l13
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl0:
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
