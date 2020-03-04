package internal

import (
	"unsafe"
)

func f650(ctx *Context, l0 int32, l1 int32) {
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
	var l21 float32
	_ = l21
	var l22 float64
	_ = l22
	var l23 float64
	_ = l23
	var l24 float64
	_ = l24
	var l25 float64
	_ = l25
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
	var s5f64 float64
	_ = s5f64
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s0f64 = float64(s0f32)
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f64 = float64(s1f32)
	s0f64 = s0f64 - s1f64
	l24 = s0f64
	s0f64 = 1
	l22 = s0f64
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l14 = s0f32
	s0f64 = float64(s0f32)
	s1i32 = l0
	s2i32 = 24
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l16 = s1f32
	s1f64 = float64(s1f32)
	s0f64 = s0f64 + s1f64
	l23 = s0f64
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f64 = float64(s1f32)
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2f64 = float64(s2f32)
	s1f64 = s1f64 - s2f64
	l25 = s1f64
	if s0f64 > s1f64 {
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
		s0f64 = l25
		s1f64 = l23
		s0f64 = s0f64 / s1f64
		l22 = s0f64
		s1f64 = 1
		s2f64 = l22
		s3f64 = 1
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f64 = s0f64
		} else {
			s0f64 = s1f64
		}
		l22 = s0f64
	}
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s0f64 = float64(s0f32)
	s1i32 = l0
	s2i32 = 36
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l15 = s1f32
	s1f64 = float64(s1f32)
	s0f64 = s0f64 + s1f64
	l23 = s0f64
	s1f64 = l24
	if s0f64 > s1f64 {
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
		s0f64 = l24
		s1f64 = l23
		s0f64 = s0f64 / s1f64
		l23 = s0f64
		s1f64 = l22
		s2f64 = l23
		s3f64 = l22
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f64 = s0f64
		} else {
			s0f64 = s1f64
		}
		l22 = s0f64
	}
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l17 = s0f32
	s0f64 = float64(s0f32)
	s1i32 = l0
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l18 = s1f32
	s1f64 = float64(s1f32)
	s0f64 = s0f64 + s1f64
	l23 = s0f64
	s1f64 = l25
	if s0f64 > s1f64 {
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
		s0f64 = l25
		s1f64 = l23
		s0f64 = s0f64 / s1f64
		l23 = s0f64
		s1f64 = l22
		s2f64 = l23
		s3f64 = l22
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f64 = s0f64
		} else {
			s0f64 = s1f64
		}
		l22 = s0f64
	}
	s0i32 = l0
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0f32
	s0f64 = float64(s0f32)
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l20 = s1f32
	s1f64 = float64(s1f32)
	s0f64 = s0f64 + s1f64
	l23 = s0f64
	s1f64 = l24
	if s0f64 > s1f64 {
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
		s0f64 = l24
		s1f64 = l23
		s0f64 = s0f64 / s1f64
		l23 = s0f64
		s1f64 = l22
		s2f64 = l23
		s3f64 = l22
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f64 = s0f64
		} else {
			s0f64 = s1f64
		}
		l22 = s0f64
	}
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0f32 = l14
	s1f32 = l16
	s0f32 = s0f32 + s1f32
	l13 = s0f32
	s1f32 = l14
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = 0
		l16 = s0f32
		s0i32 = l10
		goto lbl5
	}
	s0f32 = l13
	s1f32 = l16
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = 0
	l14 = s0f32
	s0i32 = l11
lbl5:
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl4:
	s0f32 = l12
	s1f32 = l15
	s0f32 = s0f32 + s1f32
	l21 = s0f32
	s1f32 = l12
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		l2 = s0i32
		s0f32 = 0
		l15 = s0f32
		s0f32 = l12
		l13 = s0f32
		goto lbl8
	}
	s0f32 = 0
	l13 = s0f32
	s0i32 = l4
	l2 = s0i32
	s0f32 = l21
	s1f32 = l15
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
lbl8:
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0f32 = l13
	l12 = s0f32
