package internal

import (
	"unsafe"
)

func f1201(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var l14 int64
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l4
	s1i64 = int64(s1i32)
	s2i32 = l2
	s2i64 = int64(s2i32)
	s1i64 = s1i64 + s2i64
	l14 = s1i64
	s2i64 = 2147483647
	s3i64 = l14
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
	l14 = s1i64
	s2i64 = -2147483647
	s3i64 = l14
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i64)
	s0i32 = l7
	s1i32 = l3
	s2i32 = 2
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1i64 = int64(s1i32)
	s2i32 = l1
	s2i64 = int64(s2i32)
	s1i64 = s1i64 + s2i64
	l14 = s1i64
	s2i64 = 2147483647
	s3i64 = l14
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
	l14 = s1i64
	s2i64 = -2147483647
	s3i64 = l14
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i64)
	s0i32 = l7
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l7
	s3i32 = -64
	s2i32 = s2i32 - s3i32
	s0i32 = f111(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l13 = s0i32
	lbl1:
		s0i32 = l6
		s1i32 = -1
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+52)]))
		l8 = s2i32
		s3i32 = l13
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
		l10 = s0i32
		s1i32 = l5
		s2i32 = -1
		s3i32 = l2
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
		l3 = s3i32
		s4i32 = l1
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
		l11 = s1i32
		s0i32 = s0i32 & s1i32
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		s1i32 = 255
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l4 = s0i32
			s1i32 = l3
			s2i32 = l2
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
			l9 = s2i32
			s3i32 = l8
			s4i32 = l3
			s3i32 = s3i32 - s4i32
			s4i32 = l2
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+56)]))
			s5i32 = l9
			s4i32 = s4i32 - s5i32
			s5i32 = l4
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
			goto lbl2
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l4 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l9 = s0i32
		s0i32 = l8
		s1i32 = l3
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
			s1i32 = l4
			s0i32 = s0i32 - s1i32
			l8 = s0i32
			s0i32 = l9
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l12 = s0i32
			s0i32 = l1
			s1i32 = l3
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l9
				s1i32 = l1
				s2i32 = l4
				s3i32 = l8
				s4i32 = l11
				s5i32 = 255
				s4i32 = s4i32 & s5i32
				s5i32 = l12
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
				goto lbl2
			}
			s0i32 = l9
			s1i32 = l3
			s2i32 = l4
			s3i32 = l8
			s4i32 = l10
			s5i32 = 255
			s4i32 = s4i32 & s5i32
			s5i32 = l12
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
			goto lbl2
		}
		s0i32 = l9
		s1i32 = l3
		s2i32 = l4
		s3i32 = l8
		s4i32 = -2
		s3i32 = s3i32 + s4i32
		s4i32 = l2
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+56)]))
		s5i32 = l4
		s4i32 = s4i32 - s5i32
		s5i32 = l11
		s6i32 = 255
		s5i32 = s5i32 & s6i32
		s6i32 = l10
		s7i32 = 255
		s6i32 = s6i32 & s7i32
		s7i32 = l9
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
	lbl2:
		s0i32 = l2
		f110(ctx, s0i32)
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l7
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
