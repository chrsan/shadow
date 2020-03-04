package internal

import (
	"unsafe"
)

func f1176(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l12 int32
	_ = l12
	var l13 int32
	_ = l13
	var l14 int32
	_ = l14
	var l15 int32
	_ = l15
	var l16 int32
	_ = l16
	var l17 int32
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l2
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l2
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s0i32 = f244(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l0 = s0i32
		s1i32 = l1
		s2i32 = l2
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
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
	}
	s0i32 = l6
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = l1
		s2i32 = 12
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l6
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s0i32 = f118(ctx, s0i32)
		l4 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l3 = s0i32
		s0i32 = l4
		s1i32 = 1024
		s2i32 = l4
		s3i32 = 1024
		if uint32(s2i32) > uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l8 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			l5 = s0i32
			goto lbl3
		}
		s0i32 = l0
		s1i32 = 52
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			f13(ctx, s0i32)
		}
		s0i32 = l4
		s1i32 = 1025
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l8
			s1i32 = 2
			s0i32 = f44(ctx, s0i32, s1i32)
			l5 = s0i32
		}
		s0i32 = l0
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	lbl3:
		s0i32 = l6
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l6
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s1i32 = 7
		s0i32 = s0i32 & s1i32
		l10 = s0i32
		s0i32 = l3
		s1i32 = 3
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		l11 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l14 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0i32
		s0i32 = l6
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l12 = s0i32
		s0i32 = l3
		s1i32 = 7
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l10
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl7
			}
		lbl9:
			s0i32 = l7
			s1i32 = l11
			s0i32 = s0i32 + s1i32
			s0i32 = int32(ctx.Mem[int(s0i32+0)])
			l1 = s0i32
			s0i32 = 0
			l3 = s0i32
			s0i32 = l5
			l4 = s0i32
		lbl10:
			s0i32 = l4
			s1i32 = 0
			s2i32 = l1
			s3i32 = 128
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 1
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			l1 = s0i32
			s0i32 = l4
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l4 = s0i32
			s0i32 = l3
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s1i32 = l10
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl10
			}
			s0i32 = l5
			s1i32 = l12
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l7
			s1i32 = l14
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			s0i32 = l9
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l9 = s0i32
			s1i32 = l13
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl9
			}
			goto lbl7
		}
		s0i32 = l10
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl12:
			s0i32 = 0
			l3 = s0i32
			s0i32 = l5
			l1 = s0i32
		lbl13:
			s0i32 = l1
			s1i32 = 0
			s2i32 = l3
			s3i32 = l7
			s2i32 = s2i32 + s3i32
			s2i32 = int32(ctx.Mem[int(s2i32+0)])
			l4 = s2i32
			s3i32 = 1
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			ctx.Mem[int(s0i32+7)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 2
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+6)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 4
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+5)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 8
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+4)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 16
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+3)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 32
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+2)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 64
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+1)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 0
			s2i32 = l4
			s3i32 = 128
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 - s2i32
			s2i32 = 8
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l1
			s1i32 = 8
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s0i32 = l3
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s1i32 = l11
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl13
			}
			s0i32 = l5
			s1i32 = l12
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s0i32 = l7
			s1i32 = l14
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			s0i32 = l9
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l9 = s0i32
			s1i32 = l13
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl12
			}
			goto lbl7
			panic("unreachable executed")
			panic("unreachable executed")
		}
	lbl14:
		s0i32 = 0
		l3 = s0i32
		s0i32 = l5
		l1 = s0i32
	lbl15:
		s0i32 = l1
		s1i32 = 0
		s2i32 = l3
		s3i32 = l7
		s2i32 = s2i32 + s3i32
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		l4 = s2i32
		s3i32 = 1
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+7)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 2
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+6)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 4
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+5)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 8
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+4)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 16
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+3)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 32
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+2)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 64
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+1)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 128
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l11
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl15
		}
		s0i32 = l7
		s1i32 = l11
		s0i32 = s0i32 + s1i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l4 = s0i32
		s0i32 = 0
		l3 = s0i32
	lbl16:
		s0i32 = l1
		s1i32 = 0
		s2i32 = l4
		s3i32 = 128
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l4 = s0i32
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s1i32 = l10
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl16
		}
		s0i32 = l5
		s1i32 = l12
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l7
		s1i32 = l14
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0i32 = l9
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s1i32 = l13
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl14
		}
	lbl7:
		s0i32 = l8
		l1 = s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s1i32 = s1i32 - s2i32
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = 2
		s1i32 = f44(ctx, s1i32, s2i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1076)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s2i32 = l3
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	}
	s0i32 = l1
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f202(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l11 = s0i32
	s0i32 = 0
	l9 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l1 = s0i32
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = 3
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 20960
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0i32
	}
	s0i32 = l3
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l13 = s0i32
	s0i32 = l6
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l11
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 1
	s2i32 = l1
	s3i32 = l1
	s4i32 = 2
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1076)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 4
	s0i32 = s0i32 | s1i32
	l14 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l10 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l1 = s0i32
lbl19:
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l8 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s0i32 = s0i32 - s1i32
	l12 = s0i32
	s0i32 = l8
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l16 = s0i32
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l17 = s0i32
	l5 = s0i32
lbl20:
	s0i32 = l5
	l3 = s0i32
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0i32
	s1i32 = l12
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl20
	}
	s0i32 = l10
	s1i32 = l7
	s2i32 = l15
	s1i32 = s1i32 + s2i32
	l12 = s1i32
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l15 = s0i32
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l8
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l17
	s3i32 = l16
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s1i32 = s1i32 + s2i32
	l3 = s1i32
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l7 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl22:
		s0i32 = l5
		s1i32 = l7
		s0i32 = s0i32 - s1i32
		l5 = s0i32
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+2)])
		l7 = s0i32
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		l3 = s0i32
		s0i32 = l5
		s1i32 = l7
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl22
		}
		s0i32 = l8
		l3 = s0i32
	}
	s0i32 = l12
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	s1i32 = l10
	s2i32 = l15
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l8 = s0i32
	s0i32 = l7
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l5 = s0i32
lbl23:
	s0i32 = l4
	s1i32 = l13
	s2i32 = l3
	s3i32 = l5
	s4i32 = l6
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s5i32 = l9
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
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s1i32 = l6
	s2i32 = l14
	s3i32 = l7
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
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
	s1i32 = l11
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l1
	s1i32 = l10
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl19
	}
lbl0:
	s0i32 = l6
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
