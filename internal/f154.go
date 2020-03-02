package internal

import (
	"math"
	"unsafe"
)

func f154(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
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
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l11 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l20 = s1f32
	s0f32 = s0f32 - s1f32
	l12 = s0f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l17 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l10 = s1f32
	s0f32 = s0f32 - s1f32
	l13 = s0f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l4
	s0i32 = s0i32 | s1i32
	l4 = s0i32
	s0f32 = l3
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l2
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l2
	s3f32 = 360
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l17
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	s1f32 = l10
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l10 = s0f32
	goto lbl2
lbl3:
	s0f32 = l12
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l13
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		l1 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		l1 = s0i32
		s1f32 = l10
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l11
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		goto lbl0
	}
	s0i32 = l0
	s1f32 = l11
	s2f32 = l10
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	goto lbl0
lbl1:
	s0i32 = l5
	s1f32 = 0
	s2f32 = l2
	s3f32 = 0.017453292
	s2f32 = s2f32 * s3f32
	l8 = s2f32
	s2f32 = f107(ctx, s2f32)
	l9 = s2f32
	s3f32 = l9
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = 0.00024414062
	if s3f32 <= s4f32 {
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
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)])) = s1f32
	s0i32 = l5
	s1f32 = 0
	s2f32 = l8
	s2f32 = f86(ctx, s2f32)
	l8 = s2f32
	s3f32 = l8
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = 0.00024414062
	if s3f32 <= s4f32 {
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
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = s1f32
	s0i32 = l5
	s1f32 = 0
	s2f32 = l2
	s3f32 = l3
	s2f32 = s2f32 + s3f32
	s3f32 = 0.017453292
	s2f32 = s2f32 * s3f32
	l2 = s2f32
	s2f32 = f107(ctx, s2f32)
	l18 = s2f32
	s3f32 = l18
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = 0.00024414062
	if s3f32 <= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+156)])) = s1f32
	s0i32 = l5
	s1f32 = 0
	s2f32 = l2
	s2f32 = f86(ctx, s2f32)
	l19 = s2f32
	s3f32 = l19
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = 0.00024414062
	if s3f32 <= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = s1f32
	s0f32 = l15
	s1f32 = l9
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = l14
	s1f32 = l8
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = l3
	s0f32 = float32(math.Abs(float64(s0f32)))
	l16 = s0f32
	s1f32 = 360
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
	s0f32 = l16
	s1f32 = 359
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = 0.001953125
	s1f32 = l3
	s0f32 = float32(math.Copysign(float64(s0f32), float64(s1f32)))
	l16 = s0f32
lbl6:
	s0f32 = 0
	s1f32 = l2
	s2f32 = l16
	s1f32 = s1f32 - s2f32
	l2 = s1f32
	s1f32 = f107(ctx, s1f32)
	l8 = s1f32
	s2f32 = l8
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 0.00024414062
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l8 = s0f32
	s0f32 = l15
	s1f32 = 0
	s2f32 = l2
	s2f32 = f86(ctx, s2f32)
	l9 = s2f32
	s3f32 = l9
	s3f32 = float32(math.Abs(float64(s3f32)))
	s4f32 = 0.00024414062
	if s3f32 <= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l9 = s1f32
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0f32 = l14
	s1f32 = l8
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl7:
	s0i32 = l5
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = s1f32
	s0i32 = l5
	s1f32 = l8
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+156)])) = s1f32
lbl5:
	s0f32 = l15
	s1f32 = l9
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0f32 = l14
	s1f32 = l8
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l5
	s1f32 = l17
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	s2f32 = l10
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l18
	s3f32 = l13
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l2 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+148)])) = s1f32
	s0i32 = l5
	s1f32 = l11
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	s2f32 = l20
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = l19
	s3f32 = l12
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l3 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = s1f32
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		l1 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		l1 = s0i32
		s1f32 = l2
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l3
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		goto lbl0
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l1 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = l1
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l3
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
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l2
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl10:
	s0i32 = l0
	s1f32 = l3
	s2f32 = l2
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	goto lbl0
lbl8:
	s0i32 = l5
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+200)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+192)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+184)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 168
	s0i32 = s0i32 + s1i32
	s1f32 = l12
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	s2f32 = l13
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	f240(ctx, s0i32, s1f32, s2f32)
	s0i32 = l5
	s1i32 = 168
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	f142(ctx, s0i32, s1f32, s2f32)
	s0i32 = l5
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 152
	s1i32 = s1i32 + s2i32
	s2f32 = l3
	s3f32 = 0
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	s3i32 = l5
	s4i32 = 168
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s0i32 = f686(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l6 = s0i32
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 1
		s0i32 = s0i32 | s1i32
		l1 = s0i32
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 168
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = l1
			s3i32 = l1
			s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		}
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = l5
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l7 = s0i64
			s0i32 = l5
			s1i32 = 168
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = 0
			s3i32 = 0
			s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
			l1 = s0i32
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = 0
			s2f32 = 0
			s0i32 = f49(ctx, s0i32, s1i32, s2f32)
			s1i64 = l7
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = 512
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
			goto lbl13
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l4 = s0i32
		s1i32 = 0
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l2 = s0f32
			goto lbl15
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l4
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l2 = s1f32
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
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0f32 = s0f32 - s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 0.00024414062
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
	lbl15:
		s0i32 = l0
		s1f32 = l2
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	lbl13:
		s0i32 = l6
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = 0
		l4 = s0i32
	lbl17:
		s0i32 = l0
		s1i32 = l5
		s2i32 = l4
		s3i32 = 28
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = l1
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
		s4i32 = l1
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
		s5i32 = l1
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+24)]))
		s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 168
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+152)]))
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+156)]))
	s3i32 = l5
	s4i32 = 144
	s3i32 = s3i32 + s4i32
	f239(ctx, s0i32, s1f32, s2f32, s3i32)
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l5
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)]))
		l7 = s0i64
		s0i32 = l5
		s1i32 = 168
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		l1 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		s1i64 = l7
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l4 = s0i32
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)]))
		l2 = s0f32
		goto lbl19
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = l4
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
	l2 = s1f32
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
	if s0i32 != 0 {
		goto lbl19
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl19:
	s0i32 = l0
	s1f32 = l2
	s2i32 = l5
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+148)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl0:
	s0i32 = l5
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
