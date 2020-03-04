package internal

import (
	"math"
	"unsafe"
)

func f248(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int64
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 432
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	ctx.g0 = s0i32
	s0i32 = l9
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
	s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l8 = s1i32
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l7
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l9
		s0i32 = int32(ctx.Mem[int(s0i32+8)])
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+1)])
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+2)])
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+3)])
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		s1i32 = l6
		s1i32 = int32(ctx.Mem[int(s1i32+4)])
		s2i32 = 4
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 | s1i32
		if s0i32 != 0 {
			goto lbl11
		}
	}
	s0i32 = 0
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	s0i32 = l2
	s1i32 = l0
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 3728
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i64 = int64(s2i32)
	s1i64 = s1i64 * s2i64
	l10 = s1i64
	s1i32 = int32(s1i64)
	s2i32 = 0
	s3i64 = l10
	s4i64 = 2147483648
	if uint64(s3i64) < uint64(s4i64) {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l0 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l0
	s3i32 = l5
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
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l6
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl13:
	s0i32 = l1
	s1i32 = l4
	s2i32 = l0
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl13
	}
	goto lbl0
lbl11:
	s0i32 = l7
	s1i32 = 2
	s0i32 = s0i32 | s1i32
	s1i32 = 6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l8
	s1i32 = 2
	s0i32 = s0i32 | s1i32
	s1i32 = 6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l6
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l9
	s0i32 = int32(ctx.Mem[int(s0i32+8)])
	s1i32 = l6
	s1i32 = int32(ctx.Mem[int(s1i32+2)])
	s0i32 = s0i32 | s1i32
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l6
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	if s0i32 != 0 {
		goto lbl14
	}
	s0i32 = l6
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 22112
	s1i32 = 22116
	s2i32 = l7
	s3i32 = l8
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
	s1i32 = 22108
	s2i32 = l3
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = 0
	l3 = s0i32
lbl15:
	s0i32 = l1
	s1i32 = l4
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3i32 = l6
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
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl15
	}
	goto lbl0
lbl14:
	s0i32 = l7
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l8
	s1i32 = -2
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = 18
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l7
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl1
	case 1:
		goto lbl6
	case 2:
		goto lbl5
	case 3:
		goto lbl6
	case 4:
		goto lbl7
	case 5:
		goto lbl7
	case 6:
		goto lbl5
	case 7:
		goto lbl5
	case 8:
		goto lbl5
	case 9:
		goto lbl8
	case 10:
		goto lbl8
	case 11:
		goto lbl16
	case 12:
		goto lbl5
	case 13:
		goto lbl3
	case 14:
		goto lbl5
	case 15:
		goto lbl2
	case 16:
		goto lbl5
	case 17:
		goto lbl4
	default:
		goto lbl5
	}
lbl16:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l8 = s0i32
lbl17:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl19:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s3i32 = 12
		s2i32 = s2i32 | s3i32
		s1i32 = s1i32 + s2i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2f32 = 255
		s1f32 = s1f32 * s2f32
		l11 = s1f32
		s2f32 = 4.2949673e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f32 = l11
		s3f32 = 0
		if s2f32 >= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1f32 = l11
			s1i32 = int32(uint32(math.Trunc(float64(s1f32))))
			goto lbl20
		}
		s1i32 = 0
	lbl20:
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl19
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l6 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l6
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl17
	}
	goto lbl0
lbl10:
	s0i32 = l1
	s1i32 = l4
	s2i32 = l5
	s3i32 = l6
	s2i32 = s2i32 * s3i32
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	goto lbl0
lbl9:
	s0i32 = l9
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+424)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = l5
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 3728
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = i32DivU(s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+428)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+416)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = l2
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 3728
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = i32DivU(s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+420)])) = uint32(s1i32)
	s0i32 = l9
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l1
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l9
	s2i32 = 384
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l9
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	s2i32 = 256
	s3i32 = 256
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l2 = s0i32
	s0i32 = l1
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l9
	s3i32 = 424
	s2i32 = s2i32 + s3i32
	f433(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = l1
	s2i32 = 1941503
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	s4i32 = 2097151
	s3i32 = s3i32 & s4i32
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 1
	s2i32 = s2i32 & s3i32
	f296(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = l0
	f172(ctx, s0i32, s1i32)
	s0i32 = l1
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l9
	s3i32 = 416
	s2i32 = s2i32 + s3i32
	f173(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	f272(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	f43(ctx, s0i32)
	goto lbl0
lbl8:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l8 = s0i32
lbl22:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl24:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+6)])))
		s1f32 = f497(ctx, s1i32)
		s2f32 = 255
		s1f32 = s1f32 * s2f32
		l11 = s1f32
		s2f32 = 4.2949673e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f32 = l11
		s3f32 = 0
		if s2f32 >= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1f32 = l11
			s1i32 = int32(uint32(math.Trunc(float64(s1f32))))
			goto lbl25
		}
		s1i32 = 0
	lbl25:
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl24
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l6 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l6
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl22
	}
	goto lbl0
lbl7:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l6 = s0i32
lbl27:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl29:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 30
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 85
		s1i32 = s1i32 * s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl29
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl27
	}
	goto lbl0
lbl6:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l6 = s0i32
lbl30:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl32:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+3)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl32
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl30
	}
	goto lbl0
lbl5:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = 0
	l4 = s0i32
lbl33:
	s0i32 = l1
	s1i32 = 255
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s0i32 = f27(ctx, s0i32, s1i32, s2i32)
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl33
	}
	goto lbl0
lbl4:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l6 = s0i32
lbl34:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl36:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i64 = int64(ctx.Mem[int(s1i32+7)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i64)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl36
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl34
	}
	goto lbl0
lbl3:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l8 = s0i32
lbl37:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl39:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		s1f32 = f497(ctx, s1i32)
		s2f32 = 255
		s1f32 = s1f32 * s2f32
		l11 = s1f32
		s2f32 = 4.2949673e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2f32 = l11
		s3f32 = 0
		if s2f32 >= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1f32 = l11
			s1i32 = int32(uint32(math.Trunc(float64(s1f32))))
			goto lbl40
		}
		s1i32 = 0
	lbl40:
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl39
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l6 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l8
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l6
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl37
	}
	goto lbl0
lbl2:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l6 = s0i32
lbl42:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl44:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+1)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl44
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl42
	}
	goto lbl0
lbl1:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s0i32 = 0
	l6 = s0i32
lbl45:
	s0i32 = 0
	l0 = s0i32
	s0i32 = l7
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl47:
		s0i32 = l0
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = l0
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
		s2i32 = 15
		s1i32 = s1i32 & s2i32
		l7 = s1i32
		s2i32 = 4
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l7
		s1i32 = s1i32 | s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l7 = s1i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl47
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l8 = s0i32
	}
	s0i32 = l4
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl45
	}
lbl0:
	s0i32 = l9
	s1i32 = 432
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
