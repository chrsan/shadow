package internal

import (
	"math"
	"unsafe"
)

func f1338(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32) {
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s0i64 int64
	_ = s0i64
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
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = ctx.g0
	s1i32 = 176
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l6
	s1i32 = l4
	f114(ctx, s0i32, s1i32)
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -193
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l5
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	f123(ctx, s0i32, s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l4
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
lbl1:
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l12 = s0i64
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = l12
	s1i32 = int32(s1i64)
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
	s0i32 = l5
	s1i64 = l12
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
	s2f32 = l2
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	s2f32 = l3
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
	s2f32 = l2
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+124)]))
	s2f32 = l3
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
	s0i32 = 0
	l4 = s0i32
	s0i32 = 0
	s1i32 = l6
	s1i32 = f65(ctx, s1i32)
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 6
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l7 = s1i32
	if s1i32 != 0 {
		goto lbl5
	}
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l6
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = l6
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 | s2i32
	if s1i32 != 0 {
		goto lbl5
	}
	s1i32 = l5
	s2i32 = 112
	s1i32 = s1i32 + s2i32
	goto lbl4
lbl5:
	s1i32 = l6
	s2i32 = l5
	s3i32 = 112
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 40
	s3i32 = s3i32 + s4i32
	s4i32 = l7
	s1i32 = f64(ctx, s1i32, s2i32, s3i32, s4i32)
lbl4:
	s0i32 = f70(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 0
	s1i32 = l0
	s2f32 = l2
	s3f32 = l3
	s4i32 = l1
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
	s5i32 = l1
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+28)]))
	s6i32 = l6
	s1i32 = f534(ctx, s1i32, s2f32, s3f32, s4i32, s5i32, s6i32)
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0i32 = 1
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0i32 = l5
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s2i32 = l1
	s3i32 = l4
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+180)]))
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
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l4 = s0i32
	s1i32 = 0
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
lbl3:
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l8 = s0i32
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l5
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l6
	s3i32 = l7
	s4i32 = l5
	s5i32 = 112
	s4i32 = s4i32 + s5i32
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l3
		s1f32 = 0
		s0f32 = s0f32 + s1f32
		l13 = s0f32
		s0f32 = l2
		s1f32 = 0
		s0f32 = s0f32 + s1f32
		l14 = s0f32
	lbl9:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l8 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l9 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l10 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l11 = s0i32
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l11
		s1f32 = float32(s1i32)
		l15 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
		s0i32 = l5
		s1i32 = l10
		s1f32 = float32(s1i32)
		l16 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l5
		s1f32 = l15
		s2f32 = l3
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l5
		s1f32 = l16
		s2f32 = l2
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l5
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l5
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l8
		s1i32 = l1
		s2i32 = l5
		s3i32 = 24
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = 8
		s3i32 = s3i32 + s4i32
		s4i32 = l9
		s5i32 = 0
		s6i32 = l8
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+112)]))
		if int(s6i32) < 0 || int(s6i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s6i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s6i32].numParams != 6 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl9
		}
		goto lbl7
	}
lbl10:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l8 = s0i32
	s0i32 = l1
	s1i32 = 116
	s0i32 = s0i32 + s1i32
	s1f32 = l2
	s2f32 = l3
	s3i32 = l5
	s4i32 = 24
	s3i32 = s3i32 + s4i32
	f239(ctx, s0i32, s1f32, s2f32, s3i32)
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l13 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l13
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
	l13 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l13
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
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl11
	}
	s0i32 = -2147483648
lbl11:
	l9 = s0i32
	s0i32 = l1
	s1i32 = l4
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	l13 = s2f32
	s3f32 = 2.1474835e+09
	s4f32 = l13
	s5f32 = 2.1474835e+09
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l13 = s2f32
	s3f32 = -2.1474835e+09
	s4f32 = l13
	s5f32 = -2.1474835e+09
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l13 = s2f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2f32 = l13
		s2i32 = int32(math.Trunc(float64(s2f32)))
		goto lbl13
	}
	s2i32 = -2147483648
lbl13:
	s3i32 = l9
	s4i32 = l8
	s5i32 = 0
	s6i32 = 21916
	s7i32 = l1
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+176)]))
	if int(s7i32) < 0 || int(s7i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s7i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s7i32].numParams != 7 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32, int32, int32))(table[s7i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl10
	}
lbl7:
	s0i32 = l7
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f23(ctx, s0i32)
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	}
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l4
	s1i32 = l4
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
lbl2:
	s0i32 = l6
	s0i32 = f23(ctx, s0i32)
lbl0:
	s0i32 = l5
	s1i32 = 176
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
