package internal

import (
	"math"
	"unsafe"
)

func f209(ctx *Context, l0 int32) int32 {
	var l1 float32
	_ = l1
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
	s0f32 = 0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1f32
	s2f32 = 1
	s3f32 = l1
	s4f32 = 1
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
	l1 = s1f32
	s2f32 = 255
	s1f32 = s1f32 * s2f32
	s2f32 = l1
	s3f32 = 0
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
	s0f32 = float32(math.RoundToEven(float64(s0f32)))
	l1 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 2.1474836e+09
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l1
		s0i32 = int32(math.Trunc(float64(s0f32)))
		goto lbl0
	}
	s0i32 = -2147483648
lbl0:
	s1i32 = 8
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = 65280
	s0i32 = s0i32 & s1i32
	s1f32 = 0
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l1 = s2f32
	s3f32 = 1
	s4f32 = l1
	s5f32 = 1
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
	l1 = s2f32
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = l1
	s4f32 = 0
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
	s1f32 = float32(math.RoundToEven(float64(s1f32)))
	l1 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l1
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl2
	}
	s1i32 = -2147483648
lbl2:
	s2i32 = 255
	s1i32 = s1i32 & s2i32
	s2f32 = 0
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l1 = s3f32
	s4f32 = 1
	s5f32 = l1
	s6f32 = 1
	if s5f32 < s6f32 {
		s5i32 = 1
	} else {
		s5i32 = 0
	}
	if s5i32 != 0 {
		// s3f32 = s3f32
	} else {
		s3f32 = s4f32
	}
	l1 = s3f32
	s4f32 = 255
	s3f32 = s3f32 * s4f32
	s4f32 = l1
	s5f32 = 0
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
	s2f32 = float32(math.RoundToEven(float64(s2f32)))
	l1 = s2f32
	s2f32 = float32(math.Abs(float64(s2f32)))
	s3f32 = 2.1474836e+09
	if s2f32 < s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		s2f32 = l1
		s2i32 = int32(math.Trunc(float64(s2f32)))
		goto lbl4
	}
	s2i32 = -2147483648
lbl4:
	s3i32 = 16
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 16711680
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	s0i32 = s0i32 | s1i32
	s1f32 = 0
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l1 = s2f32
	s3f32 = 1
	s4f32 = l1
	s5f32 = 1
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
	l1 = s2f32
	s3f32 = 255
	s2f32 = s2f32 * s3f32
	s3f32 = l1
	s4f32 = 0
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
	s1f32 = float32(math.RoundToEven(float64(s1f32)))
	l1 = s1f32
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 2.1474836e+09
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l1
		s1i32 = int32(math.Trunc(float64(s1f32)))
		goto lbl6
	}
	s1i32 = -2147483648
lbl6:
	s2i32 = 24
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 | s1i32
	return s0i32
}
