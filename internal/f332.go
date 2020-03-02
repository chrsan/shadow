package internal

import (
	"math"
	"unsafe"
)

func f332(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int32
	_ = l25
	var l26 int32
	_ = l26
	var l27 int32
	_ = l27
	var l28 int32
	_ = l28
	var l29 int32
	_ = l29
	var l30 int32
	_ = l30
	var l31 float32
	_ = l31
	var l32 float32
	_ = l32
	var l33 float32
	_ = l33
	var l34 float32
	_ = l34
	var l35 float32
	_ = l35
	var l36 float32
	_ = l36
	var l37 float32
	_ = l37
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
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l11 = s0i32
	s1i32 = -1
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	s0i32 = l11
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l24 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l29 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l30 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l0 = s0i32
lbl1:
	s0i32 = l10
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = 0
	l10 = s0i32
	s0i32 = l29
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1i32 = l24
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l12 = s0i32
	s1i32 = 5
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l15 = s0i32
		s0i32 = l16
		l8 = s0i32
		s0i32 = l22
		l19 = s0i32
		s0i32 = l13
		l20 = s0i32
		s0i32 = l7
		l25 = s0i32
		s0i32 = l17
		goto lbl3
	}
	s0i32 = l12
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl5
	case 1:
		goto lbl0
	case 2:
		goto lbl0
	case 3:
		goto lbl0
	case 4:
		goto lbl6
	default:
		goto lbl4
	}
lbl6:
	s0i32 = 1
	l14 = s0i32
	s0i32 = 5
	l12 = s0i32
	s0i32 = l7
	l26 = s0i32
lbl5:
	s0i32 = l21
	s1i32 = l7
	s2i32 = l12
	s3i32 = 5
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l20 = s0i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l35 = s0f32
	s1i32 = l13
	s1f32 = math.Float32frombits(uint32(s1i32))
	l36 = s1f32
	s0f32 = s0f32 - s1f32
	l33 = s0f32
	s0i32 = 0
	l10 = s0i32
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l37 = s0f32
	s1i32 = l22
	s1f32 = math.Float32frombits(uint32(s1i32))
	l31 = s1f32
	s0f32 = s0f32 - s1f32
	l34 = s0f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l33
	s3f32 = 0
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
		goto lbl0
	}
	s0f32 = l34
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1f32 = l33
	s0f32 = s0f32 * s1f32
	l32 = s0f32
	s1f32 = l32
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i32 = l7
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l8
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l25 = s0i32
	s0i32 = l23
	s1i32 = l7
	s2i32 = l8
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l23 = s0i32
	s0i32 = 1
	l15 = s0i32
	s0f32 = l31
	s1f32 = l37
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0f32 = l36
	s1f32 = l35
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l16
	l8 = s0i32
	s0i32 = l22
	l19 = s0i32
	s0i32 = l13
	l20 = s0i32
	s0i32 = l17
	goto lbl3
