package internal

import (
	"unsafe"
)

func f375(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	var l18 int32
	_ = l18
	var l19 int32
	_ = l19
	var l20 int32
	_ = l20
	var l21 int32
	_ = l21
	var l22 int64
	_ = l22
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
	var s1f64 float64
	_ = s1f64
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 55
	s0i32 = s0i32 + s1i32
	l21 = s0i32
	s0i32 = l7
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	l18 = s0i32
	s0i32 = 0
	l1 = s0i32
lbl1:
	s0i32 = l15
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = 2147483647
	s2i32 = l15
	s1i32 = s1i32 - s2i32
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 29100
		s1i32 = 61
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = -1
		l15 = s0i32
		goto lbl2
	}
	s0i32 = l1
	s1i32 = l15
	s0i32 = s0i32 + s1i32
	l15 = s0i32
lbl2:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	l11 = s0i32
	l1 = s0i32
	s0i32 = l11
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l8 = s0i32
	if s0i32 != 0 {
	lbl19:
		s0i32 = l8
		s1i32 = 255
		s0i32 = s0i32 & s1i32
		l9 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l8 = s0i32
			goto lbl22
		}
		s0i32 = l9
		s1i32 = 37
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl21
		}
		s0i32 = l1
		l8 = s0i32
	lbl24:
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+1)])
		s1i32 = 37
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl22
		}
		s0i32 = l7
		s1i32 = l1
		s2i32 = 2
		s1i32 = s1i32 + s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+2)])
		l12 = s0i32
		s0i32 = l9
		l1 = s0i32
		s0i32 = l12
		s1i32 = 37
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl24
		}
	lbl22:
		s0i32 = l8
		s1i32 = l11
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l11
			s2i32 = l1
			f77(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l1
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = -1
		l17 = s0i32
		s0i32 = 1
		l8 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
		l1 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
		s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		s1i32 = 10
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl26
		}
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+2)])
		s1i32 = 36
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl26
		}
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
		s1i32 = -48
		s0i32 = s0i32 + s1i32
		l17 = s0i32
		s0i32 = 1
		l19 = s0i32
		s0i32 = 3
		l8 = s0i32
	lbl26:
		s0i32 = l7
		s1i32 = l1
		s2i32 = l8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = 0
		l8 = s0i32
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
		l16 = s0i32
		s1i32 = -32
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s1i32 = 31
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			l9 = s0i32
			goto lbl27
		}
		s0i32 = l1
		l9 = s0i32
		s0i32 = 1
		s1i32 = l12
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l12 = s0i32
		s1i32 = 75913
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
	lbl29:
		s0i32 = l7
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l8
		s1i32 = l12
		s0i32 = s0i32 | s1i32
		l8 = s0i32
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+1)]))
		l16 = s0i32
		s1i32 = -32
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s1i32 = 31
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
		s0i32 = l9
		l1 = s0i32
		s0i32 = 1
		s1i32 = l12
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l12 = s0i32
		s1i32 = 75913
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			goto lbl29
		}
	lbl27:
		s0i32 = l16
		s1i32 = 42
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l7
			s1i32 = l9
			s1i32 = int32(int8(ctx.Mem[int(s1i32+1)]))
			s2i32 = -48
			s1i32 = s1i32 + s2i32
			s2i32 = 10
			if uint32(s1i32) >= uint32(s2i32) {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl33
			}
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
			l1 = s1i32
			s1i32 = int32(ctx.Mem[int(s1i32+2)])
			s2i32 = 36
			if s1i32 != s2i32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl33
			}
			s1i32 = l1
			s1i32 = int32(int8(ctx.Mem[int(s1i32+1)]))
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = l4
			s1i32 = s1i32 + s2i32
			s2i32 = -192
			s1i32 = s1i32 + s2i32
			s2i32 = 10
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
			s1i32 = l1
			s1i32 = int32(int8(ctx.Mem[int(s1i32+1)]))
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s2i32 = l3
			s1i32 = s1i32 + s2i32
			s2i32 = -384
			s1i32 = s1i32 + s2i32
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l13 = s1i32
			s1i32 = 1
			l19 = s1i32
			s1i32 = l1
			s2i32 = 3
			s1i32 = s1i32 + s2i32
			goto lbl32
		lbl33:
			s1i32 = l19
			if s1i32 != 0 {
				goto lbl17
			}
			s1i32 = 0
			l19 = s1i32
			s1i32 = 0
			l13 = s1i32
			s1i32 = l0
			if s1i32 != 0 {
				s1i32 = l2
				s2i32 = l2
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
				l1 = s2i32
				s3i32 = 4
				s2i32 = s2i32 + s3i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
				s1i32 = l1
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
				l13 = s1i32
			}
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
			s2i32 = 1
			s1i32 = s1i32 + s2i32
		lbl32:
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
			s0i32 = l13
			s1i32 = -1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl30
			}
			s0i32 = 0
			s1i32 = l13
			s0i32 = s0i32 - s1i32
			l13 = s0i32
			s0i32 = l8
			s1i32 = 8192
			s0i32 = s0i32 | s1i32
			l8 = s0i32
			goto lbl30
		}
		s0i32 = l7
		s1i32 = 76
		s0i32 = s0i32 + s1i32
		s0i32 = f542(ctx, s0i32)
		l13 = s0i32
		s1i32 = 0
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
		l1 = s0i32
	lbl30:
		s0i32 = -1
		l10 = s0i32
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		s1i32 = 46
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl35
		}
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+1)])
		s1i32 = 42
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s0i32 = int32(int8(ctx.Mem[int(s0i32+2)]))
			s1i32 = -48
			s0i32 = s0i32 + s1i32
			s1i32 = 10
			if uint32(s0i32) >= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl37
			}
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
			l1 = s0i32
			s0i32 = int32(ctx.Mem[int(s0i32+3)])
			s1i32 = 36
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl37
			}
			s0i32 = l1
			s0i32 = int32(int8(ctx.Mem[int(s0i32+2)]))
			s1i32 = 2
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l4
			s0i32 = s0i32 + s1i32
			s1i32 = -192
			s0i32 = s0i32 + s1i32
			s1i32 = 10
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s0i32 = int32(int8(ctx.Mem[int(s0i32+2)]))
			s1i32 = 3
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l3
			s0i32 = s0i32 + s1i32
			s1i32 = -384
			s0i32 = s0i32 + s1i32
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l10 = s0i32
			s0i32 = l7
			s1i32 = l1
			s2i32 = 4
			s1i32 = s1i32 + s2i32
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
			goto lbl35
		lbl37:
			s0i32 = l19
			if s0i32 != 0 {
				goto lbl17
			}
			s0i32 = l0
			if s0i32 != 0 {
				s0i32 = l2
				s1i32 = l2
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
				l1 = s1i32
				s2i32 = 4
				s1i32 = s1i32 + s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
				s0i32 = l1
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			} else {
				s0i32 = 0
			}
			l10 = s0i32
			s0i32 = l7
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
			s2i32 = 2
			s1i32 = s1i32 + s2i32
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
			goto lbl35
		}
		s0i32 = l7
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = 76
		s0i32 = s0i32 + s1i32
		s0i32 = f542(ctx, s0i32)
		l10 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
		l1 = s0i32
	lbl35:
		s0i32 = 0
		l9 = s0i32
	lbl39:
		s0i32 = l9
		l20 = s0i32
		s0i32 = -1
		l14 = s0i32
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
		s1i32 = -65
		s0i32 = s0i32 + s1i32
		s1i32 = 57
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l7
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l16 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
		l9 = s0i32
		s0i32 = l16
		l1 = s0i32
		s0i32 = l9
		s1i32 = l20
		s2i32 = 58
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		s1i32 = 15439
		s0i32 = s0i32 + s1i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l9 = s0i32
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl39
		}
		s0i32 = l9
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l9
		s1i32 = 19
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l17
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl42
			}
			goto lbl0
		}
		s0i32 = l17
		s1i32 = 0
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl41
		}
		s0i32 = l4
		s1i32 = l17
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l3
		s2i32 = l17
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	lbl42:
		s0i32 = 0
		l1 = s0i32
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		goto lbl40
	lbl41:
		s0i32 = l0
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l7
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = l9
		s2i32 = l2
		s3i32 = l6
		f541(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
		l16 = s0i32
	lbl40:
		s0i32 = l8
		s1i32 = -65537
		s0i32 = s0i32 & s1i32
		l12 = s0i32
		s1i32 = l8
		s2i32 = l8
		s3i32 = 8192
		s2i32 = s2i32 & s3i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l8 = s0i32
		s0i32 = 0
		l14 = s0i32
		s0i32 = 15484
		l17 = s0i32
		s0i32 = l18
		l9 = s0i32
		s0i32 = l16
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		s0i32 = int32(int8(ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s1i32 = -33
		s0i32 = s0i32 & s1i32
		s1i32 = l1
		s2i32 = l1
		s3i32 = 15
		s2i32 = s2i32 & s3i32
		s3i32 = 3
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l1
		s2i32 = l20
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		s1i32 = -88
		s0i32 = s0i32 + s1i32
		l16 = s0i32
		s1i32 = 32
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l1
		s1i32 = -65
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s1i32 = 6
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 83
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
			s0i32 = l10
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl47
			}
			s0i32 = l7
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
			goto lbl45
		}
		s0i32 = l12
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl5
		case 1:
			goto lbl46
		case 2:
			goto lbl5
		default:
			goto lbl16
		}
	lbl47:
		s0i32 = 0
		l1 = s0i32
		s0i32 = l0
		s1i32 = 32
		s2i32 = l13
		s3i32 = 0
		s4i32 = l8
		f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		goto lbl44
	lbl46:
		s0i32 = l7
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l7
		s1i32 = l7
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i64)
		s0i32 = l7
		s1i32 = l7
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = -1
		l10 = s0i32
		s0i32 = l7
		s1i32 = 8
		s0i32 = s0i32 + s1i32
	lbl45:
		l9 = s0i32
		s0i32 = 0
		l1 = s0i32
	lbl50:
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl49
		}
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l11
		s0i32 = f544(ctx, s0i32, s1i32)
		l11 = s0i32
		s1i32 = 0
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l12 = s0i32
		if s0i32 != 0 {
			goto lbl51
		}
		s0i32 = l11
		s1i32 = l10
		s2i32 = l1
		s1i32 = s1i32 - s2i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl51
		}
		s0i32 = l9
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l10
		s1i32 = l1
		s2i32 = l11
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl50
		}
		goto lbl49
	lbl51:
		s0i32 = -1
		l14 = s0i32
		s0i32 = l12
		if s0i32 != 0 {
			goto lbl0
		}
	lbl49:
		s0i32 = l0
		s1i32 = 32
		s2i32 = l13
		s3i32 = l1
		s4i32 = l8
		f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			l1 = s0i32
			goto lbl44
		}
		s0i32 = 0
		l12 = s0i32
		s0i32 = l7
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
		l9 = s0i32
	lbl53:
		s0i32 = l9
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl44
		}
		s0i32 = l7
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l11
		s0i32 = f544(ctx, s0i32, s1i32)
		l11 = s0i32
		s1i32 = l12
		s0i32 = s0i32 + s1i32
		l12 = s0i32
		s1i32 = l1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl44
		}
		s0i32 = l0
		s1i32 = l7
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		s2i32 = l11
		f77(ctx, s0i32, s1i32, s2i32)
		s0i32 = l9
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l12
		s1i32 = l1
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl53
		}
	lbl44:
		s0i32 = l0
		s1i32 = 32
		s2i32 = l13
		s3i32 = l1
		s4i32 = l8
		s5i32 = 8192
		s4i32 = s4i32 ^ s5i32
		f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l13
		s1i32 = l1
		s2i32 = l13
		s3i32 = l1
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		goto lbl1
	lbl21:
		s0i32 = l7
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+1)])
		l8 = s0i32
		s0i32 = l9
		l1 = s0i32
		goto lbl19
	lbl20:
		s0i32 = l16
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl5
		case 1:
			goto lbl5
		case 2:
			goto lbl5
		case 3:
			goto lbl5
		case 4:
			goto lbl5
		case 5:
			goto lbl5
		case 6:
			goto lbl5
		case 7:
			goto lbl5
		case 8:
			goto lbl16
		case 9:
			goto lbl5
		case 10:
			goto lbl14
		case 11:
			goto lbl13
		case 12:
			goto lbl16
		case 13:
			goto lbl16
		case 14:
			goto lbl16
		case 15:
			goto lbl5
		case 16:
			goto lbl13
		case 17:
			goto lbl5
		case 18:
			goto lbl5
		case 19:
			goto lbl5
		case 20:
			goto lbl5
		case 21:
			goto lbl9
		case 22:
			goto lbl12
		case 23:
			goto lbl11
		case 24:
			goto lbl5
		case 25:
			goto lbl5
		case 26:
			goto lbl15
		case 27:
			goto lbl5
		case 28:
			goto lbl8
		case 29:
			goto lbl5
		case 30:
			goto lbl5
		default:
			goto lbl10
		}
	}
	s0i32 = l15
	l14 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l19
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = 1
	l1 = s0i32
