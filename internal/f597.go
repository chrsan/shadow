package internal

import (
	"unsafe"
)

func f597(ctx *Context, l0 int32, l1 int32) {
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l4 = s1i32
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 20
	s0i32 = i32DivS(s0i32, s1i32)
	l7 = s0i32
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = 214748365
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		s1i32 = l2
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l4
		s2i32 = s2i32 - s3i32
		s3i32 = 20
		s2i32 = i32DivS(s2i32, s3i32)
		l3 = s2i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l6 = s2i32
		s3i32 = l6
		s4i32 = l2
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 214748364
		s3i32 = l3
		s4i32 = 107374182
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l3 = s1i32
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl2
		}
		s0i32 = l3
		s1i32 = 214748365
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l3
		s1i32 = 20
		s0i32 = s0i32 * s1i32
		s0i32 = f17(ctx, s0i32)
	lbl2:
		l6 = s0i32
		s1i32 = l7
		s2i32 = 20
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l1
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = l1
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = l5
		s2i32 = -20
		s1i32 = i32DivS(s1i32, s2i32)
		s2i32 = 20
		s1i32 = s1i32 * s2i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l5
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = l4
			s2i32 = l5
			s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l0
		s1i32 = l6
		s2i32 = l3
		s3i32 = 20
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l4
		if s0i32 != 0 {
			s0i32 = l4
			f12(ctx, s0i32)
		}
		return
	}
	f104(ctx)
	panic("unreachable executed")
lbl0:
	f63(ctx)
	panic("unreachable executed")
}
