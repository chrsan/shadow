package internal

import (
	"unsafe"
)

func f1276(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 13195221663744
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l6
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l6
	s1i32 = l4
	s2i32 = 0
	f454(ctx, s0i32, s1i32, s2i32)
	s0i32 = l6
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = -261122
	s1i32 = s1i32 & s2i32
	s2i32 = l5
	s3i32 = 10
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = 261120
	s2i32 = s2i32 & s3i32
	s3i32 = l3
	s4i32 = 15
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	s2i32 = s2i32 | s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint32(s1i32)
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l7
		s0i32 = f37(ctx, s0i32)
		l1 = s0i32
		s1i32 = l2
		s2i32 = 4
		s3i32 = 1
		f277(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l6
		s3i32 = 0
		s4i32 = l0
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+104)]))
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
		s0i32 = l1
		f34(ctx, s0i32)
		goto lbl1
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l6
	s3i32 = l0
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+80)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32))(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
lbl1:
	s0i32 = l6
	s0i32 = f23(ctx, s0i32)
	s0i32 = l7
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
}
