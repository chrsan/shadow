package internal

import (
	"math"
	"unsafe"
)

func f399(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l11 float32
	_ = l11
	var l12 float32
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
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2f32 = 0.5
	f339(ctx, s0i32, s1i32, s2f32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0f32 = s0f32 - s1f32
	l9 = s0f32
	s1f32 = l9
	s0f32 = s0f32 * s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l9 = s1f32
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s2f32 = l11
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l11 = s1f32
	s2f32 = l11
	s1f32 = s1f32 * s2f32
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
		s0i32 = l4
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		l10 = s1f32
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 - s2f32
		l8 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l4
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l12 = s1f32
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 - s2f32
		l9 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0f32 = l8
		s1f32 = l8
		s0f32 = s0f32 * s1f32
		s1f32 = l9
		s2f32 = l9
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l11 = s0f32
		s1f32 = l10
		s2i32 = l3
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		s1f32 = s1f32 - s2f32
		l10 = s1f32
		s2f32 = l10
		s1f32 = s1f32 * s2f32
		s2f32 = l12
		s3i32 = l3
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
		s2f32 = s2f32 - s3f32
		l12 = s2f32
		s3f32 = l12
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		l13 = s1f32
		if s0f32 > s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0f32 = l12
			l9 = s0f32
			s0f32 = l10
			l8 = s0f32
			s0f32 = l13
			l11 = s0f32
			goto lbl3
		}
		s0i32 = l4
		s1f32 = l12
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l4
		s1f32 = l10
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	lbl3:
		s0i32 = l4
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		s1f32 = l11
		s0i32 = f113(ctx, s0i32, s1f32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 2
			l5 = s0i32
			goto lbl0
		}
		s0i32 = 2
		l5 = s0i32
		s0f32 = l8
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s0f32 = s0f32 * s1f32
		s1f32 = l9
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = 0
		if s0f32 > s1f32 {
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
			goto lbl1
		}
		goto lbl0
	}
	s0f32 = l11
	s1f32 = l8
	s0f32 = s0f32 + s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l10 = s1f32
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l12 = s2f32
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	l13 = s3f32
	s4f32 = l12
	s5f32 = l13
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
	l14 = s2f32
	s3f32 = l10
	s4f32 = l14
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
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l8
	s1f32 = l11
	s0f32 = s0f32 - s1f32
	s1i32 = l1
	s2i32 = 16
	s3f32 = l13
	s4f32 = l12
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	l6 = s3i32
	s4i32 = l1
	s5i32 = l6
	s4i32 = s4i32 + s5i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s5f32 = l10
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s1i32 = s1i32 + s2i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l11
	s1f32 = l9
	s0f32 = s0f32 + s1f32
	s1i32 = l1
	s2i32 = 20
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l14 = s1f32
	s2i32 = l1
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	l7 = s2i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l15 = s2f32
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l16 = s3f32
	s4f32 = l15
	s5f32 = l16
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
	l17 = s2f32
	s3f32 = l14
	s4f32 = l17
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
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l9
	s1f32 = l11
	s0f32 = s0f32 - s1f32
	s1i32 = l6
	s2i32 = l7
	s3i32 = l1
	s4i32 = 4
	s3i32 = s3i32 + s4i32
	s4f32 = l16
	s5f32 = l15
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l6 = s2i32
	s3i32 = l6
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4f32 = l14
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = l8
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s1f32 = l14
	s2f32 = l9
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l10
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3f32 = l9
	s2f32 = s2f32 - s3f32
	l10 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l11
	s2f32 = l16
	s3f32 = l9
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l13
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	s3f32 = l10
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l13 = s1f32
	s2f32 = l11
	s3f32 = l15
	s4f32 = l9
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s3f32 = l12
	s4f32 = l8
	s3f32 = s3f32 - s4f32
	s4f32 = l10
	s3f32 = s3f32 * s4f32
	s2f32 = s2f32 - s3f32
	l8 = s2f32
	s3f32 = l8
	s2f32 = s2f32 + s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l8
	s2f32 = l13
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s2f32 = l8
	s1f32 = s1f32 + s2f32
	s2f32 = l13
	s3i32 = l4
	s4i32 = 8
	s3i32 = s3i32 + s4i32
	s0i32 = f122(ctx, s0f32, s1f32, s2f32, s3i32)
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	f339(ctx, s0i32, s1i32, s2f32)
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	l8 = s0f32
	s1f32 = l8
	s0f32 = s0f32 * s1f32
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s2f32 = l8
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2f32 = 1
	s3i32 = l4
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4f32 = -0.5
	s3f32 = s3f32 + s4f32
	s3f32 = float32(math.Abs(float64(s3f32)))
	l8 = s3f32
	s4f32 = l8
	s3f32 = s3f32 + s4f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	l8 = s1f32
	s2f32 = l8
	s1f32 = s1f32 * s2f32
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l10 = s1f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l4
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l12 = s1f32
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0f32 = l8
	s1f32 = l8
	s0f32 = s0f32 * s1f32
	s1f32 = l9
	s2f32 = l9
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l11 = s0f32
	s1f32 = l10
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s1f32 = s1f32 - s2f32
	l10 = s1f32
	s2f32 = l10
	s1f32 = s1f32 * s2f32
	s2f32 = l12
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
	s2f32 = s2f32 - s3f32
	l12 = s2f32
	s3f32 = l12
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	l13 = s1f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l12
		l9 = s0f32
		s0f32 = l10
		l8 = s0f32
		s0f32 = l13
		l11 = s0f32
		goto lbl6
	}
	s0i32 = l4
	s1f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l4
	s1f32 = l10
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
lbl6:
	s0i32 = l4
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	s1f32 = l11
	s0i32 = f113(ctx, s0i32, s1f32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 2
		l5 = s0i32
		goto lbl0
	}
	s0i32 = 2
	l5 = s0i32
	s0f32 = l8
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s0f32 = s0f32 * s1f32
	s1f32 = l9
	s2i32 = l4
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = 0
	l5 = s0i32
lbl0:
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l5
	return s0i32
}
