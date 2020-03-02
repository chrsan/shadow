package internal

import (
	"math"
	"unsafe"
)

func f538(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int64
	_ = l20
	var l21 float64
	_ = l21
	var l22 float64
	_ = l22
	var l23 float64
	_ = l23
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
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = ctx.g0
	s1i32 = 560
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = l2
	s2i32 = -3
	s1i32 = s1i32 + s2i32
	s2i32 = 24
	s1i32 = i32DivS(s1i32, s2i32)
	l5 = s1i32
	s2i32 = 0
	s3i32 = l5
	s4i32 = 0
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l14 = s1i32
	s2i32 = -24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 16176
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l11 = s0i32
	s1i32 = l3
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l11
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l14
		s1i32 = l8
		s0i32 = s0i32 - s1i32
		l2 = s0i32
	lbl1:
		s0i32 = l6
		s1i32 = 320
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 0
		if s1i32 < s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = 0
		} else {
			s1i32 = l2
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = 16192
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1f64 = float64(s1i32)
		}
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l12
	s1i32 = -24
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = 0
	l5 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
lbl3:
	s0i32 = l7
	if s0i32 != 0 {
		s0f64 = 0
		l21 = s0f64
		goto lbl4
	}
	s0i32 = l5
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = 0
	l2 = s0i32
	s0f64 = 0
	l21 = s0f64
lbl6:
	s0f64 = l21
	s1i32 = l0
	s2i32 = l2
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l6
	s3i32 = 320
	s2i32 = s2i32 + s3i32
	s3i32 = l10
	s4i32 = l2
	s3i32 = s3i32 - s4i32
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	l21 = s0f64
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
		goto lbl6
	}
lbl4:
	s0i32 = l6
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f64 = l21
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l5
	s1i32 = l11
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = 23
	s1i32 = l9
	s0i32 = s0i32 - s1i32
	l17 = s0i32
	s0i32 = 24
	s1i32 = l9
	s0i32 = s0i32 - s1i32
	l15 = s0i32
	s0i32 = l11
	l5 = s0i32
lbl8:
	s0i32 = l6
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l21 = s0f64
	s0i32 = 0
	l2 = s0i32
	s0i32 = l5
	l7 = s0i32
	s0i32 = l5
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l13 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl10:
		s0i32 = l6
		s1i32 = 480
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1f64 = l21
		s2f64 = l21
		s3f64 = 5.960464477539063e-08
		s2f64 = s2f64 * s3f64
		l21 = s2f64
		s2f64 = math.Abs(s2f64)
		s3f64 = 2.147483648e+09
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			s2f64 = l21
			s2i32 = int32(math.Trunc(s2f64))
			goto lbl12
		}
		s2i32 = -2147483648
	lbl12:
		s2f64 = float64(s2i32)
		l21 = s2f64
		s3f64 = -1.6777216e+07
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 + s2f64
		l22 = s1f64
		s1f64 = math.Abs(s1f64)
		s2f64 = 2.147483648e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l22
			s1i32 = int32(math.Trunc(s1f64))
			goto lbl11
		}
		s1i32 = -2147483648
	lbl11:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l7
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f64 = l21
		s0f64 = s0f64 + s1f64
		l21 = s0f64
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l7
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l10 = s0i32
		s0i32 = l8
		l7 = s0i32
		s0i32 = l10
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0f64 = l21
	s1i32 = l9
	s0f64 = f250(ctx, s0f64, s1i32)
	l21 = s0f64
	s1f64 = l21
	s2f64 = 0.125
	s1f64 = s1f64 * s2f64
	s1f64 = math.Floor(s1f64)
	s2f64 = -8
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	l21 = s0f64
	s0f64 = math.Abs(s0f64)
	s1f64 = 2.147483648e+09
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l21
		s0i32 = int32(math.Trunc(s0f64))
		goto lbl15
	}
	s0i32 = -2147483648
