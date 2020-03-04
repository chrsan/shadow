package internal

import (
	"unsafe"
)

func f1893(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int64
	_ = l11
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	s0i32 = l0
	s1i32 = l1
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	l4 = s1i32
	s2i32 = 20
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l7 = s0i32
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0i32
	s0i32 = l3
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l11 = s0i64
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l3 = s1i32
	s2i32 = l2
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l8
		s2i32 = l7
		s3i32 = l8
		s4i32 = l7
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
		l10 = s1i32
	lbl1:
		s1i32 = l0
		s2i32 = l3
		s3i32 = l2
		if uint32(s2i32) >= uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			s2i32 = l3
		} else {
			s2i32 = l0
			s3i32 = l3
			s4i32 = 20
			s3i32 = s3i32 * s4i32
			s2i32 = s2i32 + s3i32
			l4 = s2i32
			s3i32 = -20
			s2i32 = s2i32 + s3i32
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			l5 = s2i32
			s3i32 = l4
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
			l6 = s3i32
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				s2i32 = l0
				s3i32 = l3
				s4i32 = -1
				s3i32 = s3i32 + s4i32
				s4i32 = 20
				s3i32 = s3i32 * s4i32
				s2i32 = s2i32 + s3i32
				l5 = s2i32
				s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
				l6 = s2i32
				s3i32 = l5
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
				l5 = s3i32
				s4i32 = l6
				s5i32 = l5
				if s4i32 < s5i32 {
					s4i32 = 1
				} else {
					s4i32 = 0
				}
				if s4i32 != 0 {
					// s2i32 = s2i32
				} else {
					s2i32 = s3i32
				}
				s3i32 = l4
				s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
				l5 = s3i32
				s4i32 = l4
				s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+4)]))
				l4 = s4i32
				s5i32 = l5
				s6i32 = l4
				if s5i32 < s6i32 {
					s5i32 = 1
				} else {
					s5i32 = 0
				}
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
				goto lbl5
			}
			s2i32 = l5
			s3i32 = l6
			if s2i32 < s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
		lbl5:
			s3i32 = l3
			s2i32 = s2i32 | s3i32
		}
		l4 = s2i32
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		s3i32 = 20
		s2i32 = s2i32 * s3i32
		s1i32 = s1i32 + s2i32
		l3 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l5 = s1i32
		s2i32 = l9
		if s1i32 == s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l10
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
			l5 = s2i32
			s3i32 = l3
			s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
			l6 = s3i32
			s4i32 = l5
			s5i32 = l6
			if s4i32 < s5i32 {
				s4i32 = 1
			} else {
				s4i32 = 0
			}
			if s4i32 != 0 {
				// s2i32 = s2i32
			} else {
				s2i32 = s3i32
			}
			if s1i32 < s2i32 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			if s1i32 != 0 {
				goto lbl3
			}
			s1i32 = l1
			l4 = s1i32
			goto lbl2
		}
		s1i32 = l9
		s2i32 = l5
		if s1i32 < s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl3
		}
		s1i32 = l1
		l4 = s1i32
		goto lbl2
	lbl3:
		s1i32 = l1
		s2i32 = 20
		s1i32 = s1i32 * s2i32
		s2i32 = l0
		s1i32 = s1i32 + s2i32
		s2i32 = -20
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = l3
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
		s1i32 = l1
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l3
		s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
		s1i32 = l4
		l1 = s1i32
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		l3 = s1i32
		s2i32 = l2
		if uint32(s1i32) <= uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl1
		}
	lbl2:
		s1i32 = l4
		s2i32 = -1
		s1i32 = s1i32 + s2i32
	} else {
		s1i32 = l4
	}
	s2i32 = 20
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i64 = l11
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l8
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l9
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
