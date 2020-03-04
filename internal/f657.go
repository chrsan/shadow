package internal

import (
	"unsafe"
)

func f657(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
	var l15 int64
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l5 = s0i32
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
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
			goto lbl3
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
			goto lbl3
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl3:
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl1
	}
	s0i32 = 2147483647
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l6 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = -2147483648
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l7 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = l2
	s3i32 = l7
	s3i64 = int64(s3i32)
	l10 = s3i64
	s4i32 = l2
	s4i64 = int64(s4i32)
	s3i64 = s3i64 + s4i64
	s4i64 = -2147483648
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l2 = s1i32
	s2i32 = l6
	s2i64 = int64(s2i32)
	l11 = s2i64
	s3i32 = l2
	s3i64 = int64(s3i32)
	s2i64 = s2i64 + s3i64
	s3i64 = 2147483647
	if s2i64 > s3i64 {
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
	s0i32 = 2147483647
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l2 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = -2147483648
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l7 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = l1
	s3i32 = l7
	s3i64 = int64(s3i32)
	l12 = s3i64
	s4i32 = l1
	s4i64 = int64(s4i32)
	s3i64 = s3i64 + s4i64
	s4i64 = -2147483648
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l1 = s1i32
	s2i32 = l2
	s2i64 = int64(s2i32)
	l13 = s2i64
	s3i32 = l1
	s3i64 = int64(s3i32)
	s2i64 = s2i64 + s3i64
	s3i64 = 2147483647
	if s2i64 > s3i64 {
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
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l0 = s0i32
		s0i64 = l11
		s1i32 = l6
		s1i64 = int64(s1i32)
		l14 = s1i64
		s0i64 = s0i64 + s1i64
		l11 = s0i64
		s1i64 = 2147483647
		s2i64 = l11
		s3i64 = 2147483647
		if s2i64 < s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l11 = s0i64
		s1i64 = -2147483647
		s2i64 = l11
		s3i64 = -2147483647
		if s2i64 > s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l11 = s0i64
		s0i32 = int32(s0i64)
		l1 = s0i32
		s1i32 = 2147483647
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i64 = l13
		s1i32 = l7
		s1i64 = int64(s1i32)
		l15 = s1i64
		s0i64 = s0i64 + s1i64
		l13 = s0i64
		s1i64 = 2147483647
		s2i64 = l13
		s3i64 = 2147483647
		if s2i64 < s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l13 = s0i64
		s1i64 = -2147483647
		s2i64 = l13
		s3i64 = -2147483647
		if s2i64 > s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		l13 = s0i64
		s0i32 = int32(s0i64)
		l2 = s0i32
		s1i32 = 2147483647
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i64 = l13
		s1i64 = l12
		s2i64 = l15
		s1i64 = s1i64 + s2i64
		l12 = s1i64
		s2i64 = 2147483647
		s3i64 = l12
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l12 = s1i64
		s2i64 = -2147483647
		s3i64 = l12
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l12 = s1i64
		s0i64 = s0i64 - s1i64
		l13 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i64 = l11
		s1i64 = l10
		s2i64 = l14
		s1i64 = s1i64 + s2i64
		l10 = s1i64
		s2i64 = 2147483647
		s3i64 = l10
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l10 = s1i64
		s2i64 = -2147483647
		s3i64 = l10
		s4i64 = -2147483647
		if s3i64 > s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		l10 = s1i64
		s0i64 = s0i64 - s1i64
		l11 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i64 = l11
		s1i64 = l13
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967296
		if uint64(s0i64) < uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	lbl6:
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
			goto lbl7
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
			goto lbl7
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl7:
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl1
	lbl5:
		s0i64 = l10
		s0i32 = int32(s0i64)
		l5 = s0i32
		s0i64 = l12
		s0i32 = int32(s0i64)
		l6 = s0i32
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
			goto lbl8
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
			goto lbl8
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	lbl8:
		s0i32 = l3
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l3
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s1i32 = f658(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		goto lbl9
	}
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = 0
	l2 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l8 = s0i32
	s1i32 = 7
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l8
	s1i32 = 536870908
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	s0i32 = f44(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l1
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	l2 = s0i32
lbl11:
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l10 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = l10
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = l11
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
	s0i32 = l3
	s1i32 = l2
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
		goto lbl12
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l1
	f13(ctx, s0i32)
lbl12:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l1 = s0i32
lbl9:
	s0i32 = l3
	s1i32 = l7
	s1i64 = int64(s1i32)
	l10 = s1i64
	s2i32 = l3
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
	s1i64 = s1i64 + s2i64
	l11 = s1i64
	s2i64 = 2147483647
	s3i64 = l11
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l11 = s1i64
	s2i64 = -2147483647
	s3i64 = l11
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i64)
	s0i32 = l3
	s1i32 = l6
	s1i64 = int64(s1i32)
	l11 = s1i64
	s2i32 = l3
	s2i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)])))
	s1i64 = s1i64 + s2i64
	l12 = s1i64
	s2i64 = 2147483647
	s3i64 = l12
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l12 = s1i64
	s2i64 = -2147483647
	s3i64 = l12
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i64)
	s0i32 = l3
	s1i32 = l3
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])))
	s2i64 = l10
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i64)
	s0i32 = l3
	s1i32 = l3
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])))
	s2i64 = l11
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l0 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	s2i32 = 2147483647
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl13
	}
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l3 = s0i32
lbl14:
	s0i32 = l3
	s1i32 = l0
	s2i32 = l6
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l3
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	s1i32 = 2147483647
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		l2 = s0i32
		goto lbl15
	}
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l5
	l2 = s0i32
lbl17:
	s0i32 = l3
	s1i32 = l1
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l7
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l1
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
lbl15:
	s0i32 = l3
	s1i32 = 2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s1i32 = 2147483647
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 + s1i32
lbl13:
	s1i32 = 2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl1:
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	return
lbl0:
	s0i32 = l4
	s1i32 = 89
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4814
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 4770
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 4744
	s1i32 = l4
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
