package internal

import (
	"math"
	"unsafe"
)

func f1144(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		l6 = s0i32
	lbl1:
		s0i32 = l1
		s1i32 = l5
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0f32
		s0i32 = l4
		s1i32 = l7
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l10 = s1f32
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
		l8 = s2f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l4
		s1f32 = l9
		s2f32 = l8
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l4
		s1f32 = l10
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l4
		s1f32 = l9
		s2f32 = l8
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l4
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s0i32 = f135(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			s2f32 = 65536
			s1f32 = s1f32 * s2f32
			l8 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l8
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl3
			}
			s1i32 = -2147483648
		lbl3:
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			s2f32 = 65536
			s1f32 = s1f32 * s2f32
			l8 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l8
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl5
			}
			s1i32 = -2147483648
		lbl5:
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
			s2f32 = 65536
			s1f32 = s1f32 * s2f32
			l8 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l8
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl7
			}
			s1i32 = -2147483648
		lbl7:
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l4
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
			s2f32 = 65536
			s1f32 = s1f32 * s2f32
			l8 = s1f32
			s2f32 = 2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s2f32 = -2.1474835e+09
			s3f32 = l8
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
			l8 = s1f32
			s1f32 = float32(math.Abs(float64(s1f32)))
			s2f32 = 2.1474836e+09
			if s1f32 < s2f32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				s1f32 = l8
				s1i32 = int32(math.Trunc(float64(s1f32)))
				goto lbl9
			}
			s1i32 = -2147483648
		lbl9:
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
			s0i32 = l4
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			s2i32 = l3
			f1793(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
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
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
