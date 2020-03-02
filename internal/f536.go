package internal

import (
	"math"
	"unsafe"
)

func f536(ctx *Context, l0 float32) float32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
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
	var s3f32 float32
	_ = s3f32
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l2 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s1i32 = 1283457024
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1054867455
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = -1
			s1i32 = l1
			s2i32 = 964689920
			if uint32(s1i32) >= uint32(s2i32) {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl2
			}
			goto lbl1
		}
		s0f32 = l0
		s0f32 = float32(math.Abs(float64(s0f32)))
		l0 = s0f32
		s0i32 = l1
		s1i32 = 1066926079
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 1060110335
			if uint32(s0i32) <= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0f32 = l0
				s1f32 = l0
				s0f32 = s0f32 + s1f32
				s1f32 = -1
				s0f32 = s0f32 + s1f32
				s1f32 = l0
				s2f32 = 2
				s1f32 = s1f32 + s2f32
				s0f32 = s0f32 / s1f32
				l0 = s0f32
				s0i32 = 0
				goto lbl2
			}
			s0f32 = l0
			s1f32 = -1
			s0f32 = s0f32 + s1f32
			s1f32 = l0
			s2f32 = 1
			s1f32 = s1f32 + s2f32
			s0f32 = s0f32 / s1f32
			l0 = s0f32
			s0i32 = 1
			goto lbl2
		}
		s0i32 = l1
		s1i32 = 1075576831
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l0
			s1f32 = -1.5
			s0f32 = s0f32 + s1f32
			s1f32 = l0
			s2f32 = 1.5
			s1f32 = s1f32 * s2f32
			s2f32 = 1
			s1f32 = s1f32 + s2f32
			s0f32 = s0f32 / s1f32
			l0 = s0f32
			s0i32 = 2
			goto lbl2
		}
		s0f32 = -1
		s1f32 = l0
		s0f32 = s0f32 / s1f32
		l0 = s0f32
		s0i32 = 3
	lbl2:
		l1 = s0i32
		s0f32 = l0
		s1f32 = l0
		s0f32 = s0f32 * s1f32
		l4 = s0f32
		s1f32 = l4
		s0f32 = s0f32 * s1f32
		l3 = s0f32
		s1f32 = l3
		s2f32 = -0.106480174
		s1f32 = s1f32 * s2f32
		s2f32 = -0.19999158
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 * s1f32
		l5 = s0f32
		s0f32 = l4
		s1f32 = l3
		s2f32 = l3
		s3f32 = 0.061687607
		s2f32 = s2f32 * s3f32
		s3f32 = 0.14253636
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = 0.33333328
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 * s1f32
		l3 = s0f32
		s0i32 = l1
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l0
			s1f32 = l0
			s2f32 = l5
			s3f32 = l3
			s2f32 = s2f32 + s3f32
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 - s1f32
			return s0f32
		}
		s0i32 = l1
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l1 = s0i32
		s1i32 = 19024
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f32 = l0
		s2f32 = l5
		s3f32 = l3
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 * s2f32
		s2i32 = l1
		s3i32 = 19040
		s2i32 = s2i32 + s3i32
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 - s2f32
		s2f32 = l0
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 - s1f32
		l0 = s0f32
		s0f32 = -s0f32
		s1f32 = l0
		s2i32 = l2
		s3i32 = 0
		if s2i32 < s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l0 = s0f32
	lbl1:
		s0f32 = l0
		return s0f32
	}
	s0f32 = l0
	s1f32 = 1.5707963
	s2f32 = l0
	s1f32 = float32(math.Copysign(float64(s1f32), float64(s2f32)))
	s2i32 = l1
	s3i32 = 2139095040
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	return s0f32
}
