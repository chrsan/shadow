package internal

import (
	"math"
	"unsafe"
)

func f1687(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l18 float32
	_ = l18
	var l19 float32
	_ = l19
	var l20 float32
	_ = l20
	var l21 float32
	_ = l21
	var l22 float32
	_ = l22
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
	var s5f32 float32
	_ = s5f32
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	f231(ctx, s0i32)
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 0.5
	s0f32 = s0f32 * s1f32
	l12 = s0f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l13 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l5 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l14 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l7 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l15 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l6 = s0i32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l16 = s0f32
	s0i32 = int32(math.Float32bits(s0f32))
	l9 = s0i32
	s0f32 = l14
	s1f32 = l16
	s0f32 = s0f32 - s1f32
	s1f32 = 0
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1f32 = l13
	s2f32 = l15
	s1f32 = s1f32 - s2f32
	s2f32 = 0
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = 14016
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
	}
	s0i32 = l4
	s1f32 = l12
	s2i32 = l6
	s3i32 = l5
	s4f32 = l15
	s5f32 = l13
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l10 = s4i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l1 = s2i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l17 = s2f32
	s1f32 = s1f32 + s2f32
	l13 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l4
	s1f32 = l12
	s2i32 = l9
	s3i32 = l7
	s4f32 = l16
	s5f32 = l14
	if s4f32 > s5f32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	l11 = s4i32
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l8 = s2i32
	s2f32 = math.Float32frombits(uint32(s2i32))
	l18 = s2f32
	s1f32 = s1f32 + s2f32
	l14 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l4
	s1i32 = l5
	s2i32 = l6
	s3i32 = l10
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l6 = s1i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l19 = s1f32
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	l15 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l4
	s1i32 = l7
	s2i32 = l9
	s3i32 = l11
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l5 = s1i32
	s1f32 = math.Float32frombits(uint32(s1i32))
	l20 = s1f32
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	l16 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0f32 = l17
	s1f32 = l19
	s0f32 = s0f32 - s1f32
	l21 = s0f32
	s0f32 = l18
	s1f32 = l20
	s0f32 = s0f32 - s1f32
	l22 = s0f32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+13)])
	l7 = s0i32
	s1i32 = 2
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l7
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl3
	case 1:
		goto lbl4
	default:
		goto lbl5
	}
lbl5:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = 1.4142135
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l2
	s1i32 = l4
	s2i32 = l3
	f139(ctx, s0i32, s1i32, s2i32)
	goto lbl2
lbl4:
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)])) = uint32(s2i32)
		s1i32 = l4
		s2f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)])) = s2f32
		s1i32 = l4
		s2f32 = l13
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)])) = s2f32
		s1i32 = l4
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)])) = uint32(s2i32)
		s1i32 = l4
		s2f32 = l16
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)])) = s2f32
		s1i32 = l4
		s2f32 = l16
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)])) = s2f32
		s1i32 = l4
		s2i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)])) = uint32(s2i32)
		s1i32 = l4
		s2i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)])) = uint32(s2i32)
		s1i32 = l4
		s2f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)])) = s2f32
		s1i32 = l4
		s2f32 = l14
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = s2f32
		s1i32 = l4
		s2f32 = l15
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = s2f32
		s1i32 = l4
		s2i32 = l8
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint32(s2i32)
		s1i32 = l4
		s2f32 = l15
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)])) = s2f32
		s1i32 = l4
		s2i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
		s1i32 = 7
		goto lbl6
	}
	s1i32 = l4
	s2f32 = l15
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)])) = s2f32
	s1i32 = l4
	s2f32 = l15
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)])) = s2f32
	s1i32 = l4
	s2i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)])) = uint32(s2i32)
	s1i32 = l4
	s2f32 = l13
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)])) = s2f32
	s1i32 = l4
	s2f32 = l13
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)])) = s2f32
	s1i32 = l4
	s2i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)])) = uint32(s2i32)
	s1i32 = l4
	s2f32 = l14
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)])) = s2f32
	s1i32 = l4
	s2f32 = l14
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)])) = s2f32
	s1i32 = l4
	s2i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = uint32(s2i32)
	s1i32 = l4
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)])) = uint32(s2i32)
	s1i32 = l4
	s2f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = s2f32
	s1i32 = l4
	s2f32 = l16
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = s2f32
	s1i32 = 0
lbl6:
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l4
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = 8
	s3i32 = 1
	f277(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl2
lbl3:
	s0i32 = l2
	s1i32 = l4
	s2f32 = l12
	s3f32 = l12
	s4i32 = l3
	f2094(ctx, s0i32, s1i32, s2f32, s3f32, s4i32)
lbl2:
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l21
	s2f32 = l22
	s3f32 = l21
	s4f32 = l22
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
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+14)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1f32 = l17
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	s0i32 = l4
	s1f32 = l18
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = s1f32
	s0i32 = l4
	s1f32 = l12
	s2f32 = l19
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l4
	s1f32 = l12
	s2f32 = l20
	s1f32 = s1f32 + s2f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l2
	s1i32 = l4
	s2i32 = l3
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 14016
	s2i32 = s2i32 + s3i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	f139(ctx, s0i32, s1i32, s2i32)
lbl0:
	s0i32 = l4
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
