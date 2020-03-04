package internal

import (
	"unsafe"
)

func f177(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
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
	var l7 int64
	_ = l7
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l4
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+40)]))
		s1i32 = s1i32 + s2i32
	} else {
		s1i32 = 0
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l3 = s0i32
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 1
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint16(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 0
	ctx.Mem[int(s0i32+91)] = uint8(s1i32)
	s0i32 = l2
	s1i32 = l3
	s2i32 = -4
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	s3i32 = l3
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	l3 = s2i32
	s3i32 = l4
	s4i32 = l3
	if s3i32 < s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = l3
	s3i32 = l4
	s4i32 = 0
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l5 = s1i32
	s2i32 = 4
	if s1i32 >= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
		s2i32 = 72
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l6 = s1i32
		l1 = s1i32
	lbl5:
		s1i32 = l1
		l3 = s1i32
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1i32 = l2
		s2i32 = 56
		s1i32 = s1i32 + s2i32
		s2i32 = l2
		s3i32 = 96
		s2i32 = s2i32 + s3i32
		s1i32 = f131(ctx, s1i32, s2i32)
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl5
		}
		s1i32 = l3
		s2i32 = -8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = l6
		s3i32 = l5
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		s3i32 = l1
		s2i32 = s2i32 - s3i32
		s3i32 = 3
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = f2088(ctx, s1i32, s2i32)
		l1 = s1i32
		s2i32 = 2
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		switch s1i32 {
		case 0:
			goto lbl6
		case 1:
			goto lbl7
		default:
			goto lbl1
		}
	lbl7:
		s1i32 = 2
		goto lbl2
	lbl6:
		s1i32 = l2
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l1 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
		l5 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)])) = uint32(s2i32)
		s1i32 = 0
		l3 = s1i32
		s1i32 = l2
		s2i32 = l5
		if s2i32 != 0 {
			s2i32 = l5
			s3i32 = l1
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
			s2i32 = s2i32 + s3i32
		} else {
			s2i32 = 0
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		l1 = s2i32
		s3i32 = -4
		s2i32 = s2i32 + s3i32
		s3i32 = 0
		s4i32 = l1
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)])) = uint32(s2i32)
		s1i32 = l4
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l4
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l2
		s2i32 = 0
		ctx.Mem[int(s1i32+91)] = uint8(s2i32)
		s1i32 = l2
		s2i32 = 1
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)])) = uint16(s2i32)
		goto lbl3
	}
	s1i32 = l1
	s1i32 = int32(ctx.Mem[int(s1i32+84)])
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		s1i32 = int32(ctx.Mem[int(s1i32+85)])
		l3 = s1i32
		goto lbl9
	}
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s3i32 = l3
	s1i32 = f28(ctx, s1i32, s2i32, s3i32)
	l3 = s1i32
	s1i32 = l1
	s2i32 = 0
	ctx.Mem[int(s1i32+84)] = uint8(s2i32)
	s1i32 = l1
	s2i32 = l3
	ctx.Mem[int(s1i32+85)] = uint8(s2i32)
lbl9:
	s1i32 = l3
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		goto lbl3
	}
	s1i32 = 0
	l1 = s1i32
	goto lbl1
lbl3:
	s1i32 = l2
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = uint64(s2i64)
	s1i32 = l2
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
	s1i32 = l2
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
	s1i32 = l2
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
	s1i32 = l2
	s2i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
	s1i32 = l2
	s2i32 = 1
	ctx.Mem[int(s1i32+52)] = uint8(s2i32)
	s1i32 = 0
	l4 = s1i32
	s1i32 = l2
	s2i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)])) = uint32(s2i32)
	s1i32 = l2
	s2i64 = 8589934597
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)])) = uint64(s2i64)
lbl12:
	s1i32 = 2
	s2i32 = l2
	s3i32 = 56
	s2i32 = s2i32 + s3i32
	s3i32 = l2
	s4i32 = 96
	s3i32 = s3i32 + s4i32
	s2i32 = f131(ctx, s2i32, s3i32)
	l3 = s2i32
	s3i32 = 6
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = l3
	s2i32 = 1
	s1i32 = s1i32 - s2i32
	switch s1i32 {
	case 0:
		goto lbl14
	case 1:
		goto lbl15
	case 2:
		goto lbl15
	case 3:
		goto lbl17
	case 4:
		goto lbl16
	case 5:
		goto lbl11
	default:
		goto lbl18
	}
lbl18:
	s1i32 = 2
	s2i32 = l4
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = l2
	s2i32 = l2
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+96)]))
	l7 = s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
	s1i32 = l2
	s2i64 = l7
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
	s1i32 = l2
	s2i64 = l7
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	goto lbl12
lbl17:
	s1i32 = 3
	l3 = s1i32
	s1i32 = 1
	goto lbl13
lbl16:
	s1i32 = l2
	s2i32 = l2
	s1i32 = f675(ctx, s1i32, s2i32)
	if s1i32 != 0 {
		goto lbl12
	}
	s1i32 = 2
	goto lbl2
lbl15:
	s1i32 = 2
	l3 = s1i32
lbl14:
	s1i32 = 1
lbl13:
	l1 = s1i32
lbl19:
	s1i32 = l2
	s2i32 = l2
	s3i32 = 96
	s2i32 = s2i32 + s3i32
	s3i32 = l1
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s1i32 = f675(ctx, s1i32, s2i32)
	if s1i32 != 0 {
		s1i32 = l1
		s2i32 = l3
		if s1i32 == s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		l5 = s1i32
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1i32 = l5
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl19
		}
		goto lbl12
	}
	s1i32 = 2
	goto lbl2
lbl11:
	s1i32 = 1
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+9)])
	s3i32 = 2
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl2
	}
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 2
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl21
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s1i32 = int32(ctx.Mem[int(s1i32+84)])
	if s1i32 != 0 {
		s1i32 = l1
		s2i32 = 4
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s3i32 = l1
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
		s1i32 = f28(ctx, s1i32, s2i32, s3i32)
		l3 = s1i32
		s1i32 = l1
		s2i32 = 0
		ctx.Mem[int(s1i32+84)] = uint8(s2i32)
		s1i32 = l1
		s2i32 = l3
		ctx.Mem[int(s1i32+85)] = uint8(s2i32)
	}
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl21
	}
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	if s1i32 != 0 {
		goto lbl21
	}
	s1i32 = 2
	s2i32 = 1
	s3i32 = l2
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+48)]))
	s4i32 = 2
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	goto lbl2
lbl21:
	s1i32 = l0
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
	ctx.Mem[int(s1i32+9)] = uint8(s2i32)
	s1i32 = 1
lbl2:
	l1 = s1i32
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
lbl1:
	s0i32 = l2
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l1
	return s0i32
}
