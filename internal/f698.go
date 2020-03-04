package internal

import (
	"unsafe"
)

func f698(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var s5i64 int64
	_ = s5i64
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 691749
	s1i32 = l5
	s2i32 = 2097151
	s1i32 = s1i32 & s2i32
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	s0i32 = l0
	s0i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
	l11 = s0i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i64 = l10
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
	s0i32 = l4
	s1i64 = 4294967296
	s2i64 = l11
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s3i32 = l5
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l11 = s1i64
	s2i32 = l6
	s2i64 = int64(uint32(s2i32))
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+32)] = uint8(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	l5 = s0i32
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l6 = s0i32
		s0i32 = l5
		s1i32 = l4
		s2i32 = 104
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		s0i32 = f141(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			goto lbl3
		}
		goto lbl2
	}
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l4
		s2i32 = 104
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s0i32 = f1320(ctx, s0i32, s1i32, s2i32)
		l6 = s0i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 0
		l6 = s0i32
		goto lbl2
	}
	s0i64 = l11
	s1i64 = 4294967296
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l6 = s0i32
		s0i32 = l5
		s1i32 = l4
		s2i32 = 104
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+108)]))
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 3728
		s2i32 = s2i32 + s3i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i64 = int64(s2i32)
		s3i64 = l10
		s4i64 = 32
		s3i64 = s3i64 << (uint64(s4i64) & 63)
		s4i64 = 32
		s3i64 = s3i64 >> (uint64(s4i64) & 63)
		s2i64 = s2i64 * s3i64
		l10 = s2i64
		s2i32 = int32(s2i64)
		s3i32 = 0
		s4i64 = l10
		s5i64 = 2147483648
		if uint64(s4i64) < uint64(s5i64) {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s0i32 = f198(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			goto lbl3
		}
		goto lbl2
	}
	s0i32 = 0
	l6 = s0i32
	s0i32 = l5
	s1i32 = l4
	s2i32 = 104
	s1i32 = s1i32 + s2i32
	s0i32 = f460(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl3:
	s0i32 = l4
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+32)] = uint8(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	l3 = s0i32
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 8589934593
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = 0
		l0 = s0i32
		s0i32 = l4
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l4
		s0i32 = f460(ctx, s0i32, s1i32)
		l8 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l7
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l9 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l9
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
		s0i32 = l7
		f12(ctx, s0i32)
	lbl9:
		s0i32 = l8
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
	}
	s0i32 = 1248
	s0i32 = f17(ctx, s0i32)
	s1i32 = l5
	s2i32 = l1
	s3i32 = l6
	s4i32 = l3
	s5i32 = 0
	s6i32 = l2
	if s6i32 != 0 {
		// s4i32 = s4i32
	} else {
		s4i32 = s5i32
	}
	s0i32 = f450(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l0 = s0i32
lbl7:
	s0i32 = l3
	s0i32 = f41(ctx, s0i32)
	s0i32 = l0
	l6 = s0i32
lbl2:
	s0i32 = l5
	s0i32 = f41(ctx, s0i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
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
		goto lbl0
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl0:
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l6
	return s0i32
}
