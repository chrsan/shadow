package internal

import (
	"unsafe"
)

func f315(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var s2i64 int64
	_ = s2i64
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	ctx.g0 = s0i32
	s0i32 = 0
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	s1i32 = l3
	s0i32 = s0i32 & s1i32
	s1i32 = 0
	s2i32 = l2
	s1i32 = s1i32 - s2i32
	s2i32 = l2
	s1i32 = s1i32 & s2i32
	s2i32 = 0
	s3i32 = l1
	s2i32 = s2i32 - s3i32
	s3i32 = l1
	s2i32 = s2i32 & s3i32
	s3i32 = 0
	s4i32 = l0
	s3i32 = s3i32 - s4i32
	s4i32 = l0
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 | s1i32
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s1i32 = l6
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l6 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s0i32 = s0i32 ^ s1i32
	l6 = s0i32
	s1i32 = 32704
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = l8
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l8 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l8
		s0i32 = s0i32 ^ s1i32
		l8 = s0i32
		s1i32 = 32705
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = 1
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s3i32 = l0
	s4i32 = 1
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	l0 = s2i32
	s3i32 = l3
	s4i32 = 1
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	s4i32 = l1
	s5i32 = 1
	s4i32 = s4i32 >> (uint32(s5i32) & 31)
	s3i32 = s3i32 + s4i32
	l1 = s3i32
	s4i32 = l4
	s5i32 = l5
	f315(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s5i32 = l5
	f315(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	goto lbl0
lbl1:
	s0i32 = l9
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 24772
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 24796
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 24820
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 24844
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l8
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s2i32 = l2
		s3i32 = l0
		if s2i32 < s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l11 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s1i32 = 10
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l8 = s0i32
		s0i32 = l2
		s1i32 = l0
		s2i32 = l11
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l10 = s0i32
		s1i32 = 6
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l6 = s0i32
		s0i32 = l0
		s1i32 = l2
		s2i32 = l11
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l13 = s0i32
		s1i32 = 63
		s0i32 = s0i32 + s1i32
		l14 = s0i32
		s1i32 = 6
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l2 = s0i32
		s0i32 = l7
		s1i32 = l1
		s2i32 = l3
		s3i32 = l11
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l0 = s1i32
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l10
			s1i32 = 63
			s0i32 = s0i32 & s1i32
			l3 = s0i32
			s0i32 = l13
			s1i32 = l10
			s0i32 = s0i32 - s1i32
			l12 = s0i32
			s0i32 = 0
			l1 = s0i32
			s0i32 = l9
			s1i32 = 56
			s0i32 = s0i32 + s1i32
			goto lbl5
		}
		s0i32 = l0
		s1i32 = l7
		s0i32 = s0i32 - s1i32
		s1i32 = 16
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l13
		s2i32 = l10
		s1i32 = s1i32 - s2i32
		l12 = s1i32
		s0i32 = i32DivS(s0i32, s1i32)
		l1 = s0i32
		s1i32 = 32
		s2i32 = l10
		s3i32 = 63
		s2i32 = s2i32 & s3i32
		l3 = s2i32
		s1i32 = s1i32 - s2i32
		s0i32 = s0i32 * s1i32
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1i32 = 6
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l9
		s1i32 = 48
		s0i32 = s0i32 + s1i32
	lbl5:
		l0 = s0i32
		s0i32 = 0
		l7 = s0i32
		s0i32 = 0
		s1i32 = l13
		s2i32 = 63
		s1i32 = s1i32 & s2i32
		s2i32 = l2
		s3i32 = l6
		s2i32 = s2i32 - s3i32
		s3i32 = 1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l11 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l10 = s0i32
		s0i32 = l12
		s1i32 = 64
		s2i32 = l3
		s1i32 = s1i32 - s2i32
		s2i32 = l11
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l11 = s0i32
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			l3 = s0i32
			goto lbl3
		}
		s0i32 = l6
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l7 = s1i32
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l3 = s1i32
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1i32 = l6
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			l3 = s0i32
			goto lbl8
		}
		s0i32 = l3
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		s1i32 = l1
		s0i32 = s0i32 * s1i32
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = 64
		l11 = s0i32
		s0i32 = l2
		s1i32 = l3
		s0i32 = s0i32 - s1i32
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l14
		s1i32 = 63
		s0i32 = s0i32 & s1i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s0i32 = 0
		l10 = s0i32
	lbl8:
		s0i32 = l7
		s1i32 = l2
		s2i32 = l2
		s3i32 = l7
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l6 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l2 = s0i32
		s1i32 = l3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l8
		s1i32 = l2
		s2i32 = l3
		s3i32 = -1
		s2i32 = s2i32 ^ s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s1i32 = s1i32 * s2i32
		s2i32 = l8
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s2i32 = l1
		s3i32 = -1
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l12 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = -32768
		s0i32 = s0i32 + s1i32
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l13 = s0i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l14 = s1i32
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l7
		s1i32 = l8
		s2i32 = l12
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = 98303
		s0i32 = s0i32 + s1i32
		s1i32 = 16
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l7 = s0i32
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l12 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = 0
		s1i32 = l10
		s2i32 = l6
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l10 = s0i32
		s0i32 = 0
		s1i32 = l4
		s2i32 = l12
		s3i32 = l13
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
		s1i32 = l4
		s2i32 = l14
		s3i32 = l7
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		goto lbl3
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = l3
	s3i32 = l1
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l10 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l11 = s0i32
	s1i32 = 10
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l8 = s0i32
	s0i32 = l3
	s1i32 = l1
	s2i32 = l10
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l6 = s0i32
	s0i32 = l1
	s1i32 = l3
	s2i32 = l10
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s0i32 = l11
	s1i32 = l0
	s2i32 = l2
	s3i32 = l10
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l0 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l6
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l6
		s1i32 = 63
		s0i32 = s0i32 & s1i32
		l11 = s0i32
		s0i32 = l3
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l12 = s0i32
		s0i32 = 0
		l1 = s0i32
		s0i32 = l9
		s1i32 = 40
		s0i32 = s0i32 + s1i32
		goto lbl10
	}
	s0i32 = l0
	s1i32 = l11
	s0i32 = s0i32 - s1i32
	s1i32 = 16
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l3
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	l12 = s1i32
	s0i32 = i32DivS(s0i32, s1i32)
	l1 = s0i32
	s1i32 = 32
	s2i32 = l6
	s3i32 = 63
	s2i32 = s2i32 & s3i32
	l11 = s2i32
	s1i32 = s1i32 - s2i32
	s0i32 = s0i32 * s1i32
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = 6
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l9
	s1i32 = 32
	s0i32 = s0i32 + s1i32
