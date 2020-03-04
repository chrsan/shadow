package internal

import (
	"unsafe"
)

func f1345(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
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
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s0i32 = f65(ctx, s0i32)
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		s2i32 = 6
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l4 = s1i32
		if s1i32 != 0 {
			goto lbl3
		}
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		s2i32 = s2i32 | s3i32
		s1i32 = s1i32 | s2i32
		if s1i32 != 0 {
			goto lbl3
		}
		s1i32 = l1
		goto lbl2
	lbl3:
		s1i32 = l3
		s2i32 = l1
		s3i32 = l5
		s4i32 = 8
		s3i32 = s3i32 + s4i32
		s4i32 = l4
		s1i32 = f64(ctx, s1i32, s2i32, s3i32, s4i32)
	lbl2:
		s0i32 = f70(ctx, s0i32, s1i32)
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4388)]))
	l4 = s0i32
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 1
		f57(ctx, s0i32, s1i32)
	}
	s0i32 = l5
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l3
	s3i32 = 0
	s4i32 = l1
	s0i32 = f61(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	l4 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl6:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s3i32 = l4
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+56)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+100)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int32))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	if s0i32 != 0 {
		goto lbl6
	}
lbl5:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+64)])
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
		f54(ctx, s0i32)
	}
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l0 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = f23(ctx, s0i32)
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
lbl0:
	s0i32 = l5
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
