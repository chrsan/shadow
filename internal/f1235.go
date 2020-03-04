package internal

import (
	"unsafe"
)

func f1235(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+72)])
	l5 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
	l9 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	l10 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l4
		s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l5 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	lbl3:
		s0i32 = l5
		s1i32 = 65535
		s0i32 = s0i32 & s1i32
		l6 = s0i32
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l11 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l8
		s1i32 = l1
		s2i32 = l2
		s3i32 = l9
		s4i32 = l6
		s5i32 = l8
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
		s0i32 = l6
		l5 = s0i32
		s0i32 = l11
		s1i32 = 255
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l10
			s1i32 = l7
			s2i32 = l9
			s3i32 = l6
			s4i32 = 0
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
			goto lbl4
		}
	lbl6:
		s0i32 = l10
		s1i32 = l7
		s2i32 = l5
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		l11 = s2i32
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l12 = s2i32
		s1i32 = s1i32 + s2i32
		s2i32 = l9
		s3i32 = l12
		s2i32 = s2i32 + s3i32
		s3i32 = 1
		s4i32 = l3
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
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l12 = s0i32
		s0i32 = l11
		l5 = s0i32
		s0i32 = l12
		if s0i32 != 0 {
			goto lbl6
		}
	lbl4:
		s0i32 = l4
		s1i32 = l6
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l5 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l1
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l7
		s1i32 = l6
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l10 = s0i32
		goto lbl3
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l5
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l8
	s1i32 = l8
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
	s0i32 = (*(*func(*Context, int32) int32)(table[s1i32].f()))(ctx, s0i32)
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl7:
	s0i32 = l5
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
	s0i32 = l8
	s1i32 = l1
	s2i32 = l2
	s3i32 = l9
	s4i32 = l5
	s5i32 = l8
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
	s0i32 = l6
	s1i32 = 255
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = l9
		s2i32 = l5
		s3i32 = 255
		s4i32 = l0
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
		goto lbl8
	}
	s0i32 = l7
	s1i32 = l9
	s2i32 = l5
	s3i32 = l6
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+68)]))
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
lbl8:
	s0i32 = l1
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l7
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l4
	s1i32 = l5
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	goto lbl0
lbl1:
	s0i32 = l4
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl10:
	s0i32 = l5
	s1i32 = 65535
	s0i32 = s0i32 & s1i32
	l5 = s0i32
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l6
	s1i32 = 255
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l1
		s2i32 = l2
		s3i32 = l7
		s4i32 = l5
		s5i32 = l8
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
		goto lbl11
	}
	s0i32 = l8
	s1i32 = l1
	s2i32 = l2
	s3i32 = l9
	s4i32 = l5
	s5i32 = l8
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
	s0i32 = l7
	s1i32 = l9
	s2i32 = l5
	s3i32 = l6
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+68)]))
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
lbl11:
	s0i32 = l1
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l7
	s1i32 = l5
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s0i32 = l4
	s1i32 = l5
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl10
	}
lbl0:
}
