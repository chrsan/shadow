package internal

import (
	"unsafe"
)

func f220(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	var l12 int64
	_ = l12
	var l13 int64
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l12 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i64 = int64(s0i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l13 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l12
	s1i64 = l13
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
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
			l8 = s0i32
			s1i32 = l7
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
			l9 = s0i32
			s1i32 = l6
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l10 = s0i32
			s1i32 = l5
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l11 = s0i32
			s1i32 = l4
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i32 = l9
			s0i64 = int64(s0i32)
			s1i32 = l11
			s1i64 = int64(s1i32)
			s0i64 = s0i64 - s1i64
			l12 = s0i64
			s1i64 = 1
			if s0i64 < s1i64 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i32 = l8
			s0i64 = int64(s0i32)
			s1i32 = l10
			s1i64 = int64(s1i32)
			s0i64 = s0i64 - s1i64
			l13 = s0i64
			s1i64 = 1
			if s0i64 < s1i64 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl3
			}
			s0i64 = l12
			s1i64 = l13
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
				goto lbl3
			}
			s0i32 = l2
			s1i32 = l4
			s2i32 = l5
			s3i32 = l6
			s4i32 = l4
			s3i32 = s3i32 - s4i32
			s4i32 = l7
			s5i32 = l5
			s4i32 = s4i32 - s5i32
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
		lbl3:
			s0i32 = l3
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = l3
			s2i32 = l1
			s0i32 = f25(ctx, s0i32, s1i32, s2i32)
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl0
			}
			s0i32 = l2
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l0 = s1i32
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l1 = s2i32
			s3i32 = l3
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
			s4i32 = l0
			s3i32 = s3i32 - s4i32
			s4i32 = l3
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
			s5i32 = l1
			s4i32 = s4i32 - s5i32
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
		s0i32 = l3
		s1i32 = l1
		s2i32 = l0
		s0i32 = f111(ctx, s0i32, s1i32, s2i32)
		l0 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 != 0 {
			goto lbl0
		}
	lbl4:
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l1 = s1i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
		l4 = s2i32
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+52)]))
		s4i32 = l1
		s3i32 = s3i32 - s4i32
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+56)]))
		s5i32 = l4
		s4i32 = s4i32 - s5i32
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
		s0i32 = l0
		f110(ctx, s0i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+60)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l4
	s2i32 = l5
	s3i32 = l6
	s4i32 = l4
	s3i32 = s3i32 - s4i32
	s4i32 = l7
	s5i32 = l5
	s4i32 = s4i32 - s5i32
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
lbl0:
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
