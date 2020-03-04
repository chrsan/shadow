package internal

import (
	"math"
	"unsafe"
)

func f2012(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	s1i32 = 192
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 0
	ctx.Mem[int(s0i32+32)] = uint8(s1i32)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l8 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l9 = s0i32
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l9
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = s1f32
	s0i32 = l6
	s1i32 = l8
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = s1f32
	s0i32 = l6
	s1i32 = l2
	s2i32 = l6
	s3i32 = 136
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l8 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l8
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 152
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = 120
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	f978(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0f32
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l14 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l12 = s0f32
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l15 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
		l10 = s0f32
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
		l13 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0f32 = l10
		s1f32 = l11
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)]))
		l11 = s0f32
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
		l16 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0f32 = l16
		s1f32 = l15
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0f32 = l13
		s1f32 = l14
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0f32 = l11
		s1f32 = l12
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	lbl4:
		s0i32 = l6
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 136
		s1i32 = s1i32 + s2i32
		s0i32 = f135(ctx, s0i32, s1i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l6
		s1i32 = 152
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = 104
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s3i32 = 120
		s2i32 = s2i32 + s3i32
		s0i32 = f42(ctx, s0i32, s1i32, s2i32)
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
		s1f32 = 0
		s0f32 = s0f32 * s1f32
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+108)]))
		s0f32 = s0f32 * s1f32
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
		s0f32 = s0f32 * s1f32
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		s0f32 = s0f32 * s1f32
		l10 = s0f32
		s1f32 = l10
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l6
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)]))
		l13 = s0f32
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
		l10 = s0f32
	lbl3:
		s0i32 = 1
		l8 = s0i32
		s0f32 = l10
		s1f32 = l13
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+140)]))
		l11 = s0f32
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
		l14 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0f32
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l15 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0f32 = l12
		s1f32 = l10
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l10 = s0f32
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l12 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0f32 = l15
		s1f32 = l13
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		s1f32 = l10
		s2f32 = l11
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		s0i32 = s0i32 | s1i32
		s1f32 = l12
		s2f32 = l14
		if s1f32 >= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		s0i32 = s0i32 | s1i32
		l8 = s0i32
	lbl5:
		s0i32 = l5
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l8
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l4
		s0i32 = int32(ctx.Mem[int(s0i32+45)])
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl1
		}
	lbl6:
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
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
			goto lbl7
		}
		s1i32 = -2147483648
	lbl7:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
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
			goto lbl9
		}
		s1i32 = -2147483648
	lbl9:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
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
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l7
		s2i32 = l6
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s0i32 = f342(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 0
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		l1 = s2i32
		s3i32 = 1
		if s2i32 < s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = 152
			s0i32 = s0i32 + s1i32
			s1i32 = l2
			s2i32 = 0
			s3i32 = l5
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s1f32 = float32(s1i32)
			s2i32 = l1
			s3i32 = 0
			s4i32 = l1
			s5i32 = 0
			if s4i32 > s5i32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2i32 = s2i32
			} else {
				s2i32 = s3i32
			}
			s2f32 = float32(s2i32)
			f200(ctx, s0i32, s1f32, s2f32)
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)]))
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)]))
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s1f32 = float32(s1i32)
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l6
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)]))
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s1f32 = float32(s1i32)
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
	lbl17:
		s0i32 = l7
		l1 = s0i32
		goto lbl1
	lbl16:
		s0i32 = l7
		l1 = s0i32
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)]))
		l2 = s0i32
		s1i32 = 128
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = l6
			s2i32 = 152
			s1i32 = s1i32 + s2i32
			s1i32 = f24(ctx, s1i32)
			l2 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)])) = uint32(s1i32)
		}
		s0i32 = l2
		s1i32 = 14
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l6
	s3i32 = 152
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s4i32 = l4
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+244)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	goto lbl0
lbl1:
	s0i32 = l6
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = l1
	s3i32 = l6
	s4i32 = 152
	s3i32 = s3i32 + s4i32
	f300(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s0i32 = f46(ctx, s0i32, s1i32)
	l2 = s0i32
	l1 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -193
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l6
	f83(ctx, s0i32, s1i32)
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl20:
	s0i32 = l0
	s1i32 = l3
	s2i32 = l2
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+80)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s0i32 = f23(ctx, s0i32)
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl0:
	s0i32 = l7
	s0i32 = f41(ctx, s0i32)
	s0i32 = l6
	s1i32 = 192
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
