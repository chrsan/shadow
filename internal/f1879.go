package internal

import (
	"math"
	"unsafe"
)

func f1879(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	var l11 int32
	_ = l11
	var l12 int64
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
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f24(ctx, s1i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = 15
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl2
	}
	s0i32 = l3
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f24(ctx, s1i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = l4
	s2i32 = l0
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l13 = s0f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l14 = s1f32
	if s0f32 < s1f32 {
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
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l15 = s0f32
	s1i32 = l4
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l16 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0f32 = l15
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	s1f32 = l16
	s0f32 = s0f32 * s1f32
	s1f32 = l14
	s0f32 = s0f32 * s1f32
	l17 = s0f32
	s1f32 = l17
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l8 = s0i32
		s0i32 = l3
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1f32 = l14
		s2f32 = l13
		s1f32 = s1f32 - s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l13 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
		s0i32 = l2
		s1f32 = l16
		s2f32 = l15
		s1f32 = s1f32 - s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l14 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l2
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
		s0i32 = l2
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l2
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l2
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l2
		s1f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l2
		s1f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		goto lbl2
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l13 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l14 = s0f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l3 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l1
		s1i32 = f24(ctx, s1i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l3
	s1i32 = 12
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		l10 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		l9 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l5 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		l11 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		l6 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l0 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
		goto lbl7
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l14 = s0f32
	s0i32 = l2
	s1i32 = l0
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s2i32 = 3
	s3i32 = 1
	s4i32 = l1
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	l13 = s4f32
	s5f32 = 0
	if s4f32 < s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l7 = s4i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l5 = s2i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	s2i32 = l5
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	s3i32 = 2
	s2i32 = s2i32 & s3i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l9 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l10 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l9
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l9 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	s2i32 = l5
	s3i32 = 2
	s2i32 = s2i32 + s3i32
	s3i32 = 3
	s2i32 = s2i32 & s3i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l6
	s2i32 = -8
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l6 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l0 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0f32 = l13
	s0f32 = -s0f32
	s1f32 = l13
	s2i32 = l7
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l13 = s0f32
	s0f32 = l14
	s1f32 = l14
	s1f32 = -s1f32
	s2i32 = l7
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l14 = s0f32
lbl7:
	s0i32 = l2
	s1f32 = l13
	s1f32 = -s1f32
	s2f32 = l13
	s3f32 = l13
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l7 = s3i32
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l13 = s1f32
	s2i32 = l0
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	s1f32 = -s1f32
	s2f32 = l14
	s3f32 = l14
	s4f32 = 0
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l0 = s3i32
	if s3i32 != 0 {
		// s1f32 = s1f32
	} else {
		s1f32 = s2f32
	}
	l14 = s1f32
	s2i32 = l6
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
	s0i32 = l2
	s1f32 = l13
	s2i32 = l11
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	s2i32 = l5
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
	s0i32 = l2
	s1f32 = l13
	s2i32 = l9
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	s2i32 = l10
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l2
	s1f32 = l13
	s2i32 = l3
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l2
	s1f32 = l14
	s2i32 = l1
	s2f32 = math.Float32frombits(uint32(s2i32))
	s1f32 = s1f32 * s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l0
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l7
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
			l12 = s0i64
			s0i32 = l2
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
			s0i32 = l2
			s1i64 = l12
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
			s0i32 = l2
			s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
			l12 = s0i64
			s0i32 = l2
			s1i32 = l2
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
			s0i32 = l2
			s1i64 = l12
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
			goto lbl9
		}
		s0i32 = l2
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l12 = s0i64
		s0i32 = l2
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = l12
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l2
		s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		l12 = s0i64
		s0i32 = l2
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
		s0i32 = l2
		s1i64 = l12
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
		goto lbl9
	}
	s0i32 = l7
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l12 = s0i64
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l12
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l12 = s0i64
	s0i32 = l2
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l12
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
lbl9:
	s0i32 = l2
	s1i32 = l2
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s0i32 = f1878(ctx, s0i32, s1i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l2
	f650(ctx, s0i32, s1i32)
lbl2:
	s0i32 = 1
	l8 = s0i32
lbl0:
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l8
	return s0i32
}
