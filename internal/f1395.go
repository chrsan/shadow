package internal

import (
	"math"
	"unsafe"
)

func f1395(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l7 = s0i32
		s0i32 = l6
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l6
		s1i32 = l7
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l0
	s1i32 = 27304
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = 84
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l6
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = -3229614080
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+160)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 257
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+158)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = l4
	ctx.Mem[int(s0i32+157)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+156)] = uint8(s1i32)
	s0i32 = l0
	s1i64 = -4294967296
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+176)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = 27320
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+100)]))
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+108)]))
	s4f32 = 0.5
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s2f32 = s2f32 * s3f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l0
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+104)]))
	s5f32 = 0.5
	s4f32 = s4f32 * s5f32
	s5i32 = l0
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+112)]))
	s6f32 = 0.5
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 + s2f32
	s2f32 = 0.0078125
	s1f32 = s1f32 * s2f32
	l8 = s1f32
	s2f32 = 64
	s1f32 = s1f32 * s2f32
	s2f32 = 150
	s1f32 = float32(math.Min(float64(s1f32), float64(s2f32)))
	l9 = s1f32
	s2f32 = l8
	s3f32 = 0
	s2f32 = float32(math.Max(float64(s2f32), float64(s3f32)))
	s3f32 = 1
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l9
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = s1f32
	s0i32 = l5
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		l3 = s0i32
		s0i32 = 1
		goto lbl1
	}
	s0i32 = l4
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l4
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l3 = s0i32
	s0i32 = l4
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = l3
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
lbl1:
	l6 = s0i32
	s0i32 = l5
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l10 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l4 = s0i32
		s0i32 = l3
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l3
		s1i32 = l4
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l5
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
	s0i32 = l5
	s1i32 = l5
	s2i32 = 56
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 60
	s2i32 = s2i32 + s3i32
	s3i32 = l5
	s4i32 = 48
	s3i32 = s3i32 + s4i32
	s4i32 = l5
	s5i32 = 52
	s4i32 = s4i32 + s5i32
	s5f32 = l11
	s6f32 = l10
	if s5f32 < s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	l3 = s3i32
	s4f32 = l8
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3f32 = l8
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s0i32 = f1394(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = 1
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		s1i32 = 3
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)]))
		s1i32 = 2139095040
		s0i32 = s0i32 & s1i32
		s1i32 = 2139095040
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l2 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			s1i32 = l2
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl6
			}
			s0i32 = l0
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			s2i32 = l2
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = f29(ctx, s1i32, s2i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l2 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			s1i32 = l2
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
			s0i32 = l0
			s1i32 = l2
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
			s2i32 = l2
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = f29(ctx, s1i32, s2i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		s2i32 = 12
		s1i32 = s1i32 * s2i32
		l1 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			s1i32 = l1
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl4
			}
			s0i32 = l0
			s1i32 = l1
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
			s2i32 = l1
			s3i32 = 1
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = f29(ctx, s1i32, s2i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		}
		s0i32 = l4
		s0f32 = math.Float32frombits(uint32(s0i32))
		l8 = s0f32
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+158)])
		if s0i32 != 0 {
			s0i32 = l0
			s1f32 = l8
			s2f32 = l9
			s3i32 = 0
			s0i32 = f557(ctx, s0i32, s1f32, s2f32, s3i32)
			l3 = s0i32
			goto lbl8
		}
		s0i32 = l0
		s1f32 = l8
		s2f32 = l9
		s0i32 = f553(ctx, s0i32, s1f32, s2f32)
		l3 = s0i32
	lbl8:
		s0i32 = l0
		s1i32 = l3
		ctx.Mem[int(s0i32+156)] = uint8(s1i32)
	}
	s0i32 = l5
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
lbl6:
	s0i32 = l5
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l5
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl5:
	s0i32 = l5
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l5
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl4:
	s0i32 = l5
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l5
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l5
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