lbl15:
	l10 = s0i32
	s0f64 = l21
	s1i32 = l10
	s1f64 = float64(s1i32)
	s0f64 = s0f64 - s1f64
	l21 = s0f64
	s0i32 = l9
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l18 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+476)]))
		l2 = s1i32
		s2i32 = l2
		s3i32 = l15
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l2 = s2i32
		s3i32 = l15
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 - s2i32
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+476)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0i32 = l7
		s1i32 = l17
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		goto lbl20
	}
	s0i32 = l9
	if s0i32 != 0 {
		goto lbl19
	}
	s0i32 = l5
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+476)]))
	s1i32 = 23
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
lbl20:
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	goto lbl18
lbl19:
	s0i32 = 2
	l8 = s0i32
	s0f64 = l21
	s1f64 = 0.5
	if s0f64 >= s1f64 {
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
		goto lbl18
	}
	s0i32 = 0
	l8 = s0i32
	goto lbl17
lbl18:
	s0i32 = 0
	l2 = s0i32
	s0i32 = 0
	l7 = s0i32
	s0i32 = l13
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl23:
		s0i32 = l6
		s1i32 = 480
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l19 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0i32
		s0i32 = 16777215
		l16 = s0i32
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l13
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl25
			}
			s0i32 = 16777216
			l16 = s0i32
			s0i32 = 1
			l7 = s0i32
		}
		s0i32 = l19
		s1i32 = l16
		s2i32 = l13
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl24
	lbl25:
		s0i32 = 0
		l7 = s0i32
	lbl24:
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
	}
	s0i32 = l18
	if s0i32 != 0 {
		goto lbl27
	}
	s0i32 = l9
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+476)]))
		s2i32 = 8388607
		s1i32 = s1i32 & s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+476)])) = uint32(s1i32)
		goto lbl27
	}
	s0i32 = l5
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+476)]))
	s2i32 = 4194303
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+476)])) = uint32(s1i32)
lbl27:
	s0i32 = l10
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l8
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0f64 = 1
	s1f64 = l21
	s0f64 = s0f64 - s1f64
	l21 = s0f64
	s0i32 = 2
	l8 = s0i32
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0f64 = l21
	s1f64 = 1
	s2i32 = l9
	s1f64 = f250(ctx, s1f64, s2i32)
	s0f64 = s0f64 - s1f64
	l21 = s0f64
