package internal

import (
	"unsafe"
)

func f1831(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l10 int64
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l4 = s0i32
	s0i32 = l0
	s1i32 = l2
	if s1i32 != 0 {
		s1i32 = 0
		s2i32 = l4
		s2i64 = int64(uint32(s2i32))
		s3i64 = 3
		s2i64 = s2i64 * s3i64
		l10 = s2i64
		s3i64 = 32
		s2i64 = int64(uint64(s2i64) >> (uint64(s3i64) & 63))
		s2i32 = int32(s2i64)
		if s2i32 != 0 {
			goto lbl0
		}
		s1i64 = l10
		s1i32 = int32(s1i64)
		l4 = s1i32
	}
	s1i32 = l4
	s2i32 = l7
	s3i32 = 92
	s2i32 = s2i32 + s3i32
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	l8 = s0i32
	s0i32 = l4
	s1i32 = 1073741824
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+536)]))
		l6 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l5 = s1i32
		s2i32 = l4
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l4 = s2i32
		s1i32 = s1i32 | s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+540)]))
		s3i32 = l6
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l0
			s2i32 = 532
			s1i32 = s1i32 + s2i32
			s2i32 = l4
			s3i32 = 4
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = 0
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+536)]))
			l6 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
		} else {
			s1i32 = l5
		}
		s2i32 = l6
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = l5
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+536)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s0i32 = f439(ctx, s0i32, s1i32)
		l4 = s0i32
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			goto lbl4
		}
		s0i32 = l7
		s1i32 = 32
		s0i32 = s0i32 + s1i32
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
		s0i32 = l4
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		l9 = s0i32
	lbl8:
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0i32
	lbl11:
		s0i32 = l4
		s0i32 = int32(ctx.Mem[int(s0i32+36)])
		l1 = s0i32
		s0i32 = l6
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l2 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			if s0i32 != 0 {
				goto lbl13
			}
		lbl15:
			s0i32 = l4
			s1i32 = l2
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l2
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			l2 = s0i32
			s1i32 = 5
			if uint32(s0i32) > uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			s0i32 = l2
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			switch s0i32 {
			case 0:
				goto lbl12
			case 1:
				goto lbl12
			case 2:
				goto lbl12
			case 3:
				goto lbl12
			case 4:
				goto lbl16
			default:
				goto lbl17
			}
		lbl17:
			s0i32 = l4
			s1i32 = l4
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			l2 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l2
			s2i32 = 8
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		lbl16:
			s0i32 = l1
			l2 = s0i32
			s1i32 = l6
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl15
			}
			goto lbl3
		}
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		goto lbl10
	lbl13:
		s0i32 = l4
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l2 = s0i32
		s1i32 = 5
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl12
		}
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl12
		case 1:
			goto lbl12
		case 2:
			goto lbl12
		case 3:
			goto lbl12
		case 4:
			goto lbl18
		default:
			goto lbl19
		}
	lbl19:
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l1 = s1i32
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l1
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		goto lbl10
	lbl18:
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		goto lbl10
	lbl12:
		s0i32 = l4
		s1i32 = 1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l1 = s1i32
		s2i32 = l2
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = 2040
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l2
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		s3i32 = l2
		s2i32 = s2i32 & s3i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl11
		}
		s0i32 = l1
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		goto lbl9
	lbl10:
		s0i32 = l4
		s1i32 = 256
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
		s0i32 = l9
	lbl9:
		l1 = s0i32
		s0i32 = 0
		l2 = s0i32
		s0i32 = l1
		s1i32 = l7
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = l3
		s0i32 = f442(ctx, s0i32, s1i32, s2i32, s3i32)
		l1 = s0i32
		s1i32 = 0
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
	lbl20:
		s0i32 = l0
		s1i32 = l7
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = l8
		s3i32 = l5
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
		if int(s4i32) < 0 || int(s4i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s4i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s4i32].numParams != 4 {
			panic("argument count mismatch")
		}
		s0i32 = (*(*func(*Context, int32, int32, int32, int32) int32)(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		l6 = s0i32
		s1i32 = 2
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl21
		}
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl21
		case 1:
			goto lbl23
		default:
			goto lbl22
		}
	lbl23:
		s0i32 = l5
		s1i32 = -4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		goto lbl21
	lbl22:
		s0i32 = l5
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l8
		s1i32 = l7
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
		s0i32 = s0i32 + s1i32
		l8 = s0i32
	lbl21:
		s0i32 = l1
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		goto lbl8
		panic("unreachable executed")
		panic("unreachable executed")
	}
	f63(ctx)
	panic("unreachable executed")
lbl4:
lbl24:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+36)])
	l1 = s0i32
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l2 = s1i32
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l6 = s2i32
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		if s1i32 != 0 {
			goto lbl28
		}
	lbl30:
		s1i32 = l4
		s2i32 = l2
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		l1 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l2 = s1i32
		s2i32 = 5
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl27
		}
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		switch s1i32 {
		case 0:
			goto lbl27
		case 1:
			goto lbl27
		case 2:
			goto lbl27
		case 3:
			goto lbl27
		case 4:
			goto lbl31
		default:
			goto lbl32
		}
	lbl32:
		s1i32 = l4
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		l2 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = l2
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	lbl31:
		s1i32 = l1
		l2 = s1i32
		s2i32 = l6
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl30
		}
		goto lbl3
	}
	s1i32 = l1
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = -8
	s2i32 = s2i32 + s3i32
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint64(s2i64)
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = uint64(s2i64)
	goto lbl26
lbl28:
	s1i32 = l4
	s2i32 = l2
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l2
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l2 = s1i32
	s2i32 = 5
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl27
	}
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	switch s1i32 {
	case 0:
		goto lbl27
	case 1:
		goto lbl27
	case 2:
		goto lbl27
	case 3:
		goto lbl27
	case 4:
		goto lbl33
	default:
		goto lbl34
	}
lbl34:
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l1 = s2i32
	s3i32 = -8
	s2i32 = s2i32 + s3i32
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint64(s2i64)
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = uint64(s2i64)
	s1i32 = l4
	s2i32 = l1
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
	goto lbl26
lbl33:
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = -8
	s2i32 = s2i32 + s3i32
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = uint64(s2i64)
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = uint64(s2i64)
	goto lbl26
lbl27:
	s1i32 = l4
	s2i32 = 1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)])) = uint16(s2i32)
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l1 = s2i32
	s3i32 = l2
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = 2040
	s3i32 = s3i32 & s4i32
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s3i32 = l2
	s4i32 = -1
	s3i32 = s3i32 + s4i32
	s4i32 = l2
	s3i32 = s3i32 & s4i32
	s4i32 = 1
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = 1
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl24
	}
	s1i32 = l1
	s2i32 = -8
	s1i32 = s1i32 + s2i32
	goto lbl25
lbl26:
	s1i32 = l4
	s2i32 = 256
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)])) = uint16(s2i32)
	s1i32 = l3
lbl25:
	s2i32 = l8
	s3i32 = l5
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32, int32) int32)(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
	l1 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl24
	case 1:
		goto lbl36
	default:
		goto lbl35
	}
lbl36:
	s0i32 = l5
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	goto lbl24
lbl35:
	s0i32 = l5
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l8
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	goto lbl24
	panic("unreachable executed")
	panic("unreachable executed")
lbl3:
	s0i32 = l5
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
lbl0:
	l0 = s0i32
	s0i32 = l7
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
