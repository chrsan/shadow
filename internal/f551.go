package internal

import (
	"unsafe"
)

func f551(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 float32
	_ = l8
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0i32 = ctx.g0
	s1i32 = 208
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		s2i32 = l6
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		l7 = s0i32
		s0i32 = l6
		s1i32 = 0
		ctx.Mem[int(s0i32+84)] = uint8(s1i32)
		s0i32 = l6
		s1i32 = l7
		ctx.Mem[int(s0i32+85)] = uint8(s1i32)
	}
	s0i32 = l2
	s1i32 = l5
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l6
	s3i32 = 4
	s2i32 = s2i32 + s3i32
	s0i32 = f42(ctx, s0i32, s1i32, s2i32)
	s0i32 = l5
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1f32 = 0
	s0f32 = s0f32 * s1f32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0f32 = s0f32 * s1f32
	s1i32 = l5
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s0f32 = s0f32 * s1f32
	l8 = s0f32
	s1f32 = l8
	if s0f32 == s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1f32 = 0
		s0f32 = s0f32 * s1f32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s0f32 = s0f32 * s1f32
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s0f32 = s0f32 * s1f32
		l8 = s0f32
		s1f32 = l8
		if s0f32 == s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl1
lbl2:
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l2
	s3i32 = l3
	s4i32 = l4
	s0i32 = f1397(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l1 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+156)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl4
	}
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	s4i32 = l1
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+60)]))
	s5i32 = l1
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+52)]))
	f593(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
lbl4:
	s0i32 = l1
	s1i32 = 27304
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
	f13(ctx, s0i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	f13(ctx, s0i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	f13(ctx, s0i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	f13(ctx, s0i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	f13(ctx, s0i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	f13(ctx, s0i32)
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	f13(ctx, s0i32)
lbl1:
	s0i32 = l5
	s1i32 = 208
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