lbl17:
	s0f64 = l21
	s1f64 = 0
	if s0f64 == s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l7 = s0i32
		s0i32 = l5
		l2 = s0i32
		s1i32 = l11
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl30
		}
	lbl31:
		s0i32 = l6
		s1i32 = 480
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l7
		s0i32 = s0i32 | s1i32
		l7 = s0i32
		s0i32 = l2
		s1i32 = l11
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl30
		}
		s0i32 = l9
		l12 = s0i32
	lbl32:
		s0i32 = l12
		s1i32 = -24
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = l6
		s1i32 = 480
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl32
		}
		goto lbl7
	lbl30:
		s0i32 = 1
		l2 = s0i32
	lbl33:
		s0i32 = l2
		l7 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l6
		s1i32 = 480
		s0i32 = s0i32 + s1i32
		s1i32 = l11
		s2i32 = l7
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl33
		}
		s0i32 = l5
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		l7 = s0i32
	lbl34:
		s0i32 = l6
		s1i32 = 320
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l5
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s2i32 = l14
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 16192
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f64 = float64(s1i32)
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = 0
		l2 = s0i32
		s0f64 = 0
		l21 = s0f64
		s0i32 = l3
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl36:
			s0f64 = l21
			s1i32 = l0
			s2i32 = l2
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l6
			s3i32 = 320
			s2i32 = s2i32 + s3i32
			s3i32 = l8
			s4i32 = l2
			s3i32 = s3i32 - s4i32
			s4i32 = 3
			s3i32 = s3i32 << (uint32(s4i32) & 31)
			s2i32 = s2i32 + s3i32
			s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s1f64 = s1f64 * s2f64
			s0f64 = s0f64 + s1f64
			l21 = s0f64
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
				goto lbl36
			}
		}
		s0i32 = l6
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1f64 = l21
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l5
		s1i32 = l7
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl34
		}
		s0i32 = l7
		l5 = s0i32
		goto lbl8
	}
	s0f64 = l21
	s1i32 = 0
	s2i32 = l9
	s1i32 = s1i32 - s2i32
	s0f64 = f250(ctx, s0f64, s1i32)
	l21 = s0f64
	s1f64 = 1.6777216e+07
	if s0f64 >= s1f64 {
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
		s0i32 = l6
		s1i32 = 480
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1f64 = l21
		s2f64 = l21
		s3f64 = 5.960464477539063e-08
		s2f64 = s2f64 * s3f64
		l21 = s2f64
		s2f64 = math.Abs(s2f64)
		s3f64 = 2.147483648e+09
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			s2f64 = l21
			s2i32 = int32(math.Trunc(s2f64))
			goto lbl40
		}
		s2i32 = -2147483648
	lbl40:
		l2 = s2i32
		s2f64 = float64(s2i32)
		s3f64 = -1.6777216e+07
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 + s2f64
		l21 = s1f64
		s1f64 = math.Abs(s1f64)
		s2f64 = 2.147483648e+09
		if s1f64 < s2f64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f64 = l21
			s1i32 = int32(math.Trunc(s1f64))
			goto lbl39
		}
		s1i32 = -2147483648
	lbl39:
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		goto lbl37
	}
	s0f64 = l21
	s0f64 = math.Abs(s0f64)
	s1f64 = 2.147483648e+09
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = l21
		s0i32 = int32(math.Trunc(s0f64))
		goto lbl43
	}
	s0i32 = -2147483648
lbl43:
	l2 = s0i32
	s0i32 = l9
	l12 = s0i32
lbl37:
	s0i32 = l6
	s1i32 = 480
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl7:
	s0f64 = 1
	s1i32 = l12
	s0f64 = f250(ctx, s0f64, s1i32)
	l21 = s0f64
	s0i32 = l5
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl45
	}
	s0i32 = l5
	l2 = s0i32
lbl46:
	s0i32 = l6
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f64 = l21
	s2i32 = l6
	s3i32 = 480
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2f64 = float64(s2i32)
	s1f64 = s1f64 * s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0f64 = l21
	s1f64 = 5.960464477539063e-08
	s0f64 = s0f64 * s1f64
	l21 = s0f64
	s0i32 = l2
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l2
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl46
	}
	s0i32 = l5
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl45
	}
	s0i32 = l5
	l2 = s0i32
lbl47:
	s0i32 = l5
	s1i32 = l2
	l0 = s1i32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s0f64 = 0
	l21 = s0f64
	s0i32 = 0
	l2 = s0i32
lbl48:
	s0f64 = l21
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 18960
	s1i32 = s1i32 + s2i32
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l6
	s3i32 = l0
	s4i32 = l2
	s3i32 = s3i32 + s4i32
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f64 = s1f64 * s2f64
	s0f64 = s0f64 + s1f64
	l21 = s0f64
	s0i32 = l2
	s1i32 = l11
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl49
	}
	s0i32 = l2
	s1i32 = l3
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl48
	}
lbl49:
	s0i32 = l6
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f64 = l21
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l0
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l0
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl47
	}
lbl45:
	s0i32 = l4
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl50
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl52
	case 1:
		goto lbl52
	case 2:
		goto lbl54
	default:
		goto lbl53
	}
lbl54:
	s0f64 = 0
	l22 = s0f64
	s0i32 = l5
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl55
	}
	s0i32 = l6
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l21 = s0f64
	s0i32 = l5
	l2 = s0i32
