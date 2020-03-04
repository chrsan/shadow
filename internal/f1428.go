package internal

import (
	"math"
	"unsafe"
)

func f1428(ctx *Context, l0 int32, l1 float64, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 float64
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
	var s5i32 int32
	_ = s5i32
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i32 = l3
	s2i32 = 88
	s1i32 = s1i32 + s2i32
	s2i32 = l3
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s4i32 = 72
	s3i32 = s3i32 + s4i32
	s4i32 = l3
	s5i32 = -64
	s4i32 = s4i32 - s5i32
	f572(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l3
	s1i32 = l3
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	s2f64 = l1
	s1f64 = s1f64 - s2f64
	l7 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = s1f64
	s0i32 = l3
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
	s1i32 = l3
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	s2i32 = l3
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+72)]))
	s3f64 = l7
	s4i32 = l2
	s0i32 = f571(ctx, s0f64, s1f64, s2f64, s3f64, s4i32)
	l4 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l3
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l2
	s3i32 = l5
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	f573(ctx, s0i32, s1i32, s2f64)
	s0i32 = l3
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)]))
	s1f64 = l1
	s0f64 = s0f64 - s1f64
	s0f64 = math.Abs(s0f64)
	s1f64 = 1.1920928955078125e-07
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s2i32 = l6
		s3i32 = l3
		s2i32 = f568(ctx, s2i32, s3i32)
		s3f64 = l1
		s4i32 = 1
		s5i32 = l2
		s0i32 = f570(ctx, s0i32, s1i32, s2i32, s3f64, s4i32, s5i32)
		l4 = s0i32
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l4
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl0:
	s0i32 = l3
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
