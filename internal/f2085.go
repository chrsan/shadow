package internal

import (
	"math"
	"unsafe"
)

func f2085(ctx *Context, l0 int32, l1 int32, l2 float32, l3 float32, l4 int32, l5 int32) {
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
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
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = f137(ctx)
	l8 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l8 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l8
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s0i32 = f97(ctx, s0i32)
	f12(ctx, s0i32)
lbl0:
	s0i32 = l0
	s1i32 = -1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 512
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+10)])
	s2i32 = 248
	s1i32 = s1i32 & s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+10)] = uint8(s1i32)
	s0f32 = 360
	l10 = s0f32
	s0f32 = l3
	s0f32 = float32(math.Abs(float64(s0f32)))
	l11 = s0f32
	s1f32 = 360
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = 0
	s3i32 = 1
	f330(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl1
lbl2:
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l9 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l10 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l12 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l13 = s0f32
		s0i32 = l7
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		l5 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l5
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 0
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
		l5 = s0i32
		s1f32 = l12
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		s2f32 = l13
		s3f32 = 0.5
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l5
		s1f32 = l9
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		s2f32 = l10
		s3f32 = 0.5
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l0
		s1i32 = 512
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint16(s1i32)
		s0f32 = 180
		l10 = s0f32
	}
	s0i32 = l4
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	l5 = s0i32
	s0f32 = l3
	s1f32 = -360
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0f32 = l3
		l9 = s0f32
		goto lbl4
	}
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = -180
	s4i32 = l5
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = 0
	l5 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = -180
	s2f32 = s2f32 + s3f32
	l2 = s2f32
	s3f32 = -180
	s4i32 = 0
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0f32 = l2
	s1f32 = -180
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0f32 = l3
	s1f32 = 360
	s0f32 = s0f32 + s1f32
	l9 = s0f32
	s1f32 = -360
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
		goto lbl4
	}
lbl6:
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = -180
	s4i32 = 0
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = -180
	s2f32 = s2f32 + s3f32
	l2 = s2f32
	s3f32 = -180
	s4i32 = 0
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0f32 = l2
	s1f32 = -180
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0f32 = l9
	s1f32 = 360
	s0f32 = s0f32 + s1f32
	l9 = s0f32
	s1f32 = -360
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
		goto lbl6
	}
lbl4:
	s0f32 = l9
	s1f32 = 360
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = 180
	s4i32 = l5
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = 0
	l5 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = 180
	s2f32 = s2f32 + s3f32
	l2 = s2f32
	s3f32 = 180
	s4i32 = 0
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0f32 = l2
	s1f32 = 180
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0f32 = l9
	s1f32 = -360
	s0f32 = s0f32 + s1f32
	l9 = s0f32
	s1f32 = 360
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl7
	}
lbl8:
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = 180
	s4i32 = 0
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = 180
	s2f32 = s2f32 + s3f32
	l2 = s2f32
	s3f32 = 180
	s4i32 = 0
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0f32 = l2
	s1f32 = 180
	s0f32 = s0f32 + s1f32
	l2 = s0f32
	s0f32 = l9
	s1f32 = -360
	s0f32 = s0f32 + s1f32
	l9 = s0f32
	s1f32 = 360
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl8
	}
lbl7:
	s0i32 = l0
	s1i32 = l1
	s2f32 = l2
	s3f32 = l9
	s4i32 = l5
	f154(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
	s0i32 = l4
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		l4 = s0i32
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l4
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		s0i32 = s0i32 + s1i32
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		s1i32 = 4
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl10
		}
		s0i32 = l7
		s1i32 = l0
		s2i32 = 0
		s3i32 = 0
		s0i32 = f33(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 5
		s2f32 = 0
		s0i32 = f49(ctx, s0i32, s1i32, s2f32)
	lbl10:
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l1 = s1i32
		s2i32 = l1
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		s1i32 = s1i32 ^ s2i32
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1f32 = l3
	s2f32 = 0
	if s1f32 > s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	ctx.Mem[int(s0i32+9)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 1
	s2i32 = 2
	s3f32 = l11
	s4f32 = l10
	if s3f32 <= s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	ctx.Mem[int(s0i32+8)] = uint8(s1i32)
lbl1:
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
