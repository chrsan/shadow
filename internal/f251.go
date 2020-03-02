package internal

import (
	"math"
	"unsafe"
)

func f251(ctx *Context, l0 float32, l1 float32) float32 {
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
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
	var s9f32 float32
	_ = s9f32
	var s10f32 float32
	_ = s10f32
	s0f32 = 1
	l9 = s0f32
	s0f32 = l0
	s0i32 = int32(math.Float32bits(s0f32))
	l4 = s0i32
	s1i32 = 1065353216
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l1
	s0i32 = int32(math.Float32bits(s0f32))
	l5 = s0i32
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l4
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l3 = s0i32
	s1i32 = 2139095040
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l2
	s3i32 = 2139095041
	if uint32(s2i32) < uint32(s3i32) {
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
		s0f32 = l0
		s1f32 = l1
		s0f32 = s0f32 + s1f32
		return s0f32
	}
	s0i32 = 0
	s1i32 = l4
	s2i32 = -1
	if s1i32 > s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = 2
	s1i32 = l2
	s2i32 = 1266679807
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = 0
	s1i32 = l2
	s2i32 = 1065353216
	if uint32(s1i32) < uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = 0
	s1i32 = l2
	s2i32 = 150
	s3i32 = l2
	s4i32 = 23
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 - s3i32
	l6 = s2i32
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l7 = s1i32
	s2i32 = l6
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l2
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = 2
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
lbl4:
	l6 = s0i32
	s0i32 = l2
	s1i32 = 1065353216
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 2139095040
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l3
		s1i32 = 1065353216
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l3
		s1i32 = 1065353217
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l1
			s1f32 = 0
			s2i32 = l5
			s3i32 = -1
			if s2i32 > s3i32 {
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
		s0f32 = 0
		s1f32 = l1
		s1f32 = -s1f32
		s2i32 = l5
		s3i32 = -1
		if s2i32 > s3i32 {
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
	s0f32 = l0
	s1f32 = 1
	s2f32 = l0
	s1f32 = s1f32 / s2f32
	s2i32 = l5
	s3i32 = -1
	if s2i32 > s3i32 {
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
lbl5:
	s0i32 = l5
	s1i32 = 1073741824
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s1f32 = l0
		s0f32 = s0f32 * s1f32
		return s0f32
	}
	s0i32 = l4
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l5
	s1i32 = 1056964608
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0f32 = l0
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	return s0f32
lbl9:
	s0f32 = l0
	s0f32 = float32(math.Abs(float64(s0f32)))
	l8 = s0f32
	s0i32 = l3
	s1i32 = 0
	s2i32 = l3
	s3i32 = 1073741824
	s2i32 = s2i32 | s3i32
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
		s0f32 = 1
		s1f32 = l8
		s0f32 = s0f32 / s1f32
		s1f32 = l8
		s2i32 = l5
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
		l9 = s0f32
		s0i32 = l4
		s1i32 = -1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l6
		s1i32 = l3
		s2i32 = -1065353216
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 | s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l9
			s1f32 = l9
			s0f32 = s0f32 - s1f32
			l0 = s0f32
			s1f32 = l0
			s0f32 = s0f32 / s1f32
			return s0f32
		}
		s0f32 = l9
		s0f32 = -s0f32
		s1f32 = l9
		s2i32 = l6
		s3i32 = 1
		if s2i32 == s3i32 {
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
	s0i32 = l4
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l6
	s1i32 = 1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	if s0i32 != 0 {
		s0f32 = l0
		s1f32 = l0
		s0f32 = s0f32 - s1f32
		l0 = s0f32
		s1f32 = l0
		s0f32 = s0f32 / s1f32
		return s0f32
	}
	s0f32 = -1
	l9 = s0f32
lbl12:
	s0i32 = l2
	s1i32 = 1291845633
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1065353207
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l9
			s1f32 = 1e+30
			s0f32 = s0f32 * s1f32
			s1f32 = 1e+30
			s0f32 = s0f32 * s1f32
			s1f32 = l9
			s2f32 = 1e-30
			s1f32 = s1f32 * s2f32
			s2f32 = 1e-30
			s1f32 = s1f32 * s2f32
			s2i32 = l5
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
			return s0f32
		}
		s0i32 = l3
		s1i32 = 1065353224
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l9
			s1f32 = 1e+30
			s0f32 = s0f32 * s1f32
			s1f32 = 1e+30
			s0f32 = s0f32 * s1f32
			s1f32 = l9
			s2f32 = 1e-30
			s1f32 = s1f32 * s2f32
			s2f32 = 1e-30
			s1f32 = s1f32 * s2f32
			s2i32 = l5
			s3i32 = 0
			if s2i32 > s3i32 {
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
		s0f32 = l8
		s1f32 = -1
		s0f32 = s0f32 + s1f32
		l0 = s0f32
		s1f32 = 1.442688
		s0f32 = s0f32 * s1f32
		l8 = s0f32
		s1f32 = l0
		s2f32 = 7.0526075e-06
		s1f32 = s1f32 * s2f32
		s2f32 = l0
		s3f32 = l0
		s2f32 = s2f32 * s3f32
		s3f32 = 0.5
		s4f32 = l0
		s5f32 = l0
		s6f32 = -0.25
		s5f32 = s5f32 * s6f32
		s6f32 = 0.33333334
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s3f32 = -1.442695
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		l11 = s1f32
		s0f32 = s0f32 + s1f32
		s0i32 = int32(math.Float32bits(s0f32))
		s1i32 = -4096
		s0i32 = s0i32 & s1i32
		s0f32 = math.Float32frombits(uint32(s0i32))
		l0 = s0f32
		s1f32 = l8
		s0f32 = s0f32 - s1f32
		goto lbl14
	}
	s0f32 = l8
	s1f32 = 1.6777216e+07
	s0f32 = s0f32 * s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = l3
	s2i32 = l3
	s3i32 = 8388608
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l3 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s1i32 = 8388607
	s0i32 = s0i32 & s1i32
	l6 = s0i32
	s1i32 = 1065353216
	s0i32 = s0i32 | s1i32
	l2 = s0i32
	s0i32 = l4
	s1i32 = 23
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = -151
	s2i32 = -127
	s3i32 = l3
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = 0
	l4 = s0i32
	s0i32 = l6
	s1i32 = 1885298
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l6
	s1i32 = 6140887
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l4 = s0i32
		goto lbl18
	}
	s0i32 = l2
	s1i32 = -8388608
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl18:
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l6 = s0i32
	s1i32 = 19120
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l13 = s0f32
	s1i32 = l2
	s1f32 = math.Float32frombits(uint32(s1i32))
	l11 = s1f32
	s2i32 = l6
	s3i32 = 19104
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l10 = s2f32
	s1f32 = s1f32 - s2f32
	l12 = s1f32
	s2f32 = 1
	s3f32 = l10
	s4f32 = l11
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	l8 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = -4096
	s1i32 = s1i32 & s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l0 = s1f32
	s2f32 = l0
	s3f32 = l0
	s2f32 = s2f32 * s3f32
	l15 = s2f32
	s3f32 = 3
	s2f32 = s2f32 + s3f32
	s3f32 = l8
	s4f32 = l0
	s3f32 = s3f32 + s4f32
	s4f32 = l14
	s5f32 = l12
	s6f32 = l0
	s7i32 = l2
	s8i32 = 1
	s7i32 = s7i32 >> (uint32(s8i32) & 31)
	s8i32 = -536875008
	s7i32 = s7i32 & s8i32
	s8i32 = 536870912
	s7i32 = s7i32 | s8i32
	s8i32 = l4
	s9i32 = 21
	s8i32 = s8i32 << (uint32(s9i32) & 31)
	s7i32 = s7i32 + s8i32
	s8i32 = 4194304
	s7i32 = s7i32 + s8i32
	s7f32 = math.Float32frombits(uint32(s7i32))
	l12 = s7f32
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 - s6f32
	s6f32 = l0
	s7f32 = l11
	s8f32 = l12
	s9f32 = l10
	s8f32 = s8f32 - s9f32
	s7f32 = s7f32 - s8f32
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	l11 = s4f32
	s3f32 = s3f32 * s4f32
	s4f32 = l8
	s5f32 = l8
	s4f32 = s4f32 * s5f32
	l0 = s4f32
	s5f32 = l0
	s4f32 = s4f32 * s5f32
	s5f32 = l0
	s6f32 = l0
	s7f32 = l0
	s8f32 = l0
	s9f32 = l0
	s10f32 = 0.20697501
	s9f32 = s9f32 * s10f32
	s10f32 = 0.23066075
	s9f32 = s9f32 + s10f32
	s8f32 = s8f32 * s9f32
	s9f32 = 0.27272812
	s8f32 = s8f32 + s9f32
	s7f32 = s7f32 * s8f32
	s8f32 = 0.33333334
	s7f32 = s7f32 + s8f32
	s6f32 = s6f32 * s7f32
	s7f32 = 0.42857143
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 * s6f32
	s6f32 = 0.6
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	l10 = s3f32
	s2f32 = s2f32 + s3f32
	s2i32 = int32(math.Float32bits(s2f32))
	s3i32 = -4096
	s2i32 = s2i32 & s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l0 = s2f32
	s1f32 = s1f32 * s2f32
	l12 = s1f32
	s2f32 = l11
	s3f32 = l0
	s2f32 = s2f32 * s3f32
	s3f32 = l8
	s4f32 = l10
	s5f32 = l0
	s6f32 = -3
	s5f32 = s5f32 + s6f32
	s6f32 = l15
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s1f32 = s1f32 + s2f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = -4096
	s1i32 = s1i32 & s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l0 = s1f32
	s2f32 = 0.96191406
	s1f32 = s1f32 * s2f32
	l10 = s1f32
	s2i32 = l6
	s3i32 = 19112
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3f32 = l8
	s4f32 = l0
	s5f32 = l12
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 - s4f32
	s4f32 = 0.9617967
	s3f32 = s3f32 * s4f32
	s4f32 = l0
	s5f32 = -0.000117368574
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 + s3f32
	l11 = s2f32
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	s1i32 = l3
	s1f32 = float32(s1i32)
	l8 = s1f32
	s0f32 = s0f32 + s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = -4096
	s0i32 = s0i32 & s1i32
	s0f32 = math.Float32frombits(uint32(s0i32))
	l0 = s0f32
	s1f32 = l8
	s0f32 = s0f32 - s1f32
	s1f32 = l13
	s0f32 = s0f32 - s1f32
	s1f32 = l10
	s0f32 = s0f32 - s1f32
lbl14:
	l10 = s0f32
	s0f32 = l0
	s1i32 = l5
	s2i32 = -4096
	s1i32 = s1i32 & s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l13 = s1f32
	s0f32 = s0f32 * s1f32
	l8 = s0f32
	s1f32 = l11
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	s2f32 = l1
	s1f32 = s1f32 * s2f32
	s2f32 = l1
	s3f32 = l13
	s2f32 = s2f32 - s3f32
	s3f32 = l0
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l0 = s1f32
	s0f32 = s0f32 + s1f32
	l1 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l2 = s0i32
	s1i32 = 1124073473
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = 1124073472
	l4 = s0i32
	s0i32 = l2
	s1i32 = 1124073472
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l0
		s1f32 = 4.2995666e-08
		s0f32 = s0f32 + s1f32
		s1f32 = l1
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl21
		}
		goto lbl1
	}
	s0i32 = l2
	s1i32 = 2147483647
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s1i32 = 1125515265
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = -1021968384
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0f32 = l0
	s1f32 = l1
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl23
	}
	goto lbl0
