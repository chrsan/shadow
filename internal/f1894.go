package internal

import (
	"math/bits"
	"unsafe"
)

func f1894(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l3
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
		s0i32 = l1
		s1i32 = l3
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		f139(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s0i32 = f428(ctx, s0i32, s1i32)
	l2 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	if s0i32 != 0 {
		s0i32 = -20
		l0 = s0i32
		goto lbl4
	}
	s0i32 = 0
	l0 = s0i32
lbl8:
	s0i32 = l0
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l7
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 6
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			s1i32 = l5
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl7
			}
			s0i32 = l6
			s1i32 = l5
			s2i32 = 20
			s1i32 = s1i32 * s2i32
			s0i32 = f29(ctx, s0i32, s1i32)
			l6 = s0i32
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l10 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l9 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l8 = s0i32
		s0i32 = l6
		s1i32 = l0
		s2i32 = 20
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = 0
		ctx.Mem[int(s0i32+12)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l10 = s0i32
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+32)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l9
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l10
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l2
		f427(ctx, s0i32)
		s0i32 = l7
		l0 = s0i32
		s0i32 = l2
		s0i32 = int32(ctx.Mem[int(s0i32+24)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		goto lbl6
	}
	s0i32 = l3
	s1i32 = 345
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4994
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4946
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 4863
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl7:
	s0i32 = l3
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 5024
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4946
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 4863
	s1i32 = l3
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl6:
	s0i32 = l6
	s1i32 = l6
	s2i32 = l7
	s3i32 = 20
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	l9 = s1i32
	s2i32 = -20
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl4:
	s0i32 = 64
	s1i32 = l0
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	s2i32 = 20
	s1i32 = i32DivS(s1i32, s2i32)
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1i32 = int32(bits.LeadingZeros32(uint32(s1i32)))
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 - s1i32
	s1i32 = l6
	s2i32 = l0
	f653(ctx, s0i32, s1i32, s2i32)
	s0i32 = l7
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl3:
	s0i32 = l6
	l4 = s0i32
lbl11:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	l0 = s0i32
	s1i32 = 3
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l5 = s0i32
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l8 = s0i32
			s0i32 = l4
			l2 = s0i32
		lbl14:
			s0i32 = l2
			l0 = s0i32
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l0
			s0i32 = int32(ctx.Mem[int(s0i32+32)])
			l10 = s0i32
			s1i32 = 2
			s0i32 = s0i32 & s1i32
			if s0i32 != 0 {
				goto lbl14
			}
			s0i32 = l8
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl14
			}
			s0i32 = l0
			s1i32 = l10
			s2i32 = 2
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+32)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l4
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
			s0i32 = l4
			s0i32 = int32(ctx.Mem[int(s0i32+12)])
			l0 = s0i32
		}
		s0i32 = l4
		l2 = s0i32
		s0i32 = l0
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
		lbl16:
			s0i32 = l2
			l0 = s0i32
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l0
			s0i32 = int32(ctx.Mem[int(s0i32+32)])
			l8 = s0i32
			s1i32 = 1
			s0i32 = s0i32 & s1i32
			if s0i32 != 0 {
				goto lbl16
			}
			s0i32 = l5
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl16
			}
			s0i32 = l4
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l8
			s2i32 = 1
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+32)] = uint8(s1i32)
		}
		s0i32 = l4
		s1i32 = 3
		ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	}
	s0i32 = l4
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l9
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l7
	l4 = s0i32
lbl2:
	s0i32 = l1
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	f441(ctx, s0i32, s1i32)
lbl17:
	s0i32 = l6
	l0 = s0i32
lbl18:
	s0i32 = l0
	l2 = s0i32
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0i32
	s0i32 = l1
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(s1i32)
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s2f32 = float32(s2i32)
	f176(ctx, s0i32, s1f32, s2f32)
	s0i32 = l2
	s1i32 = 0
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	s0i32 = 1
	l5 = s0i32
	s0i32 = l2
	l0 = s0i32
lbl19:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	s1i32 = l7
	l0 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl20
		}
	}
	s0i32 = l1
	s1i32 = l8
	s1f32 = float32(s1i32)
	s2i32 = l9
	s2f32 = float32(s2i32)
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(s1i32)
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s2f32 = float32(s2i32)
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
lbl20:
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l7 = s0i32
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl19
	}
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1f32 = float32(s1i32)
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s2f32 = float32(s2i32)
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l1
	f230(ctx, s0i32)
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l6
	f13(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
