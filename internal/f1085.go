package internal

import (
	"unsafe"
)

func f1085(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	ctx.g0 = s0i32
	s0i32 = l9
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l2
	s2i32 = 31
	s1i32 = s1i32 + s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	if s0i32 != 0 {
	lbl3:
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l8
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l5
		f615(ctx, s0i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l8
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 8
		s0i32 = f50(ctx, s0i32, s1i32)
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0i32
		s0i32 = l0
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l5
		f13(ctx, s0i32)
	lbl2:
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l7 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl5:
		s0i32 = l5
		l6 = s0i32
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		s1i32 = -20
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l10 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l10
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
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
	lbl6:
		s0i32 = l5
		s1i32 = l7
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l10 = s0i32
	s0i32 = l9
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l2
	s2i32 = 31
	s1i32 = s1i32 + s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	l7 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		l12 = s0i32
	lbl9:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		s1i32 = 3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l7
			s2i32 = l10
			f1084(ctx, s0i32, s1i32, s2i32)
			goto lbl10
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l5 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l6 = s1i32
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l14 = s0i64
			goto lbl12
		}
		s0i32 = l0
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l14 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i64 = int64(uint32(s1i32))
		s2i32 = l6
		s3i32 = l5
		s2i32 = s2i32 - s3i32
		s3i32 = 48
		s2i32 = i32DivS(s2i32, s3i32)
		s2i64 = int64(uint32(s2i32))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 | s2i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l12
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+144)]))
		if int(s2i32) < 0 || int(s2i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s2i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s2i32].numParams != 2 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
	lbl12:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l5 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l6 = s0i32
		s0i32 = l9
		s1i64 = l14
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l9
		s1i64 = l14
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l5
		s2i32 = l5
		s3i32 = 36
		s2i32 = s2i32 + s3i32
		l5 = s2i32
		s3i32 = l5
		s4i32 = l6
		s5i32 = 1
		s4i32 = s4i32 << (uint32(s5i32) & 31)
		s5i32 = 3
		s4i32 = s4i32 + s5i32
		s5i32 = -4
		s4i32 = s4i32 & s5i32
		s3i32 = s3i32 + s4i32
		s4i32 = l6
		s5i32 = l9
		s6i32 = 8
		s5i32 = s5i32 + s6i32
		s6i32 = l1
		f1271(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = l8
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l0
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 8
		s0i32 = f50(ctx, s0i32, s1i32)
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0i32
		s0i32 = l0
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
		s0i32 = l5
		f13(ctx, s0i32)
	lbl14:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l5 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l11 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl16:
			s0i32 = l5
			l6 = s0i32
			s1i32 = -48
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l6
			s1i32 = -20
			s0i32 = s0i32 + s1i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l6 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl17
			}
			s0i32 = l6
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l13 = s1i32
			s2i32 = -1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l13
			s1i32 = 1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl17
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
		lbl17:
			s0i32 = l5
			s1i32 = l11
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl16
			}
		}
		s0i32 = l0
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	lbl10:
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l5 = s0i32
		s0i32 = l7
		f615(ctx, s0i32)
		s0i32 = l10
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		if s0i32 != 0 {
			goto lbl9
		}
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l5 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l6 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i64 = int64(uint32(s1i32))
		s2i32 = l6
		s3i32 = l5
		s2i32 = s2i32 - s3i32
		s3i32 = 48
		s2i32 = i32DivS(s2i32, s3i32)
		s2i64 = int64(uint32(s2i32))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 | s2i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l0
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+144)]))
		if int(s2i32) < 0 || int(s2i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s2i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s2i32].numParams != 2 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
	}
	s0i32 = l9
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
