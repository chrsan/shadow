package internal

import (
	"math"
	"unsafe"
)

func f1700(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
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
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
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
	var s8f32 float32
	_ = s8f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l14 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l15 = s1f32
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s0i32 = 1
	l3 = s0i32
	s0i32 = 1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l16 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l18 = s2f32
	s1f32 = s1f32 - s2f32
	l19 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0i32 = 1
	s1f32 = l11
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl0
	}
	s0f32 = l19
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l11
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
lbl0:
	l5 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l21 = s0f32
	s1f32 = l14
	s0f32 = s0f32 - s1f32
	l12 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l17 = s0f32
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	l20 = s0f32
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
		goto lbl1
	}
	s0f32 = l12
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
		goto lbl1
	}
	s0f32 = l20
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l12
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l3 = s0i32
lbl1:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l22 = s0f32
	s1f32 = l21
	s0f32 = s0f32 - s1f32
	l13 = s0f32
	s0i32 = 1
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	l23 = s1f32
	s2f32 = l17
	s1f32 = s1f32 - s2f32
	l24 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	s1f32 = l13
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = 2139095040
	s1i32 = s1i32 & s2i32
	s2i32 = 2139095040
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0f32 = l24
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l13
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
lbl2:
	l6 = s0i32
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 & s1i32
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l6
		if s1i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = 1
	s1i32 = l3
	s2i32 = l5
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s0f32 = l13
	s0f32 = float32(math.Abs(float64(s0f32)))
	l13 = s0f32
	s1f32 = l24
	s1f32 = float32(math.Abs(float64(s1f32)))
	l24 = s1f32
	s2f32 = l24
	s3f32 = l13
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
	s1f32 = l22
	s2f32 = l14
	s1f32 = s1f32 - s2f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	l14 = s1f32
	s2f32 = l23
	s3f32 = l16
	s2f32 = s2f32 - s3f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	l16 = s2f32
	s3f32 = l16
	s4f32 = l14
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
	s2f32 = l12
	s2f32 = float32(math.Abs(float64(s2f32)))
	l12 = s2f32
	s3f32 = l20
	s3f32 = float32(math.Abs(float64(s3f32)))
	l16 = s3f32
	s4f32 = l16
	s5f32 = l12
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
	l12 = s2f32
	s3f32 = l22
	s4f32 = l15
	s3f32 = s3f32 - s4f32
	s3f32 = float32(math.Abs(float64(s3f32)))
	l16 = s3f32
	s4f32 = l23
	s5f32 = l18
	s4f32 = s4f32 - s5f32
	s4f32 = float32(math.Abs(float64(s4f32)))
	l20 = s4f32
	s5f32 = l20
	s6f32 = l16
	if s5f32 < s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	l16 = s3f32
	s4f32 = l21
	s5f32 = l15
	s4f32 = s4f32 - s5f32
	s4f32 = float32(math.Abs(float64(s4f32)))
	l15 = s4f32
	s5f32 = l17
	s6f32 = l18
	s5f32 = s5f32 - s6f32
	s5f32 = float32(math.Abs(float64(s5f32)))
	l18 = s5f32
	s6f32 = l18
	s7f32 = l15
	if s6f32 < s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	if s6i32 != 0 {
		// s4f32 = s4f32
	} else {
		s4f32 = s5f32
	}
	l15 = s4f32
	s5f32 = l11
	s5f32 = float32(math.Abs(float64(s5f32)))
	l11 = s5f32
	s6f32 = l19
	s6f32 = float32(math.Abs(float64(s6f32)))
	l18 = s6f32
	s7f32 = l18
	s8f32 = l11
	if s7f32 < s8f32 {
		s7i32 = 1
	} else {
		s7i32 = 0
	}
	if s7i32 != 0 {
		// s5f32 = s5f32
	} else {
		s5f32 = s6f32
	}
	l11 = s5f32
	s6f32 = -1
	s7f32 = l11
	s8f32 = -1
	if s7f32 > s8f32 {
		s7i32 = 1
	} else {
		s7i32 = 0
	}
	l7 = s7i32
	if s7i32 != 0 {
		// s5f32 = s5f32
	} else {
		s5f32 = s6f32
	}
	l11 = s5f32
	s6f32 = l11
	s7f32 = l15
	if s6f32 < s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	l9 = s6i32
	if s6i32 != 0 {
		// s4f32 = s4f32
	} else {
		s4f32 = s5f32
	}
	l11 = s4f32
	s5f32 = l11
	s6f32 = l16
	if s5f32 < s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	l10 = s5i32
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	l11 = s3f32
	s4f32 = l11
	s5f32 = l12
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l3 = s4i32
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l11 = s2f32
	s3f32 = l11
	s4f32 = l14
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l4 = s3i32
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l11 = s1f32
	s2f32 = l11
	s3f32 = l13
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l11 = s0f32
	s1f32 = l11
	s0f32 = s0f32 * s1f32
	s1f32 = 1e-05
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s0i32 = l0
	s1i32 = 3
	s2i32 = 3
	s3i32 = 2
	s4i32 = 3
	s5i32 = 2
	s6i32 = l7
	s7i32 = l9
	if s7i32 != 0 {
		// s5i32 = s5i32
	} else {
		s5i32 = s6i32
	}
	s6i32 = l10
	if s6i32 != 0 {
		// s4i32 = s4i32
	} else {
		s4i32 = s5i32
	}
	s5i32 = l3
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	s4i32 = l4
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l6
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l7 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l18 = s0f32
	s1i32 = l0
	s2i32 = 2
	s3i32 = l3
	s4i32 = l4
	s3i32 = s3i32 | s4i32
	s4i32 = l6
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l3 = s2i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1f32
	s0f32 = s0f32 - s1f32
	l14 = s0f32
	s1i32 = l0
	s2i32 = 2
	s3i32 = l7
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	l6 = s2i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l20 = s1f32
	s2f32 = l11
	s1f32 = s1f32 - s2f32
	l19 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l9
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l21 = s1f32
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l12 = s2f32
	s1f32 = s1f32 - s2f32
	l15 = s1f32
	s2i32 = l10
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l22 = s2f32
	s3f32 = l12
	s2f32 = s2f32 - s3f32
	l17 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l14
	s2f32 = l14
	s1f32 = s1f32 * s2f32
	s2f32 = l15
	s3f32 = l15
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l23 = s1f32
	s0f32 = s0f32 / s1f32
	l13 = s0f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl7
	}
	s0f32 = l13
	s1f32 = 1
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl7
	}
	s0f32 = l21
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	s1f32 = l12
	s2f32 = 1
	s3f32 = l13
	s2f32 = s2f32 - s3f32
	l19 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l22
	s0f32 = s0f32 - s1f32
	l17 = s0f32
	s1f32 = l17
	s0f32 = s0f32 * s1f32
	l17 = s0f32
	s0f32 = l18
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	s1f32 = l11
	s2f32 = l19
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l20
	s0f32 = s0f32 - s1f32
	l13 = s0f32
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	goto lbl6
lbl7:
	s0f32 = l17
	s1f32 = l17
	s0f32 = s0f32 * s1f32
	l17 = s0f32
	s0f32 = l19
	s1f32 = l19
	s0f32 = s0f32 * s1f32
