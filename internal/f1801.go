package internal

import (
	"unsafe"
)

func f1801(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
lbl1:
	s0i32 = l2
	s1i32 = 3
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		s4i32 = l0
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
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l4
		s1i32 = 2
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		s0i32 = l4
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l4
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l13 = s0i32
	s1i32 = l2
	s2i32 = 2
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l12 = s1i32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = -1
		s2i32 = l1
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
		s2i32 = s2i32 - s3i32
		l5 = s2i32
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s3i32 = l5
		s2i32 = s2i32 & s3i32
		s3i32 = l3
		s2i32 = s2i32 + s3i32
		l3 = s2i32
		s3i32 = 0
		s4i32 = l5
		s5i32 = l5
		s6i32 = 0
		if s5i32 < s6i32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3i32 = s3i32
		} else {
			s3i32 = s4i32
		}
		l8 = s3i32
		s2i32 = s2i32 + s3i32
		l9 = s2i32
		s3i32 = 3
		s2i32 = s2i32 & s3i32
		l10 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l11 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l14 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		l6 = s1i32
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
			l5 = s0i32
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
			l15 = s0i32
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l5
				s1i32 = l5
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
					goto lbl5
				}
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l6 = s0i32
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			s2i32 = l14
			s3i32 = l15
			s4i32 = l5
			s5i32 = l6
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
			l6 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
			s2i32 = l6
			s3i32 = l0
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
			l5 = s3i32
			s4i32 = l5
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
			l6 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l6
			s2i32 = l5
			s3i32 = 1
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = 2
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
			l6 = s0i32
			s1i32 = l5
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
			s0i32 = l6
			s1i32 = l5
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
			l6 = s0i32
		lbl5:
			s0i32 = l0
			s1i32 = l6
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		}
		s0i32 = l10
		s1i32 = 4
		s2i32 = l10
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l5 = s0i32
		s0i32 = l8
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l10 = s0i32
		s0i32 = l7
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l6 = s0i32
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l8
		s1i32 = 2
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		l8 = s0i32
		s1i32 = -1
		s0i32 = s0i32 ^ s1i32
		s1i32 = l9
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l11
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l9 = s0i32
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			s2i32 = l8
			s1i32 = s1i32 + s2i32
			s2i32 = l12
			s3i32 = l7
			s4i32 = l5
			s5i32 = l10
			s4i32 = s4i32 - s5i32
			l5 = s4i32
			s5i32 = 6
			s4i32 = s4i32 << (uint32(s5i32) & 31)
			s5i32 = l5
			s6i32 = 2
			s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
			s4i32 = s4i32 - s5i32
			s5i32 = 255
			s4i32 = s4i32 & s5i32
			s5i32 = l9
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
			goto lbl7
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l11 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s2i32 = l8
		s1i32 = s1i32 + s2i32
		s2i32 = l12
		s3i32 = l9
		s4i32 = l7
		s5i32 = 4
		s6i32 = l10
		s5i32 = s5i32 - s6i32
		l7 = s5i32
		s6i32 = 6
		s5i32 = s5i32 << (uint32(s6i32) & 31)
		s6i32 = l7
		s7i32 = 2
		s6i32 = int32(uint32(s6i32) >> (uint32(s7i32) & 31))
		s5i32 = s5i32 - s6i32
		s6i32 = 255
		s5i32 = s5i32 & s6i32
		s6i32 = l5
		s7i32 = 6
		s6i32 = s6i32 << (uint32(s7i32) & 31)
		s7i32 = l5
		s8i32 = 2
		s7i32 = int32(uint32(s7i32) >> (uint32(s8i32) & 31))
		s6i32 = s6i32 - s7i32
		s7i32 = 255
		s6i32 = s6i32 & s7i32
		s7i32 = l11
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
		s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+24)]))
		if int(s7i32) < 0 || int(s7i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s7i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s7i32].numParams != 7 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32, int32))(table[s7i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	lbl7:
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l13
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
		l5 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l7 = s1i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l5
		s1i32 = l7
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		s1i32 = 0
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l4 = s0i32
	}
	s0i32 = l4
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl9:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l0
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
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l5 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l5
	if s0i32 != 0 {
		goto lbl9
	}
lbl0:
}
