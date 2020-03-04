package internal

import (
	"unsafe"
)

func f939(ctx *Context, l0 int32, l1 int32) {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int64
	_ = l5
	var l6 int64
	_ = l6
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
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i64 = int64(s0i32)
	l5 = s0i64
	s1i64 = 3
	s0i64 = s0i64 * s1i64
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	l3 = s1i32
	s1i64 = int64(s1i32)
	l6 = s1i64
	if s0i64 >= s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	l4 = s0i32
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 2
	s0i32 = s0i32 & s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l2 = s0i32
lbl0:
	s0i32 = 0
	s1i32 = l1
	s2i32 = l3
	if s1i32 <= s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l2
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i64 = l5
	s1i64 = l5
	s2i64 = 1
	s1i64 = s1i64 + s2i64
	s2i64 = 1
	s1i64 = s1i64 >> (uint64(s2i64) & 63)
	s0i64 = s0i64 + s1i64
	s1i64 = 7
	s0i64 = s0i64 + s1i64
	s1i64 = -8
	s0i64 = s0i64 & s1i64
	l5 = s0i64
	s1i64 = l6
	if s0i64 == s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s1i64 = l5
	s2i64 = 2147483647
	s3i64 = l5
	s4i64 = 2147483647
	if s3i64 < s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	l5 = s1i64
	s2i64 = -2147483647
	s3i64 = l5
	s4i64 = -2147483647
	if s3i64 > s4i64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1i64 = s1i64
	} else {
		s1i64 = s2i64
	}
	s1i32 = int32(s1i64)
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = f50(ctx, s0i32, s1i32)
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l3 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = 0
		l1 = s0i32
	lbl3:
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l4
		s1i32 = s1i32 + s2i32
		s1i32 = int32(ctx.Mem[int(s1i32+0)])
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	l1 = s0i32
	s1i32 = 1
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		f13(ctx, s0i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+12)])
		l1 = s0i32
	}
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 252
	s1i32 = s1i32 & s2i32
	s2i32 = 1
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+12)] = uint8(s1i32)
lbl1:
}
