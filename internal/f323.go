package internal

import (
	"unsafe"
)

func f323(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0i32 = s0i32 + s1i32
	l15 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	l9 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l18 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	s0i32 = l2
	l5 = s0i32
lbl0:
	s0i32 = l5
	l6 = s0i32
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l15
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l12 = s0i32
	s0i32 = l3
	l5 = s0i32
lbl1:
	s0i32 = l5
	l7 = s0i32
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1028)]))
	s1i32 = l15
	s2i32 = l6
	s3i32 = l2
	s2i32 = s2i32 - s3i32
	s3i32 = 2
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = l3
	s2i32 = s2i32 - s3i32
	s3i32 = 2
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1032)]))
		l16 = s0i32
		goto lbl2
	}
	s0i32 = 0
	s1i32 = l5
	s2i32 = 1
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l5
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)]))
	l5 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l7
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = f29(ctx, s0i32, s1i32)
lbl4:
	l6 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)]))
	l5 = s0i32
	s0i32 = l4
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)])) = uint32(s1i32)
	s0i32 = l5
	if s0i32 != 0 {
		s0i32 = l5
		f13(ctx, s0i32)
	}
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1032)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1024)]))
		s1i32 = l4
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+1028)]))
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l4
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1028)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1024)]))
	l16 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1032)])) = uint32(s1i32)
lbl2:
	s0i32 = l16
	s1i32 = l12
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s1i32 = 2147483647
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l4 = s2i32
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
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s1i32 = l18
		s0i32 = s0i32 - s1i32
		l20 = s0i32
		s0i32 = l2
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l3
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l13 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0i32
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l10 = s0i32
		s0i32 = 1
		l19 = s0i32
	lbl8:
		s0i32 = l5
		s1i32 = l4
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 1
			l14 = s0i32
			s0i32 = l6
			s1i32 = l4
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l6
				l2 = s0i32
				s0i32 = l5
				l3 = s0i32
				s0i32 = l4
				l9 = s0i32
				l7 = s0i32
				goto lbl9
			}
			s0i32 = l8
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l2 = s0i32
			s0i32 = l8
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l9 = s0i32
			s0i32 = l8
			s1i32 = 8
			s0i32 = s0i32 + s1i32
			l8 = s0i32
			s0i32 = l5
			l3 = s0i32
			s0i32 = l6
			l7 = s0i32
			goto lbl9
		}
		s0i32 = l4
		s1i32 = l5
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 2
			l14 = s0i32
			s0i32 = l6
			l2 = s0i32
			s0i32 = l4
			l3 = s0i32
			s0i32 = l10
			l7 = s0i32
			s0i32 = l10
			s1i32 = l5
			l9 = s1i32
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			s0i32 = l9
			l7 = s0i32
			l4 = s0i32
			goto lbl9
		}
		s0i32 = l10
		s1i32 = l6
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l17 = s0i32
		s0i32 = l6
		s1i32 = l10
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
		if s0i32 != 0 {
			s0i32 = 0
			s1i32 = l6
			s2i32 = l2
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s1i32 = l10
			s2i32 = l17
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l7 = s0i32
			s0i32 = l5
			s1i32 = l10
			s2i32 = l17
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			l9 = s0i32
			s0i32 = 3
			l14 = s0i32
			s0i32 = l6
			l2 = s0i32
			s0i32 = l5
			l3 = s0i32
			goto lbl12
		}
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l2 = s0i32
		s0i32 = l8
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0i32
		s0i32 = 3
		l14 = s0i32
		s0i32 = l8
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		l8 = s0i32
		s0i32 = l5
		l3 = s0i32
		s0i32 = l10
		l7 = s0i32
		s0i32 = l17
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l4
		l8 = s0i32
		s0i32 = l6
		l7 = s0i32
		l4 = s0i32
		goto lbl9
	lbl12:
		s0i32 = l13
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l10 = s0i32
		s0i32 = l13
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = l13
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l13 = s0i32
	lbl9:
		s0i32 = l14
		s1i32 = l18
		s0i32 = s0i32 - s1i32
		s1i32 = l20
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l3
		s1i32 = l7
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l19
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l11
			s1i32 = -4
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = l3
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl16
			}
		}
		s0i32 = l11
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l11
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l11
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l11 = s0i32
	lbl16:
		s0i32 = l5
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 0
		l19 = s0i32
	lbl15:
		s0i32 = l2
		l6 = s0i32
		s0i32 = l9
		l5 = s0i32
		s1i32 = 2147483647
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l4
		s1i32 = 2147483647
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
	}
	s0i32 = l11
	s1i32 = 2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l2 = s0i32
	s1i32 = l11
	s2i32 = l16
	s1i32 = s1i32 - s2i32
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	s2i32 = l12
	s1i32 = s1i32 - s2i32
	l3 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1032)]))
		l4 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l2
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s1i32 = l4
			s2i32 = l12
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = l3
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s3i32 = -4
			s2i32 = s2i32 + s3i32
			s0i32 = f102(ctx, s0i32, s1i32, s2i32)
			if s0i32 != 0 {
				goto lbl19
			}
		}
		s0i32 = l2
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l3
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
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1032)]))
	l4 = s0i32
lbl19:
	s0i32 = l4
	s1i32 = l15
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l3
	s2i32 = 1
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	return
lbl18:
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
}
