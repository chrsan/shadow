package internal

import (
	"unsafe"
)

func f48(ctx *Context) int32 {
	var l0 int32
	_ = l0
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = 29092
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
	} else {
		s0i32 = 29084
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		l0 = s0i32
		s1i32 = 2
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		switch s0i32 {
		case 0:
			goto lbl2
		case 1:
			goto lbl1
		default:
			goto lbl3
		}
	lbl3:
		s0i32 = 29084
		s1i32 = 29084
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		l0 = s1i32
		s2i32 = 1
		s3i32 = l0
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = 4
		s0i32 = f17(ctx, s0i32)
		l0 = s0i32
		s1i32 = 27268
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = 29084
		s1i32 = 2
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = 29088
		s1i32 = l0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		goto lbl1
	lbl2:
	lbl4:
		s0i32 = 29084
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		s1i32 = 2
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
	lbl1:
		s0i32 = 29088
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	}
	return s0i32
}
