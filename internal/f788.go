package internal

import (
	"unsafe"
)

func f788(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
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
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l1 = s0i32
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l3 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 1
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l1
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		if int(s1i32) < 0 || int(s1i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s1i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s1i32].numParams != 1 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32))(table[s1i32].f()))(ctx, s0i32)
	lbl2:
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		f90(ctx, s0i32)
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+32)] = uint8(s1i32)
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l4 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l5 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l6 = s0i32
	s0i32 = l2
	s1i64 = 8589934593
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l3
	s2i32 = l6
	s1i32 = s1i32 - s2i32
	s1i64 = int64(uint32(s1i32))
	s2i32 = l5
	s3i32 = l4
	s2i32 = s2i32 - s3i32
	s2i64 = int64(uint32(s2i32))
	s3i64 = 32
	s2i64 = s2i64 << (uint64(s3i64) & 63)
	s1i64 = s1i64 | s2i64
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
	s4i32 = 0
	s5i32 = 0
	s0i32 = f157(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
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
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	f12(ctx, s0i32)
lbl0:
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
