package internal

import (
	"unsafe"
)

func f1802(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var s5i32 int32
	_ = s5i32
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l6 = s0i32
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l9 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l4 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		l7 = s1i32
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
			l1 = s0i32
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
			l8 = s0i32
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l1
				s1i32 = l1
				s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
				s2i32 = 1
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				s0i32 = s0i32 + s1i32
				s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl3
				}
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l7 = s0i32
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			s2i32 = l4
			s3i32 = l8
			s4i32 = l1
			s5i32 = l7
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
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
			s1i32 = i32RemS(s1i32, s2i32)
			l4 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
			s2i32 = l4
			s3i32 = l0
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
			l1 = s3i32
			s4i32 = l1
			s5i32 = 2
			s4i32 = s4i32 + s5i32
			s5i32 = 2
			s4i32 = i32DivS(s4i32, s5i32)
			s3i32 = s3i32 + s4i32
			s4i32 = 1
			s3i32 = s3i32 + s4i32
			s2i32 = s2i32 * s3i32
			s3i32 = 1
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l4 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l4
			s2i32 = l1
			s3i32 = 1
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = 2
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
			l4 = s0i32
			s1i32 = l1
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			s0i32 = l4
			s1i32 = l1
			s2i32 = 1
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s1i32 = 0
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
			s1i32 = 0
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
			l7 = s0i32
		lbl3:
			s0i32 = l0
			s1i32 = l7
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		}
		s0i32 = l0
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	}
	s0i32 = 0
	s1i32 = l5
	s2i32 = l6
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = 31
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l5
	s0i32 = s0i32 & s1i32
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	l3 = s0i32
	s0i32 = l5
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l4
	s2i32 = 2
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l12 = s1i32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l10 = s0i32
		s0i32 = 0
		goto lbl5
	}
	s0i32 = 4
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	s1i32 = 0
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l10 = s0i32
	s0i32 = l3
	s1i32 = 4
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l13 = s0i32
	s0i32 = l5
	s1i32 = l1
	s2i32 = 0
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 - s1i32
lbl5:
	l9 = s0i32
	s0i32 = l12
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
	l1 = s1i32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l14 = s0i32
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	l3 = s0i32
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l10
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		l1 = s0i32
		goto lbl7
	}
	s0i32 = l3
	s1i32 = l12
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l7
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l8
	l3 = s0i32
	s0i32 = l5
	l6 = s0i32
	s0i32 = l7
	l4 = s0i32
lbl10:
	s0i32 = l4
	s1i32 = l3
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l1 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l3
		s1i32 = l4
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l3
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l4
		s1i32 = s1i32 - s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		goto lbl9
	}
	s0i32 = l1
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l3
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l4
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
lbl9:
	s0i32 = l10
	s1i32 = 4
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l10 = s0i32
	s0i32 = l12
	s1i32 = l14
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l11
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l1 = s0i32
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l4 = s0i32
		s0i32 = l6
		l3 = s0i32
		goto lbl13
	}
	s0i32 = l6
	l3 = s0i32
	s0i32 = 1
	l4 = s0i32
lbl15:
	s0i32 = l4
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l1
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l4
	s1i32 = l11
	s2i32 = l1
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l11 = s1i32
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l1 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
lbl13:
	s0i32 = l3
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l11
	s1i32 = l4
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l11
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l4
	s1i32 = s1i32 - s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
lbl12:
	s0i32 = l6
	s1i32 = l10
	s2i32 = l6
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = l1
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l8
	s1i32 = l3
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = 0
	l7 = s0i32
lbl7:
	s0i32 = l9
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		l4 = s0i32
		goto lbl16
	}
	s0i32 = l2
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l3 = s0i32
	s0i32 = l8
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l7
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l1
	l6 = s0i32
	s0i32 = l7
	l5 = s0i32
lbl19:
	s0i32 = l5
	s1i32 = l8
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l2 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l8
		s1i32 = l5
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l8
		s1i32 = l5
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = l5
		s1i32 = s1i32 - s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		goto lbl18
	}
	s0i32 = l2
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l8
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl19
	}
lbl18:
	s0i32 = 64
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	l8 = s0i32
	s0i32 = l1
	s1i32 = l7
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l9
	s1i32 = l4
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l3 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		l7 = s0i32
		s0i32 = l1
		l5 = s0i32
		s0i32 = l4
		l6 = s0i32
		goto lbl22
	}
	s0i32 = l4
	l6 = s0i32
	s0i32 = l1
	l5 = s0i32
	s0i32 = l9
	l7 = s0i32
lbl24:
	s0i32 = l7
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l7
	s1i32 = l6
	s2i32 = l3
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l3 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
lbl22:
	s0i32 = l5
	s1i32 = l7
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l6
	s1i32 = l7
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	s0i32 = l6
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l7
	s1i32 = s1i32 - s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
lbl21:
lbl25:
	s0i32 = l1
	s1i32 = l8
	s2i32 = l1
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s2i32 = l2
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = l4
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l2 = s1i32
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = 0
	l7 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	l5 = s0i32
	s0i32 = l9
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl25
	}
lbl16:
	s0i32 = l13
	if s0i32 != 0 {
		s0i32 = 1
		l8 = s0i32
		s0i32 = l4
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l7
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
		s0i32 = l1
		l5 = s0i32
		s0i32 = l7
		l9 = s0i32
	lbl28:
		s0i32 = l9
		s1i32 = l4
		s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l2 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = l9
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s1i32 = int32(ctx.Mem[int(s1i32+0)])
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l4
			s1i32 = l9
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			s0i32 = l4
			s1i32 = l9
			s2i32 = 1
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s1i32 = l2
			s2i32 = l9
			s1i32 = s1i32 - s2i32
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			goto lbl27
		}
		s0i32 = l2
		s1i32 = l5
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l4
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l9
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		l9 = s0i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl28
		}
	lbl27:
		s0i32 = l1
		s1i32 = l7
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l1 = s0i32
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			l4 = s0i32
			goto lbl31
		}
		s0i32 = l5
		l4 = s0i32
	lbl33:
		s0i32 = l8
		s1i32 = l1
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl30
		}
		s0i32 = l1
		s1i32 = l4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l8
		s1i32 = l6
		s2i32 = l1
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		l1 = s1i32
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl33
		}
	lbl31:
		s0i32 = l4
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l6
		s1i32 = l8
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l6
		s1i32 = l8
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l8
		s1i32 = s1i32 - s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	lbl30:
		s0i32 = l5
		s1i32 = l5
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l13
		s1i32 = s1i32 + s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
	s0i32 = l0
	s1i32 = l5
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+72)]))
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
}
