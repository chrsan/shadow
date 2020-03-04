package internal

import (
	"math/bits"
	"unsafe"
)

func f409(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var l17 int64
	_ = l17
	var l18 int64
	_ = l18
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 688
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	ctx.g0 = s0i32
	s0i32 = l8
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l18 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+680)])) = uint64(s1i64)
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l17 = s0i64
	s0i32 = l8
	s1i64 = l18
	s1i32 = int32(s1i64)
	s2i32 = l5
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+680)])) = uint32(s1i32)
	s0i32 = l8
	s1i64 = l17
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+672)])) = uint64(s1i64)
	s0i32 = l8
	s1i64 = l17
	s1i32 = int32(s1i64)
	s2i32 = l5
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+672)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l8
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+676)]))
	s2i32 = l5
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+676)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l8
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+684)]))
	s2i32 = l5
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+684)])) = uint32(s1i32)
	s0i32 = l8
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l8
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = 24540
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 636
	s0i32 = s0i32 + s1i32
	s1i32 = l8
	s2i32 = 124
	s1i32 = s1i32 + s2i32
	s2i32 = 512
	s3i32 = 512
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l15 = s0i32
	s0i32 = l8
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+668)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 24196
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 104
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = l8
	s4i32 = 672
	s3i32 = s3i32 + s4i32
	s4i32 = l6
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s0i32 = f317(ctx, s0i32, s1i32, s2i32)
	l10 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+10)])
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l6 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l0 = s1i32
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
			goto lbl0
		}
		s0i32 = l4
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		s2i32 = l7
		s3i32 = l4
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l4 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l3
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l1 = s2i32
		s3i32 = l1
		s4i32 = l3
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
		l1 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l18 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i64 = l17
		s1i64 = l18
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
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l0
		s2i32 = l5
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l1
		s3i32 = l5
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l6
		s4i32 = l0
		s3i32 = s3i32 - s4i32
		s4i32 = l5
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = l4
		s5i32 = l1
		s4i32 = s4i32 - s5i32
		s5i32 = l5
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = l2
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+20)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		goto lbl0
	}
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)]))
	l9 = s0i32
	s1i32 = l10
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = l9
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 64
		s1i32 = l7
		s2i32 = l9
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s1i32 = int32(bits.LeadingZeros32(uint32(s1i32)))
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 - s1i32
		s1i32 = l9
		s2i32 = l7
		f408(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l10
	s1i32 = 2
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l7 = s0i32
	lbl4:
		s0i32 = l9
		s1i32 = l7
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s1i32 = -4
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l11
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l11
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l12
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s1i32 = l10
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	}
	s0i32 = l10
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l9
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s0i32 = l8
	s1i32 = -2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l8
	s1i64 = -9223372032559808512
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
	s0i32 = l9
	s1i32 = l8
	s2i32 = 72
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = 0
	l11 = s0i32
	s0i32 = l8
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l8
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+676)]))
	l7 = s0i32
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+684)]))
	l9 = s0i32
	s0i32 = l8
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l8
	s1i32 = 24868
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l4 = s0i32
	s1i32 = l9
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l12 = s0i32
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l3 = s0i32
	s1i32 = l7
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l13 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+10)])
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l5
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l5
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = 940
		l11 = s0i32
		s0i32 = l8
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f177(ctx, s0i32)
		l5 = s0i32
	}
	s0i32 = l4
	s1i32 = l9
	s2i32 = l4
	s3i32 = l12
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l6
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l14 = s0i32
	s0i32 = l3
	s1i32 = l7
	s2i32 = l3
	s3i32 = l13
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l6
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l10
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l11
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l5
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	s1i32 = l2
	s2i32 = l1
	s3i32 = l14
	f624(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl7
lbl8:
	s0i32 = 1
	s1i32 = -1
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+10)])
	s3i32 = 1
	s2i32 = s2i32 & s3i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l12 = s0i32
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+680)]))
	l16 = s0i32
lbl9:
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l3 = s0i32
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l4 = s0i32
	s0i32 = l11
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l1
		s2i32 = 1
		s3i32 = l11
		if int(s3i32) < 0 || int(s3i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s3i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s3i32].numParams != 3 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = 0
	l0 = s0i32
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l5 = s0i32
	lbl13:
		s0i32 = l0
		s1i32 = l4
		l7 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = 32768
		s1i32 = s1i32 + s2i32
		s2i32 = 16
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l4 = s1i32
		s2i32 = l5
		s3i32 = l12
		s2i32 = s2i32 & s3i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l0 = s0i32
		s0i32 = l5
		s1i32 = l7
		s1i32 = int32(int8(ctx.Mem[int(s1i32+27)]))
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l12
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l4
		s1i32 = l0
		s0i32 = s0i32 - s1i32
		l4 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l2
		s1i32 = l0
		s2i32 = l1
		s3i32 = l4
		s4i32 = l2
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
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
	lbl14:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = l1
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			s0i32 = int32(int8(ctx.Mem[int(s0i32+24)]))
			l6 = s0i32
			s1i32 = 1
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l7
				s0i32 = f319(ctx, s0i32)
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl18
				}
				s0i32 = l7
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
				l10 = s0i32
				goto lbl16
			}
			s0i32 = l6
			s1i32 = -1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			s0i32 = l7
			s0i32 = f318(ctx, s0i32)
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l10 = s0i32
			goto lbl16
		lbl18:
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l6 = s0i32
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s1i32 = l6
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			goto lbl15
		}
		s0i32 = l7
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l7
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s1i32 = s1i32 + s2i32
		l10 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	lbl16:
		s0i32 = l10
		s1i32 = l3
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l10
			l3 = s0i32
			goto lbl15
		}
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0i32
		l6 = s0i32
	lbl21:
		s0i32 = l6
		l9 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0i32
		if s0i32 != 0 {
			s0i32 = l9
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			s1i32 = l10
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl21
			}
		}
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l7
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l13
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l13
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l9
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	lbl15:
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl13
		}
		s0i32 = l4
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		goto lbl11
	}
	s0i32 = 0
	l5 = s0i32
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
lbl11:
	l7 = s0i32
	s0i32 = l5
	s1i32 = l12
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l16
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l2
	s1i32 = l0
	s2i32 = l1
	s3i32 = l3
	s4i32 = l2
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
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
lbl23:
	s0i32 = l11
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l1
		s2i32 = 0
		s3i32 = l11
		if int(s3i32) < 0 || int(s3i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s3i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s3i32].numParams != 3 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l14
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l0 = s1i32
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl25:
	s0i32 = l6
	l7 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl25
		}
	}
lbl27:
	s0i32 = l4
	l0 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
lbl29:
	s0i32 = l7
	l3 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s1i32 = l0
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl28
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl29
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl28:
	s0i32 = l0
	l7 = s0i32
	s0i32 = l1
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	goto lbl9
	panic("unreachable executed")
	panic("unreachable executed")
lbl7:
	s0i32 = l8
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
lbl0:
	s0i32 = l8
	s1i32 = 24540
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
	s0i32 = l15
	f43(ctx, s0i32)
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
	f13(ctx, s0i32)
	s0i32 = l8
	s1i32 = 688
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
