package internal

import (
	"unsafe"
)

func f32(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
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
	var l23 int32
	_ = l23
	var l24 int32
	_ = l24
	var l25 int64
	_ = l25
	var l26 int64
	_ = l26
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
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l19 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 65535
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l22 = s0i32
	s1i32 = l2
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l16 = s1i32
	s0i32 = s0i32 - s1i32
	l15 = s0i32
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 255
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = l11
		s0i32 = s0i32 | s1i32
		l6 = s0i32
		s0i32 = l3
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		s1i32 = l4
		s0i32 = s0i32 - s1i32
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		s0i32 = i32DivS(s0i32, s1i32)
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		l2 = s0i32
		s0i32 = l10
		if s0i32 != 0 {
			s0i32 = l9
			s1i32 = l16
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l6
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l1
				s1i32 = l2
				ctx.Mem[int(s0i32+0)] = uint8(s1i32)
				goto lbl0
			}
			s0i32 = l1
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			s1i32 = l2
			s2i32 = 255
			s1i32 = s1i32 & s2i32
			s2i32 = l8
			s1i32 = s1i32 * s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s0i32 = s0i32 + s1i32
			l0 = s0i32
			s0i32 = l12
			if s0i32 != 0 {
				s0i32 = l1
				s1i32 = l0
				s2i32 = 255
				s3i32 = l0
				s4i32 = 255
				if uint32(s3i32) < uint32(s4i32) {
					s3i32 = 1
				} else {
					s3i32 = 0
				}
				if s3i32 != 0 {
					// s1i32 = s1i32
				} else {
					s1i32 = s2i32
				}
				ctx.Mem[int(s0i32+0)] = uint8(s1i32)
				goto lbl0
			}
			s0i32 = l1
			s1i32 = l0
			s2i32 = l0
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s1i32 = s1i32 - s2i32
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			goto lbl0
		}
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 0
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
			if int(s2i32) < 0 || int(s2i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s2i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s2i32].numParams != 2 {
				panic("argument count mismatch")
			}
			s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
			l0 = s0i32
			s1i32 = l16
			s2i32 = l1
			s3i32 = 1
			s4i32 = l2
			s5i32 = 255
			s4i32 = s4i32 & s5i32
			s5i32 = l0
			s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
			s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
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
		s0i32 = l0
		s1i32 = l16
		s2i32 = l1
		s3i32 = l2
		s4i32 = 255
		s3i32 = s3i32 & s4i32
		s4i32 = l8
		s3i32 = s3i32 * s4i32
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+64)]))
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
		goto lbl0
	}
	s0i32 = l15
	s1i32 = 32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l15
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s0i32 = f17(ctx, s0i32)
		l14 = s0i32
		s1i32 = l15
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l20 = s0i32
		s0i32 = l14
		s1i32 = l15
		s0i32 = s0i32 + s1i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l18 = s0i32
		goto lbl7
	}
	s0i32 = l15
	s1i32 = l19
	s0i32 = s0i32 + s1i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l18 = s0i32
	s0i32 = l15
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l19
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l20 = s0i32
	s0i32 = l19
	l14 = s0i32
	s0i32 = l15
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	goto lbl6
lbl7:
lbl9:
	s0i32 = 1
	l23 = s0i32
	s0i32 = l20
	s1i32 = l13
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l13
	s1i32 = l14
	s0i32 = s0i32 + s1i32
	s1i32 = l8
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l13
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s1i32 = l15
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl6:
	s0i32 = l20
	s1i32 = l15
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l4
	s1i32 = 65535
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l21 = s0i32
	s1i32 = l16
	s2i32 = 2
	s1i32 = s1i32 + s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l14
		s1i32 = l14
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l24 = s1i32
		s2i32 = l8
		s3i32 = l6
		s4i32 = 11
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		l21 = s3i32
		s4i32 = l2
		s5i32 = 65536
		s4i32 = s4i32 + s5i32
		s5i32 = -65536
		s4i32 = s4i32 & s5i32
		s5i32 = l2
		s4i32 = s4i32 - s5i32
		l17 = s4i32
		s5i32 = 11
		s4i32 = s4i32 >> (uint32(s5i32) & 31)
		l6 = s4i32
		s5i32 = l6
		s4i32 = s4i32 * s5i32
		s3i32 = s3i32 * s4i32
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 - s3i32
		l13 = s2i32
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		l6 = s2i32
		s3i32 = l24
		s4i32 = l6
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l13
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l14
		s1i32 = l14
		s1i32 = int32(ctx.Mem[int(s1i32+1)])
		l6 = s1i32
		s2i32 = l4
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		s3i32 = l17
		s2i32 = s2i32 - s3i32
		s3i32 = 11
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l2 = s2i32
		s3i32 = l2
		s2i32 = s2i32 * s3i32
		s3i32 = l21
		s2i32 = s2i32 * s3i32
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		l4 = s2i32
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		l2 = s2i32
		s3i32 = l6
		s4i32 = l2
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+1)] = uint8(s1i32)
		goto lbl10
	}
	s0i32 = l2
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	l17 = s0i32
	s0i32 = l18
	s1i32 = l4
	s2i32 = l2
	s3i32 = -65536
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 - s2i32
	l13 = s1i32
	s2i32 = 65535
	s1i32 = s1i32 + s2i32
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l4 = s1i32
	s2i32 = 1
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		if s1i32 != 0 {
			goto lbl12
		}
		goto lbl14
	}
	s1i32 = l18
	s2i32 = l4
	s3i32 = -1
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s1i32 = s1i32 + s2i32
	s2i32 = l13
	s3i32 = l2
	s4i32 = 16
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 - s3i32
	s2i64 = int64(s2i32)
	l25 = s2i64
	s3i32 = l6
	s3i64 = int64(s3i32)
	s2i64 = s2i64 * s3i64
	l26 = s2i64
	s3i64 = 16
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s3i64 = 32
	s2i64 = s2i64 >> (uint64(s3i64) & 63)
	s3i64 = l25
	s2i64 = s2i64 * s3i64
	s3i64 = 25
	s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
	ctx.Mem[int(s1i32+0)] = uint8(s2i64)
	s1i32 = l13
	s2i32 = 131073
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = -2
		s1i32 = s1i32 + s2i32
		l13 = s1i32
		s1i64 = l26
		s2i64 = 16
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		s1i32 = int32(s1i64)
		s2i32 = l6
		s3i32 = 1
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l2 = s1i32
	lbl17:
		s1i32 = l13
		s2i32 = l18
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		ctx.Mem[int(s1i32+0)] = uint8(s2i32)
		s1i32 = l2
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s1i32 = l13
		s2i32 = 1
		if s1i32 > s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		l4 = s1i32
		s1i32 = l13
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		l13 = s1i32
		s1i32 = l4
		if s1i32 != 0 {
			goto lbl17
		}
	}
	s1i32 = l8
	s2i32 = 65536
	s3i32 = l17
	s2i32 = s2i32 - s3i32
	s3i32 = 11
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	l2 = s2i32
	s3i32 = l2
	s2i32 = s2i32 * s3i32
	s3i32 = l6
	s4i32 = 11
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 * s3i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	goto lbl13
lbl14:
	s1i32 = l13
	s2i32 = l17
	s1i32 = s1i32 + s2i32
	s2i32 = 2
	s1i32 = i32DivS(s1i32, s2i32)
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = l8
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
lbl13:
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
lbl12:
	s0i32 = l16
	s1i32 = l21
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l16
	l13 = s0i32
