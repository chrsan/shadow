package internal

import (
	"math"
	"unsafe"
)

func f127(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
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
	var s8i32 int32
	_ = s8i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = l0
	s1i32 = 14
	s2i32 = -1
	s3i32 = -1
	s4i32 = -1
	s5i32 = 1
	s6i32 = l1
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s5f32 = float32(s5i32)
	s6f32 = -1
	s5f32 = s5f32 + s6f32
	s5i32 = int32(math.Float32bits(s5f32))
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l3 = s0i32
	s1i32 = l1
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l4 = s0i32
	s0i32 = l3
	s1i32 = l2
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		l6 = s0f32
		s0i32 = l4
		s1i32 = 14
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s1i32 = 14
		s2i32 = -1
		s3i32 = -1
		s4i32 = -1
		s5f32 = l6
		s6i32 = l3
		s7i32 = l1
		s8i32 = 24
		s7i32 = s7i32 * s8i32
		s6i32 = s6i32 + s7i32
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+16)]))
		s5f32 = s5f32 * s6f32
		s5i32 = int32(math.Float32bits(s5f32))
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 14
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = l1
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1f32 = 1
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	goto lbl1
lbl2:
	s0f32 = l6
	s1f32 = 1
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	l2 = s0i32
	goto lbl0
lbl1:
	s0i32 = l0
	s1i32 = 21
	s2i32 = l2
	s3i32 = l1
	s4i32 = -1
	s5i32 = 0
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l2 = s0i32
lbl0:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l2
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 14
		s2i32 = -1
		s3i32 = -1
		s4i32 = -1
		s5i32 = l1
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
		s5f32 = float32(math.RoundToEven(float64(s5f32)))
		l6 = s5f32
		s5f32 = float32(math.Abs(float64(s5f32)))
		s6f32 = 2.1474836e+09
		if s5f32 < s6f32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			s5f32 = l6
			s5i32 = int32(math.Trunc(float64(s5f32)))
			goto lbl5
		}
		s5i32 = -2147483648
	lbl5:
		s6i32 = 0
		s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
		return s0i32
	}
	s0i32 = l0
	s1i32 = 42
	s2i32 = l2
	s3i32 = -1
	s4i32 = -1
	s5i32 = 0
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	return s0i32
}
