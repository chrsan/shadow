package internal

import (
	"unsafe"
)

func f675(ctx *Context, l0 int32, l1 int32) {
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
	var l7 int32
	_ = l7
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
	var s6f32 float32
	_ = s6f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l2 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l6 = s0i32
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	s1i32 = 0
	s2i32 = l6
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l2 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l5 = s0i32
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l7 = s0i32
	}
	s0i32 = l4
	s1i32 = l6
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = l2
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl2:
	s0i32 = l1
	l2 = s0i32
	s1i32 = l4
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l5 = s1i32
	s2i32 = 4338
	s1i32 = s1i32 + s2i32
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l3 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = 4
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = 0
	s1i32 = l3
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 - s1i32
	switch s0i32 {
	case 0:
		goto lbl7
	case 1:
		goto lbl6
	case 2:
		goto lbl5
	case 3:
		goto lbl4
	default:
		goto lbl0
	}
lbl7:
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l2
	s3i32 = l3
	s4i32 = 3
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 + s3i32
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	goto lbl3
lbl6:
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l2
	s5i32 = l3
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	f130(ctx, s0i32, s1f32, s2f32, s3f32, s4f32)
	goto lbl3
lbl5:
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s4i32 = l2
	s5i32 = l3
	s6i32 = 3
	s5i32 = s5i32 << (uint32(s6i32) & 31)
	s4i32 = s4i32 + s5i32
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
	s5i32 = l7
	s6i32 = -4
	s5i32 = s5i32 + s6i32
	l7 = s5i32
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	goto lbl3
lbl4:
	s0i32 = l0
	s1i32 = l1
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l1
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s3i32 = l1
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l1
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = l1
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
	s6i32 = l2
	s7i32 = l3
	s8i32 = 3
	s7i32 = s7i32 << (uint32(s8i32) & 31)
	s6i32 = s6i32 + s7i32
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)]))
	f175(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32, s6f32)
lbl3:
	s0i32 = l4
	s1i32 = l6
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl0:
}
