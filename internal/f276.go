package internal

import (
	"unsafe"
)

func f276(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
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
	var s2i64 int64
	_ = s2i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 6
		return s0i32
	}
	s0i32 = 1
	l4 = s0i32
	s0i32 = l0
	s1i32 = l3
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	s0i32 = l0
	s1i32 = l3
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l5 = s1i32
	s2i32 = 4
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
		l3 = s1i32
		s1i32 = l5
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		switch s1i32 {
		case 0:
			goto lbl6
		case 1:
			goto lbl4
		case 2:
			goto lbl5
		case 3:
			goto lbl3
		default:
			goto lbl2
		}
	lbl6:
		s1i32 = l1
		s2i32 = l2
		s3i32 = -8
		s2i32 = s2i32 + s3i32
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		goto lbl2
	lbl5:
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
		s3i32 = 4
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
	lbl4:
		s1i32 = l1
		s2i32 = l2
		s3i32 = -8
		s2i32 = s2i32 + s3i32
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = l2
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1i32 = l2
		s2i32 = 8
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1i32 = 2
		l4 = s1i32
		goto lbl2
	lbl3:
		s1i32 = l1
		s2i32 = l2
		s3i32 = -8
		s2i32 = s2i32 + s3i32
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = l2
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = l2
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1i32 = l2
		s2i32 = 16
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1i32 = 3
		l4 = s1i32
	lbl2:
		s1i32 = l1
		s2i32 = l3
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l2
		s2i32 = l4
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
	} else {
		s1i32 = l2
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l5
	return s0i32
}