lbl54:
	s0i32 = l4
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = l1
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = l2
		s3i32 = l6
		f541(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = 1
		l14 = s0i32
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = 10
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl54
		}
		goto lbl0
	}
	s0i32 = 1
	l14 = s0i32
	s0i32 = l1
	s1i32 = 10
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl56:
	s0i32 = l4
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l1
	s1i32 = 8
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl56
	}
	goto lbl0
lbl17:
	s0i32 = -1
	l14 = s0i32
	goto lbl0
lbl16:
	s0i32 = l0
	s1i32 = l7
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	s2i32 = l13
	s3i32 = l10
	s4i32 = l8
	s5i32 = l1
	s6i32 = l5
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, float64, int32, int32, int32, int32) int32)(table[s6i32].f()))(ctx, s0i32, s1f64, s2i32, s3i32, s4i32, s5i32)
	l1 = s0i32
	goto lbl1
lbl15:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l1 = s0i32
	s1i32 = 15494
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l11 = s0i32
	s1i32 = l10
	s0i32 = f1367(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = l10
	s2i32 = l11
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l9 = s0i32
	s0i32 = l12
	l8 = s0i32
	s0i32 = l1
	s1i32 = l11
	s0i32 = s0i32 - s1i32
	s1i32 = l10
	s2i32 = l1
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l10 = s0i32
	goto lbl5
lbl14:
	s0i32 = l7
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	ctx.Mem[int(s0i32+55)] = uint8(s1i64)
	s0i32 = 1
	l10 = s0i32
	s0i32 = l21
	l11 = s0i32
	s0i32 = l12
	l8 = s0i32
	goto lbl5
lbl13:
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l22 = s0i64
	s1i64 = -1
	if s0i64 <= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i64 = 0
		s2i64 = l22
		s1i64 = s1i64 - s2i64
		l22 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
		s0i32 = 1
		l14 = s0i32
		s0i32 = 15484
		goto lbl7
	}
	s0i32 = l8
	s1i32 = 2048
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = 1
		l14 = s0i32
		s0i32 = 15485
		goto lbl7
	}
	s0i32 = 15486
	s1i32 = 15484
	s2i32 = l8
	s3i32 = 1
	s2i32 = s2i32 & s3i32
	l14 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	goto lbl7
