package internal

import (
	"unsafe"
)

func f1229(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l6
	s1i32 = l4
	ctx.Mem[int(s0i32+15)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l8 = s1i32
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+52)])
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l4 = s0i32
		s1i32 = l1
		s2i32 = l2
		s3i32 = l6
		s4i32 = 8
		s3i32 = s3i32 + s4i32
		s4i32 = 1
		s5i32 = l4
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
		s0i32 = int32(ctx.Mem[int(s0i32+72)])
		if s0i32 != 0 {
			s0i32 = l6
			s0i32 = int32(ctx.Mem[int(s0i32+15)])
			l0 = s0i32
			s1i32 = 255
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
			lbl4:
				s0i32 = l5
				s1i32 = l6
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l5
				s1i32 = l8
				s0i32 = s0i32 + s1i32
				l5 = s0i32
				s0i32 = l3
				s1i32 = 1
				if s0i32 > s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				l0 = s0i32
				s0i32 = l3
				s1i32 = -1
				s0i32 = s0i32 + s1i32
				l3 = s0i32
				s0i32 = l0
				if s0i32 != 0 {
					goto lbl4
				}
				goto lbl0
				panic("unreachable executed")
				panic("unreachable executed")
			}
			s0i32 = l0
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l0 = s0i32
		lbl5:
			s0i32 = l5
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			l1 = s1i32
			s2i32 = 24
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = l5
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			l2 = s2i32
			s3i32 = 24
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			l4 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = l0
			s1i32 = s1i32 * s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = l4
			s1i32 = s1i32 + s2i32
			s2i32 = 24
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = l1
			s3i32 = 255
			s2i32 = s2i32 & s3i32
			s3i32 = l2
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			l4 = s3i32
			s2i32 = s2i32 - s3i32
			s3i32 = l0
			s2i32 = s2i32 * s3i32
			s3i32 = 8
			s2i32 = s2i32 >> (uint32(s3i32) & 31)
			s3i32 = l4
			s2i32 = s2i32 + s3i32
			s1i32 = s1i32 | s2i32
			s2i32 = l1
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = 255
			s2i32 = s2i32 & s3i32
			s3i32 = l2
			s4i32 = 8
			s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			l4 = s3i32
			s2i32 = s2i32 - s3i32
			s3i32 = l0
			s2i32 = s2i32 * s3i32
			s3i32 = l4
			s4i32 = 8
			s3i32 = s3i32 << (uint32(s4i32) & 31)
			s2i32 = s2i32 + s3i32
			s3i32 = -256
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 | s2i32
			s2i32 = l1
			s3i32 = 16
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = 255
			s2i32 = s2i32 & s3i32
			s3i32 = l2
			s4i32 = 16
			s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			l1 = s3i32
			s2i32 = s2i32 - s3i32
			s3i32 = l0
			s2i32 = s2i32 * s3i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = l1
			s2i32 = s2i32 + s3i32
			s3i32 = 16
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l8
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l3
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l1 = s0i32
			s0i32 = l3
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l1
			if s0i32 != 0 {
				goto lbl5
			}
			goto lbl0
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l1 = s0i32
		if s0i32 != 0 {
		lbl7:
			s0i32 = l1
			s1i32 = l5
			s2i32 = l6
			s3i32 = 8
			s2i32 = s2i32 + s3i32
			s3i32 = 1
			s4i32 = l6
			s5i32 = 15
			s4i32 = s4i32 + s5i32
			s5i32 = l1
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
			s0i32 = l5
			s1i32 = l8
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l3
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l0 = s0i32
			s0i32 = l3
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l0
			if s0i32 != 0 {
				goto lbl7
			}
			goto lbl0
			panic("unreachable executed")
			panic("unreachable executed")
		}
		s0i32 = l5
		s1i32 = l6
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s3i32 = l6
		s3i32 = int32(ctx.Mem[int(s3i32+15)])
		l1 = s3i32
		s4i32 = l0
		s5i32 = 64
		s6i32 = 68
		s7i32 = l1
		s8i32 = 255
		if s7i32 == s8i32 {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i32 = s5i32
		} else {
			s5i32 = s6i32
		}
		s4i32 = s4i32 + s5i32
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		l0 = s4i32
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
		s0i32 = l3
		s1i32 = 2
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	lbl8:
		s0i32 = l5
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l6
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s3i32 = l6
		s3i32 = int32(ctx.Mem[int(s3i32+15)])
		s4i32 = l0
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
		s0i32 = l3
		s1i32 = 2
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l1 = s0i32
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l1
		if s0i32 != 0 {
			goto lbl8
		}
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+72)])
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 255
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl11:
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
			l4 = s0i32
			s1i32 = l1
			s2i32 = l2
			s3i32 = l6
			s4i32 = 4
			s3i32 = s3i32 + s4i32
			s4i32 = 1
			s5i32 = l4
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
			s0i32 = l5
			s1i32 = l6
			s1i32 = int32(ctx.Mem[int(s1i32+15)])
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			l4 = s1i32
			s2i32 = l6
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l9 = s2i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = 255
			s2i32 = s2i32 & s3i32
			s3i32 = l5
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
			l10 = s3i32
			s4i32 = 8
			s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			l7 = s3i32
			s2i32 = s2i32 - s3i32
			s1i32 = s1i32 * s2i32
			s2i32 = l7
			s3i32 = 8
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = -256
			s1i32 = s1i32 & s2i32
			s2i32 = l4
			s3i32 = l9
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			s4i32 = l10
			s5i32 = 255
			s4i32 = s4i32 & s5i32
			l7 = s4i32
			s3i32 = s3i32 - s4i32
			s2i32 = s2i32 * s3i32
			s3i32 = 8
			s2i32 = s2i32 >> (uint32(s3i32) & 31)
			s3i32 = l7
			s2i32 = s2i32 + s3i32
			s1i32 = s1i32 | s2i32
			s2i32 = l4
			s3i32 = l9
			s4i32 = 24
			s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
			s4i32 = l10
			s5i32 = 24
			s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
			l7 = s4i32
			s3i32 = s3i32 - s4i32
			s2i32 = s2i32 * s3i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = l7
			s2i32 = s2i32 + s3i32
			s3i32 = 24
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 | s2i32
			s2i32 = l4
			s3i32 = l9
			s4i32 = 16
			s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
			s4i32 = 255
			s3i32 = s3i32 & s4i32
			s4i32 = l10
			s5i32 = 16
			s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
			s5i32 = 255
			s4i32 = s4i32 & s5i32
			l4 = s4i32
			s3i32 = s3i32 - s4i32
			s2i32 = s2i32 * s3i32
			s3i32 = 8
			s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
			s3i32 = l4
			s2i32 = s2i32 + s3i32
			s3i32 = 16
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l8
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l2
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l3
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l4 = s0i32
			s0i32 = l3
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l4
			if s0i32 != 0 {
				goto lbl11
			}
			goto lbl0
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl12:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l4 = s0i32
		s1i32 = l1
		s2i32 = l2
		s3i32 = l5
		s4i32 = 1
		s5i32 = l4
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
		s0i32 = l5
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l4 = s0i32
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l4
		if s0i32 != 0 {
			goto lbl12
		}
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l9 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l10 = s0i32
	if s0i32 != 0 {
	lbl14:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l4 = s0i32
		s1i32 = l1
		s2i32 = l2
		s3i32 = l9
		s4i32 = 1
		s5i32 = l4
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
		s0i32 = l10
		s1i32 = l5
		s2i32 = l9
		s3i32 = 1
		s4i32 = l6
		s5i32 = 15
		s4i32 = s4i32 + s5i32
		s5i32 = l10
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
		s0i32 = l5
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l4 = s0i32
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l4
		if s0i32 != 0 {
			goto lbl14
		}
		goto lbl0
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l0
	s1i32 = 64
	s2i32 = 68
	s3i32 = l4
	s4i32 = 255
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l10 = s0i32
lbl15:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l7 = s0i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l9
	s4i32 = 1
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
	s0i32 = l5
	s1i32 = l9
	s2i32 = 1
	s3i32 = l4
	s4i32 = l10
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
	s0i32 = l5
	s1i32 = l8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl15
	}
lbl0:
	s0i32 = l6
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
