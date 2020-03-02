package internal

import (
	"unsafe"
)

func f121(ctx *Context, l0 int32, l1 int32) int32 {
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	s0i32 = int32(int8(ctx.Mem[int(s0i32+52)]))
	l3 = s0i32
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l1 = s0i32
		s1i32 = l1
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l1 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s0i32 = s0i32 ^ s1i32
		s1i32 = 1
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+100)]))
		l1 = s1i32
		s2i32 = l1
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l1 = s2i32
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s1i32 = s1i32 ^ s2i32
		s2i32 = l0
		s2i32 = int32(ctx.Mem[int(s2i32+81)])
		l1 = s2i32
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		l2 = s0i32
		s1i32 = l2
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l4 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s0i32 = s0i32 ^ s1i32
		s1i32 = 1
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		l4 = s1i32
		s2i32 = l4
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l5 = s2i32
		s1i32 = s1i32 + s2i32
		s2i32 = l5
		s1i32 = s1i32 ^ s2i32
		s2i32 = l1
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l2
		s1i32 = l4
		s2i32 = l1
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s0i32 = s0i32 - s1i32
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+82)])
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = 65535
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		return s0i32
	}
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = 0
		l3 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l1 = s0i32
		s1i32 = l1
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l1 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s0i32 = s0i32 ^ s1i32
		s1i32 = 1
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+100)]))
		l1 = s1i32
		s2i32 = l1
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l1 = s2i32
		s1i32 = s1i32 + s2i32
		s2i32 = l1
		s1i32 = s1i32 ^ s2i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		l1 = s0i32
		s1i32 = l1
		s2i32 = 31
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l2 = s1i32
		s0i32 = s0i32 + s1i32
		s1i32 = l2
		s0i32 = s0i32 ^ s1i32
		s1i32 = 1
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		l2 = s1i32
		s2i32 = l2
		s3i32 = 31
		s2i32 = s2i32 >> (uint32(s3i32) & 31)
		l4 = s2i32
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s1i32 = s1i32 ^ s2i32
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 - s1i32
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+81)])
		s0i32 = s0i32 >> (uint32(s1i32) & 31)
		s1i32 = 65535
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		return s0i32
	}
	s0i32 = 0
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	s1i32 = l0
	s2i32 = 31
	s1i32 = s1i32 >> (uint32(s2i32) & 31)
	l0 = s1i32
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s0i32 = s0i32 ^ s1i32
	s1i32 = 65536
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	s0i32 = s0i32 - s1i32
	s1i32 = 65535
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	l3 = s0i32
lbl0:
	s0i32 = l3
	return s0i32
}
