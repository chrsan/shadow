package internal

import (
	"unsafe"
)

func f653(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l12 int64
	_ = l12
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
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 640
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = l3
		s1i32 = 20
		s0i32 = i32DivU(s0i32, s1i32)
		l4 = s0i32
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l0 = s0i32
			s1i32 = 1
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			l2 = s0i32
			if s0i32 != 0 {
			lbl5:
				s0i32 = l1
				s1i32 = l2
				s2i32 = l0
				f1893(ctx, s0i32, s1i32, s2i32)
				s0i32 = l2
				s1i32 = -1
				s0i32 = s0i32 + s1i32
				l2 = s0i32
				if s0i32 != 0 {
					goto lbl5
				}
				s0i32 = l3
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl0
				}
			}
		lbl6:
			s0i32 = l5
			s1i32 = l1
			s2i32 = 16
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l1
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l5
			s1i32 = l1
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l1
			s1i32 = l1
			s2i32 = l4
			s3i32 = 20
			s2i32 = s2i32 * s3i32
			s1i32 = s1i32 + s2i32
			l0 = s1i32
			l2 = s1i32
			s2i32 = 16
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l1
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l2
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l5
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l0
			s1i32 = l5
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l1
			s1i32 = l4
			f1892(ctx, s0i32, s1i32)
			s0i32 = l4
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l4 = s0i32
			if s0i32 != 0 {
				goto lbl6
			}
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l4
		s2i32 = 1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 20
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l7 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s0i32 = l5
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l2
		l6 = s1i32
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		l3 = s0i32
		s1i32 = l2
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l7
			s2i32 = l4
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
			s0i32 = l1
			l4 = s0i32
		lbl8:
			s0i32 = l8
			s1i32 = l4
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l9 = s1i32
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l4
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
				l9 = s0i32
				s1i32 = l4
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
				l10 = s1i32
				s2i32 = l9
				s3i32 = l10
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
				s1i32 = l7
				if s0i32 < s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl10
				}
				goto lbl9
			}
			s0i32 = l9
			s1i32 = l8
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl9
			}
		lbl10:
			s0i32 = l5
			s1i32 = l4
			s2i32 = 16
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l4
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l5
			s1i32 = l4
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l4
			s1i32 = l3
			s2i32 = 16
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l3
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l4
			s1i32 = l3
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l5
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l5
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			l3 = s0i32
		lbl9:
			s0i32 = l4
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			l4 = s0i32
			s1i32 = l2
			if uint32(s0i32) < uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
		}
		s0i32 = l5
		s1i32 = l3
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = l3
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l3
		s3i32 = -20
		s2i32 = s2i32 + s3i32
		f653(ctx, s0i32, s1i32, s2i32)
		s0i32 = l2
		s1i32 = l3
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s1i32 = 640
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l1
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	l0 = s0i32
lbl12:
	s0i32 = l0
	l6 = s0i32
	s0i32 = l3
	l0 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l7 = s0i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		l8 = s1i32
		s2i32 = l7
		s3i32 = l8
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
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l3 = s1i32
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l4 = s2i32
		s3i32 = l3
		s4i32 = l4
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
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		goto lbl13
	}
	s0i32 = l3
	s1i32 = l4
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l7 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l8 = s0i32
lbl14:
	s0i32 = l6
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l12 = s0i64
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l9 = s0i32
	s0i32 = l0
	s1i32 = l0
	s2i32 = -20
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	l6 = s0i32
	s0i32 = l3
	s1i32 = l1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
	s0i32 = l7
	s1i32 = l8
	s2i32 = l7
	s3i32 = l8
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
	l10 = s0i32
lbl17:
	s0i32 = l3
	l4 = s0i32
	s0i32 = l6
	s1i32 = -40
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s1i32 = l9
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l10
		s1i32 = l6
		s2i32 = -32
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l3 = s1i32
		s2i32 = l6
		s3i32 = -36
		s2i32 = s2i32 + s3i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l11 = s2i32
		s3i32 = l3
		s4i32 = l11
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
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
		goto lbl16
	}
	s0i32 = l9
	s1i32 = l3
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl16
	}
lbl18:
	s0i32 = l4
	s1i32 = l4
	s2i32 = -20
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l4
	l6 = s0i32
	s0i32 = l3
	s1i32 = l1
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
lbl16:
	s0i32 = l6
	s1i32 = -20
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i64 = l12
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl13:
	s0i32 = l0
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
lbl0:
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
