package internal

import (
	"math"
)

func f1359(ctx *Context, l0 float64) float64 {
	var l1 int32
	_ = l1
	var l2 int64
	_ = l2
	var l3 float64
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
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
	var s6f64 float64
	_ = s6f64
	var s7f64 float64
	_ = s7f64
	var s8f64 float64
	_ = s8f64
	var s9f64 float64
	_ = s9f64
	s0f64 = l0
	s0i64 = int64(math.Float64bits(s0f64))
	l2 = s0i64
	s1i64 = 32
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s1i32 = 1072693248
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = l2
		s0i32 = int32(s0i64)
		s1i32 = l1
		s2i32 = -1072693248
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 | s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f64 = 3.141592653589793
			s1f64 = 0
			s2i64 = l2
			s3i64 = 0
			if s2i64 < s3i64 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0f64 = s0f64
			} else {
				s0f64 = s1f64
			}
			return s0f64
		}
		s0f64 = 0
		s1f64 = l0
		s2f64 = l0
		s1f64 = s1f64 - s2f64
		s0f64 = s0f64 / s1f64
		return s0f64
	}
	s0i32 = l1
	s1i32 = 1071644671
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = 1.5707963267948966
		s1i32 = l1
		s2i32 = 1012924417
		if uint32(s1i32) < uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl2
		}
		s0f64 = 6.123233995736766e-17
		s1f64 = l0
		s2f64 = l0
		s1f64 = s1f64 * s2f64
		l3 = s1f64
		s2f64 = l3
		s3f64 = l3
		s4f64 = l3
		s5f64 = l3
		s6f64 = l3
		s7f64 = 3.479331075960212e-05
		s6f64 = s6f64 * s7f64
		s7f64 = 0.0007915349942898145
		s6f64 = s6f64 + s7f64
		s5f64 = s5f64 * s6f64
		s6f64 = -0.04005553450067941
		s5f64 = s5f64 + s6f64
		s4f64 = s4f64 * s5f64
		s5f64 = 0.20121253213486293
		s4f64 = s4f64 + s5f64
		s3f64 = s3f64 * s4f64
		s4f64 = -0.3255658186224009
		s3f64 = s3f64 + s4f64
		s2f64 = s2f64 * s3f64
		s3f64 = 0.16666666666666666
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 * s2f64
		s2f64 = l3
		s3f64 = l3
		s4f64 = l3
		s5f64 = l3
		s6f64 = 0.07703815055590194
		s5f64 = s5f64 * s6f64
		s6f64 = -0.6882839716054533
		s5f64 = s5f64 + s6f64
		s4f64 = s4f64 * s5f64
		s5f64 = 2.0209457602335057
		s4f64 = s4f64 + s5f64
		s3f64 = s3f64 * s4f64
		s4f64 = -2.403394911734414
		s3f64 = s3f64 + s4f64
		s2f64 = s2f64 * s3f64
		s3f64 = 1
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 / s2f64
		s2f64 = l0
		s1f64 = s1f64 * s2f64
		s0f64 = s0f64 - s1f64
		s1f64 = l0
		s0f64 = s0f64 - s1f64
		s1f64 = 1.5707963267948966
		s0f64 = s0f64 + s1f64
		return s0f64
	}
	s0i64 = l2
	s1i64 = -1
	if s0i64 <= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f64 = 1.5707963267948966
		s1f64 = l0
		s2f64 = 1
		s1f64 = s1f64 + s2f64
		s2f64 = 0.5
		s1f64 = s1f64 * s2f64
		l0 = s1f64
		s1f64 = math.Sqrt(s1f64)
		l3 = s1f64
		s2f64 = l3
		s3f64 = l0
		s4f64 = l0
		s5f64 = l0
		s6f64 = l0
		s7f64 = l0
		s8f64 = l0
		s9f64 = 3.479331075960212e-05
		s8f64 = s8f64 * s9f64
		s9f64 = 0.0007915349942898145
		s8f64 = s8f64 + s9f64
		s7f64 = s7f64 * s8f64
		s8f64 = -0.04005553450067941
		s7f64 = s7f64 + s8f64
		s6f64 = s6f64 * s7f64
		s7f64 = 0.20121253213486293
		s6f64 = s6f64 + s7f64
		s5f64 = s5f64 * s6f64
		s6f64 = -0.3255658186224009
		s5f64 = s5f64 + s6f64
		s4f64 = s4f64 * s5f64
		s5f64 = 0.16666666666666666
		s4f64 = s4f64 + s5f64
		s3f64 = s3f64 * s4f64
		s4f64 = l0
		s5f64 = l0
		s6f64 = l0
		s7f64 = l0
		s8f64 = 0.07703815055590194
		s7f64 = s7f64 * s8f64
		s8f64 = -0.6882839716054533
		s7f64 = s7f64 + s8f64
		s6f64 = s6f64 * s7f64
		s7f64 = 2.0209457602335057
		s6f64 = s6f64 + s7f64
		s5f64 = s5f64 * s6f64
		s6f64 = -2.403394911734414
		s5f64 = s5f64 + s6f64
		s4f64 = s4f64 * s5f64
		s5f64 = 1
		s4f64 = s4f64 + s5f64
		s3f64 = s3f64 / s4f64
		s2f64 = s2f64 * s3f64
		s3f64 = -6.123233995736766e-17
		s2f64 = s2f64 + s3f64
		s1f64 = s1f64 + s2f64
		s0f64 = s0f64 - s1f64
		l0 = s0f64
		s1f64 = l0
		s0f64 = s0f64 + s1f64
		return s0f64
	}
	s0f64 = 1
	s1f64 = l0
	s0f64 = s0f64 - s1f64
	s1f64 = 0.5
	s0f64 = s0f64 * s1f64
	l0 = s0f64
	s1f64 = l0
	s2f64 = l0
	s3f64 = l0
	s4f64 = l0
	s5f64 = l0
	s6f64 = 3.479331075960212e-05
	s5f64 = s5f64 * s6f64
	s6f64 = 0.0007915349942898145
	s5f64 = s5f64 + s6f64
	s4f64 = s4f64 * s5f64
	s5f64 = -0.04005553450067941
	s4f64 = s4f64 + s5f64
	s3f64 = s3f64 * s4f64
	s4f64 = 0.20121253213486293
	s3f64 = s3f64 + s4f64
	s2f64 = s2f64 * s3f64
	s3f64 = -0.3255658186224009
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 * s2f64
	s2f64 = 0.16666666666666666
	s1f64 = s1f64 + s2f64
	s0f64 = s0f64 * s1f64
	s1f64 = l0
	s2f64 = l0
	s3f64 = l0
	s4f64 = l0
	s5f64 = 0.07703815055590194
	s4f64 = s4f64 * s5f64
	s5f64 = -0.6882839716054533
	s4f64 = s4f64 + s5f64
	s3f64 = s3f64 * s4f64
	s4f64 = 2.0209457602335057
	s3f64 = s3f64 + s4f64
	s2f64 = s2f64 * s3f64
	s3f64 = -2.403394911734414
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 * s2f64
	s2f64 = 1
	s1f64 = s1f64 + s2f64
	s0f64 = s0f64 / s1f64
	s1f64 = l0
	s1f64 = math.Sqrt(s1f64)
	l3 = s1f64
	s0f64 = s0f64 * s1f64
	s1f64 = l0
	s2f64 = l3
	s2i64 = int64(math.Float64bits(s2f64))
	s3i64 = -4294967296
	s2i64 = s2i64 & s3i64
	s2f64 = math.Float64frombits(uint64(s2i64))
	l0 = s2f64
	s3f64 = l0
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s2f64 = l3
	s3f64 = l0
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 / s2f64
	s0f64 = s0f64 + s1f64
	s1f64 = l0
	s0f64 = s0f64 + s1f64
	l0 = s0f64
	s1f64 = l0
	s0f64 = s0f64 + s1f64
lbl2:
	return s0f64
}
