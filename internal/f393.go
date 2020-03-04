package internal

import (
	"unsafe"
)

func f393(ctx *Context, l0 int32, l1 int32) {
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
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l3 = s0i32
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 32
		s0i32 = s0i32 + s1i32
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+48)]))
		s0i32 = f28(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l4 = s0i32
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)]))
		s1i32 = 2
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l1
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
			l3 = s0i32
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l4
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
				l5 = s0i32
				s1i32 = 3
				if s0i32 >= s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					s0i32 = l5
					s1i32 = -2
					s0i32 = s0i32 + s1i32
					l7 = s0i32
					s0i32 = l4
					s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
					l3 = s0i32
				lbl5:
					s0i32 = l3
					s1i32 = l2
					s2i32 = 6
					s1i32 = s1i32 * s2i32
					s0i32 = s0i32 + s1i32
					l6 = s0i32
					s1i32 = 0
					*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
					s0i32 = l6
					s1i32 = l2
					s2i32 = 2
					s1i32 = s1i32 + s2i32
					*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
					s0i32 = l6
					s1i32 = l2
					s2i32 = 1
					s1i32 = s1i32 + s2i32
					l2 = s1i32
					*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
					s0i32 = l2
					s1i32 = l7
					if s0i32 < s1i32 {
						s0i32 = 1
					} else {
						s0i32 = 0
					}
					if s0i32 != 0 {
						goto lbl5
					}
				}
				s0i32 = l4
				s1i32 = l5
				s2i32 = 3
				s1i32 = s1i32 * s2i32
				s2i32 = -6
				s1i32 = s1i32 + s2i32
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
				goto lbl2
			}
			s0i32 = l4
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
			l9 = s0i32
			s1i32 = 3
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l3
				s1i32 = 0
				s2i32 = l4
				if s2i32 != 0 {
					// s0i32 = s0i32
				} else {
					s0i32 = s1i32
				}
				l8 = s0i32
				s0i32 = l9
				s1i32 = -2
				s0i32 = s0i32 + s1i32
				l6 = s0i32
				s0i32 = l4
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
				l7 = s0i32
			lbl7:
				s0i32 = l7
				s1i32 = l2
				s2i32 = 6
				s1i32 = s1i32 * s2i32
				s0i32 = s0i32 + s1i32
				l5 = s0i32
				s1i32 = l8
				s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
				*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
				s0i32 = l5
				s1i32 = l8
				s2i32 = l2
				s3i32 = 1
				s2i32 = s2i32 + s3i32
				l3 = s2i32
				s3i32 = 1
				s2i32 = s2i32 << (uint32(s3i32) & 31)
				s1i32 = s1i32 + s2i32
				s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])))
				*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+2)])) = uint16(s1i32)
				s0i32 = l5
				s1i32 = l2
				s2i32 = 1
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				s2i32 = l8
				s1i32 = s1i32 + s2i32
				s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
				*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint16(s1i32)
				s0i32 = l3
				l2 = s0i32
				s1i32 = l6
				if s0i32 < s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl7
				}
			}
			s0i32 = l4
			s1i32 = l9
			s2i32 = 3
			s1i32 = s1i32 * s2i32
			s2i32 = -6
			s1i32 = s1i32 + s2i32
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
		lbl2:
			s0i32 = l4
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint32(s1i32)
		}
	lbl8:
		s0i32 = 25884
		s1i32 = 25884
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l2 = s1i32
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl8
		}
		s0i32 = l1
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
		l3 = s0i32
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l3
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		return
	}
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
}
