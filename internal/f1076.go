package internal

import (
	"unsafe"
)

func f1076(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int64
	_ = l16
	var l17 int64
	_ = l17
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
	var s2i64 int64
	_ = s2i64
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 224
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l9 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
		s1i32 = l9
		s2i32 = 48
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l9 = s0i32
	lbl1:
		s0i32 = l10
		s1i32 = l8
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l4 = s1i32
		s2i32 = l4
		s3i32 = l10
		if uint32(s2i32) < uint32(s3i32) {
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
		s0i32 = l8
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s1i32 = l9
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l10
	f1082(ctx, s0i32, s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = 4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l9
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	s1i32 = 261120
	s0i32 = s0i32 & s1i32
	s1i32 = 3072
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	goto lbl2
lbl3:
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
lbl2:
	l12 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l4 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l10 = s0i32
		s1i32 = l4
		s2i32 = 48
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l13 = s0i32
		s0i32 = l1
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l16 = s0i64
		s0i32 = l0
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		l14 = s0i32
		s0i32 = l0
		s1i32 = 104
		s0i32 = s0i32 + s1i32
		l15 = s0i32
		s0i32 = l0
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	lbl5:
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l6 = s0i32
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l7 = s0i32
		s0i32 = l10
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+124)])
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
			f13(ctx, s0i32)
		}
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+100)])
		l4 = s0i32
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
			f13(ctx, s0i32)
			s0i32 = l0
			s0i32 = int32(ctx.Mem[int(s0i32+100)])
			l4 = s0i32
		}
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l4
		s2i32 = 252
		s1i32 = s1i32 & s2i32
		ctx.Mem[int(s0i32+100)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l15
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 4
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 17179869184
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l14
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l7
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l11
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+124)])
		s2i32 = 252
		s1i32 = s1i32 & s2i32
		ctx.Mem[int(s0i32+124)] = uint8(s1i32)
		s0i32 = l9
		s1i32 = l10
		s2i32 = 28
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s2i32 = l2
		s0i32 = f1712(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 88
			s0i32 = s0i32 + s1i32
			s1i32 = l6
			s2i32 = l9
			s3i32 = l12
			s4i32 = l0
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
			f1713(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			s0i32 = l5
			s1i32 = 80
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s2i32 = 88
			s1i32 = s1i32 + s2i32
			s2i32 = f259(ctx)
			f622(ctx, s0i32, s1i32, s2i32)
			s0i32 = l5
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l1
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l5
			s1i64 = l16
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
			s0i32 = l5
			s1i64 = l16
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
			s0i32 = l8
			s1i32 = l5
			s2i32 = 24
			s1i32 = s1i32 + s2i32
			s2i32 = l5
			s3i32 = 8
			s2i32 = s2i32 + s3i32
			f1081(ctx, s0i32, s1i32, s2i32)
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			l4 = s0i32
			s1i32 = l4
			s2i32 = 24
			s1i32 = s1i32 + s2i32
			s2i32 = l8
			s3i32 = l1
			s1i32 = f422(ctx, s1i32, s2i32, s3i32)
			f402(ctx, s0i32, s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = 0
				l7 = s0i32
				s0i32 = 0
				goto lbl9
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)]))
			l7 = s0i32
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		lbl9:
			l11 = s0i32
			s0i32 = l0
			s1i32 = l7
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l11
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
			l4 = s0i32
			s0i32 = l0
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
			l17 = s1i64
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i64)
			s0i32 = l0
			s1i64 = l17
			s2i64 = 32
			s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i64)
			s0i32 = l5
			s1i32 = 24
			s0i32 = s0i32 + s1i32
			s1i32 = l9
			s0i32 = f46(ctx, s0i32, s1i32)
			l4 = s0i32
			s0i32 = l5
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)]))
			s2i32 = -2
			s1i32 = s1i32 & s2i32
			s2i32 = l6
			s2i32 = int32(ctx.Mem[int(s2i32+17)])
			s3i32 = -1
			s2i32 = s2i32 + s3i32
			s3i32 = 2
			if uint32(s2i32) < uint32(s3i32) {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			s1i32 = s1i32 | s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = l8
			s2i32 = l5
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+220)]))
			s3i32 = l4
			s4i32 = l3
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
			if int(s4i32) < 0 || int(s4i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s4i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s4i32].numParams != 4 {
				panic("argument count mismatch")
			}
			(*(*func(*Context, int32, int32, float32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2f32, s3i32)
			s0i32 = l4
			s0i32 = f23(ctx, s0i32)
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl11
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl11
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl11:
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl12
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl12:
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl13
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl13
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl13:
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl14
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl14
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl14:
			s0i32 = l5
			s1i32 = 88
			s0i32 = s0i32 + s1i32
			f159(ctx, s0i32)
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 88
			s0i32 = s0i32 + s1i32
			s1i32 = l6
			s2i32 = l9
			s3i32 = l12
			s4i32 = l0
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
			s5i32 = l2
			f1715(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
			s0i32 = l5
			s1i32 = 80
			s0i32 = s0i32 + s1i32
			s1i32 = l5
			s2i32 = 88
			s1i32 = s1i32 + s2i32
			s2i32 = f259(ctx)
			f622(ctx, s0i32, s1i32, s2i32)
			s0i32 = l5
			s1i32 = l1
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
			s0i32 = l5
			s1i32 = l1
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l5
			s1i64 = l16
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l5
			s1i64 = l16
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
			s0i32 = l8
			s1i32 = l5
			s2i32 = 24
			s1i32 = s1i32 + s2i32
			s2i32 = l5
			s3i32 = l2
			s4i32 = l5
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+80)]))
			s5i32 = 212
			s4i32 = s4i32 + s5i32
			f1080(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			l4 = s0i32
			s1i32 = l4
			s2i32 = 24
			s1i32 = s1i32 + s2i32
			s2i32 = l8
			s1i32 = f1873(ctx, s1i32, s2i32)
			f402(ctx, s0i32, s1i32)
			s0i32 = l3
			s1i32 = l8
			s2i32 = l9
			s3i32 = l3
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
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl16
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl16
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl16:
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+216)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl17
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl17
			}
			s0i32 = l4
			s1i32 = l4
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
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl18
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl18:
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+208)]))
			l4 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl19
			}
			s0i32 = l4
			s1i32 = l4
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
				goto lbl19
			}
			s0i32 = l4
			s1i32 = l4
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
		lbl19:
			s0i32 = l5
			s1i32 = 88
			s0i32 = s0i32 + s1i32
			f159(ctx, s0i32)
		}
		s0i32 = l10
		s1i32 = 48
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s1i32 = l13
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	}
	s0i32 = l8
	f1079(ctx, s0i32)
	s0i32 = l0
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	f1083(ctx, s0i32)
	s0i32 = l5
	s1i32 = 224
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
