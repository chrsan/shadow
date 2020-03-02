package internal

import (
	"math"
	"unsafe"
)

func f103(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
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
	s0i32 = l0
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 16
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	s1f32 = 0.0625
	s0f32 = s0f32 * s1f32
	l9 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 16
	s0f32 = s0f32 * s1f32
	s1f32 = 0.5
	s0f32 = s0f32 + s1f32
	s0f32 = float32(math.Floor(float64(s0f32)))
	s1f32 = 0.0625
	s0f32 = s0f32 * s1f32
	l10 = s0f32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
	l2 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = l2
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0f32
	s1f32 = l9
	s0f32 = s0f32 - s1f32
	l6 = s0f32
	s1f32 = l6
	s0f32 = s0f32 * s1f32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l8 = s1f32
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	l6 = s1f32
	s2f32 = l6
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0.00390625
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1f32 = l7
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l7 = s2f32
	s1f32 = s1f32 - s2f32
	l12 = s1f32
	s2f32 = l10
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l6 = s3f32
	s2f32 = s2f32 - s3f32
	l11 = s2f32
	s1f32 = s1f32 * s2f32
	s2f32 = l9
	s3f32 = l7
	s2f32 = s2f32 - s3f32
	l7 = s2f32
	s3f32 = l8
	s4f32 = l6
	s3f32 = s3f32 - s4f32
	l6 = s3f32
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+124)]))
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = s1f32
	s0i32 = l0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+116)]))
	s2f32 = l12
	s3f32 = l7
	s2f32 = s2f32 + s3f32
	s3f32 = l8
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = s1f32
	s0i32 = l0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+120)]))
	s2f32 = l6
	s3f32 = l11
	s2f32 = s2f32 + s3f32
	s3f32 = l8
	s2f32 = s2f32 * s3f32
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = s1f32
	s0f32 = l8
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+128)]))
	s0f32 = s0f32 * s1f32
	s1f32 = 0
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
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+158)] = uint8(s1i32)
	}
	s0f32 = l8
	s1f32 = 0
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1f32 = l8
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = s1f32
	}
	s0i32 = l2
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0f32
	s1i32 = l1
	s2i32 = l2
	s3i32 = -2
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l11 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = l10
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l6 = s2f32
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l9
	s2f32 = l7
	s1f32 = s1f32 - s2f32
	s2f32 = l6
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l7 = s3f32
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 - s1f32
	l6 = s0f32
	s0f32 = float32(math.Abs(float64(s0f32)))
	s1f32 = 0.00024414062
	if s0f32 <= s1f32 {
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
		s0f32 = l6
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+132)]))
		s0f32 = s0f32 * s1f32
		s1f32 = 0
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
			s0i32 = l0
			s1i32 = 0
			ctx.Mem[int(s0i32+158)] = uint8(s1i32)
		}
		s0f32 = l6
		s1f32 = 0
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l0
		s1f32 = l6
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = s1f32
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0f32 = l11
	s1f32 = l9
	s0f32 = s0f32 - s1f32
	l6 = s0f32
	s1f32 = l6
	s0f32 = s0f32 * s1f32
	s1f32 = l7
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	l6 = s1f32
	s2f32 = l6
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
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
lbl1:
	s0i32 = l4
	s0i32 = f45(ctx, s0i32)
	l0 = s0i32
	s1f32 = l10
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l0
	s1f32 = l9
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
lbl0:
}
