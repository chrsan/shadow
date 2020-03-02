package internal

import (
	"math"
	"unsafe"
)

func f843(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 float32, l5 float32) {
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
	var s7f32 float32
	_ = s7f32
	var s8f32 float32
	_ = s8f32
	s0f32 = l2
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	l4 = s0f32
	s1f32 = l4
	s1f32 = float32(math.Floor(float64(s1f32)))
	s0f32 = s0f32 - s1f32
	l15 = s0f32
	s1f32 = 1
	s2f32 = l3
	s3f32 = 0.5
	s2f32 = s2f32 + s3f32
	l16 = s2f32
	s3f32 = l16
	s3f32 = float32(math.Floor(float64(s3f32)))
	s2f32 = s2f32 - s3f32
	l17 = s2f32
	s1f32 = s1f32 - s2f32
	l13 = s1f32
	s0f32 = s0f32 * s1f32
	l14 = s0f32
	s0f32 = l3
	s1f32 = -0.5
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l18 = s1f32
	s0f32 = f14(ctx, s0f32, s1f32)
	l3 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l3
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl0
	}
	s0i32 = 0
lbl0:
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l7 = s1i32
	s0i32 = s0i32 * s1i32
	l10 = s0i32
	s0f32 = l4
	s1f32 = 0
	s0f32 = f15(ctx, s0f32, s1f32)
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l5 = s1f32
	s0f32 = f14(ctx, s0f32, s1f32)
	l3 = s0f32
	s1f32 = 4.2949673e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l3
	s2f32 = 0
	if s1f32 >= s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0f32 = l3
		s0i32 = int32(uint32(math.Trunc(float64(s0f32))))
		goto lbl2
	}
	s0i32 = 0
lbl2:
	l11 = s0i32
	s0f32 = l14
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1i32
	s2i32 = l10
	s3i32 = l11
	s2i32 = s2i32 + s3i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1i32
	s2i32 = 24
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s1f32 = float32(uint32(s1i32))
	s2f32 = 0.003921569
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 1
	s2f32 = l15
	s1f32 = s1f32 - s2f32
	l3 = s1f32
	s2f32 = l13
	s1f32 = s1f32 * s2f32
	l13 = s1f32
	s2i32 = l8
	s3i32 = l10
	s4f32 = l2
	s5f32 = -0.5
	s4f32 = s4f32 + s5f32
	s5f32 = 0
	s4f32 = f15(ctx, s4f32, s5f32)
	s5f32 = l5
	s4f32 = f14(ctx, s4f32, s5f32)
	l2 = s4f32
	s5f32 = 4.2949673e+09
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	s5f32 = l2
	s6f32 = 0
	if s5f32 >= s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	s4i32 = s4i32 & s5i32
	if s4i32 != 0 {
		s4f32 = l2
		s4i32 = int32(uint32(math.Trunc(float64(s4f32))))
		goto lbl4
	}
	s4i32 = 0
lbl4:
	l12 = s4i32
	s3i32 = s3i32 + s4i32
	s4i32 = 2
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l6 = s2i32
	s3i32 = 24
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s2f32 = float32(uint32(s2i32))
	s3f32 = 0.003921569
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = 0
	s1f32 = s1f32 + s2f32
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2f32 = l13
	s3i32 = l6
	s4i32 = 255
	s3i32 = s3i32 & s4i32
	s3f32 = float32(uint32(s3i32))
	s4f32 = 0.003921569
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = 0
	s2f32 = s2f32 + s3f32
	s3f32 = l14
	s4i32 = l9
	s5i32 = 255
	s4i32 = s4i32 & s5i32
	s4f32 = float32(uint32(s4i32))
	s5f32 = 0.003921569
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l17
	s4f32 = l3
	s3f32 = s3f32 * s4f32
	l4 = s3f32
	s4i32 = l8
	s5f32 = l16
	s6f32 = 0
	s5f32 = f15(ctx, s5f32, s6f32)
	s6f32 = l18
	s5f32 = f14(ctx, s5f32, s6f32)
	l3 = s5f32
	s6f32 = 4.2949673e+09
	if s5f32 < s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	s6f32 = l3
	s7f32 = 0
	if s6f32 >= s7f32 {
		s6i32 = 1
	} else {
		s6i32 = 0
	}
	s5i32 = s5i32 & s6i32
	if s5i32 != 0 {
		s5f32 = l3
		s5i32 = int32(uint32(math.Trunc(float64(s5f32))))
		goto lbl6
	}
	s5i32 = 0
