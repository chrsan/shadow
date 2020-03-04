package internal

import (
	"unsafe"
)

func f506(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
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
	s1i32 = 4848
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = -2
	s0i32 = s0i32 & s1i32
	s1i32 = l2
	s2i32 = l1
	s3i32 = 1
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
	l6 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 3740
	s0i32 = s0i32 + s1i32
	s0i32 = f152(ctx, s0i32)
	l2 = s0i32
	s0i32 = l5
	s1i32 = 3704
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l4
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+32)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+36)]))
	s0i32 = f1145(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+336)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 3672
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 336
		s1i32 = s1i32 + s2i32
		s2i32 = 4
		s1i32 = s1i32 | s2i32
		s2i32 = 3332
		s3i32 = 3332
		s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
		l7 = s0i32
		s0i32 = l5
		s1i32 = 336
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = l4
		s0i32 = f182(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l8 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+336)]))
		l4 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3716)]))
		l0 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+40)])
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = l0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3712)])) = uint32(s1i32)
			goto lbl3
		}
		s0i32 = l2
		s1i32 = l0
		s2i32 = l4
		f150(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = l5
		s2i32 = 4840
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3712)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 4844
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
	lbl3:
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3708)]))
		l2 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l9 = s0i32
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l2
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
			s1f32 = 0
			if s0f32 == s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l5
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3704)]))
				s1i32 = 2
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				s1i32 = 21180
				s0i32 = s0i32 + s1i32
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
				goto lbl5
			}
			s0i32 = 0
			s1i32 = 305
			s2i32 = l9
			s3i32 = 12
			s2i32 = s2i32 & s3i32
			s3i32 = 4
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
			goto lbl5
		}
		s0i32 = 306
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+3736)]))
		s2f32 = 0.5
		if s1f32 <= s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		if s1i32 != 0 {
			goto lbl5
		}
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+3704)]))
		l2 = s0i32
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = 307
		s1i32 = l4
		s2i32 = l5
		s3i32 = 80
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+32)]))
		if int(s3i32) < 0 || int(s3i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s3i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s3i32].numParams != 2 {
			panic("argument count mismatch")
		}
		s1i32 = (*(*func(*Context, int32, int32) int32)(table[s3i32].f()))(ctx, s1i32, s2i32)
		l0 = s1i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl5
		}
		s0i32 = 308
		s1i32 = 309
		s2i32 = 307
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
		l0 = s3i32
		s4i32 = 4
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l0
		s3i32 = 2
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
		goto lbl5
	lbl8:
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 21192
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	lbl5:
		l9 = s0i32
		s0i32 = l1
		s1i32 = 2
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l1 = s0i32
	lbl9:
		s0i32 = l8
		s1i32 = l5
		s2i32 = 80
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = l6
		s4i32 = 32
		s5i32 = l6
		s6i32 = 32
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
		l0 = s3i32
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l6
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 1
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			l10 = s0i32
			s0i32 = 0
			l2 = s0i32
			s0f32 = 0
			l12 = s0f32
		lbl12:
			s0f32 = l12
			s1i32 = l5
			s2i32 = 80
			s1i32 = s1i32 + s2i32
			s2i32 = l2
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 * s1f32
			l12 = s0f32
			s0i32 = l2
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = l10
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			s0f32 = l12
			s1f32 = 0
			if s0f32 != s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl10
			}
		}
		s0i32 = l5
		s1i32 = 3704
		s0i32 = s0i32 + s1i32
		s1i32 = l5
		s2i32 = 80
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s3i32 = l4
		s4i32 = l9
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l3
		s1i32 = l0
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l6
		s1i32 = l0
		s0i32 = s0i32 - s1i32
		l0 = s0i32
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l0
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l6 = s0i32
		if s0i32 != 0 {
			goto lbl9
		}
	lbl10:
		s0i32 = l7
		f43(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl14
	case 1:
		goto lbl13
	default:
		goto lbl15
	}
lbl15:
	s0i32 = 0
	l2 = s0i32
	s0i32 = l5
	s1i32 = 336
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s0i32 = f46(ctx, s0i32, s1i32)
	l1 = s0i32
	l4 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -193
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l13 = s0f32
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	l12 = s0f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 12
	s0i32 = s0i32 & s1i32
	s1i32 = 4
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s0i32 = f37(ctx, s0i32)
		l4 = s0i32
		s0i32 = l5
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint64(s1i64)
		s0i32 = l5
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l4
		s1f32 = 0
		s2f32 = 0
		s3f32 = l12
		f678(ctx, s0i32, s1f32, s2f32, s3f32)
		s0i32 = l6
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
	lbl17:
		s0i32 = l5
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l8
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		f117(ctx, s0i32, s1f32, s2f32)
		s0i32 = l4
		s1i32 = l4
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 251
		s1i32 = s1i32 & s2i32
		s2i32 = l2
		s3i32 = l7
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l8 = s2i32
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = l1
		s3i32 = l5
		s4i32 = 80
		s3i32 = s3i32 + s4i32
		s4i32 = l8
		f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l4
		f34(ctx, s0i32)
		s0i32 = l1
		s0i32 = f23(ctx, s0i32)
		goto lbl1
	}
