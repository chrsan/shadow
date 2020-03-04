package internal

import (
	"unsafe"
)

func f1688(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l15 int64
	_ = l15
	var l16 float32
	_ = l16
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
	var s7i32 int32
	_ = s7i32
	var s0i64 int64
	_ = s0i64
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
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = 272
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = l3
	s1i32 = 200
	s0i32 = s0i32 + s1i32
	s0i32 = f37(ctx, s0i32)
	l8 = s0i32
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+212)])) = uint32(s1i32)
	s0f32 = l16
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	l16 = s0f32
	s0i32 = l3
	s1i32 = l1
	s2i32 = l2
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l8
		l2 = s1i32
		s1i32 = 1
		goto lbl0
	}
	s1i32 = l2
	f231(ctx, s1i32)
	s1i32 = 0
lbl0:
	l4 = s1i32
	ctx.Mem[int(s0i32+216)] = uint8(s1i32)
	s0f32 = l16
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
	} else {
		s0i32 = l3
		s1i32 = 0
		ctx.Mem[int(s0i32+16)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = l3
		s2i32 = 56
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		s4i32 = 224
		s3i32 = s3i32 + s4i32
		s0i32 = f333(ctx, s0i32, s1i32, s2i32, s3i32)
		l4 = s0i32
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+16)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l0
		s1i32 = l3
		s2i32 = 56
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+224)]))
		f1687(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+10)])
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l2
		s1i32 = l2
		s1i32 = int32(ctx.Mem[int(s1i32+10)])
		s2i32 = 2
		s1i32 = s1i32 ^ s2i32
		ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		goto lbl3
	lbl4:
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+14)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s0i32 = int32(ctx.Mem[int(s0i32+90)])
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = 0
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l4 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		l6 = s1i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
		s0i32 = l6
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s0i32 = s0i32 + s1i32
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		s1i32 = 5
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
	lbl6:
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+8)])
		l4 = s0i32
		if s0i32 != 0 {
			s0i32 = l4
		} else {
			s0i32 = l1
			s0i32 = f177(ctx, s0i32)
		}
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l12 = s0i32
	lbl5:
		s0i32 = l3
		s1i32 = 56
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2f32 = l16
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s4i32 = l0
		s4i32 = int32(ctx.Mem[int(s4i32+12)])
		s5i32 = l0
		s5i32 = int32(ctx.Mem[int(s5i32+13)])
		s6i32 = l0
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+8)]))
		s7i32 = l12
		s0i32 = f1702(ctx, s0i32, s1i32, s2f32, s3f32, s4i32, s5i32, s6f32, s7i32)
		l5 = s0i32
		s1i32 = 108
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l5
		s1i32 = 96
		s0i32 = s0i32 + s1i32
		l13 = s0i32
		s0i32 = l3
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = 0
		s0i32 = f329(ctx, s0i32, s1i32, s2i32)
		l10 = s0i32
		s0i32 = l3
		s1i32 = 248
		s0i32 = s0i32 + s1i32
		l14 = s0i32
		s0i32 = l3
		s1i32 = 240
		s0i32 = s0i32 + s1i32
		l11 = s0i32
		s0i32 = l3
		s1i32 = 224
		s0i32 = s0i32 + s1i32
		s1i32 = 8
		s0i32 = s0i32 | s1i32
		l7 = s0i32
		s0i32 = 0
		l4 = s0i32
	lbl8:
		s0i32 = l4
		l6 = s0i32
		s0i32 = l10
		s1i32 = l3
		s2i32 = 224
		s1i32 = s1i32 + s2i32
		s0i32 = f131(ctx, s0i32, s1i32)
		l4 = s0i32
		s1i32 = 6
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 1
			s0i32 = s0i32 - s1i32
			switch s0i32 {
			case 0:
				goto lbl15
			case 1:
				goto lbl14
			case 2:
				goto lbl13
			case 3:
				goto lbl12
			case 4:
				goto lbl11
			case 5:
				goto lbl9
			default:
				goto lbl16
			}
		lbl16:
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			s1i32 = 1
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l5
				s1i32 = 0
				s2i32 = 0
				f400(ctx, s0i32, s1i32, s2i32)
			}
			s0i32 = l5
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
			s0i32 = l3
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)]))
			l15 = s0i64
			s0i32 = l5
			s1i32 = 0
			ctx.Mem[int(s0i32+141)] = uint8(s1i32)
			s0i32 = l5
			s1i64 = l15
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint64(s1i64)
			s0i32 = l5
			s1i64 = l15
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
			s0i32 = l6
			l4 = s0i32
			goto lbl8
		lbl15:
			s0i32 = l5
			s1i32 = l7
			s2i32 = l10
			f78(ctx, s0i32, s1i32, s2i32)
			s0i32 = 1
			l4 = s0i32
			goto lbl8
		lbl14:
			s0i32 = l5
			s1i32 = l7
			s2i32 = l11
			f1694(ctx, s0i32, s1i32, s2i32)
			s0i32 = 2
			l4 = s0i32
			goto lbl8
		lbl13:
			s0i32 = l5
			s1i32 = l7
			s2i32 = l11
			s3i32 = l10
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
			s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
			f1696(ctx, s0i32, s1i32, s2i32, s3f32)
			s0i32 = 3
			l4 = s0i32
			goto lbl8
		lbl12:
			s0i32 = l5
			s1i32 = l7
			s2i32 = l11
			s3i32 = l14
			f1689(ctx, s0i32, s1i32, s2i32, s3i32)
			s0i32 = 4
			l4 = s0i32
			goto lbl8
		lbl11:
			s0i32 = l0
			s0i32 = int32(ctx.Mem[int(s0i32+12)])
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			s0i32 = l5
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l3
				s1i32 = l5
				s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
				*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
				s0i32 = l5
				s1i32 = l3
				s2i32 = 8
				s1i32 = s1i32 + s2i32
				s2i32 = 0
				f78(ctx, s0i32, s1i32, s2i32)
				s0i32 = 1
				l4 = s0i32
				goto lbl8
			}
			s0i32 = l13
			s1i32 = 0
			s0i32 = f679(ctx, s0i32, s1i32)
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl18
			}
			s0i32 = 1
			l4 = s0i32
			s0i32 = l9
			s1i32 = l5
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
			s0i32 = f679(ctx, s0i32, s1i32)
			if s0i32 != 0 {
				goto lbl8
			}
		lbl18:
			s0i32 = l5
			s1i32 = 1
			s2i32 = l6
			s3i32 = 1
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			f400(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l6
		l4 = s0i32
		goto lbl8
	lbl9:
		s0i32 = l5
		s1i32 = 0
		s2i32 = l6
		s3i32 = 1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		f400(ctx, s0i32, s1i32, s2i32)
		s0i32 = l2
		s1i32 = l9
		f194(ctx, s0i32, s1i32)
		s0i32 = l12
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+14)])
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s0i32 = s0i32 | s1i32
		if s0i32 != 0 {
			goto lbl20
		}
		s0i32 = l3
		s1i32 = 2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l3
		s2i32 = 224
		s1i32 = s1i32 + s2i32
		f2100(ctx, s0i32, s1i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)]))
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l1
			f2092(ctx, s0i32, s1i32)
			goto lbl20
		}
		s0i32 = l3
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+232)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 69784829952
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+256)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+248)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+240)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+232)])) = uint64(s1i64)
		s0i32 = l3
		s1i64 = 1065353216
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+224)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s2i32 = l3
		s3i32 = 224
		s2i32 = s2i32 + s3i32
		f684(ctx, s0i32, s1i32, s2i32)
	lbl20:
		s0i32 = l1
		s0i32 = int32(ctx.Mem[int(s0i32+10)])
		s1i32 = 2
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l2
			s1i32 = int32(ctx.Mem[int(s1i32+10)])
			s2i32 = 2
			s1i32 = s1i32 ^ s2i32
			ctx.Mem[int(s0i32+10)] = uint8(s1i32)
		}
		s0i32 = l5
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		f34(ctx, s0i32)
		s0i32 = l9
		f34(ctx, s0i32)
		s0i32 = l13
		f34(ctx, s0i32)
	lbl3:
		s0i32 = l3
		s0i32 = int32(ctx.Mem[int(s0i32+216)])
	}
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+212)]))
		f194(ctx, s0i32, s1i32)
	}
	s0i32 = l8
	f34(ctx, s0i32)
	s0i32 = l3
	s1i32 = 272
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