lbl18:
	s0i32 = l14
	s1i32 = l13
	s2i32 = l16
	s1i32 = s1i32 - s2i32
	l4 = s1i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l2 = s1i32
	s2i32 = l4
	s3i32 = l18
	s2i32 = s2i32 + s3i32
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	l4 = s2i32
	s3i32 = l2
	s4i32 = l4
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l4
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l13
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s1i32 = l21
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
lbl10:
	s0i32 = l22
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l13 = s1i32
	s2i32 = 2
	s1i32 = s1i32 + s2i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l14
		s1i32 = l15
		s0i32 = s0i32 + s1i32
		l17 = s0i32
		s1i32 = -2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l13 = s1i32
		s2i32 = l7
		s3i32 = 11
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l7 = s2i32
		s3i32 = l3
		s4i32 = 65536
		s3i32 = s3i32 + s4i32
		s4i32 = -65536
		s3i32 = s3i32 & s4i32
		s4i32 = l3
		s3i32 = s3i32 - s4i32
		l6 = s3i32
		s4i32 = 11
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		l2 = s3i32
		s4i32 = l2
		s3i32 = s3i32 * s4i32
		s2i32 = s2i32 * s3i32
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		l4 = s2i32
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		l2 = s2i32
		s3i32 = l13
		s4i32 = l2
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l17
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l4 = s1i32
		s2i32 = l8
		s3i32 = l5
		s4i32 = l3
		s3i32 = s3i32 - s4i32
		s4i32 = l6
		s3i32 = s3i32 - s4i32
		s4i32 = 11
		s3i32 = s3i32 >> (uint32(s4i32) & 31)
		l2 = s3i32
		s4i32 = l2
		s3i32 = s3i32 * s4i32
		s4i32 = l7
		s3i32 = s3i32 * s4i32
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 - s3i32
		l3 = s2i32
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		l2 = s2i32
		s3i32 = l4
		s4i32 = l2
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l3
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		goto lbl19
	}
	s0i32 = l3
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s0i32 = l13
	s1i32 = l18
	s0i32 = s0i32 + s1i32
	s1i32 = l16
	s0i32 = s0i32 - s1i32
	l17 = s0i32
	s0i32 = l5
	s1i32 = l3
	s2i32 = -65536
	s1i32 = s1i32 & s2i32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 65535
	s0i32 = s0i32 + s1i32
	s1i32 = 16
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
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
			goto lbl21
		}
		goto lbl22
	}
	s0i32 = l17
	s1i32 = 65536
	s2i32 = l4
	s1i32 = s1i32 - s2i32
	s1i64 = int64(uint32(s1i32))
	l25 = s1i64
	s2i32 = l7
	s2i64 = int64(s2i32)
	s1i64 = s1i64 * s2i64
	l26 = s1i64
	s2i64 = 16
	s1i64 = s1i64 << (uint64(s2i64) & 63)
	s2i64 = 32
	s1i64 = s1i64 >> (uint64(s2i64) & 63)
	s2i64 = l25
	s1i64 = s1i64 * s2i64
	s2i64 = 25
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	ctx.Mem[int(s0i32+0)] = uint8(s1i64)
	s0i32 = l5
	s1i32 = l2
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s2i32 = 16
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s0i32 = l5
	s1i32 = 131073
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = l26
		s1i64 = 16
		s0i64 = int64(uint64(s0i64) >> (uint64(s1i64) & 63))
		s0i32 = int32(s0i64)
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = 1
		l6 = s0i32
	lbl25:
		s0i32 = l6
		s1i32 = l17
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l2
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = l4
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl25
		}
	}
	s0i32 = l4
	s1i32 = l17
	s0i32 = s0i32 + s1i32
	s1i32 = l8
	s2i32 = l3
	s3i32 = 11
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	l2 = s2i32
	s3i32 = l7
	s4i32 = 11
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 * s3i32
	s3i32 = l2
	s2i32 = s2i32 * s3i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	goto lbl21
