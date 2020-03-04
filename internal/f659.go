package internal

import (
	"unsafe"
)

func f659(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
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
	var s2i64 int64
	_ = s2i64
	s0i32 = ctx.g0
	s1i32 = 1168
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 5
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l3 = s0i32
			goto lbl2
		}
		s0i32 = l1
		s1i32 = l3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l3 = s0i32
			goto lbl2
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl5:
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l0 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	lbl2:
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = -1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l2
	s1i32 = 4
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = -1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l0
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s0i32 = l1
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l2 = s0i32
		s1i32 = -1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
		s0i32 = l1
		l2 = s0i32
		s0i32 = l0
		l1 = s0i32
		goto lbl7
	}
	s0i32 = l2
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl11
	case 1:
		goto lbl10
	case 2:
		goto lbl9
	default:
		goto lbl14
	}
lbl14:
	s0i32 = l7
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l6
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l6 = s0i32
	s0i32 = l0
	l2 = s0i32
	goto lbl7
lbl12:
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl15:
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	goto lbl0
lbl11:
	s0i32 = l7
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l6
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l4
	s1i32 = 1112
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l1
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		goto lbl16
	}
lbl17:
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl18:
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	goto lbl0
lbl16:
	s0i32 = l6
	s1i32 = l7
	s0i32 = s0i32 | s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1120)]))
		l2 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l4
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1112)])))
		s0i64 = s0i64 - s1i64
		l16 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1124)]))
		l6 = s1i32
		s1i64 = int64(s1i32)
		s2i32 = l4
		s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1116)])))
		s1i64 = s1i64 - s2i64
		l17 = s1i64
		s2i64 = 1
		if s1i64 < s2i64 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 | s1i32
		s1i64 = l16
		s2i64 = l17
		s1i64 = s1i64 | s2i64
		s2i64 = 2147483648
		s1i64 = s1i64 + s2i64
		s2i64 = 4294967295
		if uint64(s1i64) > uint64(s2i64) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 | s1i32
		l1 = s0i32
		s0i32 = l3
		if s0i32 != 0 {
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l0 = s0i32
			s0i32 = 0
			s1i32 = l6
			s2i32 = 2147483647
			if s1i32 != s2i32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			s2i32 = l2
			s3i32 = 2147483647
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			s3i32 = l1
			s2i32 = s2i32 | s3i32
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
				s0i32 = l0
				s1i32 = 1
				s0i32 = s0i32 + s1i32
				s1i32 = 2
				if uint32(s0i32) < uint32(s1i32) {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl22
				}
				s0i32 = l0
				s1i32 = l0
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
				l0 = s1i32
				s2i32 = -1
				s1i32 = s1i32 + s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l0
				s1i32 = 1
				if s0i32 != s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl22
				}
				s0i32 = l3
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
				f13(ctx, s0i32)
			lbl22:
				s0i32 = l3
				s1i64 = 0
				*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
				s0i32 = l3
				s1i64 = 0
				*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
				s0i32 = l3
				s1i32 = -1
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
				goto lbl0
			}
			s0i32 = 1
			l5 = s0i32
			s0i32 = l0
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			s1i32 = 2
			if uint32(s0i32) < uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl23
			}
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l0 = s1i32
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = 1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl23
			}
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			f13(ctx, s0i32)
		lbl23:
			s0i32 = l3
			s1i32 = l4
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+1112)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l4
			s2i32 = 1120
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			goto lbl0
		}
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l8 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s1i32 = l9
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s1i32 = l12
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l8
	s0i64 = int64(s0i32)
	s1i32 = l9
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l10
	s0i64 = int64(s0i32)
	s1i32 = l2
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		l3 = s0i32
		goto lbl25
	}
	s0i32 = l1
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		l3 = s0i32
		goto lbl25
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl28
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl28
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl28:
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl25:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	goto lbl0
lbl24:
	s0i32 = 1
	l12 = s0i32
	s0i32 = l6
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l8 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s1i32 = l9
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s1i32 = l6
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l8
	s0i64 = int64(s0i32)
	s1i32 = l9
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l10
	s0i64 = int64(s0i32)
	s1i32 = l2
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		l3 = s0i32
		goto lbl29
	}
	s0i32 = l0
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		l3 = s0i32
		goto lbl29
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl32:
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl29
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl29:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	goto lbl0
lbl10:
	s0i32 = l7
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl34
		}
		s0i32 = l1
		s1i32 = l3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl34
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl35
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl35
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl35:
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l6 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl34
		}
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l6 = s0i32
	lbl34:
		s0i32 = l6
		s1i32 = -1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l6
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl37
		}
		s0i32 = l0
		s1i32 = l3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl37
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l1 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl38
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl38
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl38:
		s0i32 = l3
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l7 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl37
		}
		s0i32 = l7
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l7 = s0i32
	lbl37:
		s0i32 = l7
		s1i32 = -1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		goto lbl0
	}
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l8 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s1i32 = l9
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s1i32 = l12
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l8
	s0i64 = int64(s0i32)
	s1i32 = l9
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l10
	s0i64 = int64(s0i32)
	s1i32 = l2
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = 0
	l1 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl40
	}
	s0i32 = l0
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl40
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl41
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl41
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl41:
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl40
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
lbl40:
	s0i32 = l1
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	goto lbl0
lbl39:
	s0i32 = 2
	l12 = s0i32
	s0i32 = l6
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l2 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l8 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l2 = s0i32
	s1i32 = l9
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s1i32 = l6
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l8
	s0i64 = int64(s0i32)
	s1i32 = l9
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l16 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l10
	s0i64 = int64(s0i32)
	s1i32 = l2
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l17 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i64 = l16
	s1i64 = l17
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = 0
	l2 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl42
	}
	s0i32 = l1
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl42
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl43
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl43
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl43:
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl42
	}
	s0i32 = l2
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i32
lbl42:
	s0i32 = l2
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	goto lbl0
lbl9:
	s0i32 = l7
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl45
		}
		s0i32 = l1
		s1i32 = l3
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl45
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl46
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl46
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl46:
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l6 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl45
		}
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l6 = s0i32
	lbl45:
		s0i32 = l6
		s1i32 = -1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		goto lbl0
	}
	s0i32 = 3
	l12 = s0i32
	s0i32 = l6
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl47
	}
	s0i32 = l0
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl47
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl48
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl48
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl48:
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl47
	}
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0i32
lbl47:
	s0i32 = l7
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	goto lbl0
lbl8:
	s0i32 = l0
	l2 = s0i32
	goto lbl6
