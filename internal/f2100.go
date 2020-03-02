package internal

import (
	"unsafe"
)

func f2100(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var l11 int32
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
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
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
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+9)])
	l2 = s0i32
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+9)])
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l7 = s0i32
		s0i32 = 1
		l2 = s0i32
		s0i32 = 0
		l5 = s0i32
		goto lbl2
	}
	s0i32 = l5
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l6 = s1i32
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = 1
	l2 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l7 = s0i32
	s0i32 = l6
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = 1
	l3 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 0
	l2 = s0i32
	s0i32 = l6
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl4:
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l2 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl8
	case 1:
		goto lbl7
	case 2:
		goto lbl7
	case 3:
		goto lbl6
	default:
		goto lbl2
	}
lbl8:
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl5
lbl7:
	s0i32 = l3
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl5
lbl6:
	s0i32 = l3
	s1i32 = 3
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl5:
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l8
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l8
	l5 = s0i32
	s0i32 = 0
	l2 = s0i32
lbl2:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l6 = s0i32
		s0i32 = l4
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = l6
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l19 = s0f32
lbl11:
	s0i32 = l3
	s1i32 = 3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s0i32 = 0
	l6 = s0i32
	s0i32 = 1
	l2 = s0i32
lbl13:
	s0i32 = l7
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l12 = s0f32
	s1f32 = l13
	s2f32 = l12
	s3f32 = l13
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l4 = s2i32
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l13 = s0f32
	s0i32 = l2
	s1i32 = l6
	s2i32 = l4
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l6 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l7
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l14 = s0f32
	s1f32 = l19
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l7
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = l3
	s1i32 = i32RemS(s1i32, s2i32)
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l14
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l2
	s1i32 = l3
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l6
	l4 = s0i32
	l9 = s0i32
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0f32
	l13 = s0f32
lbl16:
	s0i32 = l7
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l14
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l11
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0f32
		s1f32 = l15
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			l9 = s0i32
			s0f32 = l12
			l15 = s0f32
			goto lbl18
		}
		s0f32 = l12
		s1f32 = l13
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl18
		}
		s0f32 = l12
		l13 = s0f32
		s0i32 = l2
		l4 = s0i32
	lbl18:
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
	}
	s0i32 = l4
	s1i32 = l9
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l9
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	s0f32 = float32(s0i32)
	goto lbl14
lbl15:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l6
	l2 = s0i32
lbl20:
	s0i32 = l2
	s1i32 = l9
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s0i32 = i32RemS(s0i32, s1i32)
	l2 = s0i32
	s1i32 = l6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l7
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l16 = s0f32
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l18 = s1f32
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l14
	s3f32 = l16
	if s2f32 == s3f32 {
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
		goto lbl20
	}
	s0i32 = l6
	l2 = s0i32
lbl21:
	s0i32 = l6
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s1i32 = i32RemS(s1i32, s2i32)
	l2 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l12
		l13 = s0f32
		s0f32 = l14
		l15 = s0f32
		goto lbl22
	}
	s0i32 = l7
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l15 = s0f32
	s0f32 = l12
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l13 = s1f32
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl22
	}
	s0f32 = l14
	s1f32 = l15
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
lbl22:
	s0f32 = l12
	s1f32 = l13
	s0f32 = s0f32 - s1f32
	s1f32 = l12
	s2f32 = l18
	s1f32 = s1f32 - s2f32
	s2f32 = l15
	s3f32 = l16
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l14
	s3f32 = l16
	s2f32 = s2f32 - s3f32
	s3f32 = l13
	s4f32 = l18
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l17 = s1f32
	s2f32 = 0
	if s1f32 == s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l12
		s1f64 = float64(s1f32)
		s2f32 = l18
		s2f64 = float64(s2f32)
		l21 = s2f64
		s1f64 = s1f64 - s2f64
		s2f32 = l15
		s2f64 = float64(s2f32)
		s3f32 = l16
		s3f64 = float64(s3f32)
		l22 = s3f64
		s2f64 = s2f64 - s3f64
		s1f64 = s1f64 * s2f64
		s2f32 = l14
		s2f64 = float64(s2f32)
		s3f64 = l22
		s2f64 = s2f64 - s3f64
		s3f32 = l13
		s3f64 = float64(s3f32)
		s4f64 = l21
		s3f64 = s3f64 - s4f64
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 - s2f64
		s1f32 = float32(s1f64)
		l17 = s1f32
	}
	s1f32 = l17
	s2f32 = l15
	s3f32 = l14
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = l17
	s2f32 = l16
	s3f32 = l14
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	s1f32 = l17
	s2f32 = l17
	s3f32 = 0
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
lbl14:
	l12 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0f32 = l12
	l20 = s0f32
	s0f32 = l14
	l19 = s0f32
lbl12:
	s0i32 = l5
	s1i32 = l8
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = l3
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = 1
		l3 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l8
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
	lbl27:
		s0i32 = l5
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l2 = s0i32
		s1i32 = 4
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl28
		}
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl31
		case 1:
			goto lbl30
		case 2:
			goto lbl30
		case 3:
			goto lbl29
		default:
			goto lbl11
		}
	lbl31:
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		goto lbl28
	lbl30:
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		goto lbl28
	lbl29:
		s0i32 = l3
		s1i32 = 3
		s0i32 = s0i32 + s1i32
		l3 = s0i32
	lbl28:
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l8
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
		s0i32 = l8
		l5 = s0i32
		goto lbl11
	}
	s0f32 = l20
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l1
	s1f32 = l20
	s2f32 = 0
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
lbl10:
}
