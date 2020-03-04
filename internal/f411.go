package internal

import (
	"math"
	"unsafe"
)

func f411(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	s1i32 = 128
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		l12 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l18 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		l15 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l13 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l14 = s0f32
		s0i32 = l0
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l19 = s0f32
		goto lbl1
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	l8 = s2i32
	s3i32 = l6
	s3f32 = math.Float32frombits(uint32(s3i32))
	l14 = s3f32
	s4i32 = l8
	s4f32 = math.Float32frombits(uint32(s4i32))
	l19 = s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l7 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	l9 = s2i32
	s3i32 = l7
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4i32 = l9
	s4f32 = math.Float32frombits(uint32(s4i32))
	l12 = s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l11 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	l7 = s2i32
	s3i32 = l11
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4i32 = l7
	s4f32 = math.Float32frombits(uint32(s4i32))
	l15 = s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
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
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l6
	s2i32 = l8
	s3f32 = l14
	s4f32 = l19
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l6 = s1i32
	s2i32 = l9
	s3i32 = l6
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l12
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l6 = s1i32
	s2i32 = l7
	s3i32 = l6
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l15
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l17 = s1f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l6 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	l8 = s2i32
	s3i32 = l6
	s3f32 = math.Float32frombits(uint32(s3i32))
	l13 = s3f32
	s4i32 = l8
	s4f32 = math.Float32frombits(uint32(s4i32))
	l18 = s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l7 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l9 = s2i32
	s3i32 = l7
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4i32 = l9
	s4f32 = math.Float32frombits(uint32(s4i32))
	l20 = s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l11 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	l7 = s2i32
	s3i32 = l11
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4i32 = l7
	s4f32 = math.Float32frombits(uint32(s4i32))
	l12 = s4f32
	if s3f32 > s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l21 = s1f32
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
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l6
	s2i32 = l8
	s3f32 = l13
	s4f32 = l18
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l3 = s1i32
	s2i32 = l9
	s3i32 = l3
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l20
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l3 = s1i32
	s2i32 = l7
	s3i32 = l3
	s3f32 = math.Float32frombits(uint32(s3i32))
	s4f32 = l12
	if s3f32 < s4f32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s1f32 = math.Float32frombits(uint32(s1i32))
	l20 = s1f32
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = l16
	if s0f32 >= s1f32 {
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
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = l17
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
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1f32 = l21
	if s0f32 >= s1f32 {
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
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1f32 = l20
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
	s0i32 = 0
	l1 = s0i32
lbl1:
	s0f32 = l19
	s1f32 = l14
	s0f32 = s0f32 - s1f32
	s1f32 = l15
	s2f32 = l14
	s1f32 = s1f32 - s2f32
	l16 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l18
	s2f32 = l13
	s1f32 = s1f32 - s2f32
	s2f32 = l12
	s3f32 = l13
	s2f32 = s2f32 - s3f32
	l17 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l16
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	l16 = s1f32
	s2f32 = l14
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l17
	s2i32 = l0
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	l17 = s2f32
	s3f32 = l13
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l19
	s1f32 = l15
	s0f32 = s0f32 - s1f32
	s1f32 = l14
	s2f32 = l15
	s1f32 = s1f32 - s2f32
	l14 = s1f32
	s0f32 = s0f32 * s1f32
	s1f32 = l18
	s2f32 = l12
	s1f32 = s1f32 - s2f32
	s2f32 = l13
	s3f32 = l12
	s2f32 = s2f32 - s3f32
	l13 = s2f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l14
	s1f32 = l16
	s2f32 = l15
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = l13
	s2f32 = l17
	s3f32 = l12
	s2f32 = s2f32 - s3f32
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = 0
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l4
	s3i32 = l5
	f626(ctx, s0i32, s1i32, s2i32, s3i32)
	goto lbl0
lbl3:
	s0i32 = 0
	l3 = s0i32
	s0i32 = l0
	s1i32 = l10
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	s2i32 = l10
	s3i32 = 4
	s2i32 = s2i32 + s3i32
	s0i32 = f2111(ctx, s0i32, s1i32, s2i32)
	l0 = s0i32
	s1i32 = 0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl4:
	s0i32 = l10
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l4
	s3i32 = l5
	f626(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l0
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl4
	}
lbl0:
	s0i32 = l10
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
