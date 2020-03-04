package internal

import (
	"unsafe"
)

func f1381(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
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
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l3
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l4 = s0i32
	s0i32 = l3
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l4
	s2i32 = l5
	s1i32 = s1i32 - s2i32
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = 2
	l5 = s0i32
	s0i32 = l3
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	l1 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+60)]))
	s3i32 = l1
	s4i32 = l5
	s5i32 = l3
	s6i32 = 12
	s5i32 = s5i32 + s6i32
	s2i32 = f3(ctx, s2i32, s3i32, s4i32, s5i32)
	l4 = s2i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		goto lbl4
	}
	s1i32 = 29100
	s2i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = -1
lbl4:
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
		s1i32 = -1
		goto lbl3
	}
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
lbl3:
	l4 = s1i32
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		goto lbl2
	}
	s0i32 = l4
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 0
	s1i32 = l5
	s2i32 = 2
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl2
	}
	s0i32 = l2
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 - s1i32
lbl2:
	l0 = s0i32
	s0i32 = l3
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
lbl1:
	s0i32 = l1
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = l4
	s3i32 = l1
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
	l7 = s3i32
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
	s1i32 = l4
	s2i32 = l7
	s3i32 = 0
	s4i32 = l8
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s1i32 = s1i32 - s2i32
	l7 = s1i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s2i32 = l7
	s1i32 = s1i32 - s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = l4
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	s0i32 = l5
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	goto lbl0
	panic("unreachable executed")
	panic("unreachable executed")
}
