package internal

import (
	"unsafe"
)

func f2110(ctx *Context, l0 int32) float32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
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
	var s7f32 float32
	_ = s7f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	ctx.g0 = s0i32
	s0f32 = -1
	l7 = s0f32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l10 = s1f32
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l4 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	l11 = s3f32
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l6 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	l13 = s1f32
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l5 = s2f32
	s3i32 = l0
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+28)]))
	l9 = s3f32
	if s2f32 == s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l4
	s1f32 = l5
	s0f32 = s0f32 - s1f32
	s1f32 = l13
	s2f32 = l6
	s1f32 = s1f32 - s2f32
	l12 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l8
	s2f32 = l6
	s1f32 = s1f32 - s2f32
	s2f32 = l9
	s3f32 = l5
	s2f32 = s2f32 - s3f32
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l12
	s2f32 = l11
	s3f32 = l5
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s2f32 = l14
	s3f32 = l10
	s4f32 = l6
	s3f32 = s3f32 - s4f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l5
	s1f32 = l4
	s0f32 = s0f32 - s1f32
	s1f32 = l10
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	l5 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l6
	s2f32 = l8
	s1f32 = s1f32 - s2f32
	s2f32 = l11
	s3f32 = l4
	s2f32 = s2f32 - s3f32
	l6 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	s1f32 = l9
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	s2f32 = l5
	s1f32 = s1f32 * s2f32
	s2f32 = l13
	s3f32 = l8
	s2f32 = s2f32 - s3f32
	s3f32 = l6
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s0i32 = f446(ctx, s0i32, s1i32)
	l3 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl2:
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0f32
		s1f32 = 0
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0f32 = l7
		s1f32 = 1
		if s0f32 >= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1f32
		s0f32 = s0f32 - s1f32
		l11 = s0f32
		s1f32 = l7
		s2f32 = l5
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
		l8 = s3f32
		s4f32 = l4
		s5f32 = l4
		s4f32 = s4f32 + s5f32
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 + s3f32
		l6 = s2f32
		s3f32 = l6
		s2f32 = s2f32 + s3f32
		s3f32 = l7
		s4i32 = l0
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
		l13 = s4f32
		s5f32 = l4
		s6f32 = l8
		s5f32 = s5f32 - s6f32
		s6f32 = 3
		s5f32 = s5f32 * s6f32
		s4f32 = s4f32 + s5f32
		s5f32 = l5
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l5 = s0f32
		s1f32 = l5
		s0f32 = s0f32 * s1f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		l5 = s1f32
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l9 = s2f32
		s1f32 = s1f32 - s2f32
		l10 = s1f32
		s2f32 = l7
		s3f32 = l9
		s4i32 = l0
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
		l6 = s4f32
		s5f32 = l5
		s6f32 = l5
		s5f32 = s5f32 + s6f32
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 + s4f32
		l12 = s3f32
		s4f32 = l12
		s3f32 = s3f32 + s4f32
		s4f32 = l7
		s5i32 = l0
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+28)]))
		l12 = s5f32
		s6f32 = l5
		s7f32 = l6
		s6f32 = s6f32 - s7f32
		s7f32 = 3
		s6f32 = s6f32 * s7f32
		s5f32 = s5f32 + s6f32
		s6f32 = l9
		s5f32 = s5f32 - s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 + s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		l9 = s1f32
		s2f32 = l9
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		s1f32 = l11
		s2f32 = l11
		s1f32 = s1f32 * s2f32
		s2f32 = l10
		s3f32 = l10
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l8
		s3f32 = l4
		s2f32 = s2f32 - s3f32
		l4 = s2f32
		s3f32 = l4
		s2f32 = s2f32 * s3f32
		s3f32 = l6
		s4f32 = l5
		s3f32 = s3f32 - s4f32
		l4 = s3f32
		s4f32 = l4
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = l13
		s3f32 = l8
		s2f32 = s2f32 - s3f32
		l4 = s2f32
		s3f32 = l4
		s2f32 = s2f32 * s3f32
		s3f32 = l12
		s4f32 = l6
		s3f32 = s3f32 - s4f32
		l4 = s3f32
		s4f32 = l4
		s3f32 = s3f32 * s4f32
		s2f32 = s2f32 + s3f32
		s1f32 = s1f32 + s2f32
		s2f32 = 1e-08
		s1f32 = s1f32 * s2f32
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	lbl3:
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0f32 = -1
	l7 = s0f32
lbl0:
	s0i32 = l1
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0f32 = l7
	return s0f32
}
