package internal

import (
	"math"
	"math/bits"
	"unsafe"
)

func f640(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
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
	var l12 int32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = 1
	s2i32 = l2
	s3i32 = 6
	s2i32 = s2i32 + s3i32
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s1f32 = float32(s1i32)
	l14 = s1f32
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	l6 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl2
	}
	s0i32 = -2147483648
lbl2:
	l5 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl4
	}
	s0i32 = -2147483648
lbl4:
	l10 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl6
	}
	s0i32 = -2147483648
lbl6:
	l4 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl8
	}
	s0i32 = -2147483648
lbl8:
	l7 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl10
	}
	s0i32 = -2147483648
lbl10:
	l3 = s0i32
	s0i32 = l5
	s1i32 = l6
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l13 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l13
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl12
	}
	s0i32 = -2147483648
lbl12:
	l8 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l14
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl14
	}
	s0i32 = -2147483648
lbl14:
	l1 = s0i32
	s0i32 = l5
	s1i32 = l6
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s0i32 = l6
	s1i32 = l5
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l6 = s0i32
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l9
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 ^ s1i32
	s1i32 = 64
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
	} else {
		s0i32 = l0
		s1i32 = l7
		s2i32 = l3
		s3i32 = l2
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l5 = s1i32
		s2i32 = 10
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = -1
		s2i32 = 1
		s3i32 = l2
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		ctx.Mem[int(s0i32+27)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l9
		s2i32 = 10
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l3
		s2i32 = l7
		s3i32 = l2
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l11 = s1i32
		s2i32 = 10
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l6
		s2i32 = 10
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 6
		s2i32 = 32
		s3i32 = l8
		s4i32 = l1
		s5i32 = l2
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		l7 = s3i32
		s4i32 = 6
		s3i32 = s3i32 * s4i32
		s4i32 = l5
		s3i32 = s3i32 + s4i32
		s4i32 = l1
		s5i32 = l8
		s6i32 = l2
		if s6i32 != 0 {
			// s4i32 = s4i32
		} else {
			s4i32 = s5i32
		}
		l8 = s4i32
		s5i32 = -15
		s4i32 = s4i32 * s5i32
		s3i32 = s3i32 + s4i32
		s4i32 = l11
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		s4i32 = 19
		s3i32 = s3i32 * s4i32
		l1 = s3i32
		s4i32 = 9
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		s4i32 = l1
		s5i32 = 31
		s4i32 = s4i32 >> (uint32(s5i32) & 31)
		l1 = s4i32
		s3i32 = s3i32 + s4i32
		s4i32 = l1
		s3i32 = s3i32 ^ s4i32
		l1 = s3i32
		s4i32 = l5
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = l11
		s4i32 = s4i32 + s5i32
		s5i32 = l7
		s6i32 = -15
		s5i32 = s5i32 * s6i32
		s4i32 = s4i32 + s5i32
		s5i32 = l8
		s6i32 = 6
		s5i32 = s5i32 * s6i32
		s4i32 = s4i32 + s5i32
		s5i32 = 19
		s4i32 = s4i32 * s5i32
		l3 = s4i32
		s5i32 = 9
		s4i32 = s4i32 >> (uint32(s5i32) & 31)
		s5i32 = l3
		s6i32 = 31
		s5i32 = s5i32 >> (uint32(s6i32) & 31)
		l3 = s5i32
		s4i32 = s4i32 + s5i32
		s5i32 = l3
		s4i32 = s4i32 ^ s5i32
		l3 = s4i32
		s5i32 = l3
		s6i32 = l1
		if s5i32 < s6i32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		l1 = s3i32
		s4i32 = l10
		s5i32 = l4
		s6i32 = l2
		if s6i32 != 0 {
			// s4i32 = s4i32
		} else {
			s4i32 = s5i32
		}
		l3 = s4i32
		s5i32 = 6
		s4i32 = s4i32 * s5i32
		s5i32 = l6
		s4i32 = s4i32 + s5i32
		s5i32 = l4
		s6i32 = l10
		s7i32 = l2
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		l10 = s5i32
		s6i32 = -15
		s5i32 = s5i32 * s6i32
		s4i32 = s4i32 + s5i32
		s5i32 = l9
		s6i32 = 3
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s4i32 = s4i32 + s5i32
		s5i32 = 19
		s4i32 = s4i32 * s5i32
		l2 = s4i32
		s5i32 = 9
		s4i32 = s4i32 >> (uint32(s5i32) & 31)
		s5i32 = l2
		s6i32 = 31
		s5i32 = s5i32 >> (uint32(s6i32) & 31)
		l2 = s5i32
		s4i32 = s4i32 + s5i32
		s5i32 = l2
		s4i32 = s4i32 ^ s5i32
		l2 = s4i32
		s5i32 = l6
		s6i32 = 3
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s6i32 = l9
		s5i32 = s5i32 + s6i32
		s6i32 = l3
		s7i32 = -15
		s6i32 = s6i32 * s7i32
		s5i32 = s5i32 + s6i32
		s6i32 = l10
		s7i32 = 6
		s6i32 = s6i32 * s7i32
		s5i32 = s5i32 + s6i32
		s6i32 = 19
		s5i32 = s5i32 * s6i32
		l4 = s5i32
		s6i32 = 9
		s5i32 = s5i32 >> (uint32(s6i32) & 31)
		s6i32 = l4
		s7i32 = 31
		s6i32 = s6i32 >> (uint32(s7i32) & 31)
		l4 = s6i32
		s5i32 = s5i32 + s6i32
		s6i32 = l4
		s5i32 = s5i32 ^ s6i32
		l4 = s5i32
		s6i32 = l4
		s7i32 = l2
		if s6i32 < s7i32 {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i32 = s4i32
		} else {
			s4i32 = s5i32
		}
		l2 = s4i32
		s5i32 = 1
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s3i32 = s3i32 + s4i32
		s4i32 = l1
		s5i32 = 1
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = l2
		s4i32 = s4i32 + s5i32
		s5i32 = l1
		s6i32 = l2
		if s5i32 > s6i32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		l1 = s3i32
		s4i32 = 16
		s3i32 = s3i32 + s4i32
		s4i32 = 5
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s3i32 = int32(bits.LeadingZeros32(uint32(s3i32)))
		s2i32 = s2i32 - s3i32
		s3i32 = 1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s4i32 = 65519
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l1 = s1i32
		ctx.Mem[int(s0i32+25)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = -1
		s2i32 = l1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		ctx.Mem[int(s0i32+24)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 0
		s2i32 = l1
		s3i32 = -4
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s4i32 = 4
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l2 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		ctx.Mem[int(s0i32+26)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l11
		s2i32 = l5
		s1i32 = s1i32 - s2i32
		s2i32 = l7
		s3i32 = l8
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = 10
		s3i32 = l1
		s2i32 = s2i32 - s3i32
		s3i32 = 6
		s4i32 = l2
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l2 = s2i32
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l11 = s1i32
		s2i32 = 3
		s1i32 = s1i32 * s2i32
		s2i32 = l1
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		l4 = s2i32
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l12 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l9
		s2i32 = l6
		s1i32 = s1i32 - s2i32
		s2i32 = l3
		s3i32 = l10
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l9 = s1i32
		s2i32 = 3
		s1i32 = s1i32 * s2i32
		s2i32 = l4
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l12
		s2i32 = l5
		s3i32 = l8
		s2i32 = s2i32 + s3i32
		s3i32 = l7
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		s2i32 = s2i32 * s3i32
		s3i32 = l2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l8 = s2i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l7
		s2i32 = l5
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 * s2i32
		s2i32 = l2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l8
		s3i32 = l1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l11
		s3i32 = l1
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		l5 = s3i32
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = l6
		s3i32 = l10
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		s2i32 = s2i32 * s3i32
		s3i32 = l2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l7 = s2i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l3
		s2i32 = l6
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 * s2i32
		s2i32 = l2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l7
		s3i32 = l1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s3i32 = l5
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = 1
	}
	return s0i32
}