lbl7:
	s0i32 = l7
	l2 = s0i32
	s0f32 = l17
	s1f32 = l18
	s0f32 = s0f32 + s1f32
	l13 = s0f32
	s1f32 = l17
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		l2 = s0i32
		s0f32 = l13
		s1f32 = l18
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl10:
	s0i32 = l3
	l2 = s0i32
	s0f32 = l19
	s1f32 = l20
	s0f32 = s0f32 + s1f32
	l13 = s0f32
	s1f32 = l19
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		l2 = s0i32
		s0f32 = l13
		s1f32 = l20
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl12:
	s0f64 = l22
	s1f64 = 1
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l11
	s1f64 = l22
	s2f32 = l14
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l10
	s1f64 = l22
	s2f32 = l16
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l4
	s1f64 = l22
	s2f64 = l25
	s3f32 = l13
	s4f32 = l14
	s3f32 = s3f32 + s4f32
	s3f64 = float64(s3f32)
	if s2f64 < s3f64 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s3i32 = 1
	s2i32 = s2i32 ^ s3i32
	if s2i32 != 0 {
		s2f32 = l12
	} else {
		s2i32 = l11
		s3i32 = l10
		s4f32 = l13
		s5f32 = l14
		if s4f32 > s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		l2 = s4i32
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l9 = s2i32
		s2f64 = l25
		s3i32 = l10
		s4i32 = l11
		s5i32 = l2
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		l13 = s3f32
		s4f64 = l25
		s5f32 = l13
		s5f64 = float64(s5f32)
		s4f64 = s4f64 - s5f64
		s4f32 = float32(s4f64)
		l12 = s4f32
		s3f32 = s3f32 + s4f32
		s3f64 = float64(s3f32)
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 1
		s2i32 = s2i32 ^ s3i32
		if s2i32 != 0 {
			goto lbl16
		}
		s2f64 = l25
		s3f32 = l13
		s4f32 = l12
		s5f32 = 0
		s4f32 = f133(ctx, s4f32, s5f32)
		l12 = s4f32
		s3f32 = s3f32 + s4f32
		s3f64 = float64(s3f32)
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 1
		s2i32 = s2i32 ^ s3i32
		if s2i32 != 0 {
			goto lbl16
		}
		s2f32 = l12
		s3f32 = 0
		s2f32 = f133(ctx, s2f32, s3f32)
		l12 = s2f32
	lbl16:
		s2i32 = l9
		s3f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])) = s3f32
		s2i32 = l5
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l15 = s2f32
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	}
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l12 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l5
	s1f64 = l22
	s2f32 = l15
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0f64 = l24
	s1f32 = l12
	s2f32 = l13
	s1f32 = s1f32 + s2f32
	s1f64 = float64(s1f32)
	if s0f64 < s1f64 {
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
		s0i32 = l4
		s1i32 = l5
		s2f32 = l12
		s3f32 = l13
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l2 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l9 = s0i32
		s0f64 = l24
		s1i32 = l5
		s2i32 = l4
		s3i32 = l2
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l13 = s1f32
		s2f64 = l24
		s3f32 = l13
		s3f64 = float64(s3f32)
		s2f64 = s2f64 - s3f64
		s2f32 = float32(s2f64)
		l12 = s2f32
		s1f32 = s1f32 + s2f32
		s1f64 = float64(s1f32)
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl18
		}
		s0f64 = l24
		s1f32 = l13
		s2f32 = l12
		s3f32 = 0
		s2f32 = f133(ctx, s2f32, s3f32)
		l12 = s2f32
		s1f32 = s1f32 + s2f32
		s1f64 = float64(s1f32)
		if s0f64 < s1f64 {
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
		s1f32 = 0
		s0f32 = f133(ctx, s0f32, s1f32)
		l12 = s0f32
	lbl18:
		s0i32 = l9
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	}
	s0i32 = l6
	s1f64 = l22
	s2i32 = l6
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l12 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l7
	s1f64 = l22
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0f64 = l25
	s1f32 = l12
	s2f32 = l13
	s1f32 = s1f32 + s2f32
	s1f64 = float64(s1f32)
	if s0f64 < s1f64 {
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
		s1i32 = l7
		s2f32 = l12
		s3f32 = l13
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l2 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l9 = s0i32
		s0f64 = l25
		s1i32 = l7
		s2i32 = l6
		s3i32 = l2
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l13 = s1f32
		s2f64 = l25
		s3f32 = l13
		s3f64 = float64(s3f32)
		s2f64 = s2f64 - s3f64
		s2f32 = float32(s2f64)
		l12 = s2f32
		s1f32 = s1f32 + s2f32
		s1f64 = float64(s1f32)
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl20
		}
		s0f64 = l25
		s1f32 = l13
		s2f32 = l12
		s3f32 = 0
		s2f32 = f133(ctx, s2f32, s3f32)
		l12 = s2f32
		s1f32 = s1f32 + s2f32
		s1f64 = float64(s1f32)
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl20
		}
		s0f32 = l12
		s1f32 = 0
		s0f32 = f133(ctx, s0f32, s1f32)
		l12 = s0f32
	lbl20:
		s0i32 = l9
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	}
	s0i32 = l8
	s1f64 = l22
	s2i32 = l8
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l12 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l3
	s1f64 = l22
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2f64 = float64(s2f32)
	s1f64 = s1f64 * s2f64
	s1f32 = float32(s1f64)
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0f64 = l24
	s1f32 = l12
	s2f32 = l13
	s1f32 = s1f32 + s2f32
	s1f64 = float64(s1f32)
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l8
	s1i32 = l3
	s2f32 = l12
	s3f32 = l13
	if s2f32 > s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l2 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s0f64 = l24
	s1i32 = l3
	s2i32 = l8
	s3i32 = l2
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l13 = s1f32
	s2f64 = l24
	s3f32 = l13
	s3f64 = float64(s3f32)
	s2f64 = s2f64 - s3f64
	s2f32 = float32(s2f64)
	l12 = s2f32
	s1f32 = s1f32 + s2f32
	s1f64 = float64(s1f32)
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl21
	}
	s0f64 = l24
	s1f32 = l13
	s2f32 = l12
	s3f32 = 0
	s2f32 = f133(ctx, s2f32, s3f32)
	l12 = s2f32
	s1f32 = s1f32 + s2f32
	s1f64 = float64(s1f32)
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl21
	}
	s0f32 = l12
	s1f32 = 0
	s0f32 = f133(ctx, s0f32, s1f32)
	l12 = s0f32
