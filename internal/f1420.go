package internal

import (
	"unsafe"
)

func f1420(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 float64
	_ = l7
	var l8 float64
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f64 float64
	_ = s0f64
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	s0i32 = l1
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l5 = s0i32
	if s0i32 != 0 {
	lbl1:
		s0i32 = l0
		s1i32 = l5
		l2 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		l5 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l7 = s0f64
		s0i32 = l5
		l4 = s0i32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		l3 = s0i32
		s1i32 = l1
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	lbl3:
		s0i32 = l2
		s1i32 = -1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = l1
		if uint32(s0i32) < uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l0
			s2i32 = l3
			s3i32 = 3
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = s1i32 + s2i32
			l4 = s1i32
			s2i32 = -8
			s1i32 = s1i32 + s2i32
			s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l4
			s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
			if s1f64 < s2f64 {
				s1i32 = 1
			} else {
				s1i32 = 0
			}
			s0i32 = s0i32 | s1i32
			l3 = s0i32
		}
		s0f64 = l7
		s1i32 = l0
		s2i32 = l3
		s3i32 = -1
		s2i32 = s2i32 + s3i32
		l4 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s1f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1f64
		if s0f64 < s1f64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 1
		s0i32 = s0i32 ^ s1i32
		if s0i32 != 0 {
			s0i32 = l2
			l4 = s0i32
			goto lbl2
		}
		s0i32 = l0
		s1i32 = l2
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1f64 = l8
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l3
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
			goto lbl3
		}
	lbl2:
		s0i32 = l0
		s1i32 = l4
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1f64 = l7
		*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
		s0i32 = l5
		if s0i32 != 0 {
			goto lbl1
		}
	}
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
lbl8:
	s0i32 = l0
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l6 = s0i64
	s0i32 = l0
	s1i32 = l0
	s2i32 = l1
	s3i32 = 3
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	l2 = s1i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l2
	s1i64 = l6
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l1
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl7
	}
	s0i32 = l0
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l7 = s0f64
	s0i32 = 2
	l2 = s0i32
	s0i32 = 1
	l4 = s0i32
lbl9:
	s0i32 = l4
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s1i32 = l2
	s2i32 = l1
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l2
	} else {
		s1i32 = l2
		s2i32 = l0
		s3i32 = l2
		s4i32 = 3
		s3i32 = s3i32 << (uint32(s4i32) & 31)
		s2i32 = s2i32 + s3i32
		l3 = s2i32
		s3i32 = -8
		s2i32 = s2i32 + s3i32
		s2f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s3i32 = l3
		s3f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		if s2f64 < s3f64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s1i32 = s1i32 | s2i32
	}
	l3 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = l0
	s1i32 = s1i32 + s2i32
	s2i32 = -8
	s1i32 = s1i32 + s2i32
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
	s0i32 = l3
	l4 = s0i32
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l2 = s0i32
	s1i32 = l1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl9
	}
	s0i32 = l4
	l2 = s0i32
	s0i32 = l3
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
lbl12:
	s0i32 = l4
	l2 = s0i32
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s0f64 = *(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l8 = s0f64
	s1f64 = l7
	if s0f64 < s1f64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		s0i32 = l3
		l2 = s0i32
		goto lbl11
	}
	s0i32 = l3
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s1f64 = l8
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l2
	l3 = s0i32
	s1i32 = 1
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	l4 = s0i32
	if s0i32 != 0 {
		goto lbl12
	}
lbl11:
	s0i32 = l2
	s1i32 = 3
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l0
	s0i32 = s0i32 + s1i32
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	s1f64 = l7
	*(*float64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f64
	s0i32 = l1
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	if s0i32 != 0 {
		goto lbl8
	}
lbl7:
}
