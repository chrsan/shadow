package internal

import (
	"unsafe"
)

func f701(ctx *Context, l0 int32, l1 int32) {
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
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 204
	s0i32 = i32DivU(s0i32, s1i32)
	l5 = s0i32
	s1i32 = l1
	s2i32 = l5
	s3i32 = 204
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = 0
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l1
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l6 = s2i32
	s3i32 = 204
	s2i32 = i32DivU(s2i32, s3i32)
	l5 = s2i32
	s3i32 = l1
	s4i32 = l5
	if uint32(s3i32) < uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l7 = s1i32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l7
		s2i32 = -204
		s1i32 = s1i32 * s2i32
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	lbl3:
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0i32
		s0i32 = l0
		s1i32 = l2
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l1 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l6 = s1i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l2
			s0i32 = s0i32 - s1i32
			l1 = s0i32
			s1i32 = 2
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l5 = s0i32
			s0i32 = l2
			s1i32 = l2
			s2i32 = l6
			s1i32 = s1i32 - s2i32
			s2i32 = 2
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			s2i32 = -2
			s1i32 = i32DivS(s1i32, s2i32)
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			l6 = s1i32
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l0
			s1i32 = l1
			if s1i32 != 0 {
				s1i32 = l3
				s2i32 = l2
				s3i32 = l1
				s1i32 = f106(ctx, s1i32, s2i32, s3i32)
				s1i32 = l0
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			} else {
				s1i32 = l2
			}
			s2i32 = l6
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l3
			s2i32 = l5
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l3 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			goto lbl4
		}
		s0i32 = l1
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = 1
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = 1
		s2i32 = l1
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		s1i32 = 1073741824
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l8 = s0i32
		s0i32 = f17(ctx, s0i32)
		l5 = s0i32
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l5
		s1i32 = l1
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl8:
			s0i32 = l1
			s1i32 = l2
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l2
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = l3
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l6 = s0i32
		}
		s0i32 = l0
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l3 = s0i32
			goto lbl4
		}
		s0i32 = l6
		f12(ctx, s0i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l3 = s0i32
	lbl4:
		s0i32 = l3
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		goto lbl3
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l1 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = s1i32 - s2i32
	l6 = s1i32
	s2i32 = 2
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l3
	s3i32 = l2
	s2i32 = s2i32 - s3i32
	s3i32 = 2
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	l2 = s2i32
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl15:
			s0i32 = l4
			s1i32 = 4080
			s1i32 = f17(ctx, s1i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l4
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			f455(ctx, s0i32, s1i32)
			s0i32 = l5
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl15
			}
		}
		s0i32 = l5
		l2 = s0i32
	lbl16:
		s0i32 = l4
		s1i32 = 4080
		s1i32 = f17(ctx, s1i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		f280(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = 203
		s3i32 = 204
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
		s5i32 = l0
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
		s4i32 = s4i32 - s5i32
		s5i32 = 4
		if s4i32 == s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l5
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		goto lbl10
	}
	s0i32 = l4
	s1i32 = l0
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = 0
	l1 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l6
	s2i32 = 1
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l6 = s1i32
	s2i32 = l6
	s3i32 = l3
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
	l3 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1073741824
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s0i32 = f17(ctx, s0i32)
		l1 = s0i32
	}
	s0i32 = l4
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	s2i32 = l2
	s3i32 = l7
	s2i32 = s2i32 - s3i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	s2i32 = l3
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
lbl18:
	s0i32 = 4080
	s0i32 = f17(ctx, s0i32)
	l8 = s0i32
	s0i32 = l2
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl19
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l1 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l6 = s1i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		s1i32 = 2
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l9 = s0i32
		s0i32 = l1
		s1i32 = l1
		s2i32 = l6
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = -2
		s1i32 = i32DivS(s1i32, s2i32)
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l1
			s2i32 = l2
			s0i32 = f106(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l4
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l3
		s2i32 = l9
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl19
	}
	s0i32 = l3
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s1i32 = 1
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = 1
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
	s1i32 = 1073741824
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l10 = s0i32
	s0i32 = f17(ctx, s0i32)
	l9 = s0i32
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l9
	s1i32 = l2
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	l2 = s0i32
	s0i32 = l1
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl23:
		s0i32 = l2
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
	}
	s0i32 = l4
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl19
	}
	s0i32 = l6
	f12(ctx, s0i32)
lbl19:
	s0i32 = l2
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l3 = s0i32
	goto lbl18
	panic("unreachable executed")
	panic("unreachable executed")
lbl12:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	goto lbl10
lbl11:
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0i32
		goto lbl24
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s0i32 = l7
	l6 = s0i32
lbl26:
	s0i32 = l2
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l3 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l1 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l9 = s1i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l2 = s0i32
		s1i32 = 2
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l8 = s0i32
		s0i32 = l1
		s1i32 = l1
		s2i32 = l9
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = -2
		s1i32 = i32DivS(s1i32, s2i32)
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l1
			s2i32 = l2
			s0i32 = f106(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l4
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l3
		s2i32 = l8
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl27
	}
	s0i32 = l3
	s1i32 = l9
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = 1
	s2i32 = l3
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s1i32 = 1073741824
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l10 = s0i32
	s0i32 = f17(ctx, s0i32)
	l8 = s0i32
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l8
	s1i32 = l3
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	l3 = s0i32
	s0i32 = l1
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl31:
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
	}
	s0i32 = l4
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l9
	if s0i32 != 0 {
		s0i32 = l9
		f12(ctx, s0i32)
	}
	s0i32 = l3
	l2 = s0i32
lbl27:
	s0i32 = l2
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	if s0i32 != 0 {
		goto lbl26
	}
lbl24:
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l1 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl34:
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		f280(ctx, s0i32, s1i32)
		s0i32 = l1
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l5 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl34
		}
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l1 = s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l2 = s0i32
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l7
	s3i32 = -204
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l5
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s2i32 = l1
		s3i32 = l5
		s2i32 = s2i32 - s3i32
		s3i32 = -4
		s2i32 = s2i32 + s3i32
		s3i32 = 2
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = -1
		s2i32 = s2i32 ^ s3i32
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	f12(ctx, s0i32)
	goto lbl1
lbl10:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l7
	s3i32 = -204
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
lbl36:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l1 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl37
	}
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = 2
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l5 = s0i32
		s0i32 = l2
		s1i32 = l2
		s2i32 = l6
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = -2
		s1i32 = i32DivS(s1i32, s2i32)
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l6 = s1i32
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l0
		s1i32 = l1
		if s1i32 != 0 {
			s1i32 = l3
			s2i32 = l2
			s3i32 = l1
			s1i32 = f106(ctx, s1i32, s2i32, s3i32)
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		} else {
			s1i32 = l2
		}
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l3
		s2i32 = l5
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl37
	}
	s0i32 = l1
	s1i32 = l6
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = 1
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s1i32 = 1073741824
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l8 = s0i32
	s0i32 = f17(ctx, s0i32)
	l5 = s0i32
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l1
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	l1 = s0i32
	s0i32 = l2
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl41:
		s0i32 = l1
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl41
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
	}
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l10
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		l3 = s0i32
		goto lbl37
	}
	s0i32 = l6
	f12(ctx, s0i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
lbl37:
	s0i32 = l3
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	if s0i32 != 0 {
		goto lbl36
	}
lbl1:
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	return
lbl0:
	f63(ctx)
	panic("unreachable executed")
}
