package internal

import (
	"unsafe"
)

func f486(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var l10 int64
	_ = l10
	var l11 int64
	_ = l11
	var l12 int64
	_ = l12
	var l13 int64
	_ = l13
	var l14 int64
	_ = l14
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
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	l5 = s0i32
	s1i32 = l1
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l8 = s2i32
	s1i32 = s1i32 - s2i32
	l4 = s1i32
	s2i32 = l5
	s1i32 = s1i32 + s2i32
	l5 = s1i32
	s2i32 = l4
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1i32 = l2
	s2i32 = 1
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	l4 = s1i32
	s2i32 = l2
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1i32 = l4
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	s3i32 = l3
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+8)]))
	l9 = s3i32
	s2i32 = s2i32 - s3i32
	l4 = s2i32
	s1i32 = s1i32 + s2i32
	l6 = s1i32
	s2i32 = l4
	if uint32(s1i32) >= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	l4 = s0i32
	s0i32 = l5
	s1i32 = l6
	s0i32 = s0i32 | s1i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = 0
	s2i32 = l2
	s1i32 = s1i32 - s2i32
	s1i64 = int64(s1i32)
	l13 = s1i64
	s2i32 = l9
	s2i64 = int64(s2i32)
	l11 = s2i64
	s3i64 = 2147483647
	s4i64 = l11
	s5i64 = 2147483647
	if s4i64 < s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l10 = s2i64
	s3i64 = -2147483647
	s4i64 = l10
	s5i64 = -2147483647
	if s4i64 > s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i64)
	s0i32 = l0
	s1i32 = 0
	s2i32 = l1
	s1i32 = s1i32 - s2i32
	s1i64 = int64(s1i32)
	l14 = s1i64
	s2i32 = l8
	s2i64 = int64(s2i32)
	l12 = s2i64
	s3i64 = 2147483647
	s4i64 = l12
	s5i64 = 2147483647
	if s4i64 < s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	l10 = s2i64
	s3i64 = -2147483647
	s4i64 = l10
	s5i64 = -2147483647
	if s4i64 > s5i64 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i64 = s2i64
	} else {
		s2i64 = s3i64
	}
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i64)
	s0i32 = l0
	s1i32 = l6
	s1i64 = int64(uint32(s1i32))
	s2i64 = l11
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	s2i64 = l13
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i64)
	s0i32 = l0
	s1i32 = l5
	s1i64 = int64(uint32(s1i32))
	s2i64 = l12
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	s2i64 = l14
	s1i64 = s1i64 + s2i64
	l10 = s1i64
	s2i64 = 2147483647
	s3i64 = l10
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
	l10 = s1i64
	s2i64 = -2147483647
	s3i64 = l10
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
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i64)
	s0i32 = l5
	l7 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l7
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l4
	s1i32 = l5
	s1i64 = int64(uint32(s1i32))
	s2i32 = l6
	s2i64 = int64(uint32(s2i32))
	s1i64 = s1i64 * s2i64
	l10 = s1i64
	s2i64 = 32
	s1i64 = int64(uint64(s1i64) >> (uint64(s2i64) & 63))
	s1i32 = int32(s1i64)
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s0i32 = s0i32 & s1i32
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i64 = l10
	s1i32 = int32(s1i64)
	s2i32 = 3
	s1i32 = s1i32 + s2i32
	s2i32 = -4
	s1i32 = s1i32 & s2i32
	s2i32 = 2
	s1i32 = f44(ctx, s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
lbl2:
}
