package internal

import (
	"math"
	"unsafe"
)

func f534(ctx *Context, l0 int32, l1 float32, l2 float32, l3 int32, l4 int32, l5 int32) int32 {
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
	var l18 int64
	_ = l18
	var l19 int64
	_ = l19
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
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l7 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l8 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l9 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l10 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l11 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	l12 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l13 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l14 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
	l15 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l7 = s0i32
	s0i32 = l6
	s1i32 = 128
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l15
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l14
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l13
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	l7 = s0i32
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = l5
	s0i32 = f346(ctx, s0i32, s1i32, s2i32)
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
	s1f32 = l1
	s2f32 = l2
	s3i32 = l6
	f239(ctx, s0i32, s1f32, s2f32, s3i32)
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l1 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l1
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
	l1 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l1
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
	l1 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl2
	}
	s0i32 = -2147483648
lbl2:
	l5 = s0i32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l1 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l1
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
	l1 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l1
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
	l1 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl4
	}
	s0i32 = -2147483648
lbl4:
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l0 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l9 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l10 = s1i32
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
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l11 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l0 = s1i32
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
		goto lbl0
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l11
	s1i32 = l5
	s1i64 = int64(s1i32)
	l17 = s1i64
	s2i32 = l4
	s2i64 = int64(s2i32)
	s1i64 = s1i64 + s2i64
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
	l16 = s1i64
	s1i32 = int32(s1i64)
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l9
	s2i32 = l8
	s2i64 = int64(s2i32)
	l19 = s2i64
	s3i32 = l3
	s3i64 = int64(s3i32)
	s2i64 = s2i64 + s3i64
	l18 = s2i64
	s3i64 = 2147483647
	s4i64 = l18
	s5i64 = 2147483647
	if s4i64 < s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l18 = s2i64
	s3i64 = -2147483647
	s4i64 = l18
	s5i64 = -2147483647
	if s4i64 > s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l18 = s2i64
	s2i32 = int32(s2i64)
	if s1i32 <= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i64 = l18
	s3i64 = l19
	s2i64 = s2i64 - s3i64
	l19 = s2i64
	s3i64 = 0
	if s2i64 > s3i64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i64 = l16
	s4i64 = l17
	s3i64 = s3i64 - s4i64
	l17 = s3i64
	s4i64 = 0
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	s3i64 = l17
	s4i64 = l19
	s3i64 = s3i64 | s4i64
	s4i64 = 2147483648
	s3i64 = s3i64 + s4i64
	s4i64 = 4294967296
	if uint64(s3i64) < uint64(s4i64) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	s3i32 = l10
	s4i32 = l8
	if s3i32 >= s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	s3i32 = l0
	s4i32 = l5
	if s3i32 >= s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 & s1i32
	l7 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l7
	return s0i32
}
