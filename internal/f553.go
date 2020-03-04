package internal

import (
	"unsafe"
)

func f553(ctx *Context, l0 int32, l1 float32, l2 float32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
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
	var s3f32 float32
	_ = s3f32
	s0i32 = ctx.g0
	s1i32 = 112
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
	s0i32 = f564(ctx, s0i32, s1i32)
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
		s0i32 = l3
		s1i32 = 0
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l3
		s1i64 = 0
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint64(s1i64)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
		l5 = s0i32
		s1i32 = 1
		if s0i32 >= s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l5
			s1i32 = 4
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = 2
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			s1i32 = l5
			s0i32 = s0i32 + s1i32
			l5 = s0i32
			s1i32 = -1
			if s0i32 <= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl2
			}
			s0i32 = l3
			s1i32 = l5
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
			s0i32 = l3
			s1i32 = 0
			s2i32 = l5
			s3i32 = 2
			s2i32 = s2i32 << (uint32(s3i32) & 31)
			s1i32 = f29(ctx, s1i32, s2i32)
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
			l5 = s0i32
		}
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
		s1i32 = l5
		s2i32 = l0
		s3i32 = 100
		s2i32 = s2i32 + s3i32
		l6 = s2i32
		s3f32 = l1
		s4i32 = l3
		s5i32 = 96
		s4i32 = s4i32 + s5i32
		s5i32 = l3
		s6i32 = 80
		s5i32 = s5i32 + s6i32
		s0i32 = f560(ctx, s0i32, s1i32, s2i32, s3f32, s4i32, s5i32)
		if s0i32 != 0 {
			s0i32 = 0
			l5 = s0i32
			s0i32 = l3
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
			s0i32 = l3
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint64(s1i64)
			s0i32 = l3
			s1i32 = 0
			*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = uint32(s1i32)
			s0i32 = l3
			s1i64 = 0
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint64(s1i64)
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
			l4 = s0i32
			s1i32 = 1
			if s0i32 >= s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l4
				s1i32 = 4
				s0i32 = s0i32 + s1i32
				l4 = s0i32
				s1i32 = 2
				s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
				s1i32 = l4
				s0i32 = s0i32 + s1i32
				l4 = s0i32
				s1i32 = -1
				if s0i32 <= s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl1
				}
				s0i32 = l3
				s1i32 = l4
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
				s0i32 = l3
				s1i32 = 0
				s2i32 = l4
				s3i32 = 3
				s2i32 = s2i32 << (uint32(s3i32) & 31)
				s1i32 = f29(ctx, s1i32, s2i32)
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
				s0i32 = l3
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
				l5 = s0i32
				s0i32 = l3
				s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
				l4 = s0i32
			}
			s0i32 = l5
			s1i32 = l4
			if s0i32 < s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				s0i32 = l4
				s1i32 = 4
				s0i32 = s0i32 + s1i32
				l4 = s0i32
				s1i32 = 2
				s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
				s1i32 = l4
				s0i32 = s0i32 + s1i32
				l4 = s0i32
				s1i32 = -1
				if s0i32 <= s1i32 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				if s0i32 != 0 {
					goto lbl0
				}
				s0i32 = l3
				s1i32 = l4
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint32(s1i32)
				s0i32 = l3
				s1i32 = l3
				s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
				s2i32 = l4
				s3i32 = 2
				s2i32 = s2i32 << (uint32(s3i32) & 31)
				s1i32 = f29(ctx, s1i32, s2i32)
				*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = uint32(s1i32)
			}
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
			s1i32 = l0
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+72)]))
			s2i32 = l6
			s3f32 = l2
			s3f32 = -s3f32
			s4i32 = l3
			s5i32 = -64
			s4i32 = s4i32 - s5i32
			s5i32 = l3
			s6i32 = 48
			s5i32 = s5i32 + s6i32
			s0i32 = f560(ctx, s0i32, s1i32, s2i32, s3f32, s4i32, s5i32)
			l5 = s0i32
			s0i32 = 0
			l4 = s0i32
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)]))
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l5
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)]))
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl8
			}
			s0i32 = l0
			s1i32 = l3
			s2i32 = 96
			s1i32 = s1i32 + s2i32
			s2i32 = l3
			s3i32 = 80
			s2i32 = s2i32 + s3i32
			s3i32 = l3
			s4i32 = -64
			s3i32 = s3i32 - s4i32
			s4i32 = l3
			s5i32 = 48
			s4i32 = s4i32 + s5i32
			f1398(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
			s0i32 = 1
			l4 = s0i32
		lbl8:
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
			f13(ctx, s0i32)
			s0i32 = l3
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)]))
			f13(ctx, s0i32)
		}
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
		f13(ctx, s0i32)
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)]))
		f13(ctx, s0i32)
	}
	s0i32 = l3
	s1i32 = 112
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
lbl2:
	s0i32 = l3
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l3
	s2i32 = 32
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl1:
	s0i32 = l3
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l3
	s2i32 = 16
	s1i32 = s1i32 + s2i32
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
lbl0:
	s0i32 = l3
	s1i32 = 365
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14791
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 14743
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = 14717
	s1i32 = l3
	f19(ctx, s0i32, s1i32)
	panic("unreachable executed")
}
