package internal

import (
	"unsafe"
)

func f1846(ctx *Context, l0 int32, l1 int32) int32 {
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
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 56
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s0i32 = f641(ctx, s0i32, s1i32, s2i32)
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+83)])
		ctx.Mem[int(s0i32+55)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+96)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+100)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+100)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+104)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+108)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+108)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+88)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = -8192
		s1i32 = s1i32 - s2i32
		s2i32 = -16384
		s1i32 = s1i32 & s2i32
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+112)]))
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = -8192
		s1i32 = s1i32 - s2i32
		s2i32 = -16384
		s1i32 = s1i32 & s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+112)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+80)])
		ctx.Mem[int(s0i32+52)] = uint8(s1i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+81)])
		l3 = s0i32
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+120)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+116)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l3
		ctx.Mem[int(s0i32+53)] = uint8(s1i32)
		s0i32 = l0
		s0i32 = f265(ctx, s0i32)
	} else {
		s0i32 = 0
	}
	return s0i32
}
