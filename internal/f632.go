package internal

import (
	"math"
	"unsafe"
)

func f632(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l16 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = 256
	s0f32 = s0f32 * s1f32
	l17 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l17
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	l9 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l19 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l20 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l17 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l18 = s0f32
	s0i32 = l4
	s1i32 = l9
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l17
	s2f32 = l18
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	l17 = s2f32
	s1f32 = s1f32 - s2f32
	s2f32 = 256
	s1f32 = s1f32 * s2f32
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
		goto lbl2
	}
	s1i32 = -2147483648
lbl2:
	l7 = s1i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l10 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l16
	s2f32 = l20
	s1f32 = s1f32 + s2f32
	s2f32 = 256
	s1f32 = s1f32 * s2f32
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
		goto lbl4
	}
	s1i32 = -2147483648
lbl4:
	l8 = s1i32
	s2i32 = 255
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l11 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l4
	s1f32 = l17
	s2f32 = l19
	s1f32 = s1f32 + s2f32
	s2f32 = 256
	s1f32 = s1f32 * s2f32
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
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	l5 = s1i32
	s2i32 = 255
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l12 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 20568
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 20632
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 20696
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
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
		goto lbl8
	}
	s0i32 = l11
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l6
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
		goto lbl8
	}
	s0i32 = l4
	s1i32 = 80
	s0i32 = s0i32 + s1i32
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
		goto lbl8
	}
	s0i32 = l2
	s1i32 = l4
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s0i32 = f1901(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l4
	s1i32 = l3
	s2i32 = l2
	s3i32 = l4
	s4i32 = -64
	s3i32 = s3i32 - s4i32
	s0i32 = f363(ctx, s0i32, s1i32, s2i32, s3i32)
	l3 = s0i32
lbl9:
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l20 = s1f32
	s2f32 = l16
	s1f32 = s1f32 - s2f32
	l16 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = 256
	s0f32 = s0f32 * s1f32
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
		goto lbl10
	}
	s0i32 = -2147483648
lbl10:
	l6 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = l19
	s2f32 = l17
	s1f32 = s1f32 - s2f32
	l17 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = 256
	s0f32 = s0f32 * s1f32
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
		goto lbl12
	}
	s0i32 = -2147483648
lbl12:
	l2 = s0i32
	s0f32 = l16
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 + s1f32
	s1f32 = 256
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l16
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl14
	}
	s0i32 = -2147483648
lbl14:
	l1 = s0i32
	s0f32 = l19
	s1f32 = 1
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l10 = s0i32
	s0f32 = l17
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 + s1f32
	s1f32 = 256
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l16
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl16
	}
	s0i32 = -2147483648
lbl16:
	l0 = s0i32
	s0i32 = l10
	if s0i32 != 0 {
		goto lbl19
	}
	s0f32 = l20
	s1f32 = 1
	if s0f32 < s1f32 {
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
		goto lbl19
	}
	s0i32 = l2
	l10 = s0i32
	s0i32 = l9
	l11 = s0i32
	s0i32 = l7
	l12 = s0i32
	s0i32 = l6
	goto lbl18
lbl19:
	s0i32 = l2
	s1i32 = -256
	s0i32 = s0i32 & s1i32
	s1i32 = l2
	s2i32 = l2
	s3i32 = l5
	s2i32 = s2i32 ^ s3i32
	s3i32 = 256
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l13 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l10 = s0i32
	s0i32 = l9
	s1i32 = -256
	s0i32 = s0i32 & s1i32
	s1i32 = l9
	s2i32 = l1
	s3i32 = l9
	s2i32 = s2i32 ^ s3i32
	s3i32 = 256
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l14 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l11 = s0i32
	s0i32 = l7
	s1i32 = -256
	s0i32 = s0i32 & s1i32
	s1i32 = l7
	s2i32 = l0
	s3i32 = l7
	s2i32 = s2i32 ^ s3i32
	s3i32 = 256
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l15 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l12 = s0i32
	s0i32 = l5
	s1i32 = l2
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = 0
	s3i32 = l13
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s0i32 = l1
	s1i32 = l9
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = 0
	s3i32 = l14
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s0i32 = l0
	s1i32 = l7
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = 0
	s3i32 = l15
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s0i32 = l6
	s1i32 = l6
	s2i32 = l8
	s1i32 = s1i32 ^ s2i32
	s2i32 = 255
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl18
	}
	s0i32 = l8
	s1i32 = l6
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	s0i32 = l6
	s1i32 = -256
	s0i32 = s0i32 & s1i32
