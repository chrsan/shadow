package internal

import (
	"unsafe"
)

func f418(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
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
	s0i32 = l0
	s1i32 = 24644
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	l3 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	l1 = s1i32
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
		l2 = s0i32
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		l4 = s0i32
		s0i32 = int32(ctx.Mem[int(s0i32+0)])
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l2
			s1i32 = l2
			s1i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
			s2i32 = 1
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl1
			}
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
		l1 = s0i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		s2i32 = l3
		s3i32 = l4
		s4i32 = l2
		s5i32 = l1
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+0)]))
		s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+12)]))
		if int(s5i32) < 0 || int(s5i32) >= len(table) {
			panic("table entry out of bounds")
		}
		if table[s5i32].numParams == -1 {
			panic("table entry is nil")
		}
		if table[s5i32].numParams != 5 {
			panic("argument count mismatch")
		}
		(*(*func(*Context, int32, int32, int32, int32, int32))(table[s5i32].f()))(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
		s1i32 = i32RemS(s1i32, s2i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
		s2i32 = l1
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+20)]))
		l2 = s3i32
		s4i32 = l2
		s5i32 = 2
		s4i32 = s4i32 + s5i32
		s5i32 = 2
		s4i32 = i32DivS(s4i32, s5i32)
		s3i32 = s3i32 + s4i32
		s4i32 = 1
		s3i32 = s3i32 + s4i32
		s2i32 = s2i32 * s3i32
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l1
		s2i32 = l2
		s3i32 = 1
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		s2i32 = 2
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
		l1 = s0i32
		s1i32 = l2
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		s1i32 = 0
		ctx.Mem[int(s0i32+0)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
		l1 = s0i32
	lbl1:
		s0i32 = l0
		s1i32 = l1
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = f35(ctx, s0i32)
	s0i32 = l0
	return s0i32
}
