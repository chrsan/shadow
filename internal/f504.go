package internal

import (
	"unsafe"
)

func f504(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 int64
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
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 4512
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])))
	s1i32 = l1
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
	s0i64 = s0i64 - s1i64
	l7 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])))
	s1i32 = l1
	s1i64 = int64(*(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])))
	s0i64 = s0i64 - s1i64
	l8 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l7
	s1i64 = l8
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4480)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l4 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	s1i32 = 4480
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l4
	s3i32 = l3
	s4i32 = 4480
	s3i32 = s3i32 + s4i32
	s4i32 = l1
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+32)]))
	s6i32 = 0
	s7i32 = l4
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+0)]))
	s7i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s7i32+32)]))
	if int(s7i32) < 0 || int(s7i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s7i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s7i32].numParams != 5 {
		panic("argument count mismatch")
	}
	s2i32 = (*(*func(*Context, int32, int32, int32, int32, int32) int32)(table[s7i32].f()))(ctx, s2i32, s3i32, s4i32, s5i32, s6i32)
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4480)]))
lbl1:
	l5 = s0i32
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1112)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 4448
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = 1112
	s1i32 = s1i32 + s2i32
	s2i32 = 4
	s1i32 = s1i32 | s2i32
	s2i32 = 3332
	s3i32 = 3332
	s0i32 = f59(ctx, s0i32, s1i32, s2i32, s3i32)
	l6 = s0i32
	s0i32 = l3
	s1i32 = 1112
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 0
	s3i32 = l2
	s0i32 = f182(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1112)]))
	l4 = s0i32
	s0i32 = l3
	s0i32 = f152(ctx, s0i32)
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l0 = s0i32
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l0
		s2i32 = l4
		f150(ctx, s0i32, s1i32, s2i32)
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
		l4 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1100)]))
		l0 = s0i32
	}
	s0i32 = l4
	s1i32 = l1
	s2i32 = l0
	f1213(ctx, s0i32, s1i32, s2i32)
	s0i32 = l2
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l2
	f40(ctx, s0i32)
	s0i32 = l6
	f43(ctx, s0i32)
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	f13(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 4512
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
