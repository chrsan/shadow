package internal

import (
	"unsafe"
)

func f697(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 float32) int32 {
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
	var s2f32 float32
	_ = s2f32
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		l6 = s0i32
		goto lbl0
	}
	s0i32 = l5
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l6 = s0i32
	s0i32 = l5
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l5
	s1i32 = l6
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl0:
	s0i32 = l6
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		f231(ctx, s0i32)
		s0i32 = 0
		l0 = s0i32
		goto lbl2
	}
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2f32 = l4
	s0i32 = f617(ctx, s0i32, s1i32, s2f32)
	l6 = s0i32
	s0i32 = l7
	s0i32 = f38(ctx, s0i32)
	l5 = s0i32
	s0i32 = l6
	s1i32 = l2
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l0 = s2i32
	if s2i32 != 0 {
		s2i32 = l5
		s3i32 = l1
		s4i32 = l0
		s5i32 = l5
		s6i32 = l1
		s7i32 = l6
		s8i32 = l3
		s4i32 = f673(ctx, s4i32, s5i32, s6i32, s7i32, s8i32)
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l1 = s2i32
	}
	s2i32 = l1
	s0i32 = f615(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l1
	s1i32 = l5
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l5
		f194(ctx, s0i32, s1i32)
		goto lbl4
	}
	s0i32 = l2
	s1i32 = l1
	f334(ctx, s0i32, s1i32)
lbl4:
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+84)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+85)])
		l1 = s0i32
		goto lbl8
	}
	s0i32 = l0
	s1i32 = 4
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+84)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+85)] = uint8(s1i32)
lbl8:
	s0i32 = l1
	s1i32 = 255
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		f231(ctx, s0i32)
		s0i32 = 0
		goto lbl10
	}
	s0i32 = l6
	s0i32 = f616(ctx, s0i32)
	s1i32 = 0
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
lbl10:
	l0 = s0i32
	s0i32 = l5
	f34(ctx, s0i32)
lbl2:
	s0i32 = l7
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
