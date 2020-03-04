package internal

import (
	"math"
	"unsafe"
)

func f1004(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var l15 int32
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
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
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
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
	s1i32 = 4640
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = 2
	l7 = s0i32
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l9 = s0i32
	s1i32 = -2
	s0i32 = s0i32 & s1i32
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l18 = s0f32
	s1f32 = -32767
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l19 = s0f32
	s1f32 = -32767
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l20 = s0f32
	s1f32 = 32767
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l20
	s1f32 = l18
	s0f32 = s0f32 - s1f32
	s1f32 = 32767
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l21 = s0f32
	s1f32 = 32767
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l21
	s1f32 = l19
	s0f32 = s0f32 - s1f32
	s1f32 = 32767
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1f32 = l21
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l21 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l21
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
	l21 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l21
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
	l21 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l21
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl1
	}
	s1i32 = -2147483648
lbl1:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l6
	s1f32 = l20
	s1f32 = float32(math.Ceil(float64(s1f32)))
	l20 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l20
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
	l20 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l20
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
	l20 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l20
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl3
	}
	s1i32 = -2147483648
lbl3:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l6
	s1f32 = l19
	s1f32 = float32(math.Floor(float64(s1f32)))
	l19 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l19
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
	l19 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l19
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
	l19 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l19
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl5
	}
	s1i32 = -2147483648
lbl5:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l6
	s1i64 = 4294967296
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l6
	s1f32 = l18
	s1f32 = float32(math.Floor(float64(s1f32)))
	l18 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l18
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
	l18 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l18
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
	l18 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l18
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl7
	}
	s1i32 = -2147483648
lbl7:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l18 = s0f32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+16)])
		if s0i32 != 0 {
			s0i32 = l3
			s1f32 = l18
			s0f32 = f180(ctx, s0i32, s1f32)
			l18 = s0f32
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l9 = s0i32
		}
		s0i32 = 0
		l7 = s0i32
		s0f32 = l18
		s1f32 = 128
		s0f32 = float32(math.Min(float64(s0f32), float64(s1f32)))
		s1i32 = l6
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s3i32 = l9
		s4i32 = l6
		s5i32 = 96
		s4i32 = s4i32 + s5i32
		s5i32 = 0
		s0i32 = f480(ctx, s0f32, s1i32, s2i32, s3i32, s4i32, s5i32)
		if s0i32 != 0 {
			goto lbl9
		}
		goto lbl0
	}
	s0i32 = 0
	l7 = s0i32
	s0i32 = l0
	s1i32 = l6
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	s3i32 = l3
	s4i32 = l6
	s5i32 = 96
	s4i32 = s4i32 + s5i32
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+32)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl9:
	s0i32 = 2
	l7 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l8 = s0i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	l9 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	s2i32 = l6
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
	s1i32 = s1i32 - s2i32
	s0i32 = s0i32 + s1i32
	l14 = s0i32
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	l10 = s0i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
	l13 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = l6
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
	s1i32 = s1i32 - s2i32
	s0i32 = s0i32 + s1i32
	l15 = s0i32
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		l4 = s0i32
		s0i32 = l10
		l8 = s0i32
		goto lbl12
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s0f32 = float32(math.Ceil(float64(s0f32)))
	l18 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl14
	}
	s0i32 = -2147483648
lbl14:
	l4 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s0f32 = float32(math.Floor(float64(s0f32)))
	l18 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl16
	}
	s0i32 = -2147483648
lbl16:
	l9 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s0f32 = float32(math.Floor(float64(s0f32)))
	l18 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl18
	}
	s0i32 = -2147483648
lbl18:
	l13 = s0i32
	s0i32 = l11
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s0f32 = float32(math.Ceil(float64(s0f32)))
	l18 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l18
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
	l18 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l18
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl20
	}
	s0i32 = -2147483648
lbl20:
	l8 = s0i32
	s1i32 = l12
	s2i32 = l10
	s1i32 = s1i32 - s2i32
	s0i32 = s0i32 + s1i32
	l12 = s0i32
