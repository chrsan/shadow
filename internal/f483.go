package internal

import (
	"unsafe"
)

func f483(ctx *Context, l0 int32, l1 float64) int32 {
	var l2 int32
	_ = l2
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
	var l8 float64
	_ = l8
	var l9 float64
	_ = l9
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
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
	var s1i64 int64
	_ = s1i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0f64 = l1
	s1f64 = l1
	s0f64 = s0f64 * s1f64
	l10 = s0f64
	s1f64 = l10
	s0f64 = s0f64 * s1f64
	s1f64 = 0.25
	s0f64 = s0f64 * s1f64
	l8 = s0f64
	s0f64 = 1
	l7 = s0f64
	s0i32 = 1
	l2 = s0i32
	s0f64 = 1
	l1 = s0f64
lbl0:
	s0f64 = l1
	s1f64 = l7
	s2f64 = l8
	s3i32 = l2
	s4i32 = l2
	s3i32 = s3i32 * s4i32
	s3f64 = float64(s3i32)
	s2f64 = s2f64 / s3f64
	s1f64 = s1f64 * s2f64
	l7 = s1f64
	s0f64 = s0f64 + s1f64
	l1 = s0f64
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0f64 = l7
	s1f64 = 1e-06
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1f64 = l1
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = 1
	l2 = s0i32
	s0f64 = l10
	s0f64 = f1356(ctx, s0f64)
	l11 = s0f64
	s0f64 = l10
	s1f64 = 0.5
	s0f64 = s0f64 * s1f64
	l9 = s0f64
	s1f64 = 1e-06
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f64 = l9
		l7 = s0f64
		goto lbl1
	}
	s0f64 = l9
	l7 = s0f64
lbl3:
	s0f64 = l7
	s1f64 = l9
	s2f64 = l8
	s3i32 = l2
	s4i32 = l2
	s5i32 = 1
	s4i32 = s4i32 + s5i32
	l2 = s4i32
	s3i32 = s3i32 * s4i32
	s3f64 = float64(s3i32)
	s2f64 = s2f64 / s3f64
	s1f64 = s1f64 * s2f64
	l9 = s1f64
	s0f64 = s0f64 + s1f64
	l7 = s0f64
	s0f64 = l9
	s1f64 = 1e-06
	if s0f64 > s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
lbl1:
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l4
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l4
	s1f64 = l7
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = l0
	s1f64 = l7
	s2f64 = l11
	s1f64 = s1f64 / s2f64
	l9 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f64
	s0i32 = l0
	s1f64 = l1
	s2f64 = l11
	s1f64 = s1f64 / s2f64
	l8 = s1f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = 1
	l2 = s0i32
	s0f64 = l9
	s1f64 = 0.01
	if s0f64 > s1f64 {
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
	lbl6:
		s0i32 = l4
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		l6 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l3 = s1i32
		s0i32 = s0i32 + s1i32
		s1f64 = l1
		s2f64 = l7
		s3i32 = l2
		s4i32 = 1
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s3f64 = float64(s3i32)
		s4f64 = l10
		s3f64 = s3f64 / s4f64
		s2f64 = s2f64 * s3f64
		s1f64 = s1f64 - s2f64
		l7 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l0
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		s1f64 = l7
		s2f64 = l11
		s1f64 = s1f64 / s2f64
		l1 = s1f64
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0f64 = l1
		s1f64 = 0.01
		if s0f64 > s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0f64 = 0
			l1 = s0f64
			s0i32 = l2
			l3 = s0i32
		lbl8:
			s0f64 = l1
			s1i32 = l0
			s2i32 = l3
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l1 = s1f64
			s2f64 = l1
			s1f64 = s1f64 + s2f64
			s0f64 = s0f64 + s1f64
			l1 = s0f64
			s0i32 = l3
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l5 = s0i32
			s0i32 = l3
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l5
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l0
			s1f64 = l8
			s2f64 = l8
			s3f64 = l1
			s2f64 = s2f64 + s3f64
			l1 = s2f64
			s1f64 = s1f64 / s2f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
			s0i32 = 1
			l3 = s0i32
		lbl9:
			s0i32 = l0
			s1i32 = l3
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = l5
			s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2f64 = l1
			s1f64 = s1f64 / s2f64
			*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
			s0i32 = l2
			s1i32 = l3
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l5 = s0i32
			s0i32 = l3
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l3 = s0i32
			s0i32 = l5
			if s0i32 != 0 {
				goto lbl9
			}
			s0f64 = 0
			l1 = s0f64
		lbl10:
			s0f64 = l1
			s1i32 = l0
			s2i32 = l2
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			l1 = s1f64
			s2f64 = l1
			s1f64 = s1f64 + s2f64
			s0f64 = s0f64 + s1f64
			l1 = s0f64
			s0i32 = l2
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l3 = s0i32
			s0i32 = l2
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l3
			if s0i32 != 0 {
				goto lbl10
			}
			goto lbl4
		} else {
			s0i32 = l4
			s1i32 = l2
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l1 = s0f64
			s0i32 = l6
			l2 = s0i32
			goto lbl6
		}
		panic("unreachable executed")
		panic("unreachable executed")
		panic("unreachable executed")
	}
	s0f64 = 0
	l1 = s0f64
	s0i32 = l0
	s1f64 = l8
	s2f64 = l8
	s3f64 = 0
	s2f64 = s2f64 + s3f64
	s1f64 = s1f64 / s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = 1
	l6 = s0i32
lbl4:
	s0i32 = l0
	s1f64 = 1
	s2f64 = l1
	s1f64 = s1f64 - s2f64
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
