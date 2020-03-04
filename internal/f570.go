package internal

import (
	"math/bits"
	"unsafe"
)

func f570(ctx *Context, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32) int32 {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 float64
	_ = l9
	var l10 float64
	_ = l10
	var l11 float64
	_ = l11
	var l12 float64
	_ = l12
	var l13 float64
	_ = l13
	var l14 float64
	_ = l14
	var l15 float64
	_ = l15
	var l16 float64
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
	var s5f64 float64
	_ = s5f64
	s0i32 = l1
	s1i32 = l0
	s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l13 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	l11 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
	l9 = s3f64
	s4f64 = l9
	s3f64 = s3f64 + s4f64
	s2f64 = s2f64 - s3f64
	s1f64 = s1f64 + s2f64
	l15 = s1f64
	s2i32 = l0
	s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	l10 = s3f64
	s4i32 = l0
	s4f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s4i32+40)]))
	l12 = s4f64
	s3f64 = s3f64 - s4f64
	s4f64 = 3
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 + s3f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l14 = s3f64
	s2f64 = s2f64 - s3f64
	l16 = s2f64
	s1f64 = s1f64 * s2f64
	s2f64 = l14
	s3f64 = l12
	s4f64 = l10
	s5f64 = l10
	s4f64 = s4f64 + s5f64
	s3f64 = s3f64 - s4f64
	s2f64 = s2f64 + s3f64
	l12 = s2f64
	s3i32 = l0
	s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+48)]))
	s4f64 = l9
	s5f64 = l11
	s4f64 = s4f64 - s5f64
	s5f64 = 3
	s4f64 = s4f64 * s5f64
	s3f64 = s3f64 + s4f64
	s4f64 = l13
	s3f64 = s3f64 - s4f64
	l11 = s3f64
	s2f64 = s2f64 * s3f64
	s1f64 = s1f64 - s2f64
	s2f64 = l9
	s3f64 = l13
	s2f64 = s2f64 - s3f64
	l9 = s2f64
	s3f64 = l16
	s2f64 = s2f64 * s3f64
	s3f64 = l10
	s4f64 = l14
	s3f64 = s3f64 - s4f64
	l10 = s3f64
	s4f64 = l11
	s3f64 = s3f64 * s4f64
	s2f64 = s2f64 - s3f64
	s3f64 = l9
	s4f64 = l12
	s3f64 = s3f64 * s4f64
	s4f64 = l10
	s5f64 = l15
	s4f64 = s4f64 * s5f64
	s3f64 = s3f64 - s4f64
	s4i32 = l1
	s5i32 = l2
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	s1i32 = f567(ctx, s1f64, s2f64, s3f64, s4i32)
	s2i32 = l2
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i64 = 4607182418800017408
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 64
		s1i32 = l7
		s1i32 = int32(bits.LeadingZeros32(uint32(s1i32)))
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 - s1i32
		s1i32 = l1
		s2i32 = l2
		f569(ctx, s0i32, s1i32, s2i32)
		s0i32 = 0
		l2 = s0i32
	lbl1:
		s0i32 = l1
		s1i32 = l2
		l8 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0f64
		s1i32 = l1
		s2i32 = l2
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		l2 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l10 = s1f64
		if s0f64 == s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s1f64 = l9
		s2f64 = l10
		s3f64 = l3
		s4i32 = l4
		s0f64 = f1424(ctx, s0i32, s1f64, s2f64, s3f64, s4i32)
		l9 = s0f64
		s1f64 = 0
		if s0f64 >= s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l6
		s1i32 = 2
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 0
			return s0i32
		}
		s0i32 = l5
		s1i32 = l6
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1f64 = l9
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l6
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l6 = s0i32
	lbl2:
		s0i32 = l7
		s1i32 = l8
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l6
	return s0i32
}
