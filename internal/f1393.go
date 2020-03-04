package internal

import (
	"unsafe"
)

func f1393(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float32
	_ = l15
	var l16 float32
	_ = l16
	var l17 float32
	_ = l17
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 144
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l5 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 4
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
			goto lbl1
		}
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		s2i32 = l5
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = f29(ctx, s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l5 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 4
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
			goto lbl0
		}
		s0i32 = l0
		s1i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
		s2i32 = l5
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = f29(ctx, s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
	}
	s0i32 = l4
	s1i32 = 104
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 1
	s0i32 = f329(ctx, s0i32, s1i32, s2i32)
	l9 = s0i32
	s1i32 = l4
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s0i32 = f131(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 76
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l4
		s1i32 = 88
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s0i32 = l4
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0i32 = l4
		s1i32 = -64
		s0i32 = s0i32 - s1i32
		s1i32 = 8
		s0i32 = s0i32 | s1i32
		l7 = s0i32
	lbl6:
		s0i32 = l8
		if s0i32 != 0 {
			s0i32 = 0
			l8 = s0i32
			goto lbl4
		}
		s0i32 = 0
		l8 = s0i32
		s0i32 = l1
		s1i32 = 5
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl14
		case 1:
			goto lbl13
		case 2:
			goto lbl12
		case 3:
			goto lbl11
		case 4:
			goto lbl9
		default:
			goto lbl10
		}
	lbl14:
		s0i32 = l2
		s1i32 = l4
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l7
		s3i32 = 1
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0i32 = l4
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l1
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = -8
			s1i32 = s1i32 + s2i32
			l1 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 - s1f32
			l12 = s0f32
			s1f32 = l12
			s0f32 = s0f32 * s1f32
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
			s2i32 = l1
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s1f32 = s1f32 - s2f32
			l12 = s1f32
			s2f32 = l12
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.00390625
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl15
			}
		}
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	lbl15:
		s0i32 = l3
		s1i32 = l7
		s2i32 = l7
		s3i32 = 1
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l7
		f103(ctx, s0i32, s1i32)
		goto lbl8
	lbl13:
		s0i32 = l2
		s1i32 = l4
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		s1f32 = 0.25
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
		s2f32 = 0.25
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l12 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		s1f32 = 0.25
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		s2f32 = 0.25
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l14 = s0f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0f32 = l14
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l1
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = -8
			s1i32 = s1i32 + s2i32
			l6 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 - s1f32
			l13 = s0f32
			s1f32 = l13
			s0f32 = s0f32 * s1f32
			s1f32 = l12
			s2i32 = l6
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s1f32 = s1f32 - s2f32
			l13 = s1f32
			s2f32 = l13
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.00390625
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl19
			}
		}
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		l1 = s0i32
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl18
		}
	lbl19:
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l1
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f32 = s0f32 - s1f32
		l12 = s0f32
		s1f32 = l12
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 - s2f32
		l12 = s1f32
		s2f32 = l12
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 0.00390625
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl17
		}
	lbl18:
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	lbl17:
		s0i32 = l3
		s1i32 = l4
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s2i32 = l4
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l7
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l10
		f103(ctx, s0i32, s1i32)
		goto lbl8
	lbl12:
		s0i32 = l2
		s1i32 = l4
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0f32 = 1
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l13 = s1f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l12 = s1f32
		s2f32 = 0.5
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 / s1f32
		l15 = s0f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		s2f32 = 0.25
		s1f32 = s1f32 * s2f32
		s2f32 = l12
		s3i32 = l4
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+52)]))
		s3f32 = 0.25
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 * s1f32
		l14 = s0f32
		s0f32 = l15
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s2f32 = 0.25
		s1f32 = s1f32 * s2f32
		s2f32 = l12
		s3i32 = l4
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
		s3f32 = 0.25
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s0f32 = s0f32 * s1f32
		l12 = s0f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0f32 = l12
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l1
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = -8
			s1i32 = s1i32 + s2i32
			l6 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 - s1f32
			l15 = s0f32
			s1f32 = l15
			s0f32 = s0f32 * s1f32
			s1f32 = l14
			s2i32 = l6
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s1f32 = s1f32 - s2f32
			l15 = s1f32
			s2f32 = l15
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.00390625
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl23
			}
		}
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		l1 = s0i32
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl22
		}
	lbl23:
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l1
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f32 = s0f32 - s1f32
		l12 = s0f32
		s1f32 = l12
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 - s2f32
		l12 = s1f32
		s2f32 = l12
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 0.00390625
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl21
		}
	lbl22:
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	lbl21:
		s0i32 = l0
		s1i32 = l3
		s2i32 = l4
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3f32 = l13
		f552(ctx, s0i32, s1i32, s2i32, s3f32)
		goto lbl8
	lbl11:
		s0i32 = l2
		s1i32 = l4
		s2i32 = 32
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3i32 = 4
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l12 = s0f32
		s1f32 = 0.32495117
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l14 = s1f32
		s2f32 = 0.44311523
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
		l13 = s1f32
		s2f32 = 0.20141602
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
		l15 = s1f32
		s2f32 = 0.030517578
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l17 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l18 = s0f32
		s1f32 = 0.32495117
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		l19 = s1f32
		s2f32 = 0.44311523
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		l20 = s1f32
		s2f32 = 0.20141602
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		l21 = s1f32
		s2f32 = 0.030517578
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l22 = s0f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 != 0 {
			s0f32 = l22
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l1
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = -8
			s1i32 = s1i32 + s2i32
			l6 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 - s1f32
			l16 = s0f32
			s1f32 = l16
			s0f32 = s0f32 * s1f32
			s1f32 = l17
			s2i32 = l6
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s1f32 = s1f32 - s2f32
			l16 = s1f32
			s2f32 = l16
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.00390625
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl25
			}
		}
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		l1 = s0i32
		s1f32 = l17
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l22
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		l15 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		l13 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		l14 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l12 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		l21 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l20 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l19 = s0f32
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l18 = s0f32
	lbl25:
		s0f32 = l12
		s1f32 = 0.030517578
		s0f32 = s0f32 * s1f32
		s1f32 = l14
		s2f32 = 0.20141602
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l13
		s2f32 = 0.44311523
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l15
		s2f32 = 0.32495117
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l12 = s0f32
		s0f32 = l18
		s1f32 = 0.030517578
		s0f32 = s0f32 * s1f32
		s1f32 = l19
		s2f32 = 0.20141602
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l20
		s2f32 = 0.44311523
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l21
		s2f32 = 0.32495117
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l14 = s0f32
		s0i32 = l1
		if s0i32 != 0 {
			s0f32 = l14
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l1
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s2i32 = -8
			s1i32 = s1i32 + s2i32
			l6 = s1i32
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s0f32 = s0f32 - s1f32
			l13 = s0f32
			s1f32 = l13
			s0f32 = s0f32 * s1f32
			s1f32 = l12
			s2i32 = l6
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			s1f32 = s1f32 - s2f32
			l13 = s1f32
			s2f32 = l13
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 + s1f32
			s1f32 = 0.00390625
			if s0f32 < s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl29
			}
		}
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		l1 = s0i32
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l1
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l1 = s0i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl28
		}
	lbl29:
		s0i32 = l4
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l1
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f32 = s0f32 - s1f32
		l12 = s0f32
		s1f32 = l12
		s0f32 = s0f32 * s1f32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 - s2f32
		l12 = s1f32
		s2f32 = l12
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 0.00390625
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl27
		}
	lbl28:
		s0i32 = l5
		s0i32 = f45(ctx, s0i32)
		s1i32 = l4
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	lbl27:
		s0i32 = l3
		s1i32 = l4
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s2i32 = l4
		s3i32 = -64
		s2i32 = s2i32 - s3i32
		s3i32 = 4
		f55(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l7
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l10
		f103(ctx, s0i32, s1i32)
		s0i32 = l0
		s1i32 = l11
		f103(ctx, s0i32, s1i32)
		goto lbl8
	lbl10:
		s0i32 = l6
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		goto lbl4
	lbl9:
		s0i32 = 1
		l8 = s0i32
	lbl8:
		s0i32 = 1
		l6 = s0i32
		s0i32 = l9
		s1i32 = l4
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		s0i32 = f131(ctx, s0i32, s1i32)
		l1 = s0i32
		s1i32 = 6
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl6
		}
	}
	s0i32 = l0
	f559(ctx, s0i32)
	s0i32 = 1
	l8 = s0i32
lbl4:
	s0i32 = l4
	s1i32 = 144
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l8
	return s0i32
lbl1:
	s0i32 = l4
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l4
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl0:
	s0i32 = l4
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l4
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