lbl12:
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l18
	s0i32 = f1373(ctx, s0i64, s1i32)
	l11 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l10
	s1i32 = l18
	s2i32 = l11
	s1i32 = s1i32 - s2i32
	l1 = s1i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l10
	s3i32 = l1
	if s2i32 > s3i32 {
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
	goto lbl6
lbl11:
	s0i32 = l10
	s1i32 = 8
	s2i32 = l10
	s3i32 = 8
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
	l10 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = s0i32 | s1i32
	l8 = s0i32
	s0i32 = 120
	l1 = s0i32
lbl10:
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l18
	s2i32 = l1
	s3i32 = 32
	s2i32 = s2i32 & s3i32
	s0i32 = f1372(ctx, s0i64, s1i32, s2i32)
	l11 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l1
	s1i32 = 4
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 15484
	s0i32 = s0i32 + s1i32
	l17 = s0i32
	s0i32 = 2
	l14 = s0i32
	goto lbl6
lbl9:
	s0i32 = 0
	l1 = s0i32
	s0i32 = l20
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l9 = s0i32
	s1i32 = 7
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l9
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl64
	case 1:
		goto lbl63
	case 2:
		goto lbl62
	case 3:
		goto lbl61
	case 4:
		goto lbl1
	case 5:
		goto lbl60
	case 6:
		goto lbl59
	default:
		goto lbl65
	}