lbl56:
	s0i32 = l6
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f64 = l21
	s2i32 = l6
	s3i32 = 160
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = -1
	s3i32 = s3i32 + s4i32
	l0 = s3i32
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	l3 = s2i32
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l23 = s2f64
	s3f64 = l23
	s4f64 = l21
	s3f64 = s3f64 + s4f64
	l21 = s3f64
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l3
	s1f64 = l21
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l2
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l3 = s0i32
	s0i32 = l0
	l2 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl56
	}
	s0i32 = l5
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl55
	}
	s0i32 = l6
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l21 = s0f64
	s0i32 = l5
	l2 = s0i32
lbl57:
	s0i32 = l6
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1f64 = l21
	s2i32 = l6
	s3i32 = 160
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = -1
	s3i32 = s3i32 + s4i32
	l0 = s3i32
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	l3 = s2i32
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l22 = s2f64
	s3f64 = l22
	s4f64 = l21
	s3f64 = s3f64 + s4f64
	l21 = s3f64
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 + s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l3
	s1f64 = l21
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l2
	s1i32 = 2
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l3 = s0i32
	s0i32 = l0
	l2 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl57
	}
	s0f64 = 0
	l22 = s0f64
	s0i32 = l5
	s1i32 = 1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl55
	}
lbl58:
	s0f64 = l22
	s1i32 = l6
	s2i32 = 160
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f64 = s0f64 + s1f64
	l22 = s0f64
	s0i32 = l5
	s1i32 = 2
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l5
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl58
	}
lbl55:
	s0i32 = l6
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)]))
	l21 = s0f64
	s0i32 = l8
	if s0i32 != 0 {
		goto lbl51
	}
	s0i32 = l1
	s1f64 = l21
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l6
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
	l20 = s0i64
	s0i32 = l1
	s1f64 = l22
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f64
	s0i32 = l1
	s1i64 = l20
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	goto lbl50
lbl53:
	s0f64 = 0
	l21 = s0f64
	s0i32 = l5
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl60:
		s0f64 = l21
		s1i32 = l6
		s2i32 = 160
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f64 = s0f64 + s1f64
		l21 = s0f64
		s0i32 = l5
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l5
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl60
		}
	}
	s0i32 = l1
	s1f64 = l21
	s1f64 = -s1f64
	s2f64 = l21
	s3i32 = l8
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	goto lbl50
lbl52:
	s0f64 = 0
	l21 = s0f64
	s0i32 = l5
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		l2 = s0i32
	lbl62:
		s0f64 = l21
		s1i32 = l6
		s2i32 = 160
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f64 = s0f64 + s1f64
		l21 = s0f64
		s0i32 = l2
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l2
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl62
		}
	}
	s0i32 = l1
	s1f64 = l21
	s1f64 = -s1f64
	s2f64 = l21
	s3i32 = l8
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l6
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)]))
	s1f64 = l21
	s0f64 = s0f64 - s1f64
	l21 = s0f64
	s0i32 = 1
	l2 = s0i32
	s0i32 = l5
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl64:
		s0f64 = l21
		s1i32 = l6
		s2i32 = 160
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f64 = s0f64 + s1f64
		l21 = s0f64
		s0i32 = l2
		s1i32 = l5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl64
		}
	}
	s0i32 = l1
	s1f64 = l21
	s1f64 = -s1f64
	s2f64 = l21
	s3i32 = l8
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	goto lbl50
lbl51:
	s0i32 = l1
	s1f64 = l21
	s1f64 = -s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l6
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)]))
	l21 = s0f64
	s0i32 = l1
	s1f64 = l22
	s1f64 = -s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f64
	s0i32 = l1
	s1f64 = l21
	s1f64 = -s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
lbl50:
	s0i32 = l6
	s1i32 = 560
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l10
	s1i32 = 7
	s0i32 = s0i32 & s1i32
	return s0i32
}