lbl6:
	s1f32 = l17
	s0f32 = s0f32 + s1f32
	s1f32 = l16
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
		s0f32 = l14
		s1i32 = l0
		s2i32 = l3
		s3i32 = l7
		s2i32 = s2i32 ^ s3i32
		s3i32 = l6
		s2i32 = s2i32 ^ s3i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l19 = s1f32
		s2f32 = l11
		s1f32 = s1f32 - s2f32
		l14 = s1f32
		s0f32 = s0f32 * s1f32
		s1f32 = l15
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l17 = s2f32
		s3f32 = l12
		s2f32 = s2f32 - s3f32
		l15 = s2f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l23
		s0f32 = s0f32 / s1f32
		l13 = s0f32
		s1f32 = 0
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl10
		}
		s0f32 = l13
		s1f32 = 1
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
		s0f32 = l21
		s1f32 = l13
		s0f32 = s0f32 * s1f32
		s1f32 = l12
		s2f32 = 1
		s3f32 = l13
		s2f32 = s2f32 - s3f32
		l14 = s2f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l17
		s0f32 = s0f32 - s1f32
		l12 = s0f32
		s1f32 = l12
		s0f32 = s0f32 * s1f32
		l12 = s0f32
		s0f32 = l18
		s1f32 = l13
		s0f32 = s0f32 * s1f32
		s1f32 = l11
		s2f32 = l14
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l19
		s0f32 = s0f32 - s1f32
		l11 = s0f32
		s1f32 = l11
		s0f32 = s0f32 * s1f32
		goto lbl9
	lbl10:
		s0f32 = l15
		s1f32 = l15
		s0f32 = s0f32 * s1f32
		l12 = s0f32
		s0f32 = l14
		s1f32 = l14
		s0f32 = s0f32 * s1f32
	lbl9:
		s1f32 = l12
		s0f32 = s0f32 + s1f32
		s1f32 = l16
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 2
	goto lbl3
lbl5:
	s0i32 = 0
	l5 = s0i32
	s0i32 = l0
	s1i32 = l8
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s0i32 = f446(ctx, s0i32, s1i32)
	l4 = s0i32
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = 0
	l3 = s0i32
lbl12:
	s0i32 = l8
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l11 = s0f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0f32 = l11
	s1f32 = 1
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l0
	s1f32 = l11
	s2i32 = l1
	s3i32 = l3
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = 0
	f337(ctx, s0i32, s1f32, s2i32, s3i32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l11
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
	}
	s0i32 = l3
	s1f32 = l12
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	if s1f32 != s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2f32 = l11
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl13:
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l3
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	goto lbl3
lbl11:
	s0i32 = 1
lbl3:
	l4 = s0i32
	s0i32 = l8
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
