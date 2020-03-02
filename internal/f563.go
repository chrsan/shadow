package internal

import (
	"unsafe"
)

func f563(ctx *Context, l0 int32, l1 int32) int32 {
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
	var l12 int32
	_ = l12
	var l13 int64
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
	var l22 float32
	_ = l22
	var l23 float32
	_ = l23
	var l24 float32
	_ = l24
	var l25 float32
	_ = l25
	var l26 float32
	_ = l26
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
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s1i32 = 3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l17 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l18 = s1f32
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s1f32 = l14
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l16 = s0f32
	s1f32 = l17
	s0f32 = s0f32 - s1f32
	l19 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l15 = s0f32
	s1f32 = l18
	s0f32 = s0f32 - s1f32
	l20 = s0f32
	s0f32 = l17
	s1i32 = l0
	s2i32 = l10
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0f32 = s0f32 - s1f32
	l21 = s0f32
	s0f32 = l18
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	l23 = s0f32
	s0f32 = l17
	s1f32 = l17
	s0f32 = s0f32 - s1f32
	l22 = s0f32
	s0f32 = l18
	s1f32 = l18
	s0f32 = s0f32 - s1f32
	l24 = s0f32
	s0i32 = 1
	l3 = s0i32
lbl5:
	s0f32 = l25
	s1f32 = l23
	s2f32 = l19
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l21
	s3f32 = l20
	l23 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l19 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = l26
	s1f32 = l16
	s2f32 = l17
	s1f32 = s1f32 - s2f32
	l21 = s1f32
	s2f32 = l24
	s1f32 = s1f32 * s2f32
	s2f32 = l15
	s3f32 = l18
	s2f32 = s2f32 - s3f32
	l24 = s2f32
	s3f32 = l22
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l20 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = 1
	l9 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s0i32 = i32RemS(s0i32, s1i32)
	l3 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l19
	s1f32 = l25
	s2f32 = l19
	s3f32 = 0
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l25 = s0f32
	s0f32 = l20
	s1f32 = l26
	s2f32 = l20
	s3f32 = 0
	if s2f32 != s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l26 = s0f32
	s0i32 = l0
	s1i32 = l3
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l22 = s0f32
	s1f32 = l15
	s0f32 = s0f32 - s1f32
	l20 = s0f32
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l15 = s0f32
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	l19 = s0f32
	s0f32 = l15
	l16 = s0f32
	s0f32 = l22
	l15 = s0f32
	s0f32 = l21
	l22 = s0f32
	s0f32 = l14
	l21 = s0f32
	s0i32 = l0
	s1i32 = l4
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = 0
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s1f32 = l14
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl4:
	s0i32 = 0
	l9 = s0i32
	s0i32 = l1
	s1i32 = 2048
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = 0
	l4 = s0i32
	s0i32 = l1
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = 2
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 0
		s2i32 = l3
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = f29(ctx, s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = 0
		l3 = s0i32
	lbl8:
		s0i32 = l0
		s1i32 = l3
		l5 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0f32
		s1f32 = 0
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l16 = s1f32
		s0f32 = s0f32 * s1f32
		l15 = s0f32
		s1f32 = l15
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = 0
		s1i32 = l5
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s2i32 = l1
		s3i32 = l3
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l8 = s0i32
		s0i32 = l4
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0i64
		s0i32 = l0
		s1i32 = l5
		s2i32 = l10
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s1i32 = i32RemS(s1i32, s2i32)
		l12 = s1i32
		s2i32 = 65535
		s1i32 = s1i32 & s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l15 = s0f32
		s1f32 = l14
		if s0f32 < s1f32 {
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
			l6 = s0i32
			s0f32 = l15
			s1f32 = l14
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl9
			}
			s0i32 = l4
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			s1f32 = l16
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl9
			}
		}
		s0i32 = 1
		l6 = s0i32
	lbl9:
		s0i32 = l0
		s1i32 = l8
		s2i32 = 65535
		s1i32 = s1i32 & s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l15 = s0f32
		s1f32 = l14
		if s0f32 < s1f32 {
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
			s0f32 = l15
			s1f32 = l14
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl11
			}
			s0i32 = l4
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			s1f32 = l16
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl11
			}
		}
		s0i32 = l6
		s1i32 = 2
		s0i32 = s0i32 | s1i32
		l6 = s0i32
	lbl11:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l7 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)]))
		s1i32 = l11
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
			l4 = s0i32
			goto lbl13
		}
		s0i32 = l7
		s1i32 = 5
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = 2
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
		s2i32 = l4
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = f29(ctx, s1i32, s2i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
	lbl13:
		s0i32 = l2
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l7
		s2i32 = 4
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l6
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+14)])) = uint16(s1i32)
		s0i32 = l4
		s1i32 = l8
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint16(s1i32)
		s0i32 = l4
		s1i32 = l12
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+10)])) = uint16(s1i32)
		s0i32 = l4
		s1i32 = l5
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		s0i32 = l4
		s1i64 = l13
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
		l6 = s0i32
	lbl16:
		s0i32 = l6
		s1i32 = l7
		s2i32 = 4
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0f32
		s1i32 = l6
		s2i32 = l7
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		s3i32 = 1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l7 = s2i32
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l16 = s1f32
		if s0f32 < s1f32 {
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
			s0f32 = l14
			s1f32 = l16
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl15
			}
			s0i32 = l5
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			if s0f32 > s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl15
			}
		}
		s0i32 = l2
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l8
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l7
		if s0i32 != 0 {
			goto lbl16
		}
	lbl15:
		s0i32 = l1
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l4 = s0i32
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 40
	s0i32 = s0i32 * s1i32
	s1i32 = 2
	s0i32 = f44(ctx, s0i32, s1i32)
	l3 = s0i32
	s0i32 = l2
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl19:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
		l1 = s0i32
		s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+14)])))
		l3 = s0i32
		s1i32 = 3
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl22
		}
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl22
		case 1:
			goto lbl22
		case 2:
			goto lbl23
		default:
			goto lbl24
		}
	lbl24:
		s0i32 = l2
		s1i32 = 56
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l0
		s3i32 = l1
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+10)])))
		l3 = s3i32
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)])))
		s4i32 = l3
		s0i32 = f562(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l2
		s1i32 = 56
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l0
		s3i32 = l1
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)])))
		l3 = s3i32
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)])))
		s4i32 = l3
		s0i32 = f562(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl21
		}
		goto lbl20
	lbl23:
		s0i32 = l2
		s1i32 = 56
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+10)])))
		l3 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s3i32 = l3
		s4i32 = l1
		s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)])))
		s0i32 = f561(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l2
		s1i32 = 56
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)])))
		l3 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s3i32 = l3
		s4i32 = l1
		s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)])))
		s0i32 = f561(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl21
		}
		goto lbl20
	lbl22:
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = 56
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = l1
			s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+10)])))
			l3 = s2i32
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = l1
			s3i32 = l0
			s4i32 = l1
			s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)])))
			l5 = s4i32
			s5i32 = 3
			s4i32 = s4i32 << (uint32(s5i32) & 31)
			s3i32 = s3i32 + s4i32
			s4i32 = l3
			s5i32 = l1
			s5i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s5i32+8)])))
			s6i32 = l5
			s0i32 = f560(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
			if s0i32 != 0 {
				goto lbl21
			}
			goto lbl20
		}
		s0i32 = l2
		s1i32 = 56
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l1
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)])))
		l3 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s3i32 = l0
		s4i32 = l1
		s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+10)])))
		l5 = s4i32
		s5i32 = 3
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s3i32 = s3i32 + s4i32
		s4i32 = l3
		s5i32 = l1
		s5i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s5i32+8)])))
		s6i32 = l5
		s0i32 = f560(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
	lbl21:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l3 = s0i32
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
			goto lbl20
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
		l3 = s0i32
		s1i32 = l3
		s2i32 = l1
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 112
		s0i32 = s0i32 + s1i32
		f1405(ctx, s0i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
	lbl20:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l4 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		l3 = s0i32
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l3
	f13(ctx, s0i32)
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l9 = s0i32
lbl6:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
	f13(ctx, s0i32)
lbl3:
	s0i32 = l2
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l9
	return s0i32
lbl2:
	s0i32 = l2
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 14655
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 14607
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = 14581
	s1i32 = l2
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl1:
	s0i32 = l2
	s1i32 = 345
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 14687
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 14607
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 14581
	s1i32 = l2
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl0:
	s0i32 = l2
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 14655
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 14607
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14581
	s1i32 = l2
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
