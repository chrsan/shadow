package internal

import (
	"unsafe"
)

func f1830(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 272
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+268)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	l7 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s0i32 = f439(ctx, s0i32, s1i32)
	l6 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	lbl4:
		s0i32 = l6
		s0i32 = int32(ctx.Mem[int(s0i32+36)])
		l2 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l3 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			if s0i32 != 0 {
				goto lbl8
			}
		lbl10:
			s0i32 = l6
			s1i32 = l4
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			l2 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l4
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			l4 = s0i32
			s1i32 = 5
			if uint32(s0i32) > uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl7
			}
			s0i32 = l4
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			switch s0i32 {
			case 0:
				goto lbl7
			case 1:
				goto lbl7
			case 2:
				goto lbl7
			case 3:
				goto lbl7
			case 4:
				goto lbl11
			default:
				goto lbl12
			}
		lbl12:
			s0i32 = l6
			s1i32 = l6
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			l4 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
			s0i32 = l6
			s1i32 = l4
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		lbl11:
			s0i32 = l2
			l4 = s0i32
			s1i32 = l3
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl10
			}
			goto lbl2
		}
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		goto lbl6
	lbl8:
		s0i32 = l6
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l4 = s0i32
		s1i32 = 5
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl7
		case 1:
			goto lbl7
		case 2:
			goto lbl7
		case 3:
			goto lbl7
		case 4:
			goto lbl13
		default:
			goto lbl14
		}
	lbl14:
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l2 = s1i32
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		goto lbl6
	lbl13:
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		goto lbl6
	lbl7:
		s0i32 = l6
		s1i32 = 1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l2 = s1i32
		s2i32 = l4
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = 2040
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l4
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		l3 = s2i32
		s3i32 = l4
		s2i32 = s2i32 & s3i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 3
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl17
		case 1:
			goto lbl16
		case 2:
			goto lbl15
		default:
			goto lbl5
		}
	lbl17:
		s0i32 = 0
		l4 = s0i32
		s0i32 = l2
		s1i32 = l5
		s0i32 = f448(ctx, s0i32, s1i32)
		l2 = s0i32
		s1i32 = 0
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	lbl18:
		s0i32 = l0
		s1i32 = l5
		s2i32 = l4
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
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
		s0i32 = l2
		s1i32 = l4
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l3 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l3
		if s0i32 != 0 {
			goto lbl18
		}
		goto lbl4
	lbl16:
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 1
		s2i32 = l5
		s2i32 = f233(ctx, s2i32)
		l2 = s2i32
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+268)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
		l4 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			f13(ctx, s0i32)
		}
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 1
		s0i32 = s0i32 | s1i32
		l3 = s0i32
		s1i32 = 18
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = l3
			s2i32 = 8
			s1i32 = f50(ctx, s1i32, s2i32)
			l3 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint32(s1i32)
			goto lbl20
		}
		s0i32 = l5
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint32(s1i32)
		s0i32 = l7
		l3 = s0i32
	lbl20:
		s0i32 = l5
		s1i32 = l5
		s2i32 = l3
		s3i32 = l2
		s1i32 = f232(ctx, s1i32, s2i32, s3i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+268)])) = uint32(s1i32)
		s0i32 = 0
		l2 = s0i32
		s0i32 = l4
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	lbl22:
		s0i32 = 0
		l4 = s0i32
		s0i32 = l3
		s1i32 = l5
		s0i32 = f448(ctx, s0i32, s1i32)
		l8 = s0i32
		s1i32 = 0
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl24:
			s0i32 = l0
			s1i32 = l5
			s2i32 = l4
			s3i32 = 4
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
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
			s0i32 = l4
			s1i32 = l8
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l9 = s0i32
			s0i32 = l4
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l4 = s0i32
			s0i32 = l9
			if s0i32 != 0 {
				goto lbl24
			}
		}
		s0i32 = l3
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+268)]))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl22
		}
		goto lbl4
	lbl15:
		s0i32 = 0
		l4 = s0i32
		s0i32 = l2
		s1i32 = l5
		s0i32 = f691(ctx, s0i32, s1i32)
		l2 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	lbl25:
		s0i32 = l0
		s1i32 = l5
		s2i32 = l4
		s3i32 = 24
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
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
		s0i32 = l2
		s1i32 = l4
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l3 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl25
		}
		goto lbl4
	lbl6:
		s0i32 = l6
		s1i32 = 256
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l1
		l2 = s0i32
	lbl5:
		s0i32 = l0
		s1i32 = l2
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
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
		goto lbl4
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0i32 = l5
	s1i32 = l0
	s2i32 = l2
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s1i32 = 1
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l5
	s2i32 = l3
	s3i32 = 870
	s4i32 = l5
	s5i32 = 80
	s4i32 = s4i32 + s5i32
	f685(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l5
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	l1 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	goto lbl1
lbl2:
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
lbl1:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l4 = s0i32
lbl0:
	s0i32 = l7
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	l0 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		f13(ctx, s0i32)
	}
	s0i32 = l5
	s1i32 = 272
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
