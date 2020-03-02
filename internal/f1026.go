package internal

import (
	"unsafe"
)

func f1026(ctx *Context, l0 int32, l1 int32, l2 float32, l3 int32, l4 int32) int32 {
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
	var l17 int64
	_ = l17
	var l18 int64
	_ = l18
	var l19 float64
	_ = l19
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
	var s2i64 int64
	_ = s2i64
	var s1f32 float32
	_ = s1f32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l11 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l8 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l8
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l11
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1f32 = l2
	s1f64 = float64(s1f32)
	l19 = s1f64
	s2f64 = l19
	s0i32 = f1042(ctx, s0i32, s1f64, s2f64)
	l8 = s0i32
	l9 = s0i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f64 = 3
	s0f64 = s0f64 * s1f64
	s1f64 = 1
	if s0f64 <= s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0i32 = 0
	} else {
		s0i32 = l9
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1f64 = 3
		s0f64 = s0f64 * s1f64
		s1f64 = 1
		if s0f64 <= s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 4294967296
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = 1
		l7 = s0i32
		s0i32 = l4
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l4
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		goto lbl1
	}
	s0i32 = l11
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l8
	s2i32 = l1
	s3i32 = l0
	f1041(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 != 0 {
		goto lbl4
	}
	goto lbl1
lbl4:
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l11
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
	}
	s0i32 = l8
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l7 = s0i32
		s0i32 = l3
		s1i32 = 3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l17 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l18 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = l0
		s1i64 = l18
		s1i32 = int32(s1i64)
		s2i64 = l17
		s2i32 = int32(s2i64)
		s1i32 = s1i32 - s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		goto lbl1
	}
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl9
	case 1:
		goto lbl8
	default:
		goto lbl10
	}
lbl10:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l11
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l6 = s2i32
	s3i32 = l11
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+60)]))
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl14
	case 1:
		goto lbl11
	case 2:
		goto lbl13
	case 3:
		goto lbl12
	default:
		goto lbl15
	}
lbl15:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	l12 = s0i32
lbl16:
	s0i32 = l9
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l3
	l0 = s0i32
	s0i32 = l8
	l4 = s0i32
	s0i32 = 7
	l1 = s0i32
lbl17:
	s0i32 = l5
	s1i32 = l5
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l7 = s1i32
	s2i32 = l4
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	s3i32 = l1
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 1
	s2i32 = s2i32 & s3i32
	l13 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 0
	s3i32 = l13
	s2i32 = s2i32 - s3i32
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s3i32 = l7
	s2i32 = s2i32 * s3i32
	s3i32 = 128
	s2i32 = s2i32 + s3i32
	l7 = s2i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = l7
	s2i32 = s2i32 + s3i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 7
	s2i32 = l1
	s3i32 = 0
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l7 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l4
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l0
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl17
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l9
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l6
	l9 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl16
	}
	goto lbl1
lbl14:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l6 = s0i32
lbl18:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l0
	l4 = s0i32
	s0i32 = l8
	l1 = s0i32
lbl19:
	s0i32 = l5
	s1i32 = l5
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l7 = s1i32
	s2i32 = l1
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	l12 = s2i32
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = l12
	s2i32 = s2i32 * s3i32
	s3i32 = 128
	s2i32 = s2i32 + s3i32
	l7 = s2i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = l7
	s2i32 = s2i32 + s3i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl19
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l9
	l3 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl18
	}
	goto lbl1
lbl13:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l6 = s0i32
lbl20:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l0
	l4 = s0i32
	s0i32 = l8
	l1 = s0i32
lbl21:
	s0i32 = l5
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 24
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	l7 = s1i32
	s2i32 = l5
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	l12 = s2i32
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	s3i32 = l12
	s2i32 = s2i32 * s3i32
	s3i32 = 128
	s2i32 = s2i32 + s3i32
	l7 = s2i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = l7
	s2i32 = s2i32 + s3i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl21
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l9
	l3 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl20
	}
	goto lbl1
lbl12:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	l12 = s0i32
lbl22:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l4
	l1 = s0i32
	s0i32 = l8
	l0 = s0i32