lbl7:
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l2
		s2i32 = l1
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			goto lbl49
		}
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l3 = s0i32
		goto lbl51
	}
	s0i32 = l2
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl51
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl53
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl53
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl53:
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl51
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl51:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	goto lbl0
lbl49:
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = 0
	l5 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl54
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl54
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl54:
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	goto lbl0
lbl6:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 2147483647
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1072)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = 1072
			s0i32 = s0i32 + s1i32
			goto lbl55
		}
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1072)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l0 = s0i32
		s0i32 = l4
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1080)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1084)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l0 = s0i32
		s0i32 = l4
		s1i64 = 9223372034707292159
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1092)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1088)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1072
		s0i32 = s0i32 + s1i32
		goto lbl55
	}
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
lbl55:
	l7 = s0i32
	s0i32 = l11
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 2147483647
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1040)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = 1040
			s0i32 = s0i32 + s1i32
			goto lbl58
		}
		s0i32 = l4
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1040)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l0 = s0i32
		s0i32 = l4
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1048)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1044)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1052)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l0 = s0i32
		s0i32 = l4
		s1i64 = 9223372034707292159
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1060)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1056)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 1040
		s0i32 = s0i32 + s1i32
		goto lbl58
	}
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
lbl58:
	l11 = s0i32
	s0i32 = l4
	s1i64 = 1099511627776
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1032)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4832
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1152)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 4840
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1160)])) = uint64(s1i64)
	s0i32 = l11
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = l11
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1144)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 4294967296
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l12
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 4848
	s1i32 = s1i32 + s2i32
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+1128)])) = uint16(s1i32)
	s0i32 = l4
	s1i32 = l2
	s2i32 = l1
	s3i32 = l2
	s4i32 = l1
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1148)])) = uint32(s1i32)
	s0i32 = 2147483647
	l6 = s0i32
	s0i32 = l4
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1132)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l5
	s2i32 = 2147483647
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 0
	s3i32 = l0
	s4i32 = 2147483647
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
	if s1i32 != 0 {
		goto lbl62
	}
	s0i32 = l11
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s0i32 = l7
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l4
	s1i32 = 1152
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 | s1i32
	l12 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl64:
		s0i32 = l1
		s1i32 = l2
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l5
			s2i32 = l5
			s3i32 = l2
			if s2i32 > s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l8 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s0i32 = l2
			s1i32 = l1
			s2i32 = l8
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l8 = s0i32
			s0i32 = l5
			s1i32 = l2
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l13 = s0i32
			s0i32 = 0
			l14 = s0i32
			s0i32 = l2
			l9 = s0i32
			s0i32 = l11
			l15 = s0i32
			s0i32 = l12
			goto lbl65
		}
		s0i32 = l2
		s1i32 = l1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l0
			s2i32 = l0
			s3i32 = l1
			if s2i32 > s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l8 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s0i32 = l1
			s1i32 = l2
			s2i32 = l8
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l9 = s0i32
			s0i32 = l0
			s1i32 = l1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l14 = s0i32
			s0i32 = 0
			l13 = s0i32
			s0i32 = l1
			l8 = s0i32
			s0i32 = l2
			l1 = s0i32
			s0i32 = l12
			goto lbl67
		}
		s0i32 = 0
		s1i32 = l5
		s2i32 = l5
		s3i32 = l0
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l9 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l0
		s2i32 = l0
		s3i32 = l5
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l8 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l1
		s1i32 = l0
		s2i32 = l8
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l8 = s0i32
		s0i32 = l2
		s1i32 = l5
		s2i32 = l9
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l9 = s0i32
		s0i32 = l0
		s1i32 = l5
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l14 = s0i32
		s0i32 = l5
		s1i32 = l0
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l13 = s0i32
		s0i32 = l11
	lbl67:
		l15 = s0i32
		s0i32 = l10
	lbl65:
		l2 = s0i32
		s0i32 = l1
		s1i32 = l6
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 1128
			s0i32 = s0i32 + s1i32
			s1i32 = l1
			s2i32 = l12
			s3i32 = l12
			f323(ctx, s0i32, s1i32, s2i32, s3i32)
		}
		s0i32 = l4
		s1i32 = 1128
		s0i32 = s0i32 + s1i32
		s1i32 = l7
		s2i32 = l15
		s3i32 = l2
		f323(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = -1
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1144)]))
		if s1i32 != 0 {
			goto lbl61
		}
		s0i32 = l13
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l8
			l1 = s0i32
			goto lbl70
		}
		s0i32 = 2147483647
		s1i32 = l5
		s2i32 = l11
		s3i32 = l11
		s4i32 = -4
		s3i32 = s3i32 + s4i32
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = 4
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 + s3i32
		l2 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l5 = s2i32
		s3i32 = 2147483647
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
		l1 = s0i32
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l11 = s0i32
	lbl70:
		s0i32 = l14
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l9
			l2 = s0i32
			goto lbl72
		}
		s0i32 = 2147483647
		s1i32 = l0
		s2i32 = l10
		s3i32 = l10
		s4i32 = -4
		s3i32 = s3i32 + s4i32
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = 4
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 + s3i32
		l6 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l0 = s2i32
		s3i32 = 2147483647
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
		l2 = s0i32
		s0i32 = l6
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l10 = s0i32
	lbl72:
		s0i32 = l7
		l6 = s0i32
		s0i32 = l5
		s1i32 = 2147483647
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl64
		}
		s0i32 = l0
		s1i32 = 2147483647
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl64
		}
		s0i32 = 0
		goto lbl62
	}
