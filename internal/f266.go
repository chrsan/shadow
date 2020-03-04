package internal

import (
	"unsafe"
)

func f266(ctx *Context, l0 int32) int32 {
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
	var l11 int32
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	s0i32 = ctx.g0
	s1i32 = 32
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)]))
	l6 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	l7 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+54)])
	l9 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+53)])
	l10 = s0i32
	s0i32 = l0
	s0i32 = int32(int8(ctx.Mem[int(s0i32+52)]))
	l2 = s0i32
lbl0:
	s0i32 = l2
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = l2
	s1i32 = -1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)]))
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
		l1 = s0i32
		l2 = s0i32
		goto lbl2
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+100)]))
	l3 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+108)]))
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
	l2 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+112)]))
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
	l1 = s1i32
	s2i32 = l3
	s3i32 = l10
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
	l3 = s1i32
	s2i32 = l2
	s3i32 = l10
	s2i32 = s2i32 >> (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l9
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l6
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l1
	s1i32 = l9
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	s1i32 = l7
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)]))
	l1 = s0i32
lbl2:
	s0i32 = 0
	s1i32 = l5
	s2i32 = l1
	s3i32 = l6
	s4i32 = l2
	s5i32 = l2
	s6i32 = l6
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
	l6 = s3i32
	s4i32 = -8192
	s3i32 = s3i32 - s4i32
	s4i32 = -16384
	s3i32 = s3i32 & s4i32
	l5 = s3i32
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	l8 = s2i32
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l2 = s0i32
	s0i32 = 2147483647
	s1i32 = l1
	s2i32 = l5
	s3i32 = l8
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	l5 = s1i32
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+124)]))
	l11 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 10
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l1 = s1i32
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s1i32 = l7
	s0i32 = s0i32 - s1i32
	s1i32 = 10
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l8 = s0i32
	s1i32 = 32768
	s0i32 = s0i32 + s1i32
	s1i32 = 65536
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l8
		s1i64 = int64(s1i32)
		s2i64 = 16
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i32 = l1
		s2i64 = int64(s2i32)
		s1i64 = i64DivS(s1i64, s2i64)
		l12 = s1i64
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = -2147483647
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
		s0i32 = l4
		s1i64 = 2147483647
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l4
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s2i32 = 24
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		s3i64 = l12
		s4i64 = 2147483647
		if s3i64 < s4i64 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i64 = l12
		s3i64 = -2147483647
		if s2i64 < s3i64 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		goto lbl4
	}
	s0i32 = l8
	s1i32 = 16
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l1
	s0i32 = i32DivS(s0i32, s1i32)
lbl4:
	l1 = s0i32
	s0i32 = l0
	s1i32 = l7
	s2i32 = l11
	s3i32 = l3
	s4i32 = l5
	s5i32 = l1
	s0i32 = f640(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	l1 = s0i32
	s0i32 = l0
	s1i32 = l5
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+124)])) = uint32(s1i32)
	s0i32 = l2
	s1i32 = -1
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l3
	l7 = s0i32
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
lbl1:
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	ctx.Mem[int(s0i32+52)] = uint8(s1i32)
	s0i32 = l4
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l1
	return s0i32
}
