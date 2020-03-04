package internal

import (
	"math"
	"unsafe"
)

func f562(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var l23 float32
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
	s0i32 = l0
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l20 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l21 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l22 = s0f32
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l23 = s0f32
	s0i32 = l12
	l7 = s0i32
	s0i32 = 0
	l2 = s0i32
	s0i32 = 1
	l1 = s0i32
lbl1:
	s0i32 = l1
	l9 = s0i32
	s0i32 = l2
	l11 = s0i32
	s0i32 = l0
	l2 = s0i32
	s0i32 = l5
	l0 = s0i32
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+18)])))
	l8 = s0i32
	s0i32 = l3
	s1i32 = l0
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l4
		s2i32 = l8
		if s1i32 == s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l16 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l18 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l19 = s0f32
	s0i32 = l4
	s1i32 = l8
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l18
		s1f32 = l20
		s2f32 = l17
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 * s1f32
		s1f32 = l16
		s2f32 = l22
		s3f32 = l19
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 - s1f32
		l16 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 5.9604645e-08
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = 0
		return s0i32
	}
	s0f32 = l18
	s1f32 = l21
	s2f32 = l17
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l16
	s2f32 = l23
	s3f32 = l19
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l16 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 5.9604645e-08
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
		goto lbl4
	}
	s0i32 = 0
	return s0i32
lbl4:
	s0i32 = l6
	l5 = s0i32
	s0f32 = l16
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
lbl2:
	l1 = s0i32
	s0i32 = l5
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l8 = s1i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		if s0i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l10
	s1i32 = l5
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	l2 = s0i32
	goto lbl6
lbl8:
	s0i32 = l2
	s1i32 = l9
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l14 = s1i32
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l5
	s1i32 = l14
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l10
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l5
	s1i32 = l9
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l7 = s0i32
	goto lbl9
lbl10:
	s0i32 = l5
	s1i32 = l9
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		if s0i32 != 0 {
			goto lbl9
		}
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	goto lbl6
lbl9:
	s0i32 = l9
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l15 = s0i32
	s0i32 = l11
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = l2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l13 = s0i32
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l5
	s1i32 = l9
	s2i32 = 0
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l7 = s1i32
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	l9 = s1i32
	s2i32 = l14
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l9
	s2i32 = l15
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = l13
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	goto lbl12
lbl13:
	s0i32 = l10
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l10
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l8
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l15
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = l13
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
lbl12:
	s0i32 = l11
	s1i32 = l13
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l5 = s0i32
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
lbl6:
	s0i32 = l0
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l6
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l3 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l5 = s0i32
		if s0i32 != 0 {
			s0i32 = 0
			s1i32 = l5
			s2i32 = l6
			s3i32 = l6
			s4i32 = 8
			s3i32 = s3i32 + s4i32
			s4i32 = l6
			s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)])))
			s5i32 = l6
			s5i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s5i32+18)])))
			s1i32 = f166(ctx, s1i32, s2i32, s3i32, s4i32, s5i32)
			if s1i32 != 0 {
				goto lbl14
			}
		}
		s0i32 = l3
		if s0i32 != 0 {
			s0i32 = 0
			s1i32 = l3
			s2i32 = l6
			s3i32 = l6
			s4i32 = 8
			s3i32 = s3i32 + s4i32
			s4i32 = l6
			s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+16)])))
			s5i32 = l6
			s5i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s5i32+18)])))
			s1i32 = f166(ctx, s1i32, s2i32, s3i32, s4i32, s5i32)
			if s1i32 != 0 {
				goto lbl14
			}
		}
		s0i32 = l0
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l6
			s1i32 = l0
			s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])))
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint16(s1i32)
			s0i32 = l6
			s1i32 = l0
			s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+18)])))
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+18)])) = uint16(s1i32)
			s0i32 = l6
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			l5 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
			goto lbl20
		}
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l3
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	lbl20:
		s0i32 = l5
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s2i32 = l0
		if s1i32 == s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		if s2i32 == 0 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = -2401053088876216593
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l12
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		goto lbl16
	}
	s0i32 = l12
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
lbl16:
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
lbl15:
	s0i32 = 1
lbl14:
	return s0i32
}
