package internal

import (
	"unsafe"
)

func f1294(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 27996
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l1 = s0i32
		s1i32 = 2
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl3
		case 1:
			goto lbl2
		default:
			goto lbl4
		}
	lbl4:
		s0i32 = 27996
		s1i32 = 27996
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l1 = s1i32
		s2i32 = 1
		s3i32 = l1
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l1
		if s0i32 != 0 {
			goto lbl3
		}
		s0i32 = 20
		s0i32 = f17(ctx, s0i32)
		l2 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
		s0i32 = l2
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
		s0i32 = 27996
		s1i32 = 2
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = 28000
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl0
	lbl3:
	lbl5:
		s0i32 = 27996
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		s1i32 = 2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl5
		}
	lbl2:
		s0i32 = 28000
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l2 = s0i32
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l2
	s1i32 = -20
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		s0i32 = f17(ctx, s0i32)
		l3 = s0i32
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 1
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = l3
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l1
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = l1
			s2i32 = l2
			s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		}
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	f524(ctx)
	panic("unreachable executed")
lbl0:
	s0i32 = l2
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
