package internal

import (
	"unsafe"
)

func f921(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l16 int32
	_ = l16
	var l17 int32
	_ = l17
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int32
	_ = l22
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
	s0i32 = ctx.g0
	l4 = s0i32
	l14 = s0i32
	s0i32 = l4
	s1i32 = 5280
	s0i32 = s0i32 - s1i32
	s1i32 = -32
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l6 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l7 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l5
	s2i32 = l6
	s3i32 = l7
	s2i32 = s2i32 + s3i32
	s3i32 = 5
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 - s2i32
	s2i32 = 32768
	s1i32 = s1i32 + s2i32
	l13 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s2i32 = 640
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s2i32 = 1152
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s2i32 = 3200
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 12
	s0i32 = s0i32 | s1i32
	l15 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 8
	s0i32 = s0i32 | s1i32
	l16 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	l17 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 28
	s0i32 = s0i32 | s1i32
	l18 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 24
	s0i32 = s0i32 | s1i32
	l19 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 20
	s0i32 = s0i32 | s1i32
	l20 = s0i32
	s0i32 = l4
	s1i32 = 640
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = s0i32 | s1i32
	l21 = s0i32
lbl1:
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1176)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1168)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1160)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1152)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+640)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+648)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+656)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = -1
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+664)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l10
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l8 = s0i32
	s0i32 = l5
	s1i32 = 8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l12 = s0i32
	if s0i32 != 0 {
	lbl3:
		s0i32 = l5
		s1i32 = 3
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 1152
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s2i32 = 1
			if s1i32 == s2i32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			l9 = s1i32
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = 8
			s3i32 = 12
			s4i32 = l5
			s5i32 = 2
			if s4i32 == s5i32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			l7 = s4i32
			if s4i32 != 0 {
				// s2i32 = s2i32
			} else {
				s2i32 = s3i32
			}
			s3i32 = l5
			s4i32 = 2
			if s3i32 < s4i32 {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			l6 = s3i32
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s0i32 = s0i32 | s1i32
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l16
			s1i32 = l15
			s2i32 = l7
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s0i32 = l17
			s1i32 = l4
			s2i32 = 640
			s1i32 = s1i32 + s2i32
			s2i32 = l9
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			goto lbl4
		}
		s0i32 = l4
		s1i32 = 1152
		s0i32 = s0i32 + s1i32
		s1i32 = 20
		s2i32 = 16
		s3i32 = l5
		s4i32 = -4
		s3i32 = s3i32 + s4i32
		l6 = s3i32
		s4i32 = 1
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l9 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 24
		s3i32 = 28
		s4i32 = l6
		s5i32 = 2
		if s4i32 == s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		l7 = s4i32
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s3i32 = l5
		s4i32 = 6
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l6 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = s0i32 | s1i32
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l19
		s1i32 = l18
		s2i32 = l7
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l7 = s0i32
		s0i32 = l20
		s1i32 = l21
		s2i32 = l9
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
	lbl4:
		s1i32 = l7
		s2i32 = l6
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 7
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l8
	s1i32 = 32
	s2i32 = l12
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l12 = s0i32
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = 0
	l5 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl7:
		s0i32 = l13
		s1i32 = l5
		s2i32 = 5
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l5
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l10
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l12
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
	}
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l11
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l11
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l11
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)]))
	l5 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
	l6 = s1i32
	s2i32 = 204
	s1i32 = i32DivU(s1i32, s2i32)
	l8 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l4
	s1i32 = l5
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+104)]))
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5276)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5272)])) = uint32(s2i32)
		s1i32 = l5
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+116)]))
		s3i32 = l6
		s2i32 = s2i32 + s3i32
		s3i32 = 204
		s2i32 = i32DivU(s2i32, s3i32)
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s1i32 = 0
		goto lbl8
	}
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1i32
	s1i32 = l4
	s2i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5272)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l9
	s3i32 = l6
	s4i32 = l8
	s5i32 = 204
	s4i32 = s4i32 * s5i32
	s3i32 = s3i32 - s4i32
	s4i32 = 20
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+5276)])) = uint32(s2i32)
	s1i32 = l5
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+116)]))
	s3i32 = l6
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s3i32 = 204
	s2i32 = i32DivU(s2i32, s3i32)
	l7 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l6
	s3i32 = l7
	s4i32 = 204
	s3i32 = s3i32 * s4i32
	s2i32 = s2i32 - s3i32
	s3i32 = 20
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
lbl8:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5268)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+5264)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+5272)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+5264)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l11
	s1i32 = l4
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	f705(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = l4
	s3i32 = 24
	s2i32 = s2i32 + s3i32
	s3i32 = l10
	s4i32 = l4
	s5i32 = 124
	s4i32 = s4i32 + s5i32
	s0i32 = f704(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l5 = s0i32
	s0i32 = l4
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	f456(ctx, s0i32)
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l8 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l6 = s0i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l9 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl11:
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = l6
		s0i32 = int32(ctx.Mem[int(s0i32+4)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l7
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l13
		s1i32 = l8
		s2i32 = 5
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l22 = s0i32
		s0i32 = 0
		l5 = s0i32
	lbl13:
		s0i32 = l3
		s1i32 = l5
		s2i32 = l8
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l10
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l22
		s2i32 = l5
		s3i32 = 5
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l12
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
	lbl12:
		s0i32 = l7
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l6
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l9
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
	}
	s0i32 = l10
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = l2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl0:
	s0i32 = l4
	s1i32 = 72
	s0i32 = s0i32 + s1i32
	f456(ctx, s0i32)
	s0i32 = l14
	ctx.g0 = s0i32
}
