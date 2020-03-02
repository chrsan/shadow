package internal

import (
	"unsafe"
)

func f1282(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var l5 int32
	_ = l5
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
	s0i32 = ctx.g0
	s1i32 = 160
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l5
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	s3i32 = l2
	s4i32 = l3
	s0i32 = f467(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l2 = s0i32
	s1i32 = l5
	s2i32 = 144
	s1i32 = s1i32 + s2i32
	s2i32 = l5
	s3i32 = 0
	s4i32 = 0
	s0i32 = f158(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 != 0 {
	lbl1:
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+156)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = s1f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+152)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = s1f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+148)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
		s0i32 = l5
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+144)]))
		s1f32 = float32(s1i32)
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = s1f32
		s0i32 = l0
		s1i32 = l1
		s2i32 = l5
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		s3i32 = l5
		s4i32 = l4
		s5i32 = 0
		s6i32 = l0
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+112)]))
		if int(s6i32) < 0 || int(s6i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s6i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s6i32].numParams != 6 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32, int32))(table[s6i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
		s0i32 = l2
		s1i32 = l5
		s2i32 = 144
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s3i32 = 0
		s4i32 = 0
		s0i32 = f158(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+92)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+76)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+60)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+44)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+28)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		f13(ctx, s0i32)
	}
	s0i32 = l2
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f13(ctx, s0i32)
	}
	s0i32 = l5
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
