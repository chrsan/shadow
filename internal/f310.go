package internal

import (
	"math"
	"unsafe"
)

func f310(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l13 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1f32
	s0f32 = s0f32 - s1f32
	l3 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	l16 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l7 = s2f32
	s1f32 = s1f32 - s2f32
	l4 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l14 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l10 = s2f32
	s1f32 = s1f32 - s2f32
	l6 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	l17 = s2f32
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	l11 = s3f32
	s2f32 = s2f32 - s3f32
	l5 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l15 = s0f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l15
	s2i32 = int32(math.Float32bits(s2f32))
	s3i32 = 2139095040
	s2i32 = s2i32 & s3i32
	s3i32 = 2139095040
	if s2i32 != s3i32 {
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
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 0
	ctx.Mem[int(s0i32+54)] = uint8(s1i32)
	s0f32 = l3
	s1f32 = l10
	s2f32 = l7
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l6
	s2f32 = l9
	s3f32 = l11
	s2f32 = s2f32 - s3f32
	l12 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l5
	s2f32 = l8
	s1f32 = s1f32 * s2f32
	s2f32 = l12
	s3f32 = l4
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l18 = s1f32
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l12
		s1f32 = l5
		s0f32 = s0f32 * s1f32
		s1f32 = l8
		s2f32 = l4
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l5
		s2f32 = l5
		s1f32 = s1f32 * s2f32
		s2f32 = l4
		s3f32 = l4
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 / s1f32
		l4 = s0f32
		s1f32 = 0
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl5
		}
		s0f32 = l4
		s1f32 = 1
		if s0f32 <= s1f32 {
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
		s1f32 = l4
		s0f32 = s0f32 * s1f32
		s1f32 = l7
		s2f32 = 1
		s3f32 = l4
		s2f32 = s2f32 - s3f32
		l8 = s2f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l10
		s0f32 = s0f32 - s1f32
		l5 = s0f32
		s1f32 = l5
		s0f32 = s0f32 * s1f32
		l5 = s0f32
		s0f32 = l17
		s1f32 = l4
		s0f32 = s0f32 * s1f32
		s1f32 = l11
		s2f32 = l8
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l9
		s0f32 = s0f32 - s1f32
		l4 = s0f32
		s1f32 = l4
		s0f32 = s0f32 * s1f32
		goto lbl4
	lbl5:
		s0f32 = l8
		s1f32 = l8
		s0f32 = s0f32 * s1f32
		l5 = s0f32
		s0f32 = l12
		s1f32 = l12
		s0f32 = s0f32 * s1f32
	lbl4:
		s1f32 = l5
		s0f32 = s0f32 + s1f32
		l4 = s0f32
		s0f32 = l11
		s1f32 = l9
		s0f32 = s0f32 - s1f32
		l5 = s0f32
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l7
		s2f32 = l10
		s1f32 = s1f32 - s2f32
		l8 = s1f32
		s2f32 = l6
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l3
		s2f32 = l3
		s1f32 = s1f32 * s2f32
		s2f32 = l6
		s3f32 = l6
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 / s1f32
		l3 = s0f32
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
		s0f32 = l3
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
		s0f32 = l14
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l10
		s2f32 = 1
		s3f32 = l3
		s2f32 = s2f32 - s3f32
		l5 = s2f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l7
		s0f32 = s0f32 - s1f32
		l6 = s0f32
		s1f32 = l6
		s0f32 = s0f32 * s1f32
		l6 = s0f32
		s0f32 = l13
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		s1f32 = l9
		s2f32 = l5
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l11
		s0f32 = s0f32 - s1f32
		l3 = s0f32
		s1f32 = l3
		s0f32 = s0f32 * s1f32
		goto lbl6
	lbl7:
		s0f32 = l8
		s1f32 = l8
		s0f32 = s0f32 * s1f32
		l6 = s0f32
		s0f32 = l5
		s1f32 = l5
		s0f32 = s0f32 * s1f32
	lbl6:
		s1f32 = l6
		s0f32 = s0f32 + s1f32
		l3 = s0f32
		s1f32 = l4
		s2f32 = l4
		s3f32 = l3
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
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		return s0i32
	}
	s0f32 = l18
	s1f32 = l15
	s0f32 = s0f32 / s1f32
	l7 = s0f32
	s1f32 = l7
	s2f32 = -1
	s1f32 = s1f32 + s2f32
	if s0f32 > s1f32 {
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
		s0i32 = 2
		s1i32 = l2
		if s1i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1f32 = 1
		s2f32 = l7
		s1f32 = s1f32 - s2f32
		l3 = s1f32
		s2f32 = l10
		s1f32 = s1f32 * s2f32
		s2f32 = l7
		s3f32 = l14
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l1
		s1f32 = l13
		s2f32 = l7
		s1f32 = s1f32 * s2f32
		s2f32 = l3
		s3f32 = l9
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = 2
		return s0i32
	}
lbl1:
	s0i32 = l1
	s1f32 = l3
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	s2f32 = l6
	s3f32 = l4
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = 0
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	ctx.Mem[int(s0i32+54)] = uint8(s1i32)
	s0i32 = 1
lbl0:
	return s0i32
}
