package internal

import (
	"unsafe"
)

func f1595(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s2i64 int64
	_ = s2i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
	if int(s2i32) < 0 || int(s2i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s2i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s2i32].numParams != 2 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32))(table[s2i32].f()))(ctx, s0i32, s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l2 = s1i32
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l2
		s2i32 = 0
		ctx.Mem[int(s1i32+32)] = uint8(s2i32)
		s1i32 = l2
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
		s1i32 = l2
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
		s1i32 = l2
		s2i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l2
		s2i32 = 0
		s3i32 = l1
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+36)]))
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
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l2
			f1384(ctx, s0i32, s1i32)
			goto lbl1
		}
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	lbl1:
		s0i32 = l2
		s0i32 = f41(ctx, s0i32)
	}
	s0i32 = l3
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
