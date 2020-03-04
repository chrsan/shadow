package internal

import (
	"unsafe"
)

func f1479(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+280)]))
	s1i32 = l1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l5 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+284)]))
	s1i32 = l2
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	l6 = s1i32
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l2
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s2i32 = l6
		s1i32 = i32RemS(s1i32, s2i32)
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		goto lbl0
	}
	s0i32 = l2
	s1i32 = l6
	s0i32 = i32RemS(s0i32, s1i32)
	l2 = s0i32
lbl0:
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l2
	s1i32 = s1i32 * s2i32
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l1
	s1i32 = l5
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l1
	s1i32 = -1
	if s0i32 <= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l5
		s1i32 = l1
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s2i32 = l5
		s1i32 = i32RemS(s1i32, s2i32)
		s2i32 = -1
		s1i32 = s1i32 ^ s2i32
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		goto lbl2
	}
	s0i32 = l1
	s1i32 = l5
	s0i32 = i32RemS(s0i32, s1i32)
	l1 = s0i32
lbl2:
	s0i32 = l3
	s1i32 = l2
	s2i32 = l1
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = l5
	s4i32 = l1
	s3i32 = s3i32 - s4i32
	l0 = s3i32
	s4i32 = l0
	s5i32 = l4
	if s4i32 > s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	l0 = s2i32
	s3i32 = 2
	s2i32 = s2i32 << (uint32(s3i32) & 31)
	l1 = s2i32
	s0i32 = f22(ctx, s0i32, s1i32, s2i32)
	l3 = s0i32
	s0i32 = l4
	s1i32 = l0
	s0i32 = s0i32 - s1i32
	l0 = s0i32
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	lbl5:
		s0i32 = l1
		s1i32 = l2
		s2i32 = l0
		s3i32 = l5
		s4i32 = l0
		s5i32 = l5
		if s4i32 < s5i32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
		if s4i32 != 0 {
			// s2i32 = s2i32
		} else {
			s2i32 = s3i32
		}
		l3 = s2i32
		s3i32 = 2
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l1 = s2i32
		s0i32 = f22(ctx, s0i32, s1i32, s2i32)
		s1i32 = l1
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l0
		s1i32 = l3
		s0i32 = s0i32 - s1i32
		l0 = s0i32
		if s0i32 != 0 {
			goto lbl5
		}
	}
}