lbl12:
	s0i32 = l13
	s1i32 = l8
	s2i32 = l15
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 - s1i32
	s1i32 = -3
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l9
	s2i32 = l4
	s3i32 = l14
	s2i32 = s2i32 + s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = -3
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s0i32 = s0i32 | s1i32
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l19 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l20 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l18 = s0f32
	s0i32 = l6
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l21 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l6
	s1f32 = l18
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l6
	s1f32 = l20
	s2i32 = l4
	s2f32 = float32(s2i32)
	l20 = s2f32
	s1f32 = s1f32 - s2f32
	l22 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l6
	s1f32 = l19
	s2i32 = l8
	s2f32 = float32(s2i32)
	l19 = s2f32
	s1f32 = s1f32 - s2f32
	l23 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0f32 = l23
	s1f32 = l18
	s0f32 = s0f32 - s1f32
	s1f32 = 2
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l22
	s1f32 = l21
	s0f32 = s0f32 - s1f32
	s1f32 = 2
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l18 = s0f32
		s0i32 = l1
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l16 = s0i64
		s0i32 = l6
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2f32 = l20
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l6
		s1f32 = l18
		s2f32 = l19
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l6
		s1i64 = l16
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l18 = s0f32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+16)])
	if s0i32 != 0 {
		s0i32 = l3
		s1f32 = l18
		s0f32 = f180(ctx, s0i32, s1f32)
	} else {
		s0f32 = l18
	}
	s1f32 = 128
	s0f32 = float32(math.Min(float64(s0f32), float64(s1f32)))
	l19 = s0f32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l6
	s3i32 = l2
	s4i32 = l5
	s0i32 = f1021(ctx, s0f32, s1i32, s2i32, s3i32, s4i32)
	l7 = s0i32
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l2
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l18 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l18
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl27
		}
		s1i32 = -2147483648
	lbl27:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = float32(math.Floor(float64(s1f32)))
		l18 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l18
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl29
		}
		s1i32 = -2147483648
	lbl29:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = float32(math.Floor(float64(s1f32)))
		l18 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l18
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl31
		}
		s1i32 = -2147483648
	lbl31:
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f32 = float32(math.Ceil(float64(s1f32)))
		l18 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s2f32 = -2.1474835e+09
		s3f32 = l18
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
		l18 = s1f32
		s1f32 = float32(math.Abs(float64(s1f32)))
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l18
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl33
		}
		s1i32 = -2147483648
	lbl33:
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l4
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 + s2i32
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s1i32 = f118(ctx, s1i32)
		s2i32 = 1
		s1i32 = f203(ctx, s1i32, s2i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l7 = s0i32
			goto lbl0
		}
		s0i32 = l6
		s1i32 = 4600
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = 0
		ctx.Mem[int(s0i32+32)] = uint8(s1i32)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		l4 = s0i32
		s0i32 = l6
		s1i64 = 8589934593
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+68)]))
		s1i32 = s1i32 - s2i32
		s1i64 = int64(uint32(s1i32))
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+80)]))
		s3i32 = l6
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+72)]))
		s2i32 = s2i32 - s3i32
		s2i64 = int64(uint32(s2i32))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 | s2i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+180)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l6
		s2i32 = 168
		s1i32 = s1i32 + s2i32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+64)]))
		s3i32 = l6
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+84)]))
		s4i32 = 0
		s5i32 = 0
		s0i32 = f157(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl36
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl36
		}
		s0i32 = l1
		f12(ctx, s0i32)
	lbl36:
		s0i32 = l6
		s1i32 = 168
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s0i32 = f406(ctx, s0i32, s1i32)
		l10 = s0i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)]))
		s1f32 = float32(s1i32)
		s1f32 = -s1f32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+72)]))
		s2f32 = float32(s2i32)
		s2f32 = -s2f32
		f169(ctx, s0i32, s1f32, s2f32)
		s0i32 = l6
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 13195221663744
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l1
		l8 = s0i32
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		s0i32 = f37(ctx, s0i32)
		l1 = s0i32
		s1i32 = l6
		s2i32 = 0
		f139(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		s1i32 = l6
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		f139(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 252
		s1i32 = s1i32 & s2i32
		s2i32 = 1
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l10
		s1i32 = l1
		s2i32 = l8
		f576(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		f34(ctx, s0i32)
		s0i32 = l8
		s0i32 = f23(ctx, s0i32)
		s0i32 = l10
		s0i32 = f258(ctx, s0i32)
		s0i32 = l4
		s0i32 = f41(ctx, s0i32)
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
		l1 = s0i32
		s0i32 = l0
		s1i32 = l5
		s2i32 = l6
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3i32 = l3
		s4i32 = l6
		s5i32 = 96
		s4i32 = s4i32 + s5i32
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+32)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		l3 = s0i32
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l1
			f13(ctx, s0i32)
		}
		s0i32 = 0
		l7 = s0i32
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl25
		}
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l18 = s0f32
	s0i32 = 0
	l7 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+16)])
	if s0i32 != 0 {
		s0i32 = l3
		s1f32 = l18
		s0f32 = f180(ctx, s0i32, s1f32)
	} else {
		s0f32 = l18
	}
	s1f32 = 128
	s0f32 = float32(math.Min(float64(s0f32), float64(s1f32)))
	s1i32 = l5
	s2i32 = l6
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l6
	s5i32 = 96
	s4i32 = s4i32 + s5i32
	s5i32 = 2
	s0i32 = f480(ctx, s0f32, s1i32, s2i32, s3i32, s4i32, s5i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl25:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	s0i32 = l5
	s0i32 = f487(ctx, s0i32)
	l1 = s0i32
	s0i32 = f652(ctx, s0i32)
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l7 = s0i32
		goto lbl23
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l1
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	f13(ctx, s0i32)
	s0i32 = l5
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0f32 = l19
	s1i32 = l0
	s2i32 = l6
	s3i32 = l2
	s4i32 = l5
	s5i32 = l7
	f1019(ctx, s0f32, s1i32, s2i32, s3i32, s4i32, s5i32)
lbl23:
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
	l16 = s0i64
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])))
	l17 = s0i64
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])))
	s2i64 = l17
	s1i64 = s1i64 - s2i64
	l17 = s1i64
	s2i64 = 2147483647
	s3i64 = l17
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l17 = s1i64
	s2i64 = -2147483647
	s3i64 = l17
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = l5
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])))
	s2i64 = l16
	s1i64 = s1i64 - s2i64
	l16 = s1i64
	s2i64 = 2147483647
	s3i64 = l16
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l16 = s1i64
	s2i64 = -2147483647
	s3i64 = l16
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
	s0i32 = l5
	s1i32 = l6
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	l0 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l12
	s1i64 = int64(uint32(s1i32))
	s2i32 = l11
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = 1
	l7 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = 4640
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l7
	return s0i32
}