lbl74:
	s0i32 = l1
	s1i32 = l2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			s1i32 = l5
			s2i32 = l5
			s3i32 = l0
			if s2i32 > s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l8 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s1i32 = l0
			s2i32 = l0
			s3i32 = l5
			if s2i32 > s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			l9 = s2i32
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s0i32 = l1
			s1i32 = l0
			s2i32 = l9
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l9 = s0i32
			s0i32 = l2
			s1i32 = l5
			s2i32 = l8
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l8 = s0i32
			s0i32 = l0
			s1i32 = l5
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l15 = s0i32
			s0i32 = l5
			s1i32 = l0
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l13 = s0i32
			s0i32 = l10
			l14 = s0i32
			s0i32 = l11
			goto lbl75
		}
		s0i32 = l1
		s1i32 = l0
		s2i32 = l0
		s3i32 = l1
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l8 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l1
		s1i32 = l2
		s2i32 = l8
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l8 = s0i32
		s0i32 = l0
		s1i32 = l1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l15 = s0i32
		s0i32 = 0
		l13 = s0i32
		s0i32 = l1
		l9 = s0i32
		s0i32 = l2
		l1 = s0i32
		s0i32 = l10
		l14 = s0i32
		s0i32 = l12
		goto lbl75
	}
	s0i32 = l2
	s1i32 = l5
	s2i32 = l5
	s3i32 = l2
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l7 = s0i32
	s0i32 = l2
	s1i32 = l1
	s2i32 = l8
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s0i32 = l5
	s1i32 = l2
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l13 = s0i32
	s0i32 = 0
	l15 = s0i32
	s0i32 = l2
	l8 = s0i32
	s0i32 = l12
	l14 = s0i32
	s0i32 = l11
