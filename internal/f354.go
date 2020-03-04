package internal

import (
	"unsafe"
)

func f354(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1i32
	s0i32 = s0i32 - s1i32
	s1i32 = 48
	s0i32 = i32DivS(s0i32, s1i32)
	l10 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 89478486
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l7
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l8
		s2i32 = s2i32 - s3i32
		s3i32 = 48
		s2i32 = i32DivS(s2i32, s3i32)
		l9 = s2i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l11 = s2i32
		s3i32 = l11
		s4i32 = l7
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 89478485
		s3i32 = l9
		s4i32 = 44739242
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l7 = s1i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl4
		}
		s0i32 = l7
		s1i32 = 89478486
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l7
		s1i32 = 48
		s0i32 = s0i32 * s1i32
		s0i32 = f17(ctx, s0i32)
	lbl4:
		l9 = s0i32
		s0i32 = l2
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0i64
		s0i32 = l3
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l12 = s0i64
		s0i32 = l4
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0i64
		s0i32 = l9
		s1i32 = l10
		s2i32 = 48
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l5
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = l14
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = l12
		s2i64 = 32
		s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i64)
		s0i32 = l2
		s1i64 = l13
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i64)
		s0i32 = l2
		s1i64 = l12
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i64)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l3
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l8 = s0i32
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l6 = s0i32
		}
		s0i32 = l7
		s1i32 = 48
		s0i32 = s0i32 * s1i32
		s1i32 = l9
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+11)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+39)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l6
		s1i32 = l8
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	lbl6:
		s0i32 = l2
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l6
		s2i32 = -48
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l3
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l4 = s0i32
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l4
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		}
		s0i32 = l1
		s1i32 = l4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = -9
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = -9
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = -16
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s2i32 = -16
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		l2 = s0i32
		s0i32 = l3
		l6 = s0i32
		s1i32 = l8
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		goto lbl0
	}
	f104(ctx)
	panic("unreachable executed")
lbl2:
	f63(ctx)
	panic("unreachable executed")
lbl1:
	s0i32 = l6
lbl0:
	l1 = s0i32
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl9:
		s0i32 = l6
		l0 = s0i32
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l0
		s1i32 = -20
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l0 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l2 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l0
		s1i32 = l0
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
	lbl10:
		s0i32 = l1
		s1i32 = l6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl9
		}
	}
	s0i32 = l1
	if s0i32 != 0 {
		s0i32 = l1
		f12(ctx, s0i32)
	}
}