lbl65:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl1
lbl64:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl1
lbl63:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	s1i64 = int64(s1i32)
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	goto lbl1
lbl62:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
	goto lbl1
lbl61:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	goto lbl1
lbl60:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl1
lbl59:
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l15
	s1i64 = int64(s1i32)
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	goto lbl1
lbl8:
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l22 = s0i64
	s0i32 = 15484
lbl7:
	l17 = s0i32
	s0i64 = l22
	s1i32 = l18
	s0i32 = f211(ctx, s0i64, s1i32)
	l11 = s0i32
lbl6:
	s0i32 = l8
	s1i32 = -65537
	s0i32 = s0i32 & s1i32
	s1i32 = l8
	s2i32 = l10
	s3i32 = -1
	if s2i32 > s3i32 {
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
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	l22 = s0i64
	s0i32 = l10
	if s0i32 != 0 {
		goto lbl67
	}
	s0i64 = l22
	if s0i64 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl67
	}
	s0i32 = l18
	l11 = s0i32
	s0i32 = 0
	goto lbl66
lbl67:
	s0i32 = l10
	s1i64 = l22
	if s1i64 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l18
	s3i32 = l11
	s2i32 = s2i32 - s3i32
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s2i32 = l10
	s3i32 = l1
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
lbl66:
	l10 = s0i32
