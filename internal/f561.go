package internal

import (
	"math"
	"unsafe"
)

func f561(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int64
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l14 = s2f32
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l7
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l2
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l15 = s2f32
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l16 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l17 = s0f32
lbl2:
	s0i32 = l4
	s1i32 = l3
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2i32 = l3
	s2i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+18)])))
	l0 = s2i32
	s3i32 = l5
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l10 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l11 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l12 = s0f32
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l13 = s0f32
		s0i32 = l0
		s1i32 = l5
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0f32 = l12
			s1f32 = l16
			s2f32 = l11
			s1f32 = s1f32 - s2f32
			s0f32 = s0f32 * s1f32
			s1f32 = l10
			s2f32 = l17
			s3f32 = l13
			s2f32 = s2f32 - s3f32
			s1f32 = s1f32 * s2f32
			s0f32 = s0f32 - s1f32
			l10 = s0f32
			s0f32 = float32(math.Abs(float64(s0f32)))
			s1f32 = 5.9604645e-08
			if s0f32 <= s1f32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			s1i32 = 1
			s0i32 = s0i32 ^ s1i32
			if s0i32 != 0 {
				goto lbl4
			}
			goto lbl0
		}
		s0f32 = l12
		s1f32 = l15
		s2f32 = l11
		s1f32 = s1f32 - s2f32
		s0f32 = s0f32 * s1f32
		s1f32 = l10
		s2f32 = l14
		s3f32 = l13
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s0f32 = s0f32 - s1f32
		l10 = s0f32
		s0f32 = float32(math.Abs(float64(s0f32)))
		s1f32 = 5.9604645e-08
		if s0f32 <= s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	lbl4:
		s0i32 = l3
		s1f32 = l10
		s2f32 = 0
		if s1f32 > s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l3 = s0i32
		if s0i32 != 0 {
			goto lbl2
		}
		goto lbl0
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l0 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l1 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		s2i32 = l3
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l4
		s4i32 = l5
		s0i32 = f166(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l2
		s2i32 = l7
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = l6
		s0i32 = f166(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l0
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s2i32 = l3
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l3
		s3i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)])))
		s4i32 = l3
		s4i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s4i32+18)])))
		s0i32 = f166(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s1i32 = l2
		s2i32 = l7
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = l6
		s0i32 = f166(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l7
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i64
	s0i32 = l3
	s1i32 = l6
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+18)])) = uint16(s1i32)
	s0i32 = l3
	s1i32 = l5
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint16(s1i32)
	s0i32 = l3
	s1i64 = l9
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = 1
	l8 = s0i32
lbl0:
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l8
	return s0i32
}
