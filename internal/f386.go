package internal

import (
	"math"
	"unsafe"
)

func f386(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
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
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1f32 = 0.5
	s2f32 = 0.5
	s3i32 = l1
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+260)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, float32, float32, int32))(table[s4i32].f()))(ctx, s0i32, s1f32, s2f32, s3i32)
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0f32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		s2f32 = 0
		if s1f32 > s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		l2 = s1i32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s2f32 = 0
		if s1f32 > s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		goto lbl1
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+284)]))
	s2i32 = 1
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l2 = s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+280)]))
	s2i32 = 1
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
lbl1:
	l4 = s1i32
	s1f32 = float32(s1i32)
	s2f32 = 1.5258789e-05
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 1.0737418e+09
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0f32
	s1i32 = l2
	s1f32 = float32(s1i32)
	s2f32 = 1.5258789e-05
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 1.0737418e+09
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l6
	s1f32 = 4.2949673e+09
	s0f32 = s0f32 * s1f32
	l6 = s0f32
	s1f32 = 9.2233715e+18
	s2f32 = l6
	s3f32 = 9.2233715e+18
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l6 = s0f32
	s0i32 = l2
	s0i64 = int64(s0i32)
	s1i64 = 16
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	l5 = s0i64
	s0i32 = l0
	s1f32 = l7
	s2f32 = 4.2949673e+09
	s1f32 = s1f32 * s2f32
	l7 = s1f32
	s2f32 = 9.2233715e+18
	s3f32 = l7
	s4f32 = 9.2233715e+18
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l7 = s1f32
	s2f32 = -9.2233715e+18
	s3f32 = l7
	s4f32 = -9.2233715e+18
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l7 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 9.223372e+18
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l7
		s1i64 = int64(math.Trunc(float64(s1f32)))
		goto lbl3
	}
	s1i64 = -9223372036854775808
lbl3:
	s2i64 = l5
	s1i64 = s1i64 - s2i64
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)])) = uint32(s1i64)
	s0i32 = l4
	s0i64 = int64(s0i32)
	s1i64 = 16
	s0i64 = s0i64 << (uint64(s1i64) & 63)
	l5 = s0i64
	s0i32 = l0
	s1f32 = l6
	s2f32 = -9.2233715e+18
	s3f32 = l6
	s4f32 = -9.2233715e+18
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l6 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 9.223372e+18
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l6
		s1i64 = int64(math.Trunc(float64(s1f32)))
		goto lbl5
	}
	s1i64 = -9223372036854775808
lbl5:
	s2i64 = l5
	s1i64 = s1i64 - s2i64
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint32(s1i64)
	s0i32 = 1
	l3 = s0i32
lbl0:
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l3
	return s0i32
}