lbl23:
	s0i32 = l5
	s1i32 = l0
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
	l6 = s1i32
	s2i32 = 31
	s1i32 = s1i32 & s2i32
	l7 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l7
	s3i32 = 2
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 | s2i32
	s2i32 = l6
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 248
	s2i32 = s2i32 & s3i32
	s3i32 = l6
	s4i32 = 13
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = 5
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 63
	s2i32 = s2i32 & s3i32
	l6 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = l6
	s4i32 = 4
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 + s2i32
	s2i32 = 3
	s1i32 = i32DivU(s1i32, s2i32)
	l6 = s1i32
	s2i32 = l5
	s2i32 = int32(ctx.Mem[int(s2i32+0)])
	l7 = s2i32
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = 255
	s2i32 = s2i32 & s3i32
	s3i32 = l7
	s2i32 = s2i32 * s3i32
	s3i32 = 128
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = l6
	s2i32 = s2i32 + s3i32
	s3i32 = 8
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 - s2i32
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l1
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l6 = s0i32
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	if s0i32 != 0 {
		goto lbl23
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l9
	l3 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl22
	}
	goto lbl1
lbl11:
	s0i32 = l11
	s1i32 = 184
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = 4012
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = 3968
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 3942
	s1i32 = l11
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl9:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl24
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l11
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l6 = s2i32
	s3i32 = l11
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+60)]))
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl27
	case 1:
		goto lbl24
	case 2:
		goto lbl26
	case 3:
		goto lbl25
	default:
		goto lbl28
	}
lbl28:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l9
	s0i32 = s0i32 - s1i32
	l12 = s0i32
lbl29:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = l9
	l0 = s0i32
	s0i32 = l8
	l4 = s0i32
	s0i32 = 7
	l1 = s0i32
lbl30:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = l1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 0
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 7
	s2i32 = l1
	s3i32 = 0
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l7 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l4
	s1i32 = l4
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	s2i32 = l7
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l0
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l0
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl30
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l6
	l3 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl29
	}
	goto lbl1
lbl27:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l6 = s0i32
lbl32:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l0
	l4 = s0i32
	s0i32 = l8
	l1 = s0i32
lbl33:
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l5
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l7
		s3i32 = 255
		s2i32 = s2i32 ^ s3i32
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl33
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l9
	l3 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl32
	}
	goto lbl1
lbl26:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l6 = s0i32
lbl35:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l0
	l4 = s0i32
	s0i32 = l8
	l1 = s0i32
lbl36:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 24
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l5
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = 256
		s3i32 = l7
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l4
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	s0i32 = l4
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl36
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s0i32 = l9
	l3 = s0i32
	s0i32 = l1
	if s0i32 != 0 {
		goto lbl35
	}
	goto lbl1
lbl25:
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l6
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	l12 = s0i32
lbl38:
	s0i32 = l3
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l9 = s0i32
	s0i32 = l4
	l1 = s0i32
	s0i32 = l8
	l0 = s0i32
lbl39:
	s0i32 = l0
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l6 = s0i32
	s1i32 = 31
	s0i32 = s0i32 & s1i32
	l7 = s0i32
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l7
	s2i32 = 2
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s0i32 = s0i32 | s1i32
	s1i32 = l6
	s2i32 = 8
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 248
	s1i32 = s1i32 & s2i32
	s2i32 = l6
	s3i32 = 13
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	s2i32 = 5
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 63
	s1i32 = s1i32 & s2i32
	l6 = s1i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l6
	s3i32 = 4
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = 3
	s0i32 = i32DivU(s0i32, s1i32)
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	l6 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l5
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = 256
		s3i32 = l6
		s2i32 = s2i32 - s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
	s0i32 = l0
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l1
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l6 = s0i32
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	if s0i32 != 0 {
		goto lbl39
	}
	s0i32 = l8
	s1i32 = l10
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s0i32 = l5
	s1i32 = l12
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = 1
	l7 = s0i32
	s0i32 = l3
	s1i32 = 1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l0 = s0i32
	s0i32 = l9
	l3 = s0i32
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl38
	}
	goto lbl1