lbl75:
	l2 = s0i32
	s0i32 = l1
	s1i32 = l6
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 1128
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l12
		s3i32 = l12
		f323(ctx, s0i32, s1i32, s2i32, s3i32)
	}
	s0i32 = l4
	s1i32 = 1128
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	s2i32 = l2
	s3i32 = l14
	f323(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l13
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		l1 = s0i32
		goto lbl79
	}
	s0i32 = 2147483647
	s1i32 = l5
	s2i32 = l11
	s3i32 = l11
	s4i32 = -4
	s3i32 = s3i32 + s4i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s4i32 = 4
	s3i32 = s3i32 | s4i32
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l5 = s2i32
	s3i32 = 2147483647
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
	l1 = s0i32
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l11 = s0i32
lbl79:
	s0i32 = l15
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		l2 = s0i32
		goto lbl81
	}
	s0i32 = 2147483647
	s1i32 = l0
	s2i32 = l10
	s3i32 = l10
	s4i32 = -4
	s3i32 = s3i32 + s4i32
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s4i32 = 4
	s3i32 = s3i32 | s4i32
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l0 = s2i32
	s3i32 = 2147483647
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
	l2 = s0i32
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l10 = s0i32
lbl81:
	s0i32 = l7
	l6 = s0i32
	s0i32 = l5
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl74
	}
	s0i32 = l0
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl74
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1144)]))
lbl62:
	l10 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1132)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1032)]))
	l0 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1136)]))
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1148)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1140)]))
	s2i32 = l10
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l10
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1140)]))
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1136)]))
	s0i32 = s0i32 - s1i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
lbl61:
	l1 = s0i32
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1032)]))
		s2i32 = l1
		s0i32 = f1903(ctx, s0i32, s1i32, s2i32)
		goto lbl83
	}
	s0i32 = l1
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l1
	s2i32 = 2
	if s1i32 > s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 | s1i32
lbl83:
	l5 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)]))
	l0 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)])) = uint32(s1i32)
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	f13(ctx, s0i32)
lbl0:
	s0i32 = l4
	s1i32 = 1168
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
