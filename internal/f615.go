package internal

import (
	"unsafe"
)

func f615(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l2 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		l0 = s1i32
		s2i32 = 4
		s1i32 = s1i32 & s2i32
		if s1i32 != 0 {
			s1i32 = 0
		} else {
			s1i32 = l2
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
			l1 = s1i32
			s1i32 = l2
			s2i32 = l0
			s3i32 = 8
			s2i32 = s2i32 & s3i32
			if s2i32 == 0 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				s2i32 = l1
				s3i32 = l0
				s4i32 = 3
				s3i32 = s3i32 & s4i32
				s4i32 = 14024
				s3i32 = s3i32 + s4i32
				s3i32 = int32(ctx.Mem[int(s3i32+0)])
				s2i32 = s2i32 * s3i32
				s3i32 = 2
				s2i32 = s2i32 << (uint32(s3i32) & 31)
				l0 = s2i32
				s2i32 = l1
				s3i32 = 1
				s2i32 = s2i32 << (uint32(s3i32) & 31)
				s3i32 = 39
				s2i32 = s2i32 + s3i32
				l3 = s2i32
				goto lbl3
			}
			s2i32 = l1
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			l4 = s2i32
			s3i32 = l0
			s4i32 = 3
			s3i32 = s3i32 & s4i32
			s4i32 = 14024
			s3i32 = s3i32 + s4i32
			s3i32 = int32(ctx.Mem[int(s3i32+0)])
			l5 = s3i32
			s2i32 = s2i32 * s3i32
			l0 = s2i32
			s2i32 = l1
			s3i32 = 1
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			l6 = s2i32
			s3i32 = 39
			s2i32 = s2i32 + s3i32
			l3 = s2i32
			s2i32 = l2
			s3i32 = l6
			s4i32 = 3
			s3i32 = s3i32 + s4i32
			s4i32 = -4
			s3i32 = s3i32 & s4i32
			s2i32 = s2i32 + s3i32
			s3i32 = l1
			s4i32 = l5
			s3i32 = s3i32 * s4i32
			s4i32 = 2
			s3i32 = s3i32 << (uint32(s4i32) & 31)
			s2i32 = s2i32 + s3i32
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
			l1 = s2i32
			if s2i32 == 0 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				goto lbl3
			}
			s2i32 = l1
			s3i32 = l4
			s2i32 = s2i32 + s3i32
			s3i32 = 4
			s2i32 = s2i32 + s3i32
			goto lbl2
		lbl3:
			s2i32 = 0
		lbl2:
			s3i32 = l0
			s4i32 = l3
			s5i32 = 3
			s4i32 = s4i32 | s5i32
			s3i32 = s3i32 + s4i32
			s2i32 = s2i32 + s3i32
			s3i32 = -4
			s2i32 = s2i32 & s3i32
			s1i32 = s1i32 + s2i32
		}
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
}
