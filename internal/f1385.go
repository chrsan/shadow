package internal

import (
	"unsafe"
)

func f1385(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = ctx.g0
	s1i32 = 96
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s0i32 = f65(ctx, s0i32)
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 2
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			l8 = s1f32
			s2i32 = l3
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			l9 = s2f32
			s3f32 = l9
			s4f32 = l8
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
			l10 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = s1f32
			s0i32 = l5
			s1f32 = l8
			s2f32 = l9
			s3f32 = l8
			s4f32 = l9
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
			l11 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = s1f32
			s0i32 = l5
			s1i32 = l3
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			l8 = s1f32
			s2i32 = l3
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l9 = s2f32
			s3f32 = l9
			s4f32 = l8
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
			l12 = s1f32
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = s1f32
			s0i32 = l5
			s1f32 = l8
			s2f32 = l9
			s3f32 = l8
			s4f32 = l9
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
			*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = s1f32
			goto lbl2
		}
		s0i32 = l5
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l2
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l12 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
		l10 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l8 = s0f32
		s0i32 = l5
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		l11 = s0f32
	lbl2:
		s0f32 = l11
		s1f32 = 0
		s0f32 = s0f32 * s1f32
		s1f32 = l8
		s0f32 = s0f32 * s1f32
		s1f32 = l10
		s0f32 = s0f32 * s1f32
		s1f32 = l12
		s0f32 = s0f32 * s1f32
		l8 = s0f32
		s1f32 = l8
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l5
		s1i32 = 80
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l0
		s1i32 = l4
		s2i32 = l5
		s3i32 = 80
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = 8
		s3i32 = s3i32 + s4i32
		s4i32 = 1
		s1i32 = f64(ctx, s1i32, s2i32, s3i32, s4i32)
		s0i32 = f70(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l7 = s0i32
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l4
	s3i32 = 0
	s4i32 = l6
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l4 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl6:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l6
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+56)]))
	s5i32 = l6
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+76)]))
	if int(s5i32) < 0 || int(s5i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s5i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s5i32].numParams != 5 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	s0i32 = l0
	if s0i32 != 0 {
		goto lbl6
	}
lbl5:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
lbl0:
	s0i32 = l5
	s1i32 = 96
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
