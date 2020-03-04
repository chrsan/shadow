package internal

import (
	"unsafe"
)

func f644(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
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
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s0i32 = f1624(ctx, s0i32)
	l3 = s0i32
	s0i32 = l2
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l5 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l1 = s1i32
		if s1i32 != 0 {
			goto lbl2
		}
		s1i32 = 52
		l1 = s1i32
		s1i32 = 20
		goto lbl0
	}
	s1i32 = l3
	s2i32 = l5
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 2 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32))(table[s3i32].f()))(ctx, s1i32, s2i32)
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
lbl2:
	s1i32 = l3
	s2i32 = l1
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+44)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 2 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32))(table[s3i32].f()))(ctx, s1i32, s2i32)
lbl1:
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 52
	s1i32 = s1i32 + s2i32
	l1 = s1i32
	s1i32 = 28
lbl0:
	s2i32 = l1
	s1i32 = s1i32 + s2i32
	f1092(ctx, s0i32, s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l1 = s0i32
	s1i32 = 1936876899
	s2i32 = 52
	s3i32 = l0
	s0i32 = f499(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1701208948
		s2i32 = l0
		s3i32 = 0
		s0i32 = f499(ctx, s0i32, s1i32, s2i32, s3i32)
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+44)]))
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l1
	s1i32 = l1
	s2i32 = 4
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	s3i32 = -4
	s2i32 = s2i32 + s3i32
	s3i32 = 0
	s4i32 = 22140
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s1i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s4i32].f()))(ctx, s1i32, s2i32, s3i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s0i32 = l3
	s0i32 = f591(ctx, s0i32)
	s0i32 = l4
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
