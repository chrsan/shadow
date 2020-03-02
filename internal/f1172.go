package internal

import (
	"unsafe"
)

func f1172(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)]))
	s1i32 = l2
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l5 = s0i32
	s1i32 = -2147483646
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l2
	s1i32 = l5
	s0i32 = s0i32 - s1i32
	s1i32 = 2
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l5 = s1i32
	s2i32 = l2
	s3i32 = -1
	s2i32 = s2i32 + s3i32
	s3i32 = 0
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+24)]))
	s5i32 = l5
	s4i32 = s4i32 - s5i32
	f100(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
lbl1:
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l4
	s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
	l5 = s0i32
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
	lbl3:
		s0i32 = l5
		s1i32 = 65535
		s0i32 = s0i32 & s1i32
		l6 = s0i32
		s1i32 = l1
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		l8 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 0
		s3i32 = l1
		s4i32 = l8
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		l9 = s3i32
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		l7 = s1i32
		s2i32 = l1
		s3i32 = l6
		s2i32 = s2i32 + s3i32
		l5 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 0
		s3i32 = l7
		s4i32 = l5
		if s3i32 < s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s0i32 = s0i32 + s1i32
		l7 = s0i32
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)]))
			s1i32 = l8
			s2i32 = l1
			s3i32 = l9
			if s3i32 != 0 {
				// s1i32 = s1i32
			} else {
				s1i32 = s2i32
			}
			s2i32 = l2
			s3i32 = l3
			s3i32 = int32(ctx.Mem[int(s3i32+0)])
			s4i32 = l7
			f100(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		}
		s0i32 = l3
		s1i32 = l6
		s0i32 = s0i32 + s1i32
		l3 = s0i32
		s0i32 = l5
		l1 = s0i32
		s0i32 = l4
		s1i32 = l6
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = int32(*(*int16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])))
		l5 = s0i32
		s1i32 = 0
		if s0i32 > s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl3
		}
	}
}
