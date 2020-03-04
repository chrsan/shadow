package internal

import (
	"unsafe"
)

func f1864(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = l0
	s1i32 = l3
	s1i32 = f1093(ctx, s1i32)
	l3 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = l0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
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
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
		goto lbl0
	}
	s0i32 = l0
	s1i32 = -16777216
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+46)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 16448
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint16(s1i32)
	s0i32 = -16777216
lbl0:
	l3 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+48)])
	s2i32 = 1
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 16
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s2i32 = 255
		s1i32 = s1i32 & s2i32
		s2i32 = 54
		s1i32 = s1i32 * s2i32
		s2i32 = l3
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = 19
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s2i32 = l3
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 255
		s2i32 = s2i32 & s3i32
		s3i32 = 183
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 130816
		s1i32 = s1i32 & s2i32
		s2i32 = l1
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		l1 = s2i32
		s1i32 = s1i32 | s2i32
		s2i32 = l1
		s3i32 = 16
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 | s2i32
		s2i32 = -16777216
		s1i32 = s1i32 | s2i32
		l3 = s1i32
	}
	s1i32 = l3
	s2i32 = 5
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 7
	s1i32 = s1i32 & s2i32
	l0 = s1i32
	s2i32 = 5
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l0
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = l0
	s3i32 = 1
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s3i32 = 13
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 7
	s2i32 = s2i32 & s3i32
	l0 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = l0
	s4i32 = 5
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 | s3i32
	s3i32 = l0
	s4i32 = 1
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s3i32 = 8
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s3i32 = 21
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	s3i32 = 7
	s2i32 = s2i32 & s3i32
	l0 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s3i32 = l0
	s4i32 = 5
	s3i32 = s3i32 << (uint32(s4i32) & 31)
	s2i32 = s2i32 | s3i32
	s3i32 = l0
	s4i32 = 1
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	s2i32 = s2i32 | s3i32
	s3i32 = 16
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 | s2i32
	s2i32 = -16777216
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
}
