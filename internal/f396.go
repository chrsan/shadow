package internal

import (
	"unsafe"
)

func f396(ctx *Context, l0 int32, l1 int32, l2 int32) {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
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
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s9i32 int32
	_ = s9i32
	var s10i32 int32
	_ = s10i32
	var s11i32 int32
	_ = s11i32
	var s12i32 int32
	_ = s12i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = 5
	s3i32 = l1
	s4i32 = l2
	s5i32 = 11
	s6i32 = l1
	s7i32 = 14
	s8i32 = -1
	s9i32 = -1
	s10i32 = -1
	s11i32 = 31
	s12i32 = 0
	s6i32 = f21(ctx, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32)
	s3i32 = f168(ctx, s3i32, s4i32, s5i32, s6i32)
	s1i32 = f96(ctx, s1i32, s2i32, s3i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 6
	s3i32 = l1
	s4i32 = l2
	s5i32 = 5
	s6i32 = l1
	s7i32 = 14
	s8i32 = -1
	s9i32 = -1
	s10i32 = -1
	s11i32 = 63
	s12i32 = 0
	s6i32 = f21(ctx, s6i32, s7i32, s8i32, s9i32, s10i32, s11i32, s12i32)
	s3i32 = f168(ctx, s3i32, s4i32, s5i32, s6i32)
	s1i32 = f96(ctx, s1i32, s2i32, s3i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 14
	s2i32 = -1
	s3i32 = -1
	s4i32 = -1
	s5i32 = 31
	s6i32 = 0
	s0i32 = f21(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32)
	l3 = s0i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	s1i32 = l3
	s2i32 = 24
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = 14
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
		s1i32 = -1
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl0
		}
	}
	s0i32 = l1
	s1i32 = l3
	s2i32 = l2
	s0i32 = f79(ctx, s0i32, s1i32, s2i32)
	l2 = s0i32
lbl0:
	s0i32 = l0
	s1i32 = l1
	s2i32 = 5
	s3i32 = l2
	s1i32 = f96(ctx, s1i32, s2i32, s3i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s2i32 = 14
	s3i32 = -1
	s4i32 = -1
	s5i32 = -1
	s6i32 = 1065353216
	s7i32 = 0
	s1i32 = f21(ctx, s1i32, s2i32, s3i32, s4i32, s5i32, s6i32, s7i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
}
