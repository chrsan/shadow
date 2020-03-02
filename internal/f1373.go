package internal

import (
	"unsafe"
)

func f1373(ctx *Context, l0 int32, l1 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l0
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 127
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
		s0i32 = 27872
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s1i32 = -128
			s0i32 = s0i32 & s1i32
			s1i32 = 57216
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl0
			}
			goto lbl2
		}
		s0i32 = l1
		s1i32 = 2047
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l1
			s2i32 = 63
			s1i32 = s1i32 & s2i32
			s2i32 = 128
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+1)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l1
			s2i32 = 6
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = 192
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = 2
			return s0i32
		}
		s0i32 = l1
		s1i32 = 55296
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		s1i32 = 0
		s2i32 = l1
		s3i32 = -8192
		s2i32 = s2i32 & s3i32
		s3i32 = 57344
		if s2i32 != s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l1
			s2i32 = 63
			s1i32 = s1i32 & s2i32
			s2i32 = 128
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+2)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l1
			s2i32 = 12
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = 224
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l1
			s2i32 = 6
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = 63
			s1i32 = s1i32 & s2i32
			s2i32 = 128
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+1)] = uint8(s1i32)
			s0i32 = 3
			return s0i32
		}
		s0i32 = l1
		s1i32 = -65536
		s0i32 = s0i32 + s1i32
		s1i32 = 1048575
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l1
			s2i32 = 63
			s1i32 = s1i32 & s2i32
			s2i32 = 128
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+3)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l1
			s2i32 = 18
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = 240
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+0)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l1
			s2i32 = 6
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = 63
			s1i32 = s1i32 & s2i32
			s2i32 = 128
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+2)] = uint8(s1i32)
			s0i32 = l0
			s1i32 = l1
			s2i32 = 12
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			s2i32 = 63
			s1i32 = s1i32 & s2i32
			s2i32 = 128
			s1i32 = s1i32 | s2i32
			ctx.Mem[int(s0i32+1)] = uint8(s1i32)
			s0i32 = 4
			return s0i32
		}
	lbl2:
		s0i32 = 29100
		s1i32 = 25
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = -1
	} else {
		s0i32 = 1
	}
	return s0i32
lbl0:
	s0i32 = l0
	s1i32 = l1
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = 1
	return s0i32
}
