package internal

import (
	"math"
	"unsafe"
)

func f956(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 float32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l1 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = f24(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l1
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = f24(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = 2
	l2 = s0i32
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = 1
	l4 = s0f32
	s0f32 = 1
	l3 = s0f32
	s0i32 = l1
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0f32
	s0i32 = l1
	s1i32 = 4
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l6
		s0f32 = float32(math.Abs(float64(s0f32)))
		l5 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s1f32 = float32(math.Abs(float64(s1f32)))
		l3 = s1f32
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0f32 = l3
			l4 = s0f32
			s0f32 = l5
			l3 = s0f32
			goto lbl3
		}
		s0f32 = l5
		l4 = s0f32
		goto lbl3
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0f32
	s1f32 = l7
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l8 = s1f32
	s2f32 = l8
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l5 = s0f32
	s0f32 = l6
	s1f32 = l6
	s0f32 = s0f32 * s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l3 = s1f32
	s2f32 = l3
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l6
	s1f32 = l7
	s0f32 = s0f32 * s1f32
	s1f32 = l3
	s2f32 = l8
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l3 = s0f32
	s1f32 = l3
	s0f32 = s0f32 * s1f32
	l7 = s0f32
	s1f32 = 5.9604645e-08
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l4
		s1f32 = l5
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0f32 = l4
			l3 = s0f32
			s0f32 = l5
			l4 = s0f32
			goto lbl6
		}
		s0f32 = l5
		l3 = s0f32
		goto lbl6
	}
	s0f32 = l4
	s1f32 = l5
	s0f32 = s0f32 + s1f32
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	l8 = s0f32
	s1f32 = l4
	s2f32 = l5
	s1f32 = s1f32 - s2f32
	l3 = s1f32
	s2f32 = l3
	s1f32 = s1f32 * s2f32
	s2f32 = l7
	s3f32 = 4
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	s1f32 = float32(math.Sqrt(float64(s1f32)))
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l3 = s1f32
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s0f32 = l8
	s1f32 = l3
	s0f32 = s0f32 - s1f32
	l3 = s0f32
lbl6:
	s0f32 = l3
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l4
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l3
	s1f32 = 0
	s0f32 = float32(math.Max(float64(s0f32), float64(s1f32)))
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	l3 = s0f32
	s0f32 = l4
	s0f32 = float32(math.Sqrt(float64(s0f32)))
	l4 = s0f32
lbl3:
	s0f32 = l3
	s1f32 = 1
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l4
	s1f32 = 1
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 3
		return s0i32
	}
	s0i32 = l1
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = f24(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l3 = s0f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	s1f32 = l3
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = 0
	l2 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l3 = s0f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	s1f32 = l3
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl11:
	s0i32 = 1
	l2 = s0i32
lbl1:
	s0i32 = l2
	return s0i32
}
