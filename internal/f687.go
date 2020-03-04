package internal

import (
	"math"
	"unsafe"
)

func f687(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
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
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l11 = s0f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l12 = s1f32
	s0f32 = s0f32 * s1f32
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l13 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l14 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l10 = s0f32
	s0f32 = l11
	s1f32 = l13
	s0f32 = s0f32 * s1f32
	s1f32 = l14
	s2f32 = l12
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l11 = s0f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l10
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l10
	s3f32 = 0
	if s2f32 >= s3f32 {
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
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0f32 = l10
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = 2
	l6 = s0i32
	s0f32 = l10
	s0f32 = -s0f32
	s1f32 = l10
	s2i32 = l2
	s3i32 = 1
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	l10 = s0f32
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l11
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		s1i32 = 3
		s2f32 = l10
		s3f32 = 0
		if s2f32 > s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		l6 = s0i32
		goto lbl3
	}
	s0f32 = l10
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1f32 = l11
	s2f32 = 0
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l1
	s1i32 = s1i32 ^ s2i32
	s0i32 = s0i32 | s1i32
	l6 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl3:
	s0i32 = 0
	l1 = s0i32
lbl5:
	s0i32 = l4
	s1i32 = l1
	s2i32 = 28
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 4208
	s1i32 = s1i32 + s2i32
	l8 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l8
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l8
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = 1060439283
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl2:
	s0f32 = l11
	s1i32 = l6
	s2i32 = 4
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 4208
	s1i32 = s1i32 + s2i32
	l7 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l12 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l10
	s2i32 = l7
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l13 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	l14 = s0f32
	s1f32 = 1
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0i32 = l5
	s1f32 = l10
	s2f32 = l13
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l5
	s1f32 = l11
	s2f32 = l12
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1f32 = 1
	s2f32 = l14
	s3f32 = 1
	s2f32 = s2f32 + s3f32
	s3f32 = 0.5
	s2f32 = s2f32 * s3f32
	s2f32 = float32(math.Sqrt(float64(s2f32)))
	l14 = s2f32
	s1f32 = s1f32 / s2f32
	s0i32 = f113(ctx, s0i32, s1f32)
	s0f32 = l12
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s0f32 = s0f32 - s1f32
	l12 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0f32 = l13
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 - s1f32
	l13 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	s1i32 = 2139095040
	s0i32 = s0i32 & s1i32
	s1i32 = 2139095040
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
	s0f32 = l12
	s1f32 = 0
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 0
	s2f32 = l13
	s3f32 = 0
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
		goto lbl6
	}
	s0i32 = l4
	s1i32 = l6
	s2i32 = 28
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l7
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l5
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l9 = s0i64
	s0i32 = l1
	s1f32 = l14
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
	s0i32 = l1
	s1f32 = l10
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	s0i32 = l1
	s1f32 = l11
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
	s0i32 = l1
	s1i64 = l9
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l6 = s0i32
lbl6:
	s0i32 = l5
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l5
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	f980(ctx, s0i32, s1f32, s2f32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1f32 = 1
		s2f32 = -1
		f476(ctx, s0i32, s1f32, s2f32)
	}
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		f116(ctx, s0i32, s1i32)
	}
	s0i32 = 0
	l1 = s0i32
	s0i32 = l6
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl9:
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = l1
	s3i32 = 28
	s2i32 = s2i32 * s3i32
	s1i32 = s1i32 + s2i32
	l0 = s1i32
	s2i32 = l0
	s3i32 = 3
	f55(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l6
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
lbl0:
	s0i32 = l5
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l6
	return s0i32
}
