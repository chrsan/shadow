package internal

import (
	"math"
	"unsafe"
)

func f1545(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = 0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	l5 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l6 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
lbl1:
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l1
		s2i32 = l2
		s3i32 = l7
		s4i32 = l5
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+40)]))
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l5 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
lbl0:
	l6 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l7 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l8 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l9 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l10 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	l11 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l12 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l13 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	l14 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l15 = s0i32
	s0i32 = l4
	s1i32 = 128
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l14
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l13
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l15
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl3
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l7 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l8 = s0i32
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l8
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = s1f32
	s0i32 = l4
	s1i32 = l7
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = s1f32
	s0i32 = l6
	s1i32 = 36
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s2i32 = l4
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
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
		goto lbl5
	}
	s1i32 = -2147483648
lbl5:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
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
		goto lbl7
	}
	s1i32 = -2147483648
lbl7:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)]))
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
		goto lbl9
	}
	s1i32 = -2147483648
lbl9:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
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
		goto lbl11
	}
	s1i32 = -2147483648
lbl11:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
lbl3:
	s0i32 = l5
	s1i32 = l1
	s2i32 = l4
	s3i32 = 24
	s2i32 = s2i32 + s3i32
	s3i32 = l4
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = l2
	f1935(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l3 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l6 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l1 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0f32 = 0
	s1i64 = l16
	s2i64 = l17
	s1i64 = s1i64 | s2i64
	s2i64 = 2147483648
	s1i64 = s1i64 + s2i64
	s2i64 = 4294967295
	if uint64(s1i64) > uint64(s2i64) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl13
	}
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	l20 = s0f32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	l19 = s0f32
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	l21 = s0f32
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s0f32 = float32(s0i32)
	goto lbl13
lbl14:
	s0f32 = 0
lbl13:
	l18 = s0f32
	s0i32 = l0
	s1f32 = l21
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4412)])) = s1f32
	s0i32 = l0
	s1i32 = 4424
	s0i32 = s0i32 + s1i32
	s1f32 = l20
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 4420
	s0i32 = s0i32 + s1i32
	s1f32 = l19
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 4416
	s0i32 = s0i32 + s1i32
	s1f32 = l18
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
