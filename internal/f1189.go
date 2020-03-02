package internal

import (
	"unsafe"
)

func f1189(ctx *Context, l0 int32) int32 {
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
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l4 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		return s0i32
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 - s1i32
	l9 = s0i32
	s0i32 = l4
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l6 = s1i32
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l7 = s0i32
	s1i32 = l3
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	l8 = s0i32
lbl2:
	s0i32 = l7
	s1i32 = l8
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l9
	l5 = s0i32
lbl3:
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l10 = s0i32
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = l10
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l8
	s1i32 = 8
	s0i32 = s0i32 + s1i32
	l8 = s0i32
	s1i32 = l7
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
lbl1:
	s0i32 = l2
	s1i32 = l6
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l4
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l1 = s1i32
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			f13(ctx, s0i32)
		}
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
		s0i32 = l0
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = 0
		return s0i32
	}
	s0i32 = l2
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = 3
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l8 = s0i32
		s1i32 = -8
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s0i32 = l6
		s1i32 = l2
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			l1 = s0i32
		lbl8:
			s0i32 = l3
			s1i32 = l1
			s2i32 = 3
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			l7 = s0i32
			s1i32 = l7
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
			s2i32 = l5
			s1i32 = s1i32 - s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
			s0i32 = l1
			s1i32 = 1
			s0i32 = s0i32 + s1i32
			l1 = s0i32
			s1i32 = l6
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
		}
		s0i32 = l3
		s1i32 = l8
		s2i32 = l6
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		l6 = s2i32
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l1 = s2i32
		s3i32 = l4
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
		s2i32 = s2i32 + s3i32
		s0i32 = f106(ctx, s0i32, s1i32, s2i32)
		l2 = s0i32
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = l5
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		l7 = s0i32
	}
	s0i32 = l3
	s1i32 = l6
	s2i32 = 3
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	l1 = s0i32
lbl9:
	s0i32 = l1
	s1i32 = -8
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l7
	s1i32 = l1
	s2i32 = -4
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l9
	l5 = s0i32
lbl11:
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 != 0 {
		goto lbl10
	}
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	l8 = s0i32
	s0i32 = l1
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l5
	s1i32 = l8
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	s1i32 = 0
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl11
	}
	s0i32 = l2
	l1 = s0i32
	goto lbl9
lbl10:
	s0i32 = l3
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l1 = s0i32
	s1i32 = 9
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1
		s2i32 = l1
		s3i32 = 3
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 - s2i32
		l1 = s1i32
		s2i32 = 3
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s0i32 = f106(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		s1i32 = s1i32 + s2i32
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
		s0i32 = l4
		s1i32 = l1
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	}
	s0i32 = 1
	return s0i32
}
