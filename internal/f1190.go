package internal

import (
	"unsafe"
)

func f1190(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
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
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l11 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		return
	}
	s0i32 = l11
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l11
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l4
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l13 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l14 = s1i32
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	l9 = s0i32
	s0i32 = l10
	l7 = s0i32
lbl2:
	s0i32 = l6
	s1i32 = l8
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l8
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = 0
	l5 = s0i32
	s0i32 = l10
	l3 = s0i32
lbl3:
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 != 0 {
		s0i32 = l1
		l2 = s0i32
		goto lbl4
	}
	s0i32 = l5
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l12 = s1i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = l12
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl4:
	s0i32 = l5
	l1 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = 0
	l1 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl7:
	s0i32 = 0
	s1i32 = l1
	s2i32 = l2
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	l12 = s2i32
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s2i32 = int32(ctx.Mem[int(s2i32+1)])
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s1i32 = l12
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
lbl6:
	s0i32 = l6
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l1
	s1i32 = l7
	s2i32 = l1
	s3i32 = l7
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
	l7 = s0i32
	s1i32 = l5
	s2i32 = l9
	s3i32 = l5
	s4i32 = l9
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
	l9 = s1i32
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	return
lbl1:
	s0i32 = l9
	s1i32 = l10
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l11
		s1i32 = l11
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			f13(ctx, s0i32)
		}
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		return
	}
	s0i32 = l0
	s1i32 = l13
	s2i32 = l7
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l9
	s2i32 = l14
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l4
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l9
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl12:
		s0i32 = l8
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l10
		l1 = s0i32
		s0i32 = l9
		l5 = s0i32
		s0i32 = 0
		l6 = s0i32
	lbl14:
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = l3
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l0 = s1i32
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s0i32 = l5
		s1i32 = l0
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l6
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l2
		l3 = s0i32
		s0i32 = l5
		s1i32 = l0
		s0i32 = s0i32 - s1i32
		l5 = s0i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		goto lbl13
	lbl15:
		s0i32 = l3
		s1i32 = l0
		s2i32 = l5
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	lbl13:
		s0i32 = l7
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l1
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl18:
			s0i32 = l2
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			l0 = s0i32
			s0i32 = l2
			s1i32 = 2
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l1
			s1i32 = l0
			s0i32 = s0i32 - s1i32
			l1 = s0i32
			s1i32 = 0
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
		}
		s0i32 = l7
		l1 = s0i32
	lbl19:
		s0i32 = l1
		s1i32 = l2
		s2i32 = -2
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l0 = s1i32
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l0
			s0i32 = s0i32 - s1i32
			l1 = s0i32
			s1i32 = 0
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl19
			}
			goto lbl16
		}
		s0i32 = l2
		s1i32 = l0
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	lbl16:
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l8
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		return
	}
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l10
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l3 = s0i32
lbl21:
	s0i32 = l8
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l10
	l1 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl23:
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l0 = s0i32
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = l0
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl23
		}
	}
	s0i32 = l7
	l1 = s0i32
lbl24:
	s0i32 = l1
	s1i32 = l2
	s2i32 = -2
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l0 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		goto lbl25
	}
	s0i32 = l1
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
lbl25:
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l8
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
lbl10:
}
