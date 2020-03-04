package internal

import (
	"math"
	"unsafe"
)

func f1596(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float32
	_ = l4
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 10
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 5
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	l2 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl0
	case 2:
		goto lbl4
	case 3:
		goto lbl3
	case 4:
		goto lbl1
	default:
		goto lbl6
	}
lbl6:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	s2f32 = float32(math.Floor(float64(s2f32)))
	l4 = s2f32
	s3f32 = 2.1474835e+09
	s4f32 = l4
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
	l4 = s2f32
	s3f32 = -2.1474835e+09
	s4f32 = l4
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
	l4 = s2f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2f32 = l4
		s2i32 = int32(math.Trunc(float64(s2f32)))
		goto lbl7
	}
	s2i32 = -2147483648
lbl7:
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s3i32 = 255
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 & s1i32
	return s0i32
lbl5:
	s0i32 = 2
	return s0i32
lbl4:
	s0i32 = 2
	s1i32 = 0
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	return s0i32
lbl3:
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1f32 = 255
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	l4 = s0f32
	s1f32 = 2.1474835e+09
	s2f32 = l4
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
	l4 = s0f32
	s1f32 = -2.1474835e+09
	s2f32 = l4
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
	l4 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l4
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl9
	}
	s0i32 = -2147483648
lbl9:
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	s1i32 = 255
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl2:
	s0i32 = 0
	return s0i32
lbl1:
	s0i32 = 0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2f32 = 255
	s1f32 = s1f32 * s2f32
	s2f32 = 0.5
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Floor(float64(s1f32)))
	l4 = s1f32
	s2f32 = 2.1474835e+09
	s3f32 = l4
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
	l4 = s1f32
	s2f32 = -2.1474835e+09
	s3f32 = l4
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
	l4 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l4
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl11
	}
	s1i32 = -2147483648
lbl11:
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = 255
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s2i32 = s2i32 | s3i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
lbl0:
	s0i32 = l2
	return s0i32
}
