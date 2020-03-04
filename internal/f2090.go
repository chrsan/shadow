package internal

import (
	"math"
	"unsafe"
)

func f2090(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
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
	var l14 float64
	_ = l14
	var l15 float64
	_ = l15
	var l16 float64
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l4 = s0i32
	s1i32 = 192
	s0i32 = s0i32 & s1i32
	s1i32 = 128
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f67(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s0f64 = float64(s0f32)
	l14 = s0f64
	s1f64 = l14
	s0f64 = s0f64 * s1f64
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s1f64 = float64(s1f32)
	l15 = s1f64
	s2f64 = l15
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	s0f64 = math.Sqrt(s0f64)
	l16 = s0f64
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 1
		l7 = s0f32
		goto lbl5
	}
	s0f64 = 1
	s1f64 = l16
	s0f64 = s0f64 / s1f64
	l16 = s0f64
	s1f64 = l14
	s0f64 = s0f64 * s1f64
	l14 = s0f64
	s0f32 = float32(s0f64)
	l8 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f64 = l16
	s1f64 = l15
	s0f64 = s0f64 * s1f64
	l15 = s0f64
	s0f32 = float32(s0f64)
	l6 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f64 = l16
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2f32 = -0.05
	s1f32 = s1f32 + s2f32
	s1f64 = float64(s1f32)
	s0f64 = s0f64 * s1f64
	s0f32 = float32(s0f64)
	l7 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f64 = l14
	s1f64 = 0
	if s0f64 != s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0f64 = l15
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
lbl5:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l4 = s0i32
		s0i32 = l1
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l4
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0f32 = l7
	s1f32 = l8
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l9 = s2f32
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	l10 = s3f32
	s4f32 = l8
	s5f32 = 0
	if s4f32 >= s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l4 = s4i32
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	s1f32 = s1f32 * s2f32
	s2f32 = l6
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l11 = s3f32
	s4i32 = l1
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)]))
	l12 = s4f32
	s5f32 = l6
	s6f32 = 0
	if s5f32 >= s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	l1 = s5i32
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l13 = s0f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	s1i32 = 2
	s2f32 = l13
	s3f32 = l7
	s4f32 = l8
	s5f32 = l10
	s6f32 = l9
	s7i32 = l4
	if s7i32 != 0 {
		// s5f32 = s5f32
	} else {
		s5f32 = s6f32
	}
	s4f32 = s4f32 * s5f32
	s5f32 = l6
	s6f32 = l12
	s7f32 = l11
	s8i32 = l1
	if s8i32 != 0 {
		// s6f32 = s6f32
	} else {
		s6f32 = s7f32
	}
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = 0
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i64 = 550821167104
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l6
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = s1f32
	s0i32 = l3
	s1f32 = l6
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = s1f32
	s0i32 = l3
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l3
	s1f32 = l7
	s2f32 = l6
	s2f32 = -s2f32
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = s1f32
	s0i32 = l3
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
	s0i32 = l3
	s1f32 = l8
	s1f32 = -s1f32
	l6 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
	s0i32 = l3
	s1f32 = l7
	s2f32 = l6
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = s1f32
	s0i32 = l3
	s1i32 = l3
	s2i32 = 88
	s1i32 = s1i32 + s2i32
	s1i32 = f24(ctx, s1i32)
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		goto lbl3
	}
	s0i32 = l3
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s0i32 = f84(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = f137(ctx)
	l1 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l0
	s0i32 = f97(ctx, s0i32)
	f12(ctx, s0i32)
lbl9:
	s0i32 = l2
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	goto lbl2
lbl4:
	s0i32 = f137(ctx)
	l1 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l0
	s0i32 = f97(ctx, s0i32)
	f12(ctx, s0i32)
lbl10:
	s0i32 = l2
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	goto lbl2
lbl3:
	s0i32 = f137(ctx)
	l1 = s0i32
	s0i32 = l3
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint16(s1i32)
	s0i32 = l3
	s1i32 = l3
	s1i32 = int32(ctx.Mem[int(s1i32+42)])
	s2i32 = 248
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = 48
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 32
	s2i32 = s2i32 + s3i32
	f138(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		l4 = s0i32
		goto lbl11
	}
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s0i32 = l1
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = l4
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl11:
	s0i32 = l4
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = f137(ctx)
		l1 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s0i32 = l2
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l0
		s0i32 = f97(ctx, s0i32)
		f12(ctx, s0i32)
	lbl15:
		s0i32 = l2
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 252
		s1i32 = s1i32 & s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		goto lbl13
	}
	s0i32 = l3
	s1i64 = 9187343237679939583
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 4286578687
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	s3i32 = 718
	s4i32 = l3
	f685(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l2
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+10)])
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = l3
	s2i32 = 88
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	f138(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		l4 = s0i32
		goto lbl16
	}
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l4
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl16:
	s0i32 = l4
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = f137(ctx)
	l1 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l0
	s0i32 = f97(ctx, s0i32)
	f12(ctx, s0i32)
lbl18:
	s0i32 = l2
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
lbl13:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = f97(ctx, s0i32)
	f12(ctx, s0i32)
lbl2:
	s0i32 = 1
	l5 = s0i32
lbl1:
	s0i32 = l3
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
