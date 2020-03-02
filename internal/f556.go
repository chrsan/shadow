package internal

import (
	"unsafe"
)

func f556(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
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
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	var s7f32 float32
	_ = s7f32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+159)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l11 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
		s2f32 = l11
		s1f32 = s1f32 - s2f32
		s2f32 = 0.95
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l12 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
		s2f32 = l11
		s1f32 = s1f32 - s2f32
		s2f32 = 0.95
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l13 = s0f32
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l14 = s0f32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l7 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+168)]))
	l5 = s2i32
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l8 = s3i32
	s2i32 = i32RemS(s2i32, s3i32)
	l3 = s2i32
	s3i32 = l5
	s4f32 = l14
	s5i32 = l7
	s6i32 = l3
	s7i32 = 3
	s6i32 = s6i32 << (uint32(s7i32) & 31)
	s5i32 = s5i32 + s6i32
	l3 = s5i32
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s4f32 = s4f32 - s5f32
	l11 = s4f32
	s5f32 = l11
	s4f32 = s4f32 * s5f32
	s5i32 = l1
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
	l15 = s5f32
	s6i32 = l3
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	s5f32 = s5f32 - s6f32
	l11 = s5f32
	s6f32 = l11
	s5f32 = s5f32 * s6f32
	s4f32 = s4f32 + s5f32
	l16 = s4f32
	s5f32 = l14
	s6i32 = l7
	s7i32 = l5
	s8i32 = 3
	s7i32 = s7i32 << (uint32(s8i32) & 31)
	s6i32 = s6i32 + s7i32
	l3 = s6i32
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
	s5f32 = s5f32 - s6f32
	l11 = s5f32
	s6f32 = l11
	s5f32 = s5f32 * s6f32
	s6f32 = l15
	s7i32 = l3
	s7f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s7i32+4)]))
	s6f32 = s6f32 - s7f32
	l11 = s6f32
	s7f32 = l11
	s6f32 = s6f32 * s7f32
	s5f32 = s5f32 + s6f32
	l13 = s5f32
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l9 = s4i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l5 = s2i32
	s3i32 = 1
	s4i32 = l8
	s5i32 = -1
	s4i32 = s4i32 + s5i32
	s5i32 = l9
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	l10 = s3i32
	s2i32 = s2i32 + s3i32
	s3i32 = l8
	s2i32 = i32RemS(s2i32, s3i32)
	l3 = s2i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s1f32 = l11
	s0f32 = s0f32 * s1f32
	s1f32 = l15
	s2i32 = l6
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s2f32 = l11
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l12 = s0f32
	s1f32 = l16
	s2f32 = l13
	s3i32 = l9
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
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl3:
		s0f32 = l14
		s1i32 = l7
		s2i32 = l3
		l5 = s2i32
		s3i32 = l10
		s2i32 = s2i32 + s3i32
		s3i32 = l8
		s2i32 = i32RemS(s2i32, s3i32)
		l3 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0f32 = s0f32 - s1f32
		l11 = s0f32
		s1f32 = l11
		s0f32 = s0f32 * s1f32
		s1f32 = l15
		s2i32 = l6
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1f32 = s1f32 - s2f32
		l11 = s1f32
		s2f32 = l11
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 + s1f32
		l11 = s0f32
		s1f32 = l12
		if s0f32 < s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l6 = s0i32
		s0f32 = l11
		l12 = s0f32
		s0i32 = l6
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+168)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = l5
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l12 = s0f32
lbl0:
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+188)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+164)]))
	l3 = s0i32
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = l12
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l5 = s1i32
	s2i32 = l3
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s1f32 = l11
	s0f32 = s0f32 * s1f32
	s1f32 = l13
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s2f32 = l11
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0.00390625
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s1i32 = 0
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0f32 = l12
	s1i32 = l5
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+136)]))
	l3 = s2i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0f32 = s0f32 - s1f32
	l11 = s0f32
	s1f32 = l11
	s0f32 = s0f32 * s1f32
	s1f32 = l13
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s1f32 = s1f32 - s2f32
	l11 = s1f32
	s2f32 = l11
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0.00390625
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	return s0i32
lbl5:
	s0i32 = l4
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	s0i32 = f45(ctx, s0i32)
	l1 = s0i32
	s1f32 = l13
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l1
	s1f32 = l12
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = 40
	s0i32 = s0i32 + s1i32
	s0i32 = f88(ctx, s0i32)
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	return s0i32
lbl4:
	s0i32 = l4
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 1
	return s0i32
}
