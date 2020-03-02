package internal

import (
	"math"
	"unsafe"
)

func f166(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])))
	l6 = s0i32
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+18)])))
	l7 = s0i32
	s1i32 = l4
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l7
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l10 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l14 = s1f32
	s0f32 = s0f32 + s1f32
	l19 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l11 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l15 = s1f32
	s0f32 = s0f32 + s1f32
	l20 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l16 = s1f32
	s0f32 = s0f32 + s1f32
	l12 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l17 = s1f32
	s0f32 = s0f32 + s1f32
	l13 = s0f32
	s0f32 = l9
	s1f32 = l8
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
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = 1
		s2i32 = -1
		s3f32 = l17
		s4f32 = l10
		s5f32 = l11
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 * s4f32
		s4f32 = l15
		s5f32 = l8
		s6f32 = l9
		s5f32 = s5f32 - s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 - s4f32
		l18 = s3f32
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
		s2f32 = l18
		s2f32 = float32(math.Abs(float64(s2f32)))
		s3f32 = 5.9604645e-08
		if s2f32 <= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l0 = s0i32
		s0f32 = l12
		s1f32 = l13
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
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 0
			s2i32 = 1
			s3i32 = -1
			s4f32 = l19
			s5f32 = l11
			s4f32 = s4f32 - s5f32
			s5f32 = l17
			s4f32 = s4f32 * s5f32
			s5f32 = l12
			s6f32 = l9
			s5f32 = s5f32 - s6f32
			s6f32 = l15
			s5f32 = s5f32 * s6f32
			s4f32 = s4f32 - s5f32
			l8 = s4f32
			s5f32 = 0
			if s4f32 > s5f32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2i32 = s2i32
			} else {
				s2i32 = s3i32
			}
			s3f32 = l8
			s3f32 = float32(math.Abs(float64(s3f32)))
			s4f32 = 5.9604645e-08
			if s3f32 <= s4f32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s0i32 = s0i32 * s1i32
			s1i32 = 31
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			return s0i32
		}
		s0i32 = 0
		s1i32 = 1
		s2i32 = -1
		s3f32 = l20
		s4f32 = l10
		s3f32 = s3f32 - s4f32
		s4f32 = l16
		s3f32 = s3f32 * s4f32
		s4f32 = l13
		s5f32 = l8
		s4f32 = s4f32 - s5f32
		s5f32 = l14
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 - s4f32
		l8 = s3f32
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
		s2f32 = l8
		s2f32 = float32(math.Abs(float64(s2f32)))
		s3f32 = 5.9604645e-08
		if s2f32 <= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l0
		s0i32 = s0i32 * s1i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		return s0i32
	}
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3f32 = l16
	s4f32 = l11
	s5f32 = l10
	s4f32 = s4f32 - s5f32
	s3f32 = s3f32 * s4f32
	s4f32 = l14
	s5f32 = l9
	s6f32 = l8
	s5f32 = s5f32 - s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l18 = s3f32
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
	s2f32 = l18
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 5.9604645e-08
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l0 = s0i32
	s0f32 = l13
	s1f32 = l12
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
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		s2i32 = 1
		s3i32 = -1
		s4f32 = l20
		s5f32 = l10
		s4f32 = s4f32 - s5f32
		s5f32 = l16
		s4f32 = s4f32 * s5f32
		s5f32 = l13
		s6f32 = l8
		s5f32 = s5f32 - s6f32
		s6f32 = l14
		s5f32 = s5f32 * s6f32
		s4f32 = s4f32 - s5f32
		l8 = s4f32
		s5f32 = 0
		if s4f32 > s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s3f32 = l8
		s3f32 = float32(math.Abs(float64(s3f32)))
		s4f32 = 5.9604645e-08
		if s3f32 <= s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = s0i32 * s1i32
		s1i32 = 31
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		return s0i32
	}
	s0i32 = 0
	s1i32 = 1
	s2i32 = -1
	s3f32 = l19
	s4f32 = l11
	s3f32 = s3f32 - s4f32
	s4f32 = l17
	s3f32 = s3f32 * s4f32
	s4f32 = l12
	s5f32 = l9
	s4f32 = s4f32 - s5f32
	s5f32 = l15
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 - s4f32
	l8 = s3f32
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
	s2f32 = l8
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 5.9604645e-08
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l0
	s0i32 = s0i32 * s1i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
lbl0:
	s0i32 = l5
	return s0i32
}
