package internal

import (
	"unsafe"
)

func f606(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
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
	var s4i32 int32
	_ = s4i32
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s10i32 int32
	_ = s10i32
	var s11i32 int32
	_ = s11i32
	var s1i64 int64
	_ = s1i64
	var s3i64 int64
	_ = s3i64
	var s5i64 int64
	_ = s5i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	ctx.g0 = s0i32
	s0i32 = 27944
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 27944
		s1i32 = f48(ctx)
		l3 = s1i32
		s2i32 = 1667
		s3i32 = l3
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		if int(s3i32) < 0 || int(s3i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s3i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s3i32].numParams != 2 {
			panic("argument count mismatch")
		}
		s1i32 = (*(*func(*Context, int32, int32) int32)(table[s3i32].f()))(ctx, s1i32, s2i32)
		l3 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l2
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	s1i32 = 5
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = f48(ctx)
		l4 = s1i32
		s2i32 = 88
		s3i32 = l3
		s4i32 = 1875
		s5i64 = 0
		s6i32 = 0
		s7i32 = 0
		s8i32 = 0
		s9i32 = 0
		s10i32 = 0
		s11i32 = l4
		s11i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s11i32+0)]))
		s11i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s11i32+16)]))
		if int(s11i32) < 0 || int(s11i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s11i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s11i32].numParams != 10 {
			panic("argument count mismatch")
		}
		s1i64 = (*(*func(*Context, int32, int32, int32, int32, int64, int32, int32, int32, int32, int32) int64)(table[s11i32].f()))(ctx, s1i32, s2i32, s3i32, s4i32, s5i64, s6i32, s7i32, s8i32, s9i32, s10i32)
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = 1875
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l1
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+92)]))
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
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = f48(ctx)
	l0 = s0i32
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s3i32 = l2
	s3i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s3i32+24)]))
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
	if int(s4i32) < 0 || int(s4i32) >= len(table) {
		panic("table entry out of bounds")
	}
	if table[s4i32].numParams == -1 {
		panic("table entry is nil")
	}
	if table[s4i32].numParams != 4 {
		panic("argument count mismatch")
	}
	(*(*func(*Context, int32, int32, int32, int64))(table[s4i32].f()))(ctx, s0i32, s1i32, s2i32, s3i64)
lbl2:
	s0i32 = l2
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
