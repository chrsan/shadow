package internal

import (
	"unsafe"
)

func f252(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l3 = s0i32
	s1i32 = 1
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s1i32 = 536870911
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	s1i32 = 536870910
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s1i32 = -15
	s0i32 = s0i32 + s1i32
	s1i32 = 6
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	goto lbl0
lbl1:
	s0i32 = l4
	s1i32 = l1
	s2i32 = 0
	f241(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l5 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl0
	}
	s0i32 = 76
	s0i32 = f17(ctx, s0i32)
	l3 = s0i32
	s0i32 = l4
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s3i32 = l2
	f1571(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l3
	s1i32 = 26392
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 36
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = 0
	ctx.Mem[int(s0i32+32)] = uint8(s1i32)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s1i32 = l1
	s2i32 = l5
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
	s0i32 = f141(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = l4
	s2i32 = 8
	s1i32 = s1i32 + s2i32
	s2i32 = 0
	s3i32 = 0
	f344(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l2 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
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
lbl5:
	s0i32 = l3
	s1i32 = 1
	ctx.Mem[int(s0i32+72)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
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
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l1 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
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
	s1i32 = l0
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
lbl0:
	s0i32 = l4
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