lbl8:
	s0f32 = l34
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 2
	s2f32 = l33
	s3f32 = 0
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3f32 = l34
	s4f32 = 0
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 | s1i32
	l9 = s0i32
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l11
		s1i32 = l9
		ctx.Mem[int(s0i32+8)] = uint8(s1i32)
		s0i32 = 0
		l18 = s0i32
		s0i32 = l16
		l8 = s0i32
		s0i32 = l17
		l9 = s0i32
		s0i32 = 1
		l12 = s0i32
		goto lbl2
	}
	s0i32 = l18
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l14
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l9
	s1i32 = l11
	s1i32 = int32(int8(ctx.Mem[int(s1i32+8)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = 0
	l18 = s0i32
	s0i32 = l16
	l8 = s0i32
	s0i32 = l22
	l19 = s0i32
	s0i32 = l13
	l20 = s0i32
	s0i32 = 1
	l14 = s0i32
	s0i32 = l17
	goto lbl3
lbl11:
	s0i32 = l0
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	s0i32 = int32(int8(ctx.Mem[int(s0i32+7)]))
	s1i32 = l9
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 3
		s1i32 = l0
		s2i32 = l0
		s3i32 = 3
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = l12
		s4i32 = 1
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		s2i32 = s2i32 & s3i32
		l0 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l12 = s0i32
		s0i32 = l20
		s1i32 = l27
		s2i32 = l0
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l27 = s0i32
		s0i32 = l19
		s1i32 = l28
		s2i32 = l0
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l28 = s0i32
		s0i32 = l16
		l8 = s0i32
		s0i32 = l17
		l9 = s0i32
		s0i32 = l14
		l18 = s0i32
		goto lbl2
	}
	s0i32 = l11
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = l9
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = 2
	l12 = s0i32
	s0i32 = l0
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	l15 = s0i32
	s0i32 = l22
	l8 = s0i32
	s0i32 = l13
	l9 = s0i32
	s0i32 = l14
	l18 = s0i32
	s0i32 = l7
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl14
	case 1:
		goto lbl13
	default:
		goto lbl2
	}
lbl14:
	s0i32 = 3
	l12 = s0i32
	s0i32 = 1
	l15 = s0i32
	s0i32 = l19
	l28 = s0i32
	s0i32 = l20
	l27 = s0i32
	s0i32 = l16
	l8 = s0i32
	s0i32 = l17
	l9 = s0i32
	s0i32 = l11
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = l11
	s1i32 = int32(ctx.Mem[int(s1i32+8)])
	s0i32 = s0i32 ^ s1i32
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	goto lbl0
lbl13:
	s0i32 = 4
	l12 = s0i32
	s0i32 = 1
	l15 = s0i32
	s0i32 = l16
	l8 = s0i32
	s0i32 = l17
	l9 = s0i32
	s0i32 = l11
	s0i32 = int32(ctx.Mem[int(s0i32+11)])
	s1i32 = l11
	s1i32 = int32(ctx.Mem[int(s1i32+9)])
	s0i32 = s0i32 ^ s1i32
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	goto lbl0
lbl4:
	s0i32 = l14
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = -1
	l15 = s0i32
	s0i32 = l11
	s0i32 = int32(int8(ctx.Mem[int(s0i32+8)]))
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l16
	l8 = s0i32
	s0i32 = l22
	l19 = s0i32
	s0i32 = l13
	l20 = s0i32
	s0i32 = l7
	l25 = s0i32
	s0i32 = 1
	l10 = s0i32
	s0i32 = l17
	goto lbl3
lbl15:
	s0i32 = l7
	l13 = s0i32
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l21
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l23
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l21
		l13 = s0i32
		goto lbl16
	}
	s0i32 = l21
	l13 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l23
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl16:
	s0i32 = l7
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l25 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l20 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0i32
	s0i32 = 1
	l15 = s0i32
	s0i32 = l16
	l8 = s0i32
	s0i32 = l13
	l21 = s0i32
	s0i32 = 1
	l18 = s0i32
	s0i32 = l17
lbl3:
	l9 = s0i32
	s0i32 = l0
	l12 = s0i32
lbl2:
	s0i32 = l2
	s1i32 = l15
	s2i32 = l24
	s1i32 = s1i32 + s2i32
	l24 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l14
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l12
		l0 = s0i32
		s0i32 = l25
		l7 = s0i32
		s0i32 = l20
		l13 = s0i32
		s0i32 = l19
		l22 = s0i32
		s0i32 = l9
		l17 = s0i32
		s0i32 = l8
		l16 = s0i32
		s0i32 = l24
		s1i32 = l30
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = 0
	l10 = s0i32
	s0i32 = l12
	s1i32 = -3
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
	s0i32 = l26
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l26
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l21
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l23
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l21
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s1i32 = l23
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0f32 = s0f32 - s1f32
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l6
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l28
		s1f32 = math.Float32frombits(uint32(s1i32))
		l31 = s1f32
		s2i32 = l8
		s2f32 = math.Float32frombits(uint32(s2i32))
		l32 = s2f32
		s3f32 = l32
		s4f32 = l31
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
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l6
		s1f32 = l31
		s2f32 = l32
		s3f32 = l31
		s4f32 = l32
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
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l6
		s1i32 = l27
		s1f32 = math.Float32frombits(uint32(s1i32))
		l31 = s1f32
		s2i32 = l9
		s2f32 = math.Float32frombits(uint32(s2i32))
		l32 = s2f32
		s3f32 = l32
		s4f32 = l31
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
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l6
		s1f32 = l31
		s2f32 = l32
		s3f32 = l31
		s4f32 = l32
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
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	}
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l14
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
	s0i32 = 1
	l10 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l11
	s1i32 = int32(int8(ctx.Mem[int(s1i32+8)]))
	s2i32 = l11
	s2i32 = int32(ctx.Mem[int(s2i32+9)])
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl0:
	s0i32 = l10
	return s0i32
}
