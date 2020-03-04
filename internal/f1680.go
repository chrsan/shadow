package internal

import (
	"math"
	"unsafe"
)

func f1680(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float32, l6 float32, l7 int32, l8 int32) {
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var l12 int64
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
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
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	ctx.g0 = s0i32
	s0i32 = 1
	l10 = s0i32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l13 = s0f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l14 = s1f32
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l15 = s0f32
	s1f32 = 0
	if s0f32 >= s1f32 {
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
		s0i32 = l9
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l12 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0f32 = 1
		s1f32 = l15
		s0f32 = s0f32 - s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 0.00024414062
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
			goto lbl0
		}
		s0i32 = 2
		l10 = s0i32
		s0i32 = l9
		l4 = s0i32
		goto lbl5
	}
	s0i32 = l9
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l9
	l4 = s0i32
	s0f32 = l15
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
lbl5:
	s0f32 = l13
	s1i64 = l12
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	s1f32 = math.Float32frombits(uint32(s1i32))
	l16 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l14
	s2i64 = l12
	s2i32 = int32(s2i64)
	s2f32 = math.Float32frombits(uint32(s2i32))
	l17 = s2f32
	s1f32 = s1f32 * s2f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l11 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s1f32 = l16
		s1f32 = -s1f32
		l16 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l9
		s1f32 = l17
		s1f32 = -s1f32
		l17 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0f32 = l14
		s0f32 = -s0f32
		l14 = s0f32
		s0f32 = l13
		s0f32 = -s0f32
		l13 = s0f32
		s0i32 = l1
		l2 = s0i32
		s0i32 = l0
		l1 = s0i32
		s0i32 = l2
		l0 = s0i32
	}
	s0f32 = l6
	s1f32 = 0.70710677
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl9
	}
	s0f32 = l15
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l9
	s1f32 = l13
	s2f32 = l17
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	l6 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1f32 = l14
	s2f32 = l16
	s1f32 = s1f32 + s2f32
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	goto lbl8
lbl9:
	s0f32 = l15
	s1f32 = 1
	s0f32 = s0f32 + s1f32
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	l15 = s0f32
	s1f32 = l6
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l10
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s1f32 = l13
		s2f32 = l17
		s1f32 = s1f32 - s2f32
		l6 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l9
		s1f32 = l16
		s2f32 = l14
		s1f32 = s1f32 - s2f32
		l13 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l11
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l9
		s1f32 = l6
		s1f32 = -s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l9
		s1f32 = l13
		s1f32 = -s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		goto lbl10
	}
	s0i32 = l9
	s1f32 = l14
	s2f32 = l16
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l9
	s1f32 = l13
	s2f32 = l17
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
lbl10:
	s0i32 = l9
	s1f32 = l5
	s2f32 = l15
	s1f32 = s1f32 / s2f32
	s0i32 = f113(ctx, s0i32, s1f32)
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
lbl8:
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l13
	s0f32 = s0f32 + s1f32
	l13 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l6
	s0f32 = s0f32 + s1f32
	l6 = s0f32
	s0i32 = l7
	if s0i32 != 0 {
		s0i32 = l0
		s1f32 = l6
		s2f32 = l13
		f682(ctx, s0i32, s1f32, s2f32)
		goto lbl12
	}
	s0i32 = l0
	s1f32 = l6
	s2f32 = l13
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl12:
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0f32
	s0i32 = l9
	s1i32 = l9
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l9
	s1f32 = l6
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l8
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	goto lbl1
lbl4:
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0f32
	s0i32 = l9
	s1i32 = l9
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l9
	s1f32 = l6
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	goto lbl2
lbl3:
	s0i32 = l9
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0f32
	s0i32 = l9
	s1i32 = l9
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l9
	s1f32 = l6
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
lbl2:
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l9
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f32 = s1f32 + s2f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l9
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s2f32 = s2f32 + s3f32
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl1:
	s0i32 = l1
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l1
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l9
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f32 = s1f32 - s2f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l9
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s2f32 = s2f32 - s3f32
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl0:
	s0i32 = l9
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
