package internal

import (
	"math"
	"unsafe"
)

func f501(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l8 int64
	_ = l8
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s0f32 float32
	_ = s0f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s0i32 = l1
	s0i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])))
	l7 = s0i64
	s0i32 = l1
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+6)])))
	l4 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = l7
	s2i32 = l6
	s2i64 = int64(s2i32)
	l9 = s2i64
	s1i64 = s1i64 + s2i64
	l7 = s1i64
	s2i64 = 2147483647
	s3i64 = l7
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l7 = s1i64
	s2i64 = -2147483647
	s3i64 = l7
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l7 = s1i64
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
	s0i32 = l0
	s1i32 = l4
	s1i64 = int64(s1i32)
	l10 = s1i64
	s2i32 = l5
	s2i64 = int64(uint32(s2i32))
	s1i64 = s1i64 + s2i64
	l8 = s1i64
	s2i64 = 2147483647
	s3i64 = l8
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l8 = s1i64
	s2i64 = -2147483647
	s3i64 = l8
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l8 = s1i64
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+24)])
	l4 = s0i32
	s1i32 = 5
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l1 = s0i32
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl1
		case 1:
			goto lbl1
		case 2:
			goto lbl2
		case 3:
			goto lbl4
		case 4:
			goto lbl1
		default:
			goto lbl5
		}
	lbl5:
		s0i32 = l5
		s1i32 = 7
		s0i32 = s0i32 + s1i32
		s1i32 = 3
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		goto lbl0
	lbl4:
		s0i32 = 2
		l1 = s0i32
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 60
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3323
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 3282
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 3256
	s1i32 = l3
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl2:
	s0i32 = 4
	l1 = s0i32
lbl1:
	s0i32 = l1
	s1i32 = l5
	s0i32 = s0i32 * s1i32
lbl0:
	l1 = s0i32
	s0i32 = l0
	s1i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
	s0i32 = l0
	s1i64 = l7
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s2f32 = float32(math.Floor(float64(s2f32)))
	l13 = s2f32
	s3f32 = 2.1474835e+09
	s4f32 = l13
	s5f32 = 2.1474835e+09
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l13 = s2f32
	s3f32 = -2.1474835e+09
	s4f32 = l13
	s5f32 = -2.1474835e+09
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l13 = s2f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2f32 = l13
		s2i32 = int32(math.Trunc(float64(s2f32)))
		goto lbl6
	}
	s2i32 = -2147483648
lbl6:
	s2i64 = int64(s2i32)
	l11 = s2i64
	s1i64 = s1i64 + s2i64
	l7 = s1i64
	s2i64 = 2147483647
	s3i64 = l7
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l7 = s1i64
	s2i64 = -2147483647
	s3i64 = l7
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
	s0i32 = l0
	s1i64 = l8
	s2f32 = l12
	s2f32 = float32(math.Floor(float64(s2f32)))
	l12 = s2f32
	s3f32 = 2.1474835e+09
	s4f32 = l12
	s5f32 = 2.1474835e+09
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l12 = s2f32
	s3f32 = -2.1474835e+09
	s4f32 = l12
	s5f32 = -2.1474835e+09
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2f32 = s2f32
	} else {
		s2f32 = s3f32
	}
	l12 = s2f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2f32 = l12
		s2i32 = int32(math.Trunc(float64(s2f32)))
		goto lbl8
	}
	s2i32 = -2147483648
lbl8:
	s2i64 = int64(s2i32)
	l7 = s2i64
	s1i64 = s1i64 + s2i64
	l8 = s1i64
	s2i64 = 2147483647
	s3i64 = l8
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l8 = s1i64
	s2i64 = -2147483647
	s3i64 = l8
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
	s0i32 = l0
	s1i64 = l9
	s2i64 = l11
	s1i64 = s1i64 + s2i64
	l9 = s1i64
	s2i64 = 2147483647
	s3i64 = l9
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l9 = s1i64
	s2i64 = -2147483647
	s3i64 = l9
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i64)
	s0i32 = l0
	s1i64 = l7
	s2i64 = l10
	s1i64 = s1i64 + s2i64
	l7 = s1i64
	s2i64 = 2147483647
	s3i64 = l7
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l7 = s1i64
	s2i64 = -2147483647
	s3i64 = l7
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i64)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
