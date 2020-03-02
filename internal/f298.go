package internal

import (
	"math"
	"unsafe"
)

func f298(ctx *Context, l0 float32, l1 float32) float32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
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
	s0f32 = l1
	s0i32 = int32(math.Float32bits(s0f32))
	l4 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	s1i32 = 2139095040
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s0i32 = int32(math.Float32bits(s0f32))
		l2 = s0i32
		s1i32 = 2147483647
		s0i32 = s0i32 & s1i32
		l3 = s0i32
		s1i32 = 2139095041
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0f32 = l0
	s1f32 = l1
	s0f32 = s0f32 + s1f32
	return s0f32
lbl0:
	s0i32 = l4
	s1i32 = 1065353216
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s0f32 = f536(ctx, s0f32)
		return s0f32
	}
	s0i32 = l4
	s1i32 = 30
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s1i32 = l2
	s2i32 = 31
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s0i32 = s0i32 | s1i32
	l2 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl5
		case 1:
			goto lbl7
		default:
			goto lbl4
		}
	lbl7:
		s0f32 = -3.1415927
		return s0f32
	}
	s0i32 = l5
	s1i32 = 2139095040
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = 1.5707964
			s1f32 = l0
			s0f32 = float32(math.Copysign(float64(s0f32), float64(s1f32)))
			return s0f32
		}
		s0i32 = l3
		s1i32 = 2139095040
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 0
		s2i32 = l5
		s3i32 = 218103808
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		if uint32(s2i32) >= uint32(s3i32) {
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
			s0f32 = 1.5707964
			s1f32 = l0
			s0f32 = float32(math.Copysign(float64(s0f32), float64(s1f32)))
			return s0f32
		}
		s0i32 = l3
		s1i32 = 218103808
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = 0
			s1i32 = l4
			if s1i32 != 0 {
				goto lbl11
			}
		}
		s0f32 = l0
		s1f32 = l1
		s0f32 = s0f32 / s1f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s0f32 = f536(ctx, s0f32)
	lbl11:
		l0 = s0f32
		s0i32 = l2
		s1i32 = 2
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			switch s0i32 {
			case 0:
				goto lbl15
			case 1:
				goto lbl14
			default:
				goto lbl4
			}
		lbl15:
			s0f32 = l0
			s0f32 = -s0f32
			return s0f32
		lbl14:
			s0f32 = 3.1415927
			s1f32 = l0
			s2f32 = 8.742278e-08
			s1f32 = s1f32 + s2f32
			s0f32 = s0f32 - s1f32
			return s0f32
		}
		s0f32 = l0
		s1f32 = 8.742278e-08
		s0f32 = s0f32 + s1f32
		s1f32 = -3.1415927
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0i32 = l3
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 19072
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	return s0f32
lbl5:
	s0f32 = 3.1415927
	l0 = s0f32
lbl4:
	s0f32 = l0
	return s0f32
lbl3:
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 19056
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	return s0f32
}