lbl5:
	s0i32 = l0
	s1i32 = 32
	s2i32 = l14
	s3i32 = l9
	s4i32 = l11
	s3i32 = s3i32 - s4i32
	l12 = s3i32
	s4i32 = l10
	s5i32 = l10
	s6i32 = l12
	if s5i32 < s6i32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	l16 = s3i32
	s2i32 = s2i32 + s3i32
	l9 = s2i32
	s3i32 = l13
	s4i32 = l13
	s5i32 = l9
	if s4i32 < s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l1 = s2i32
	s3i32 = l9
	s4i32 = l8
	f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s1i32 = l17
	s2i32 = l14
	f77(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 48
	s2i32 = l1
	s3i32 = l9
	s4i32 = l8
	s5i32 = 65536
	s4i32 = s4i32 ^ s5i32
	f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s1i32 = 48
	s2i32 = l16
	s3i32 = l12
	s4i32 = 0
	f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	s1i32 = l11
	s2i32 = l12
	f77(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = 32
	s2i32 = l1
	s3i32 = l9
	s4i32 = l8
	s5i32 = 8192
	s4i32 = s4i32 ^ s5i32
	f87(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	goto lbl1
lbl4:
	s0i32 = 0
	l14 = s0i32
lbl0:
	s0i32 = l7
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l14
	return s0i32
}