lbl18:
	l2 = s0i32
	s0i32 = l12
	s1i32 = l11
	s2i32 = l5
	s3i32 = l8
	s4i32 = l3
	s5i32 = 0
	f148(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l4
	s1i32 = l8
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l9 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l11
	s2i32 = 255
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l12
	s2i32 = 255
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l10
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l1
	s3i32 = l2
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
		s0i32 = l9
		s1i32 = l6
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l7
		s1i32 = l5
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l3
		s1i32 = l5
		s2i32 = l6
		s3i32 = l7
		s4i32 = l5
		s3i32 = s3i32 - s4i32
		s4i32 = l9
		s5i32 = l6
		s4i32 = s4i32 - s5i32
		s5i32 = l3
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
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
		goto lbl8
	}
	s0i32 = l2
	s1i32 = 255
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l9 = s0i32
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l8 = s0i32
	s1i32 = l6
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l7
	s1i32 = l5
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l3
	s1i32 = l5
	s2i32 = l6
	s3i32 = l7
	s4i32 = l5
	s3i32 = s3i32 - s4i32
	s4i32 = l8
	s5i32 = l6
	s4i32 = s4i32 - s5i32
	s5i32 = l3
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
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
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l5 = s0i32
lbl21:
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l6 = s0i32
	s0i32 = l9
	s1i32 = l8
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	if s0i32 != 0 {
		goto lbl22
	}
	s0i32 = l6
	s1i32 = l5
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl22
	}
	s0i32 = l3
	s1i32 = l5
	s2i32 = l8
	s3i32 = l6
	s4i32 = l5
	s3i32 = s3i32 - s4i32
	s4i32 = l9
	s5i32 = l8
	s4i32 = s4i32 - s5i32
	s5i32 = l3
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
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
lbl22:
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l5 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l5
	s1i32 = l10
	s2i32 = 255
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l7 = s1i32
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l3
	s1i32 = l7
	s2i32 = l8
	s3i32 = l5
	s4i32 = l7
	s3i32 = s3i32 - s4i32
	s4i32 = l9
	s5i32 = l8
	s4i32 = s4i32 - s5i32
	s5i32 = l3
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
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
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l5 = s0i32
lbl23:
	s0i32 = l5
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	l7 = s1i32
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	l11 = s0i32
	s1i32 = l9
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l3
	s1i32 = l7
	s2i32 = l9
	s3i32 = l5
	s4i32 = l7
	s3i32 = s3i32 - s4i32
	s4i32 = l11
	s5i32 = l9
	s4i32 = s4i32 - s5i32
	s5i32 = l3
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
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
lbl24:
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l8
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		s1i32 = 256
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l0
		s1i32 = l8
		s2i32 = l10
		s3i32 = l1
		s4i32 = l3
		f414(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		goto lbl8
	}
	s0i32 = l1
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l8
		s2i32 = l10
		s3i32 = l1
		s4i32 = l3
		f414(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l8
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l8 = s0i32
	}
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l9 = s0i32
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	s0i32 = l0
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l6
		s2i32 = l8
		s3i32 = l1
		s4i32 = l7
		s5i32 = l3
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
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
	}
	s0i32 = l10
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	s0i32 = l3
	s1i32 = l10
	s2i32 = 8
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l8
	s3i32 = l1
	s4i32 = l10
	s5i32 = -1
	s4i32 = s4i32 ^ s5i32
	s5i32 = 255
	s4i32 = s4i32 & s5i32
	s5i32 = l3
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
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
lbl27:
	s0i32 = l2
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s1i32 = l9
	s2i32 = l10
	s3i32 = l1
	s4i32 = 255
	s3i32 = s3i32 ^ s4i32
	s4i32 = l3
	f414(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl8:
	s0i32 = l4
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
