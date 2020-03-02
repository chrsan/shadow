package internal

import (
	"unsafe"
)

func f617(ctx *Context, l0 int32, l1 int32, l2 float32) int32 {
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
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l3 = s0i32
	s0i32 = l0
	s1f32 = l2
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = l0
	s1i32 = l3
	s2i32 = 6
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l3 = s1i32
	s2i32 = 3
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		switch s1i32 {
		case 0:
			goto lbl3
		case 1:
			goto lbl2
		default:
			goto lbl4
		}
	lbl4:
		s1i32 = l0
		s2i32 = -1082130432
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = 2147483647
		s1i32 = s1i32 & s2i32
		goto lbl0
	lbl3:
		s1i32 = l0
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = 2147483647
		s1i32 = s1i32 & s2i32
		goto lbl0
	lbl2:
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		l2 = s1f32
		s2f32 = 0
		if s1f32 == s2f32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l0
			s2i32 = -1082130432
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
			s2i32 = 2147483647
			s1i32 = s1i32 & s2i32
			goto lbl0
		}
		s1i32 = l0
		s2f32 = l2
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = s2f32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		s2i32 = -2147483648
		s1i32 = s1i32 | s2i32
		goto lbl0
	}
	s1i32 = l0
	s2i32 = -1082130432
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = 2147483647
	s1i32 = s1i32 & s2i32
lbl0:
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 2
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	s2i32 = l3
	s3i32 = -65536
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	l3 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2i32 = 12
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 196608
	s1i32 = s1i32 & s2i32
	s2i32 = l3
	s3i32 = -2147483645
	s2i32 = s2i32 & s3i32
	s1i32 = s1i32 | s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l0
	return s0i32
}