lbl6:
	s6i32 = l7
	s5i32 = s5i32 * s6i32
	l0 = s5i32
	s6i32 = l12
	s5i32 = s5i32 + s6i32
	s6i32 = 2
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l7 = s4i32
	s5i32 = 255
	s4i32 = s4i32 & s5i32
	s4f32 = float32(uint32(s4i32))
	s5f32 = 0.003921569
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l15
	s4f32 = l17
	s3f32 = s3f32 * s4f32
	l3 = s3f32
	s4i32 = l8
	s5i32 = l0
	s6i32 = l11
	s5i32 = s5i32 + s6i32
	s6i32 = 2
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	l0 = s4i32
	s5i32 = 255
	s4i32 = s4i32 & s5i32
	s4f32 = float32(uint32(s4i32))
	s5f32 = 0.003921569
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 + s3f32
	s3f32 = l13
	s4i32 = l6
	s5i32 = 8
	s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
	s5i32 = 255
	s4i32 = s4i32 & s5i32
	s4f32 = float32(uint32(s4i32))
	s5f32 = 0.003921569
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 * s4f32
	s4f32 = 0
	s3f32 = s3f32 + s4f32
	s4f32 = l14
	s5i32 = l9
	s6i32 = 8
	s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
	s6i32 = 255
	s5i32 = s5i32 & s6i32
	s5f32 = float32(uint32(s5i32))
	s6f32 = 0.003921569
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l4
	s5i32 = l7
	s6i32 = 8
	s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
	s6i32 = 255
	s5i32 = s5i32 & s6i32
	s5f32 = float32(uint32(s5i32))
	s6f32 = 0.003921569
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l3
	s5i32 = l0
	s6i32 = 8
	s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
	s6i32 = 255
	s5i32 = s5i32 & s6i32
	s5f32 = float32(uint32(s5i32))
	s6f32 = 0.003921569
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 * s5f32
	s3f32 = s3f32 + s4f32
	s4f32 = l13
	s5i32 = l6
	s6i32 = 16
	s5i32 = int32(uint32(s5i32) >> (uint32(s6i32) & 31))
	s6i32 = 255
	s5i32 = s5i32 & s6i32
	s5f32 = float32(uint32(s5i32))
	s6f32 = 0.003921569
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 * s5f32
	s5f32 = 0
	s4f32 = s4f32 + s5f32
	s5f32 = l14
	s6i32 = l9
	s7i32 = 16
	s6i32 = int32(uint32(s6i32) >> (uint32(s7i32) & 31))
	s7i32 = 255
	s6i32 = s6i32 & s7i32
	s6f32 = float32(uint32(s6i32))
	s7f32 = 0.003921569
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l4
	s6i32 = l7
	s7i32 = 16
	s6i32 = int32(uint32(s6i32) >> (uint32(s7i32) & 31))
	s7i32 = 255
	s6i32 = s6i32 & s7i32
	s6f32 = float32(uint32(s6i32))
	s7f32 = 0.003921569
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l3
	s6i32 = l0
	s7i32 = 16
	s6i32 = int32(uint32(s6i32) >> (uint32(s7i32) & 31))
	s7i32 = 255
	s6i32 = s6i32 & s7i32
	s6f32 = float32(uint32(s6i32))
	s7f32 = 0.003921569
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	s5f32 = l3
	s6i32 = l0
	s7i32 = 24
	s6i32 = int32(uint32(s6i32) >> (uint32(s7i32) & 31))
	s6f32 = float32(uint32(s6i32))
	s7f32 = 0.003921569
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 * s6f32
	s6f32 = l4
	s7i32 = l7
	s8i32 = 24
	s7i32 = int32(uint32(s7i32) >> (uint32(s8i32) & 31))
	s7f32 = float32(uint32(s7i32))
	s8f32 = 0.003921569
	s7f32 = s7f32 * s8f32
	s6f32 = s6f32 * s7f32
	s7f32 = l2
	s6f32 = s6f32 + s7f32
	s5f32 = s5f32 + s6f32
	s6i32 = l1
	s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	if int(s6i32) < 0 || int(s6i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s6i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s6i32].numParams != 6 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, float32, float32, float32, float32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2f32, s3f32, s4f32, s5f32)
}