lbl21:
	s0i32 = l9
	s1f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
lbl14:
	s0i32 = l11
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f32 = 0
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl22
		}
	}
	s0i32 = l11
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1
lbl22:
	l2 = s0i32
	s0i32 = l10
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f32 = 0
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl24
		}
	}
	s0i32 = l10
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
lbl24:
	l3 = s0i32
	s0i32 = l6
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f32 = 0
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl26
		}
	}
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
lbl26:
	l2 = s0i32
	s0i32 = l7
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f32 = 0
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl28
		}
	}
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl28
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l13 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l15 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l14 = s1f32
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s1f32 = l16
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		return
	}
	s0i32 = l0
	s1f32 = l14
	s2f32 = l13
	s3f32 = l13
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
	l16 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l0
	s1f32 = l15
	s2f32 = l12
	s3f32 = l12
	s4f32 = l15
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
	l17 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l0
	s1f32 = l14
	s2f32 = l13
	s3f32 = l14
	s4f32 = l13
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
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s1f32 = l15
	s2f32 = l12
	s3f32 = l15
	s4f32 = l12
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
	l12 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0f32 = l12
	s1f32 = l17
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
	s1i32 = 0
	s2f32 = l13
	s3f32 = l16
	if s2f32 < s3f32 {
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
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		return
	}
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	return
lbl28:
	s0i32 = l0
	f1880(ctx, s0i32)
}