lbl23:
	s0i32 = 0
	l3 = s0i32
	s0i32 = l4
	s1i32 = 1056964609
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl20
	}
lbl21:
	s0i32 = 0
	s1i32 = 8388608
	s2i32 = l4
	s3i32 = 23
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = -126
	s2i32 = s2i32 + s3i32
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s2i32 = 8388607
	s1i32 = s1i32 & s2i32
	s2i32 = 8388608
	s1i32 = s1i32 | s2i32
	s2i32 = 150
	s3i32 = l5
	s4i32 = 23
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s4i32 = 255
	s3i32 = s3i32 & s4i32
	l4 = s3i32
	s2i32 = s2i32 - s3i32
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l3 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = l3
	s2i32 = l2
	s3i32 = 0
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s0f32 = l0
	s1f32 = l8
	s2i32 = -8388608
	s3i32 = l4
	s4i32 = -127
	s3i32 = s3i32 + s4i32
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s3i32 = l5
	s2i32 = s2i32 & s3i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s0f32 = s0f32 + s1f32
	s0i32 = int32(math.Float32bits(s0f32))
	l2 = s0i32
lbl20:
	s0f32 = l9
	s1i32 = l2
	s2i32 = -32768
	s1i32 = s1i32 & s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l1 = s1f32
	s2f32 = 0.69314575
	s1f32 = s1f32 * s2f32
	l9 = s1f32
	s2f32 = l1
	s3f32 = 1.4286065e-06
	s2f32 = s2f32 * s3f32
	s3f32 = l0
	s4f32 = l1
	s5f32 = l8
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 - s4f32
	s4f32 = 0.6931472
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	l8 = s2f32
	s1f32 = s1f32 + s2f32
	l0 = s1f32
	s2f32 = l0
	s3f32 = l0
	s4f32 = l0
	s5f32 = l0
	s4f32 = s4f32 * s5f32
	l1 = s4f32
	s5f32 = l1
	s6f32 = l1
	s7f32 = l1
	s8f32 = l1
	s9f32 = 4.138137e-08
	s8f32 = s8f32 * s9f32
	s9f32 = -1.6533902e-06
	s8f32 = s8f32 + s9f32
	s7f32 = s7f32 * s8f32
	s8f32 = 6.613756e-05
	s7f32 = s7f32 + s8f32
	s6f32 = s6f32 * s7f32
	s7f32 = -0.0027777778
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 * s6f32
	s6f32 = 0.16666667
	s5f32 = s5f32 + s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l1 = s3f32
	s2f32 = s2f32 * s3f32
	s3f32 = l1
	s4f32 = -2
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 / s3f32
	s3f32 = l8
	s4f32 = l0
	s5f32 = l9
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 - s4f32
	l1 = s3f32
	s4f32 = l0
	s5f32 = l1
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 - s2f32
	s2f32 = 1
	s1f32 = s1f32 + s2f32
	l0 = s1f32
	s1i32 = int32(math.Float32bits(s1f32))
	s2i32 = l3
	s3i32 = 23
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = 8388607
	if s1i32 <= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l0
		s2i32 = l3
		s1f32 = f1355(ctx, s1f32, s2i32)
		goto lbl24
	}
	s1i32 = l2
	s1f32 = math.Float32frombits(uint32(s1i32))
lbl24:
	s0f32 = s0f32 * s1f32
	l9 = s0f32
lbl2:
	s0f32 = l9
	return s0f32
lbl1:
	s0f32 = l9
	s1f32 = 1e+30
	s0f32 = s0f32 * s1f32
	s1f32 = 1e+30
	s0f32 = s0f32 * s1f32
	return s0f32
lbl0:
	s0f32 = l9
	s1f32 = 1e-30
	s0f32 = s0f32 * s1f32
	s1f32 = 1e-30
	s0f32 = s0f32 * s1f32
	return s0f32
}
