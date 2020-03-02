package internal

import (
	"unsafe"
)

func f244(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var l12 int64
	_ = l12
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s0i64 = int64(s0i32)
	s1i32 = l1
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l11 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s0i64 = int64(s0i32)
	s1i32 = l2
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l12 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l11
	s1i64 = l12
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
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l7 = s0i32
	s1i32 = l4
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0i32
	s1i32 = l3
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0i32
	s1i32 = l2
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l9 = s0i32
	s1i32 = l1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l6
	s0i64 = int64(s0i32)
	s1i32 = l9
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l11 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l7
	s0i64 = int64(s0i32)
	s1i32 = l5
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l12 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l11
	s1i64 = l12
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
		goto lbl0
	}
	s0i32 = 0
	l6 = s0i32
	s0i32 = 0
	l0 = s0i32
	s0i32 = l2
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l10 = s0i32
	s1i32 = l7
	s2i32 = l5
	s1i32 = s1i32 - s2i32
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 12
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		l0 = s0i32
	lbl2:
		s0i32 = l0
		l2 = s0i32
		s1i32 = 8
		s0i32 = s0i32 + s1i32
		l0 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l6 = s0i32
		s1i32 = l10
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
		s0i32 = l5
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l6 = s0i32
		s0i32 = l2
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
		s1i32 = l7
		s2i32 = l8
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s0i32 = s0i32 + s1i32
		l0 = s0i32
	}
	s0i32 = l6
	s1i32 = l4
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s1i32 = l9
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+0)])
	l5 = s1i32
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		l4 = s0i32
		goto lbl3
	}
lbl5:
	s0i32 = l2
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	l5 = s0i32
	s0i32 = l0
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	l0 = s0i32
	s0i32 = l2
	s1i32 = l5
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
lbl3:
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	s1i32 = 255
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l2
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s0i32 = l3
	s1i32 = l1
	s0i32 = s0i32 - s1i32
	l2 = s0i32
lbl6:
	s0i32 = l2
	s1i32 = l0
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 1
		return s0i32
	}
	s0i32 = l2
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	l1 = s0i32
	s0i32 = l4
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	l0 = s0i32
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = l1
	s1i32 = 255
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl6
	}
lbl0:
	s0i32 = 0
	return s0i32
}
