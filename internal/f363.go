package internal

import (
	"unsafe"
)

func f363(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l6 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		l0 = s0i32
		goto lbl0
	}
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l4 = s0i32
	s1i32 = -1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l2
		s2i32 = l3
		s0i32 = f25(ctx, s0i32, s1i32, s2i32)
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	} else {
		s0i32 = l4
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l4 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l7 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l9 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l5 = s0i32
		s0i64 = int64(s0i32)
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l3 = s1i32
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l10 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i64 = l9
		s1i64 = l10
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967295
		if uint64(s0i64) > uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l8 = s0i32
		s1i32 = l5
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
		l5 = s0i32
		s1i32 = l4
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		l4 = s0i32
		s1i32 = l3
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s1i32 = l7
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l5
		s0i64 = int64(s0i32)
		s1i32 = l3
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l9 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l8
		s0i64 = int64(s0i32)
		s1i32 = l4
		s1i64 = int64(s1i32)
		s0i64 = s0i64 - s1i64
		l10 = s0i64
		s1i64 = 1
		if s0i64 < s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i64 = l9
		s1i64 = l10
		s0i64 = s0i64 | s1i64
		s1i64 = 2147483648
		s0i64 = s0i64 + s1i64
		s1i64 = 4294967296
		if uint64(s0i64) >= uint64(s1i64) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l1
		l0 = s0i32
		goto lbl0
	lbl4:
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = l2
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint64(s1i64)
		s0i32 = l0
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 44
	s0i32 = s0i32 + s1i32
	l0 = s0i32
lbl0:
	s0i32 = l6
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
