package internal

import (
	"unsafe"
)

func f1767(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
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
	var s6i64 int64
	_ = s6i64
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = l5
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
	s1i32 = l6
	s2i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)])) = uint64(s2i64)
	s1i32 = l6
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
	s1i32 = l6
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
	s1i32 = l6
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
	s1i32 = l6
	s2i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = uint64(s2i64)
	s1i32 = l6
	l10 = s1i32
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l6 = s0i32
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l5
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l2
	s3i32 = 1
	s4i32 = 0
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl4:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l3 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l6 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
		s0i32 = l5
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		f101(ctx, s0i32, s1i32)
		s0i32 = l5
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		f101(ctx, s0i32, s1i32)
		s0i32 = l5
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
		s1i32 = l5
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])))
		s0i64 = s0i64 - s1i64
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
		s0i32 = int32(s0i64)
		l7 = s0i32
		s0i32 = l5
		s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])))
		s1i32 = l5
		s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])))
		s0i64 = s0i64 - s1i64
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
		s0i32 = int32(s0i64)
		l9 = s0i32
		s0i32 = l8
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 16
			s0i32 = s0i32 + s1i32
			s1i32 = l1
			f520(ctx, s0i32, s1i32)
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l8 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
			s0i32 = l8
			s1i32 = l8
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
			if int(s1i32) < 0 || int(s1i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s1i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s1i32].numParams != 1 {
				panic("argument count mismatch")
			}
			s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
			s0i32 = l3
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			s2i32 = l9
			s3i32 = l7
			s4i32 = l6
			s5i32 = 0
			s6i32 = l4
			s7i32 = l3
			s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
			s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+176)]))
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
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			l3 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l6 = s1i32
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = 1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			if int(s1i32) < 0 || int(s1i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s1i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s1i32].numParams != 1 {
				panic("argument count mismatch")
			}
			(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
			goto lbl5
		}
		s0i32 = l3
		s1i32 = l1
		s2i32 = l9
		s3i32 = l7
		s4i32 = l6
		s5i32 = l3
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+168)]))
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
	lbl5:
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl4
		}
		goto lbl2
		panic("unreachable executed")
		panic("unreachable executed")
	}
lbl7:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l9 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	f101(ctx, s0i32, s1i32)
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	f101(ctx, s0i32, s1i32)
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
	l11 = s0i64
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])))
	l12 = s0i64
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])))
	l13 = s0i64
	s0i32 = l5
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])))
	l14 = s0i64
	s0i32 = l5
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	f520(ctx, s0i32, s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l7
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	s0i32 = l6
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i64 = l12
	s3i64 = l11
	s2i64 = s2i64 - s3i64
	l11 = s2i64
	s3i64 = 2147483647
	s4i64 = l11
	s5i64 = 2147483647
	if s4i64 < s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l11 = s2i64
	s3i64 = -2147483647
	s4i64 = l11
	s5i64 = -2147483647
	if s4i64 > s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	s2i32 = int32(s2i64)
	s3i64 = l14
	s4i64 = l13
	s3i64 = s3i64 - s4i64
	l11 = s3i64
	s4i64 = 2147483647
	s5i64 = l11
	s6i64 = 2147483647
	if s5i64 < s6i64 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i64 = s3i64
	} else {
		s3i64 = s4i64
	}
	l11 = s3i64
	s4i64 = -2147483647
	s5i64 = l11
	s6i64 = -2147483647
	if s5i64 > s6i64 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i64 = s3i64
	} else {
		s3i64 = s4i64
	}
	s3i32 = int32(s3i64)
	s4i32 = l9
	s5i32 = l3
	s6i32 = l4
	s7i32 = l6
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+176)]))
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
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if int(s1i32) < 0 || int(s1i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s1i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s1i32].numParams != 1 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
lbl8:
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl7
	}
lbl2:
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = f23(ctx, s0i32)
		s0i32 = l2
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	}
	s0i32 = l10
	s0i32 = f23(ctx, s0i32)
	s0i32 = l5
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
