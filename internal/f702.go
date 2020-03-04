package internal

import (
	"unsafe"
)

func f702(ctx *Context, l0 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	s1i32 = 204
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = -204
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l1 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s0i32 = l0
		s1i32 = l1
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l1 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l3
			s0i32 = s0i32 - s1i32
			l1 = s0i32
			s1i32 = 2
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l6 = s0i32
			s0i32 = l3
			s1i32 = l3
			s2i32 = l5
			s1i32 = s1i32 - s2i32
			s2i32 = 2
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			s2i32 = -2
			s1i32 = i32DivS(s1i32, s2i32)
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			l5 = s1i32
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l0
			s1i32 = l1
			if s1i32 != 0 {
				s1i32 = l2
				s2i32 = l3
				s3i32 = l1
				s1i32 = f106(ctx, s1i32, s2i32, s3i32)
				s1i32 = l0
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			} else {
				s1i32 = l3
			}
			s2i32 = l5
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l2
			s2i32 = l6
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			goto lbl3
		}
		s0i32 = l1
		s1i32 = l5
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
		l6 = s0i32
		s0i32 = f17(ctx, s0i32)
		l7 = s0i32
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l7
		s1i32 = l1
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			l1 = s0i32
		lbl7:
			s0i32 = l1
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l3
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s1i32 = l2
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl7
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l5 = s0i32
		}
		s0i32 = l0
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l2 = s0i32
			goto lbl3
		}
		s0i32 = l5
		f12(ctx, s0i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
	lbl3:
		s0i32 = l2
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l5 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l6 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = s1i32 - s2i32
	l2 = s1i32
	s2i32 = 2
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 4080
			s1i32 = f17(ctx, s1i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l4
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			f455(ctx, s0i32, s1i32)
			goto lbl1
		}
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
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l1 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s0i32 = l0
		s1i32 = l1
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l1 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l3
			s0i32 = s0i32 - s1i32
			l1 = s0i32
			s1i32 = 2
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l6 = s0i32
			s0i32 = l3
			s1i32 = l3
			s2i32 = l5
			s1i32 = s1i32 - s2i32
			s2i32 = 2
			s1i32 = s1i32 >> (uint32(s2i32) & 31)
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			s2i32 = -2
			s1i32 = i32DivS(s1i32, s2i32)
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			l5 = s1i32
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l0
			s1i32 = l1
			if s1i32 != 0 {
				s1i32 = l2
				s2i32 = l3
				s3i32 = l1
				s1i32 = f106(ctx, s1i32, s2i32, s3i32)
				s1i32 = l0
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			} else {
				s1i32 = l3
			}
			s2i32 = l5
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l2
			s2i32 = l6
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			goto lbl11
		}
		s0i32 = l1
		s1i32 = l5
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
		l6 = s0i32
		s0i32 = f17(ctx, s0i32)
		l7 = s0i32
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l7
		s1i32 = l1
		s2i32 = -4
		s1i32 = s1i32 & s2i32
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		l1 = s0i32
		s0i32 = l2
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l6
			l1 = s0i32
		lbl15:
			s0i32 = l1
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l3
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s1i32 = l2
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl15
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l5 = s0i32
		}
		s0i32 = l0
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l2 = s0i32
			goto lbl11
		}
		s0i32 = l5
		f12(ctx, s0i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l2 = s0i32
	lbl11:
		s0i32 = l2
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = l4
	s1i32 = l0
	s2i32 = 12
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = 1
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l8 = s0i32
	s1i32 = 1073741824
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l8
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l7 = s1i32
	s1i32 = f17(ctx, s1i32)
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	s2i32 = l5
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l6
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = 4080
	s0i32 = f17(ctx, s0i32)
	l9 = s0i32
	s0i32 = l5
	s1i32 = l8
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l2
	s1i32 = l6
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l2
		s2i32 = l2
		s3i32 = l6
		s2i32 = s2i32 - s3i32
		s3i32 = 2
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s3i32 = -2
		s2i32 = i32DivS(s2i32, s3i32)
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl17
	}
	s0i32 = l7
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
	s0i32 = l4
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l2 = s1i32
	s1i32 = f17(ctx, s1i32)
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	s2i32 = l5
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	s2i32 = l1
	s3i32 = -4
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	f12(ctx, s0i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
	s0i32 = l5
	l6 = s0i32
lbl17:
	s0i32 = l2
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl20:
		s0i32 = l4
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		f280(ctx, s0i32, s1i32)
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l1 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l3 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l7 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l6 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l2 = s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l2 = s0i32
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l3
		s2i32 = l3
		s3i32 = l1
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
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l5
	f12(ctx, s0i32)
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
