package internal

import (
	"unsafe"
)

func f1800(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l0
	s2i32 = -64
	s1i32 = s1i32 - s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 - s1i32
		l1 = s0i32
		s1i32 = 31
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = l1
		s0i32 = s0i32 & s1i32
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		s2i32 = l1
		s3i32 = l1
		s4i32 = 0
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		l1 = s1i32
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l3 = s0i32
		s0i32 = l1
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l4 = s0i32
		s0i32 = l1
		s1i32 = 2
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		l1 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+56)]))
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+76)]))
		s3i32 = l5
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 ^ s1i32
		s1i32 = l6
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = -1
		if s0i32 <= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l4
			s0i32 = s0i32 - s1i32
			s1i32 = 4
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			l3 = s0i32
			goto lbl1
		}
		s0i32 = l0
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l4
		s3i32 = 4
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 - s2i32
		s2i32 = -64
		s1i32 = s1i32 - s2i32
		l4 = s1i32
		s2i32 = l4
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = 64
		s1i32 = l2
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		s0i32 = s0i32 - s1i32
		l4 = s0i32
		s0i32 = l3
		s1i32 = 4
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l3 = s0i32
		s0i32 = l1
		s1i32 = 16
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			if s0i32 != 0 {
			lbl5:
				s0i32 = l0
				s1i32 = l0
				s1i32 = int32(ctx.Mem[int(s1i32+0)])
				s2i32 = l4
				s1i32 = s1i32 + s2i32
				ctx.Mem[int(s0i32+0)] = uint8(s1i32)
				s0i32 = l1
				s1i32 = -1
				s0i32 = s0i32 + s1i32
				l1 = s0i32
				s0i32 = l0
				s1i32 = 1
				s0i32 = s0i32 + s1i32
				l0 = s0i32
				s1i32 = 3
				s0i32 = s0i32 & s1i32
				if s0i32 != 0 {
					goto lbl5
				}
			}
			s0i32 = l4
			s1i32 = 8
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l4
			s0i32 = s0i32 | s1i32
			l2 = s0i32
			s1i32 = 16
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			s1i32 = l2
			s0i32 = s0i32 | s1i32
			l5 = s0i32
			s0i32 = l1
			s1i32 = 2
			s0i32 = s0i32 >> (uint32(s1i32) & 31)
			l2 = s0i32
		lbl6:
			s0i32 = l0
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l5
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l0
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l0 = s0i32
			s0i32 = l2
			s1i32 = 1
			if s0i32 > s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			l6 = s0i32
			s0i32 = l2
			s1i32 = -1
			s0i32 = s0i32 + s1i32
			l2 = s0i32
			s0i32 = l6
			if s0i32 != 0 {
				goto lbl6
			}
			s0i32 = l1
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			l1 = s0i32
		}
		s0i32 = l1
		s1i32 = 1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	lbl7:
		s0i32 = l0
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l1
		s1i32 = 1
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l2 = s0i32
		s0i32 = l1
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l2
		if s0i32 != 0 {
			goto lbl7
		}
	lbl1:
		s0i32 = l0
		s1i32 = l3
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+0)])
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s2i32 = l0
		s3i32 = 8
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 - s2i32
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	}
}
