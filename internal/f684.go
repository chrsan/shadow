package internal

import (
	"unsafe"
)

func f684(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var l15 int64
	_ = l15
	var l16 int64
	_ = l16
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 752
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+700)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l5
	s2i32 = 560
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	l9 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+560)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 520
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s0i32 = f440(ctx, s0i32, s1i32)
	l6 = s0i32
	s0i32 = l5
	s1i32 = l2
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
	s0i32 = l6
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l13 = s0i32
	s0i32 = l5
	s1i32 = 444
	s0i32 = s0i32 + s1i32
	l11 = s0i32
	s0i32 = l5
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l12 = s0i32
lbl0:
	s0i32 = l6
	s0i32 = int32(ctx.Mem[int(s0i32+36)])
	l0 = s0i32
	s0i32 = l6
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl5
		}
	lbl7:
		s0i32 = l6
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l0 = s1i32
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
			goto lbl4
		}
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl4
		case 1:
			goto lbl4
		case 2:
			goto lbl4
		case 3:
			goto lbl4
		case 4:
			goto lbl8
		default:
			goto lbl9
		}
	lbl9:
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	lbl8:
		s0i32 = l0
		l2 = s0i32
		s1i32 = l7
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		goto lbl1
	}
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
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
	goto lbl3
lbl5:
	s0i32 = l6
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
		goto lbl4
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl4
	case 1:
		goto lbl4
	case 2:
		goto lbl4
	case 3:
		goto lbl4
	case 4:
		goto lbl10
	default:
		goto lbl11
	}
lbl11:
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l0 = s1i32
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
	s1i32 = l0
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	goto lbl3
lbl10:
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
	goto lbl3
lbl4:
	s0i32 = l6
	s1i32 = 1
	ctx.Mem[int(s0i32+36)] = uint8(s1i32)
	s0i32 = l6
	s0i32 = int32(ctx.Mem[int(s0i32+37)])
	l7 = s0i32
	s0i32 = l6
	s1i32 = 0
	ctx.Mem[int(s0i32+37)] = uint8(s1i32)
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l8 = s1i32
	s2i32 = l2
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
	s2i32 = l2
	s3i32 = -1
	s2i32 = s2i32 + s3i32
	l0 = s2i32
	s3i32 = l2
	s2i32 = s2i32 & s3i32
	s3i32 = 1
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l10 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 3
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l8
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl14
	case 1:
		goto lbl13
	case 2:
		goto lbl12
	default:
		goto lbl2
	}
lbl14:
	s0i32 = l5
	s1i32 = l2
	s2i32 = l1
	s0i32 = f685(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l7
	s2i32 = 0
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l4
	s3i32 = l3
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
	goto lbl0
lbl13:
	s0i32 = l10
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+736)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+728)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+720)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+744)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 1
	s2i32 = l5
	s3i32 = 720
	s2i32 = s2i32 + s3i32
	s2i32 = f233(ctx, s2i32)
	l0 = s2i32
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+700)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+560)]))
	l8 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 1
	s0i32 = s0i32 | s1i32
	l2 = s0i32
	s1i32 = 18
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l2
		s2i32 = 8
		s1i32 = f50(ctx, s1i32, s2i32)
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+560)])) = uint32(s1i32)
		goto lbl16
	}
	s0i32 = l5
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+560)])) = uint32(s1i32)
	s0i32 = l9
	l2 = s0i32
lbl16:
	s0i32 = l5
	s1i32 = l5
	s2i32 = 720
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = l0
	s1i32 = f232(ctx, s1i32, s2i32, s3i32)
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+700)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l0 = s0i32
lbl18:
	s0i32 = l5
	s1i32 = l2
	s2i32 = l1
	s0i32 = f685(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l7
		s2i32 = 0
		if s1i32 != s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = l4
		s3i32 = l3
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
	}
	s0i32 = l2
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+700)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	goto lbl0
lbl12:
	s0i32 = l5
	s1i32 = l2
	s2i32 = l1
	s0i32 = f2102(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l7
	s2i32 = 0
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l4
	s3i32 = l3
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
	goto lbl0
lbl3:
	s0i32 = l6
	s1i32 = 256
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint16(s1i32)
	s0i32 = 0
	l7 = s0i32
	s0i32 = l13
	l2 = s0i32
lbl2:
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0i64
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l16 = s0i64
	s0i32 = l5
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i64 = l16
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+712)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = l15
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+704)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 704
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l5
	s3i32 = 720
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s3i32 = int32(ctx.Mem[int(s3i32+8)])
	s0i32 = f443(ctx, s0i32, s1i32, s2i32, s3i32)
	l10 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l2 = s0i32
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
	lbl21:
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l14 = s1i32
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		s1i32 = 720
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l15 = s0i64
		s0i32 = l5
		s1i32 = 720
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l16 = s0i64
		s0i32 = l14
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l8
		s1i64 = l16
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		s1i64 = l15
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l5
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l8 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		l2 = s0i32
		s1i32 = l10
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl21
		}
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = 6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l12
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+444)]))
	s1i32 = 6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l7
	s2i32 = 0
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l4
	s3i32 = l3
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
	goto lbl0
lbl1:
	s0i32 = l9
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+560)]))
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
	s1i32 = 752
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
