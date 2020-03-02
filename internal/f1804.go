package internal

import (
	"unsafe"
)

func f1804(ctx *Context, l0 int32, l1 int32) {
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0i32
	s0i32 = 1
	l2 = s0i32
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l3 = s0i32
	s1i32 = l1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		l4 = s0i32
	lbl1:
		s0i32 = l4
		s1i32 = 2
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l0
		s0i32 = s0i32 + s1i32
		s1i32 = -4
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l1
		if uint32(s1i32) >= uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l3
		} else {
			s1i32 = l3
			s2i32 = l0
			s3i32 = l3
			s4i32 = 2
			s3i32 = s3i32 << (uint32(s4i32) & 31)
			s2i32 = s2i32 + s3i32
			l4 = s2i32
			s3i32 = -4
			s2i32 = s2i32 + s3i32
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			l2 = s2i32
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
			s3i32 = l2
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
			s4i32 = l2
			s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+28)]))
			l5 = s4i32
			s5i32 = l5
			s6i32 = l4
			s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
			l2 = s6i32
			s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+28)]))
			l4 = s6i32
			if s5i32 == s6i32 {
				s5i32 = 1
			} else {
				s5i32 = 0
			}
			l5 = s5i32
			if s5i32 != 0 {
				// s3i32 = s3i32
			} else {
				s3i32 = s4i32
			}
			l6 = s3i32
			s4i32 = l6
			s5i32 = l2
			s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
			s6i32 = l4
			s7i32 = l5
			if s7i32 != 0 {
				// s5i32 = s5i32
			} else {
				s5i32 = s6i32
			}
			l4 = s5i32
			if s4i32 == s5i32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			l5 = s4i32
			if s4i32 != 0 {
				// s2i32 = s2i32
			} else {
				s2i32 = s3i32
			}
			s3i32 = l2
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
			s4i32 = l4
			s5i32 = l5
			if s5i32 != 0 {
				// s3i32 = s3i32
			} else {
				s3i32 = s4i32
			}
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			s1i32 = s1i32 | s2i32
		}
		l2 = s1i32
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s2i32 = l0
		s1i32 = s1i32 + s2i32
		s2i32 = -4
		s1i32 = s1i32 + s2i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		l4 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l3 = s0i32
		s1i32 = l1
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l3 = s0i32
	s1i32 = 1
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l1 = s0i32
		goto lbl3
	}
	s0i32 = l7
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	l4 = s0i32
lbl5:
	s0i32 = l3
	l1 = s0i32
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
	l5 = s2i32
	s3i32 = l4
	s4i32 = l5
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l5 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l6 = s1i32
	s2i32 = l6
	s3i32 = l7
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+12)]))
	s4i32 = l4
	s5i32 = l5
	if s5i32 != 0 {
		// s3i32 = s3i32
	} else {
		s3i32 = s4i32
	}
	l5 = s3i32
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l6 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	s1i32 = l7
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s2i32 = l5
	s3i32 = l6
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l2
		l1 = s0i32
		goto lbl3
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	l2 = s0i32
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l3 = s0i32
	s1i32 = 1
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl3:
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