lbl18:
	s0i32 = l5
	s1i32 = l3
	s2i32 = l2
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = s1f32
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l15 = s0f32
	s0i32 = l5
	s1f32 = l13
	s2f32 = l14
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = s1f32
	s0i32 = l5
	s1f32 = l15
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = s1f32
	s0i32 = l5
	s1f32 = l13
	s2f32 = l14
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
	s0i32 = l0
	s1i32 = l5
	s2i32 = 80
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s3i32 = 0
	s4i32 = 0
	f160(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l1
	s0i32 = f23(ctx, s0i32)
	goto lbl1
lbl14:
	s0i32 = l6
	s1i32 = 2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l5
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = l4
	s2f32 = 1
	s0i32 = f618(ctx, s0i32, s1i32, s2f32)
	l11 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+344)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+336)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 372
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l10 = s0i32
	s0i32 = l5
	s1i32 = 384
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l7 = s0i32
	s0i32 = l5
	s1i32 = 396
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l8 = s0i32
	s0i32 = l5
	s1i64 = 4575657222473777152
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+348)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l9 = s0i32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	f176(ctx, s0i32, s1f32, s2f32)
	s0i32 = l9
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	l2 = s1i32
	s2i32 = l2
	s3i32 = 20
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s3i32 = int32(ctx.Mem[int(s3i32+40)])
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l2 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l5
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l5
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
	s0i32 = l5
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s1f32 = float32(s1i32)
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s1i32 = l5
	s2i32 = 336
	s1i32 = s1i32 + s2i32
	s2i32 = l9
	s3i32 = l11
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+32)]))
	s5i32 = l5
	s6i32 = 32
	s5i32 = s5i32 + s6i32
	s6i32 = l2
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+36)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32, int32, int32) int32)(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s0i32 = f46(ctx, s0i32, s1i32)
		l1 = s0i32
		s0i32 = l5
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l5
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		f123(ctx, s0i32, s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l2 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l3 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
	lbl20:
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = -193
		s1i32 = s1i32 & s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l7
			s2i32 = l1
			s3i32 = 0
			s4i32 = 0
			f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		}
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l8
			s2i32 = l1
			s3i32 = 0
			s4i32 = 0
			f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		}
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+348)]))
		l12 = s0f32
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+352)]))
		l13 = s1f32
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+336)]))
			s1i32 = 1
			s0i32 = s0i32 & s1i32
			l2 = s0i32
			s1i32 = 2
			if uint32(s0i32) <= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l1
				s1i32 = l1
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
				s2i32 = -13
				s1i32 = s1i32 & s2i32
				s2i32 = l2
				s3i32 = 2
				s2i32 = s2i32 << (uint32(s3i32) & 31)
				s3i32 = 12
				s2i32 = s2i32 & s3i32
				s1i32 = s1i32 | s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
			}
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+340)]))
			l2 = s0i32
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+344)]))
			l3 = s0i32
			goto lbl24
		}
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+344)]))
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
		s0i32 = 0
		l2 = s0i32
	lbl27:
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+340)]))
		s1i32 = l2
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0f32
		s0i32 = l5
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l15 = s1f32
		s2f32 = l13
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l5
		s1f32 = l14
		s2f32 = l12
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l5
		s1f32 = l15
		s2f32 = l13
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
		s0i32 = l5
		s1f32 = l14
		s2f32 = l12
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
		s0i32 = l0
		s1i32 = l5
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s3i32 = 0
		s4i32 = 0
		f160(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+344)]))
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+352)]))
		l13 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+348)]))
		l12 = s0f32
		goto lbl27
		panic("unreachable executed")
		panic("unreachable executed")
	lbl24:
		s0i32 = l0
		s1i32 = 0
		s2i32 = l3
		s3i32 = l2
		s4i32 = l1
		f506(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	lbl23:
		s0i32 = l1
		s0i32 = f23(ctx, s0i32)
		s0i32 = l9
		f34(ctx, s0i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+340)]))
		l0 = s0i32
		if s0i32 != 0 {
			s0i32 = l0
			f12(ctx, s0i32)
		}
		s0i32 = l8
		f34(ctx, s0i32)
		s0i32 = l7
		f34(ctx, s0i32)
		s0i32 = l10
		f34(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l9
	f34(ctx, s0i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+340)]))
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l2
		f12(ctx, s0i32)
	}
	s0i32 = l8
	f34(ctx, s0i32)
	s0i32 = l7
	f34(ctx, s0i32)
	s0i32 = l10
	f34(ctx, s0i32)
lbl13:
	s0i32 = l5
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l2 = s0i32
	s0i32 = l5
	s1i32 = 336
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s0i32 = f46(ctx, s0i32, s1i32)
	l4 = s0i32
	l7 = s0i32
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -193
	s1i32 = s1i32 & s2i32
	s2i32 = 64
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	s0i32 = l6
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = 2
		s1i32 = 1
		s2i32 = l1
		s3i32 = 1
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
		s0i32 = 0
		l6 = s0i32
	lbl31:
		s0i32 = l2
		s1i32 = l3
		s2i32 = l6
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		f176(ctx, s0i32, s1f32, s2f32)
		s0i32 = l2
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = l4
		s3i32 = 0
		s4i32 = 1
		f132(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l2
		f278(ctx, s0i32)
		s0i32 = l6
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l7
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
	}
	s0i32 = l4
	s0i32 = f23(ctx, s0i32)
	s0i32 = l2
	f34(ctx, s0i32)
lbl1:
	s0i32 = l5
	s1i32 = 3760
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l5
	s1i32 = 3740
	s0i32 = s0i32 + s1i32
	f40(ctx, s0i32)
lbl0:
	s0i32 = l5
	s1i32 = 4848
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
