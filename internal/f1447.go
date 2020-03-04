package internal

import (
	"unsafe"
)

func f1447(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 80
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
	l4 = s0i32
	s1i32 = l3
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	s2i32 = l2
	s3i32 = l4
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+40)]))
	if int(s3i32) < 0 || int(s3i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s3i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s3i32].numParams != 3 {
		panic("argument count mismatch")
	}
	s0i32 = (*(*func(*Context, int32, int32, int32) int32)(table[s3i32].f()))(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i64 = 69784829952
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l3
	s1i64 = 1065353216
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l3
	s2i32 = 40
	s1i32 = s1i32 + s2i32
	s2i32 = l0
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	f69(ctx, s0i32, s1i32, s2i32)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
lbl0:
	s0i32 = l3
	s1i32 = 80
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l2
	return s0i32
}