lbl24:
	s0i32 = l11
	s1i32 = 217
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = 4012
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = 3968
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = 3942
	s1i32 = l11
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl8:
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l17 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l18 = s1i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l10 = s0i32
	s0i32 = l0
	s1i64 = l18
	s1i32 = int32(s1i64)
	s2i64 = l17
	s2i32 = int32(s2i64)
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0i32
	s0i32 = l0
	s0i32 = f118(ctx, s0i32)
	l3 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s2i32 = 0
		s1i32 = f203(ctx, s1i32, s2i32)
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l3 = s0i32
		s1i32 = 4
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l12
		s1i32 = l11
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		s2i32 = l11
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+60)]))
		s3i32 = l10
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl46
		case 1:
			goto lbl0
		case 2:
			goto lbl45
		case 3:
			goto lbl44
		default:
			goto lbl47
		}
	lbl47:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l10
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l14 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l15 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0i32
	lbl48:
		s0i32 = l8
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l10 = s0i32
		s0i32 = l6
		l3 = s0i32
		s0i32 = l9
		l1 = s0i32
		s0i32 = 7
		l0 = s0i32
	lbl49:
		s0i32 = l5
		s1i32 = l4
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = 256
		s3i32 = 1
		s4i32 = l1
		s4i32 = int32(ctx.Mem[int(s4i32+0)])
		s5i32 = l0
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s5i32 = 1
		s4i32 = s4i32 & s5i32
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		s1i32 = 7
		s2i32 = l0
		s3i32 = 0
		if s2i32 > s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		l16 = s2i32
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l0 = s0i32
		s0i32 = l1
		s1i32 = l1
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = l16
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l1 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l16 = s0i32
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l16
		if s0i32 != 0 {
			goto lbl49
		}
		s0i32 = l4
		s1i32 = l13
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l9
		s1i32 = l15
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l5
		s1i32 = l14
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l8
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l10
		l8 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl48
		}
		goto lbl43
	lbl46:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l9 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l10
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l10 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l14 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
	lbl50:
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l9
		l1 = s0i32
		s0i32 = l8
		l0 = s0i32
	lbl51:
		s0i32 = l5
		s1i32 = l4
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l1
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l15 = s0i32
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l15
		if s0i32 != 0 {
			goto lbl51
		}
		s0i32 = l4
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l8
		s1i32 = l14
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l5
		s1i32 = l13
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l6
		l3 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl50
		}
		goto lbl43
	lbl45:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0i32 = s0i32 - s1i32
		l3 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l9 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l10
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l10 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l9
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l14 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l8 = s0i32
	lbl52:
		s0i32 = l3
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l9
		l1 = s0i32
		s0i32 = l8
		l0 = s0i32
	lbl53:
		s0i32 = l5
		s1i32 = l4
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+3)])
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l1
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l15 = s0i32
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l15
		if s0i32 != 0 {
			goto lbl53
		}
		s0i32 = l4
		s1i32 = l10
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l8
		s1i32 = l14
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s0i32 = l5
		s1i32 = l13
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l3
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l6
		l3 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl52
		}
		goto lbl43
	lbl44:
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0i32 = s0i32 - s1i32
		l8 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0i32 = s0i32 - s1i32
		l6 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl43
		}
		s0i32 = l10
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l13 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		s1i32 = l6
		s0i32 = s0i32 - s1i32
		l14 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l15 = s0i32
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0i32
	lbl54:
		s0i32 = l8
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l6
		l0 = s0i32
		s0i32 = l9
		l3 = s0i32
	lbl55:
		s0i32 = l5
		s1i32 = l4
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l3
		s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l10 = s2i32
		s3i32 = 31
		s2i32 = s2i32 & s3i32
		l16 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = l16
		s4i32 = 2
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s2i32 = s2i32 | s3i32
		s3i32 = l10
		s4i32 = 8
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 248
		s3i32 = s3i32 & s4i32
		s4i32 = l10
		s5i32 = 13
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = l10
		s4i32 = 5
		s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
		s4i32 = 63
		s3i32 = s3i32 & s4i32
		l10 = s3i32
		s4i32 = 2
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s4i32 = l10
		s5i32 = 4
		s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
		s3i32 = s3i32 | s4i32
		s2i32 = s2i32 + s3i32
		s3i32 = 3
		s2i32 = i32DivU(s2i32, s3i32)
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		s1i32 = s1i32 * s2i32
		s2i32 = 8
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l0
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l10 = s0i32
		s0i32 = l0
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l10
		if s0i32 != 0 {
			goto lbl55
		}
		s0i32 = l4
		s1i32 = l13
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l9
		s1i32 = l15
		s0i32 = s0i32 + s1i32
		l9 = s0i32
		s0i32 = l5
		s1i32 = l14
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l8
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l0 = s0i32
		s0i32 = l1
		l8 = s0i32
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl54
		}
	lbl43:
		s0i32 = l12
		if s0i32 != 0 {
			goto lbl41
		}
		goto lbl1
	}
	s0i32 = 0
	l7 = s0i32
	s0i32 = l12
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl41:
	s0i32 = l12
	f13(ctx, s0i32)
lbl1:
	s0i32 = l11
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l7
	return s0i32
lbl0:
	s0i32 = l11
	s1i32 = 264
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = 4012
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l11
	s1i32 = 3968
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 3942
	s1i32 = l11
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
