package internal

import (
	"unsafe"
)

func f1049(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	s0i32 = 1
	l4 = s0i32
	s0i32 = l2
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = 0
	l4 = s0i32
	s0i32 = l1
	s1i32 = l5
	s0i32 = s0i32 + s1i32
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l6 = s0i32
	s1i32 = 5
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l7 = s0i32
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
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
	case 4:
		goto lbl3
	default:
		goto lbl8
	}
lbl8:
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	f176(ctx, s0i32, s1f32, s2f32)
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl2
lbl7:
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s0i32 = f30(ctx, s0i32, s1f32, s2f32)
	s0i32 = l3
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl2
lbl6:
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	f130(ctx, s0i32, s1f32, s2f32, s3f32, s4f32)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl2
lbl5:
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
	s0i32 = f73(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32)
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl2
lbl4:
	s0i32 = l0
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = l3
	s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = l3
	s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	s4i32 = l3
	s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+12)]))
	s5i32 = l3
	s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+16)]))
	s6i32 = l3
	s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+20)]))
	f175(ctx, s0i32, s1f32, s2f32, s3f32, s4f32, s5f32, s6f32)
	s0i32 = l3
	s1i32 = 24
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	goto lbl2
lbl3:
	s0i32 = l0
	f230(ctx, s0i32)
lbl2:
	s0i32 = 1
	l4 = s0i32
	s0i32 = l5
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s1i32 = l2
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl0:
	s0i32 = l4
	return s0i32
}