lbl22:
	s0i32 = l17
	s1i32 = 0
	s2i32 = l4
	s3i32 = l5
	s2i32 = s2i32 + s3i32
	s1i32 = s1i32 - s2i32
	s2i32 = 9
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2i32 = l8
	s1i32 = s1i32 * s2i32
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
lbl21:
	s0i32 = l13
	s1i32 = l22
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl19
	}
lbl26:
	s0i32 = l14
	s1i32 = l13
	s2i32 = l16
	s1i32 = s1i32 - s2i32
	l3 = s1i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l2 = s1i32
	s2i32 = l3
	s3i32 = l18
	s2i32 = s2i32 + s3i32
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	l3 = s2i32
	s3i32 = l2
	s4i32 = l3
	if uint32(s3i32) > uint32(s4i32) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l3
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l13
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s1i32 = l22
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl26
	}
lbl19:
	s0i32 = l10
	if s0i32 != 0 {
		s0i32 = l23
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
		s0i32 = 0
		l13 = s0i32
		s0i32 = l12
		if s0i32 != 0 {
		lbl30:
			s0i32 = l9
			s1i32 = l13
			s2i32 = l16
			s1i32 = s1i32 + s2i32
			s0i32 = s0i32 + s1i32
			l0 = s0i32
			s1i32 = l0
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
			s2i32 = l13
			s3i32 = l14
			s2i32 = s2i32 + s3i32
			s2i32 = int32(ctx.Mem[int(s2i32+0)])
			s1i32 = s1i32 + s2i32
			l0 = s1i32
			s2i32 = 255
			s3i32 = l0
			s4i32 = 255
			if uint32(s3i32) < uint32(s4i32) {
				s3i32 = 1
			} else {
				s3i32 = 0
			}
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l13
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l13 = s0i32
			s1i32 = l15
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl30
			}
			goto lbl27
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl31:
		s0i32 = l9
		s1i32 = l13
		s2i32 = l16
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l13
		s3i32 = l14
		s2i32 = s2i32 + s3i32
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s2i32 = l0
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l13
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l13 = s0i32
		s1i32 = l15
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl31
		}
		goto lbl27
	}
	s0i32 = l8
	s1i32 = 255
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = l11
	if s0i32 != 0 {
		goto lbl32
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32) int32)(table[s2i32].f()))(ctx, s0i32, s1i32)
	l0 = s0i32
	s1i32 = l16
	s2i32 = l1
	s3i32 = l14
	s4i32 = l20
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
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
	goto lbl27
lbl32:
	s0i32 = l0
	s1i32 = l16
	s2i32 = l1
	s3i32 = l14
	s4i32 = l15
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+60)]))
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
lbl27:
	s0i32 = l15
	s1i32 = 32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l14
	f12(ctx, s0i32)
lbl0:
	s0i32 = l19
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
