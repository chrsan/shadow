package internal

import (
	"unsafe"
)

func f1179(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l1
	s4i32 = 1
	s3i32 = s3i32 + s4i32
	s4i32 = l2
	s5i32 = l3
	s4i32 = s4i32 + s5i32
	s0i32 = f244(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = 0
		l10 = s0i32
		s0i32 = 0
		l5 = s0i32
		s0i32 = l2
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l11 = s1i32
		s0i32 = s0i32 - s1i32
		l12 = s0i32
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l11
		s1i32 = s1i32 - s2i32
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l13 = s0i32
			s1i32 = 12
			s0i32 = s0i32 + s1i32
			l8 = s0i32
			l5 = s0i32
		lbl4:
			s0i32 = l5
			l6 = s0i32
			s1i32 = 8
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l9 = s0i32
			s1i32 = l12
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			s0i32 = l9
			s1i32 = l11
			s0i32 = s0i32 + s1i32
			l10 = s0i32
			s0i32 = l6
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			s1i32 = l8
			s2i32 = l13
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s0i32 = s0i32 + s1i32
			l5 = s0i32
		}
		s0i32 = l10
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l6
		s3i32 = l3
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
		l9 = s0i32
		s0i32 = l1
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = l5
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l6 = s1i32
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl6:
			s0i32 = l8
			s1i32 = l6
			s0i32 = s0i32 - s1i32
			l8 = s0i32
			s0i32 = l5
			s0i32 = int32(ctx.Mem[int(s0i32+2)])
			l6 = s0i32
			s0i32 = l5
			s1i32 = 2
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l8
			s1i32 = l6
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl6
			}
		}
		s0i32 = l3
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s0i32 = l5
		s0i32 = int32(ctx.Mem[int(s0i32+1)])
		s1i32 = l4
		s0i32 = s0i32 * s1i32
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l5 = s0i32
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l6 = s0i32
			s1i32 = l1
			s2i32 = l2
			s3i32 = l9
			s4i32 = l5
			s5i32 = l6
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
		}
		s0i32 = l3
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l10
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l7 = s0i32
		goto lbl2
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l0 = s0i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
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
lbl0:
}
