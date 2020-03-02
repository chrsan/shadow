package internal

import (
	"math"
	"unsafe"
)

func f1136(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
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
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl1:
		s0i32 = l1
		s1i32 = l4
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s0f32 = float32(math.Floor(float64(s0f32)))
		l7 = s0f32
		s1f32 = 2.1474835e+09
		s2f32 = l7
		s3f32 = 2.1474835e+09
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
		l7 = s0f32
		s1f32 = -2.1474835e+09
		s2f32 = l7
		s3f32 = -2.1474835e+09
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0f32 = s0f32
		} else {
			s0f32 = s1f32
		}
		l7 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 2.1474836e+09
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l7
			s0i32 = int32(math.Trunc(float64(s0f32)))
			goto lbl2
		}
		s0i32 = -2147483648
	lbl2:
		l6 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		s1i32 = l5
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1f32 = float32(math.Floor(float64(s1f32)))
		l7 = s1f32
		s2f32 = 2.1474835e+09
		s3f32 = l7
		s4f32 = 2.1474835e+09
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
		s2f32 = -2.1474835e+09
		s3f32 = l7
		s4f32 = -2.1474835e+09
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
		s2f32 = 2.1474836e+09
		if s1f32 < s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1f32 = l7
			s1i32 = int32(math.Trunc(float64(s1f32)))
			goto lbl4
		}
		s1i32 = -2147483648
	lbl4:
		l5 = s1i32
		s2i32 = l6
		s0i32 = f1901(ctx, s0i32, s1i32, s2i32)
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l5
			s2i32 = l6
			s3i32 = 1
			s4i32 = l3
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+8)]))
			if int(s4i32) < 0 || int(s4i32) >= len(table) {
				panic("table entry out of bounds")
			}
			if table[s4i32].numParams == -1 {
				panic("table entry is nil")
			}
			if table[s4i32].numParams != 4 {
				panic("argument count mismatch")
			}
			(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
		}
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s1i32 = l2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
