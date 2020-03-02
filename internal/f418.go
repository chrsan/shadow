package internal

import (
	"math"
	"unsafe"
)

func f418(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
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
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l6 = s0i32
		s0i32 = l5
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l5
		s1i32 = l6
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l14 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l15 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l16 = s0f32
	s0i32 = l4
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f32 = float32(math.Floor(float64(s1f32)))
	l17 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l17
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
	l17 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l17
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
	l17 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l17
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl2
	}
	s1i32 = -2147483648
lbl2:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l16
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l16 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l16
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
	l16 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l16
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
	l16 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l16
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl4
	}
	s1i32 = -2147483648
lbl4:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l15
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l15 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l15
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
	l15 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l15
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
	l15 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l15
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	l6 = s0i32
	s0i32 = l4
	s1f32 = l14
	s1f32 = float32(math.Floor(float64(s1f32)))
	l14 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l14
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
	l14 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l14
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl8
	}
	s1i32 = -2147483648
lbl8:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 2305843005455597567
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -2305843001160630271
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])))
	s1i32 = l4
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)])))
	s0i64 = s0i64 - s1i64
	l12 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l4
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])))
	s1i32 = l4
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)])))
	s0i64 = s0i64 - s1i64
	l13 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i64 = l12
	s1i64 = l13
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967296
	if uint64(s0i64) < uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
lbl11:
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l1
	f365(ctx, s0i32, s1i32)
	goto lbl0
lbl10:
	s0i32 = l6
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		goto lbl12
	}
	s0i32 = l4
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl12:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
	l5 = s0i32
	s1i32 = 18
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 18
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
	l5 = s1i32
	s2i32 = 18
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 18
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l5
	s1i32 = s1i32 - s2i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+100)]))
	l5 = s2i32
	s3i32 = 18
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 18
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s3i32 = l5
	s2i32 = s2i32 - s3i32
	s3i32 = l4
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+96)]))
	l5 = s3i32
	s4i32 = 18
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s4i32 = 18
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	s4i32 = l5
	s3i32 = s3i32 - s4i32
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		f261(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 72
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	l7 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = 32767
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = 32768
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
	}
	s0i32 = l4
	s1i32 = 13828
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 13820
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l1
	s2i32 = l4
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = 1
	f268(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l7
	l1 = s0i32
lbl15:
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = l1
	s3i32 = l4
	s4i32 = 112
	s3i32 = s3i32 + s4i32
	s4i32 = 0
	s5i32 = 0
	s0i32 = f407(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	l9 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l2
		s1i32 = l1
		f365(ctx, s0i32, s1i32)
		goto lbl17
	}
	s0i32 = l6
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l4
		s2i32 = 112
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		f1771(ctx, s0i32, s1i32, s2i32)
	}
	s0f32 = -1
	l14 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l11 = s0i32
	s1i32 = 8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l2
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l5
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l4
	s1i32 = 136
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = 1
	l10 = s0i32
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 1
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l14 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 2
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l15 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 3
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l16 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 4
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l17 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 5
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l18 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 6
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l19 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 7
	f155(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 - s1f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+140)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
	s1f32 = s1f32 - s2f32
	s0f32 = f62(ctx, s0f32, s1f32)
	l20 = s0f32
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0f32 = l20
	s1f32 = l19
	s2f32 = l18
	s3f32 = l17
	s4f32 = l16
	s5f32 = l15
	s6f32 = l14
	s7f32 = 0
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 7
	s0f32 = s0f32 / s1f32
	l14 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0f32 = s0f32 - s1f32
		l15 = s0f32
		s1f32 = l15
		s0f32 = s0f32 * s1f32
		l15 = s0f32
		s0i32 = l2
		l5 = s0i32
		goto lbl22
	}
	s0i32 = l2
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l2
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = l5
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	l15 = s0f32
	s1f32 = l15
	s0f32 = s0f32 * s1f32
	l15 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l5 = s0i32
		goto lbl22
	}
	s0i32 = l2
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l5 = s0i32
	s0i32 = l2
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = l5
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l10 = s0i32
lbl22:
	s0i32 = l11
	s0f32 = float32(s0i32)
	l16 = s0f32
	s1f32 = l16
	s0f32 = s0f32 * s1f32
	s1f32 = l14
	s2f32 = l14
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l15
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s3i32 = l2
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s2f32 = s2f32 - s3f32
	l14 = s2f32
	s3f32 = l14
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 / s1f32
	l14 = s0f32
	s0i32 = l10
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l5
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l2 = s0i32
		s0i32 = l5
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l5
		s1i32 = l2
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0f32 = -1
	s1f32 = l14
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s3i32 = l5
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 / s2f32
	l14 = s1f32
	s2f32 = l14
	s3f32 = l14
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l14 = s0f32
lbl20:
	s0i32 = 28844
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl28
	}
	s0i32 = 24184
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = 0
	s3i32 = 0
	s0i32 = f333(ctx, s0i32, s1i32, s2i32, s3i32)
	if s0i32 != 0 {
		goto lbl28
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s0f32 = float32(s0i32)
	l15 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l5 = s0i32
		s0i32 = l2
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l5
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 - s1f32
	s1f32 = l15
	if s0f32 <= s1f32 {
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
		goto lbl27
	}
	s0f32 = l14
	s1f32 = 0.25
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl27
	}
lbl28:
	s0i32 = l0
	s1i32 = l8
	s2i32 = l4
	s3i32 = 112
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s4i32 = l3
	f1815(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	goto lbl26
lbl27:
	s0i32 = l0
	s1i32 = l8
	s2i32 = l4
	s3i32 = 112
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s4i32 = l3
	f1799(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl26:
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l8
	s1i32 = l4
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	f1770(ctx, s0i32, s1i32, s2i32)
lbl17:
	s0i32 = l9
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l9
	s0i32 = f35(ctx, s0i32)
	s0i32 = l7
	f40(ctx, s0i32)
lbl0:
	s0i32 = l4
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
