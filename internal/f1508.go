package internal

import (
	"unsafe"
)

func f1508(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
		l3 = s0i32
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = l1
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
		s0i32 = 1210
		s1i32 = 1211
		s2i32 = 1212
		s3i32 = l2
		s4i32 = 1
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l2
		s3i32 = 2
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		return s0i32
	lbl1:
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
		l1 = s0i32
		s1i32 = 128
		s0i32 = s0i32 & s1i32
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l0
			s2i32 = 32
			s1i32 = s1i32 + s2i32
			s1i32 = f24(ctx, s1i32)
			l1 = s1i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
			l2 = s0i32
		}
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i64 = 281474976776192
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint64(s1i64)
			s0i32 = 26800
			goto lbl3
		}
		s0i32 = l0
		s1i32 = 65536
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
		s1i32 = i32DivS(s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = 65536
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+28)]))
		s1i32 = i32DivS(s1i32, s2i32)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)])) = uint32(s1i32)
		s0i32 = 26816
		s1i32 = 26832
		s2i32 = l2
		s3i32 = 1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
	lbl3:
		s1i32 = l3
		s2i32 = 0
		if s1i32 > s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		l3 = s1i32
		s2i32 = 2
		s1i32 = s1i32 | s2i32
		s2i32 = l3
		s3i32 = l1
		s4i32 = 12
		s3i32 = s3i32 & s4i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = 2
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	} else {
		s0i32 = 0
	}
	return s0i32
}
