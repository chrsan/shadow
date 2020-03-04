package internal

import (
	"math"
	"unsafe"
)

func f1695(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
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
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	f1692(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 0
	s0i32 = f310(ctx, s0i32, s1i32, s2i32)
	l4 = s0i32
	s1i32 = 2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = l3
		s3i32 = 8
		s2i32 = s2i32 | s3i32
		s3i32 = l3
		s4i32 = 24
		s3i32 = s3i32 + s4i32
		f336(ctx, s0i32, s1f32, s2i32, s3i32)
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l3
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
		s1f32 = 0
		if s0f32 != s1f32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
		l6 = s0f32
		s0i32 = l1
		s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l5 = s0f32
		s0i32 = l3
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l1
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l3
		s1f32 = l6
		s2f32 = l5
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
	lbl1:
		s0i32 = l3
		s1i32 = 24
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s0i32 = f113(ctx, s0i32, s1f32)
		if s0i32 != 0 {
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
			l5 = s0f32
			s0i32 = l3
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
			goto lbl2
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0f32 = 0
		l5 = s0f32
		s0i32 = l1
		s0f32 = math.Float32frombits(uint32(s0i32))
	lbl2:
		l6 = s0f32
		s0i32 = l3
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2f32 = l5
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+132)]))
		s3f32 = float32(s3i32)
		l5 = s3f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 + s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
		s0i32 = l3
		s1i32 = l3
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2f32 = l6
		s3f32 = l5
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 - s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
		s0i32 = l0
		s1i32 = l2
		s2i32 = l3
		s3i32 = l2
		s0i32 = f399(ctx, s0i32, s1i32, s2i32, s3i32)
		l4 = s0i32
	}
	s0i32 = l3
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