lbl10:
	l0 = s0i32
	s0i32 = 0
	s1i32 = l3
	s2i32 = 63
	s1i32 = s1i32 & s2i32
	s2i32 = l3
	s3i32 = 63
	s2i32 = s2i32 + s3i32
	l13 = s2i32
	s3i32 = 6
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	l2 = s2i32
	s3i32 = l6
	s4i32 = 6
	s3i32 = s3i32 >> (uint32(s4i32) & 31)
	l6 = s3i32
	s2i32 = s2i32 - s3i32
	s3i32 = 1
	if s2i32 == s3i32 {
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
	l10 = s0i32
	s0i32 = l12
	s1i32 = 64
	s2i32 = l11
	s1i32 = s1i32 - s2i32
	s2i32 = l3
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l11 = s0i32
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		l3 = s0i32
		goto lbl3
	}
	s0i32 = l6
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l7 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l3 = s1i32
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l6
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		l3 = s0i32
		goto lbl13
	}
	s0i32 = l3
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	s1i32 = l1
	s0i32 = s0i32 * s1i32
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = 64
	l11 = s0i32
	s0i32 = l2
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l13
	s1i32 = 63
	s0i32 = s0i32 & s1i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = 0
	l10 = s0i32
lbl13:
	s0i32 = l7
	s1i32 = l2
	s2i32 = l2
	s3i32 = l7
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s1i32 = l2
	s2i32 = l3
	s3i32 = -1
	s2i32 = s2i32 ^ s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s1i32 = s1i32 * s2i32
	s2i32 = l8
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s2i32 = l1
	s3i32 = -1
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l12 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = -32768
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l13 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l14 = s1i32
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i32 = l8
	s2i32 = l12
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = 98303
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l7 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	s1i32 = l10
	s2i32 = l6
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l10 = s0i32
	s0i32 = 0
	s1i32 = l4
	s2i32 = l12
	s3i32 = l13
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
	s1i32 = l4
	s2i32 = l14
	s3i32 = l7
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
lbl3:
	s0i32 = l9
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l9
	s1i32 = 20632
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	if s1i32 != 0 {
		s1i32 = l9
		s2i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
		s1i32 = l9
		s2i32 = l7
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
		s1i32 = l9
		s2i32 = l7
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
		s1i32 = l9
	} else {
		s1i32 = l5
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = l8
	s3i32 = l1
	s4i32 = l11
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+8)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l4 = s0i32
	s0i32 = l2
	s1i32 = l3
	s2i32 = -1
	s1i32 = s1i32 ^ s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l10
	s2i32 = 0
	if s1i32 > s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	l5 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = -1
		s4i32 = 0
		s5i32 = l5
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s4i32 = l1
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		l4 = s0i32
	}
	s0i32 = l10
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = l1
		s4i32 = l10
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+8)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	}
	s0i32 = l9
	s0i32 = f35(ctx, s0i32)
lbl0:
	s0i32 = l9
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
