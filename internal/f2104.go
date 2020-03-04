package internal

import (
	"math"
	"unsafe"
)

func f2104(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 int64
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
	var s0i64 int64
	_ = s0i64
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
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	var s8f32 float32
	_ = s8f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = 1
	l5 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	if s0f32 > s1f32 {
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
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i64 = l8
		s1i64 = 32
		s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
		s0i32 = int32(s0i64)
		l1 = s0i32
		s0f32 = math.Float32frombits(uint32(s0i32))
		l11 = s0f32
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l11 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l1 = s0i32
	s0i32 = 0
	l5 = s0i32
lbl0:
	s0f32 = l11
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l13 = s1f32
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l14 = s0f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l12 = s1f32
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l14
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
		s0i32 = l3
		s1f32 = l13
		s2i32 = l3
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s0i32 = f445(ctx, s0i32, s1f32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l13
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l11 = s1f32
			s0f32 = s0f32 - s1f32
			l15 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l12 = s0f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l16 = s0f32
			s0f32 = l11
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			l13 = s1f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l17 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
			s1f32 = l12
			s2f32 = l13
			s1f32 = s1f32 - s2f32
			s2f32 = 3
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			l18 = s0f32
			s0f32 = 3.4028235e+38
			l12 = s0f32
			s0f32 = 0
			l13 = s0f32
			s0f32 = 0.25
			l14 = s0f32
			s0f32 = 0.5
			l11 = s0f32
		lbl5:
			s0f32 = l11
			s1f32 = l13
			s2f32 = l12
			s3f32 = l11
			s4f32 = l16
			s5f32 = l11
			s6f32 = l17
			s7f32 = l18
			s8f32 = l11
			s7f32 = s7f32 * s8f32
			s6f32 = s6f32 + s7f32
			s5f32 = s5f32 * s6f32
			s4f32 = s4f32 + s5f32
			s3f32 = s3f32 * s4f32
			l19 = s3f32
			s4f32 = l15
			s3f32 = s3f32 - s4f32
			s3f32 = float32(math.Abs(float64(s3f32)))
			l20 = s3f32
			if s2f32 > s3f32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l1 = s2i32
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l13 = s0f32
			s0f32 = l11
			s1f32 = l11
			s2f32 = l14
			s3f32 = l14
			s3f32 = -s3f32
			s4f32 = l19
			s5f32 = l15
			if s4f32 < s5f32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2f32 = s2f32
			} else {
				s2f32 = s3f32
			}
			s1f32 = s1f32 + s2f32
			l11 = s1f32
			if s0f32 == s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl6
			}
			s0f32 = l14
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			l14 = s0f32
			s0f32 = l20
			s1f32 = l12
			s2i32 = l1
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l12 = s0f32
			s1f32 = 0.25
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
		lbl6:
			s0i32 = l3
			s1i32 = l3
			s2i32 = -64
			s1i32 = s1i32 - s2i32
			s2f32 = l13
			f234(ctx, s0i32, s1i32, s2f32)
		}
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l11 = s1f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)]))
		s1f32 = l11
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
		s1f32 = l11
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l3
		s1i32 = l3
		s2i32 = 88
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1f32 = l11
		s2i32 = l3
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s0i32 = f445(ctx, s0i32, s1f32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l11
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
			l11 = s1f32
			s0f32 = s0f32 - s1f32
			l15 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
			l12 = s0f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l16 = s0f32
			s0f32 = l11
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
			l13 = s1f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l17 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
			s1f32 = l12
			s2f32 = l13
			s1f32 = s1f32 - s2f32
			s2f32 = 3
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			l18 = s0f32
			s0f32 = 3.4028235e+38
			l12 = s0f32
			s0f32 = 0
			l13 = s0f32
			s0f32 = 0.25
			l14 = s0f32
			s0f32 = 0.5
			l11 = s0f32
		lbl9:
			s0f32 = l11
			s1f32 = l13
			s2f32 = l12
			s3f32 = l11
			s4f32 = l16
			s5f32 = l11
			s6f32 = l17
			s7f32 = l18
			s8f32 = l11
			s7f32 = s7f32 * s8f32
			s6f32 = s6f32 + s7f32
			s5f32 = s5f32 * s6f32
			s4f32 = s4f32 + s5f32
			s3f32 = s3f32 * s4f32
			l19 = s3f32
			s4f32 = l15
			s3f32 = s3f32 - s4f32
			s3f32 = float32(math.Abs(float64(s3f32)))
			l20 = s3f32
			if s2f32 > s3f32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l1 = s2i32
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l13 = s0f32
			s0f32 = l11
			s1f32 = l11
			s2f32 = l14
			s3f32 = l14
			s3f32 = -s3f32
			s4f32 = l19
			s5f32 = l15
			if s4f32 < s5f32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2f32 = s2f32
			} else {
				s2f32 = s3f32
			}
			s1f32 = s1f32 + s2f32
			l11 = s1f32
			if s0f32 == s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl10
			}
			s0f32 = l14
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			l14 = s0f32
			s0f32 = l20
			s1f32 = l12
			s2i32 = l1
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l12 = s0f32
			s1f32 = 0.25
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl9
			}
		lbl10:
			s0i32 = l3
			s1i32 = 32
			s0i32 = s0i32 + s1i32
			s1i32 = l3
			s2i32 = -64
			s1i32 = s1i32 - s2i32
			s2f32 = l13
			f234(ctx, s0i32, s1i32, s2f32)
		}
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
	lbl7:
		s0i32 = l3
		s1f32 = l11
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)]))
		s1f32 = l11
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
			s0i32 = l3
			s1f32 = l11
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = s1f32
		}
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l12 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l11 = s0f32
		s0i32 = int32(math.Float32bits(s0f32))
		l1 = s0i32
	}
	s0f32 = l11
	s1f32 = l12
	if s0f32 > s1f32 {
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
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		goto lbl12
	}
	s0i32 = l3
	s1f32 = l12
	s2i32 = l3
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	s0i32 = f445(ctx, s0i32, s1f32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l12
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l11 = s1f32
		s0f32 = s0f32 - s1f32
		l15 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l12 = s0f32
		s1f32 = l11
		s0f32 = s0f32 - s1f32
		s1f32 = 3
		s0f32 = s0f32 * s1f32
		l16 = s0f32
		s0f32 = l11
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l13 = s1f32
		s2f32 = l12
		s1f32 = s1f32 - s2f32
		s2f32 = l12
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 3
		s0f32 = s0f32 * s1f32
		l17 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1f32 = l12
		s2f32 = l13
		s1f32 = s1f32 - s2f32
		s2f32 = 3
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l11
		s0f32 = s0f32 - s1f32
		l18 = s0f32
		s0f32 = 3.4028235e+38
		l12 = s0f32
		s0f32 = 0
		l13 = s0f32
		s0f32 = 0.25
		l14 = s0f32
		s0f32 = 0.5
		l11 = s0f32
	lbl15:
		s0f32 = l11
		s1f32 = l13
		s2f32 = l12
		s3f32 = l11
		s4f32 = l16
		s5f32 = l11
		s6f32 = l17
		s7f32 = l18
		s8f32 = l11
		s7f32 = s7f32 * s8f32
		s6f32 = s6f32 + s7f32
		s5f32 = s5f32 * s6f32
		s4f32 = s4f32 + s5f32
		s3f32 = s3f32 * s4f32
		l19 = s3f32
		s4f32 = l15
		s3f32 = s3f32 - s4f32
		s3f32 = float32(math.Abs(float64(s3f32)))
		l20 = s3f32
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l1 = s2i32
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l13 = s0f32
		s0f32 = l11
		s1f32 = l11
		s2f32 = l14
		s3f32 = l14
		s3f32 = -s3f32
		s4f32 = l19
		s5f32 = l15
		if s4f32 < s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2f32 = s2f32
		} else {
			s2f32 = s3f32
		}
		s1f32 = s1f32 + s2f32
		l11 = s1f32
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0f32 = l14
		s1f32 = 0.5
		s0f32 = s0f32 * s1f32
		l14 = s0f32
		s0f32 = l20
		s1f32 = l12
		s2i32 = l1
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l12 = s0f32
		s1f32 = 0.25
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
	lbl16:
		s0i32 = l3
		s1i32 = l3
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s2f32 = l13
		f234(ctx, s0i32, s1i32, s2f32)
	}
	s0i32 = l3
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	s1i32 = l1
	s1f32 = math.Float32frombits(uint32(s1i32))
	if s0f32 > s1f32 {
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
		s0i32 = l3
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	l8 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i64 = l8
	s1i64 = 32
	s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
	s0i32 = int32(s0i64)
	l1 = s0i32
	s0i64 = l8
	s0i32 = int32(s0i64)
	s0f32 = math.Float32frombits(uint32(s0i32))
lbl12:
	l11 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s1f32 = l11
	if s0f32 > s1f32 {
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
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l8 = s0i64
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l9 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = l8
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l10 = s0i64
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = l10
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		l5 = s0i32
		s0i64 = l8
		s0i32 = int32(s0i64)
		s0f32 = math.Float32frombits(uint32(s0i32))
		l12 = s0f32
		s0i64 = l9
		s0i32 = int32(s0i64)
		s0f32 = math.Float32frombits(uint32(s0i32))
		l11 = s0f32
		s0i64 = l9
		s1i64 = 32
		s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
		s0i32 = int32(s0i64)
		l1 = s0i32
	}
	s0f32 = l11
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l14 = s1f32
	if s0f32 <= s1f32 {
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
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l4 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l2 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l2
		s1i32 = l1
		s2i32 = l4
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l2
		s1i32 = l4
		s2i32 = l1
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl2
	}
	s0f32 = l12
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l13 = s1f32
	if s0f32 >= s1f32 {
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
		s0i32 = int32(ctx.Mem[int(s0i32+8)])
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l4 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l2 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l2
		s1i32 = l1
		s2i32 = l4
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l2
		s1i32 = l4
		s2i32 = l1
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl2
	}
	s0f32 = l12
	s1f32 = l14
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l11
	} else {
		s0i32 = l3
		s1f32 = l14
		s2i32 = l3
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s0i32 = f689(ctx, s0i32, s1f32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l14
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l11 = s1f32
			s0f32 = s0f32 - s1f32
			l15 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l12 = s0f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l16 = s0f32
			s0f32 = l11
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			l13 = s1f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l17 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
			s1f32 = l12
			s2f32 = l13
			s1f32 = s1f32 - s2f32
			s2f32 = 3
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			l18 = s0f32
			s0f32 = 3.4028235e+38
			l12 = s0f32
			s0f32 = 0
			l13 = s0f32
			s0f32 = 0.25
			l14 = s0f32
			s0f32 = 0.5
			l11 = s0f32
		lbl23:
			s0f32 = l11
			s1f32 = l13
			s2f32 = l12
			s3f32 = l11
			s4f32 = l16
			s5f32 = l11
			s6f32 = l17
			s7f32 = l18
			s8f32 = l11
			s7f32 = s7f32 * s8f32
			s6f32 = s6f32 + s7f32
			s5f32 = s5f32 * s6f32
			s4f32 = s4f32 + s5f32
			s3f32 = s3f32 * s4f32
			l19 = s3f32
			s4f32 = l15
			s3f32 = s3f32 - s4f32
			s3f32 = float32(math.Abs(float64(s3f32)))
			l20 = s3f32
			if s2f32 > s3f32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l1 = s2i32
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l13 = s0f32
			s0f32 = l11
			s1f32 = l11
			s2f32 = l14
			s3f32 = l14
			s3f32 = -s3f32
			s4f32 = l19
			s5f32 = l15
			if s4f32 < s5f32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2f32 = s2f32
			} else {
				s2f32 = s3f32
			}
			s1f32 = s1f32 + s2f32
			l11 = s1f32
			if s0f32 == s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl24
			}
			s0f32 = l14
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			l14 = s0f32
			s0f32 = l20
			s1f32 = l12
			s2i32 = l1
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l12 = s0f32
			s1f32 = 0.25
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl23
			}
		lbl24:
			s0i32 = l3
			s1i32 = l3
			s2i32 = -64
			s1i32 = s1i32 - s2i32
			s2f32 = l13
			f234(ctx, s0i32, s1i32, s2f32)
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l6 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
		l7 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l4 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l6
		s2i32 = l7
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l7
		s2i32 = l6
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		s1i32 = l1
		s1f32 = math.Float32frombits(uint32(s1i32))
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
			s0i32 = l3
			s1i32 = l1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
		}
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l2
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l13 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	}
	s1f32 = l13
	if s0f32 > s1f32 {
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
		s0i32 = l3
		s1f32 = l13
		s2i32 = l3
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s0i32 = f689(ctx, s0i32, s1f32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l13
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l11 = s1f32
			s0f32 = s0f32 - s1f32
			l15 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l12 = s0f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l16 = s0f32
			s0f32 = l11
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			l13 = s1f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s2f32 = l12
			s1f32 = s1f32 - s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 3
			s0f32 = s0f32 * s1f32
			l17 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
			s1f32 = l12
			s2f32 = l13
			s1f32 = s1f32 - s2f32
			s2f32 = 3
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = l11
			s0f32 = s0f32 - s1f32
			l18 = s0f32
			s0f32 = 3.4028235e+38
			l12 = s0f32
			s0f32 = 0
			l13 = s0f32
			s0f32 = 0.25
			l14 = s0f32
			s0f32 = 0.5
			l11 = s0f32
		lbl28:
			s0f32 = l11
			s1f32 = l13
			s2f32 = l12
			s3f32 = l11
			s4f32 = l16
			s5f32 = l11
			s6f32 = l17
			s7f32 = l18
			s8f32 = l11
			s7f32 = s7f32 * s8f32
			s6f32 = s6f32 + s7f32
			s5f32 = s5f32 * s6f32
			s4f32 = s4f32 + s5f32
			s3f32 = s3f32 * s4f32
			l19 = s3f32
			s4f32 = l15
			s3f32 = s3f32 - s4f32
			s3f32 = float32(math.Abs(float64(s3f32)))
			l20 = s3f32
			if s2f32 > s3f32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l1 = s2i32
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l13 = s0f32
			s0f32 = l11
			s1f32 = l11
			s2f32 = l14
			s3f32 = l14
			s3f32 = -s3f32
			s4f32 = l19
			s5f32 = l15
			if s4f32 < s5f32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2f32 = s2f32
			} else {
				s2f32 = s3f32
			}
			s1f32 = s1f32 + s2f32
			l11 = s1f32
			if s0f32 == s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl29
			}
			s0f32 = l14
			s1f32 = 0.5
			s0f32 = s0f32 * s1f32
			l14 = s0f32
			s0f32 = l20
			s1f32 = l12
			s2i32 = l1
			if s2i32 != 0 {
				// s0f32 = s0f32
			} else {
				s0f32 = s1f32
			}
			l12 = s0f32
			s1f32 = 0.25
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl28
			}
		lbl29:
			s0i32 = l3
			s1i32 = l3
			s2i32 = -64
			s1i32 = s1i32 - s2i32
			s2f32 = l13
			f234(ctx, s0i32, s1i32, s2f32)
		}
		s0i32 = l3
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		s1i32 = l1
		s1f32 = math.Float32frombits(uint32(s1i32))
		if s0f32 > s1f32 {
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
			s0i32 = l3
			s1i32 = l1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l1 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			goto lbl31
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	lbl31:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)]))
		l4 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l6 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l7 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l7
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l6
		s2i32 = l4
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l4
		s2i32 = l6
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = 48
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl2
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		goto lbl33
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
lbl33:
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl2:
	s0i32 = l3
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
