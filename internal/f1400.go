package internal

import (
	"unsafe"
)

func f1400(ctx *Context, l0 int32) {
	var l1 int32
	_ = l1
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
	var l8 float32
	_ = l8
	var l9 float32
	_ = l9
	var l10 float32
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
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
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	var s6f32 float32
	_ = s6f32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+172)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)]))
	l1 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l9 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l8 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l11 = s0f32
	s0i32 = l1
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l10 = s0f32
	s0i32 = 1
	l2 = s0i32
	s0i32 = l0
	s1i32 = 88
	s0i32 = s0i32 + s1i32
	l5 = s0i32
	s0i32 = f45(ctx, s0i32)
	l1 = s0i32
	s1f32 = l9
	s2f32 = l10
	s1f32 = s1f32 - s2f32
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = s1f32
	s0i32 = l1
	s1f32 = l8
	s2f32 = l11
	s1f32 = s1f32 - s2f32
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = s1f32
	s0i32 = 0
	l1 = s0i32
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	l3 = s1i32
	s2i32 = 1
	if s1i32 > s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l8
		s2i32 = l0
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+120)]))
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+76)]))
		l1 = s3i32
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+4)]))
		s2f32 = s2f32 - s3f32
		s1f32 = s1f32 * s2f32
		s2f32 = l9
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+116)]))
		s4i32 = l1
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s1f32 = s1f32 - s2f32
		l9 = s1f32
		s1i32 = 1
		l4 = s1i32
	lbl1:
		s1i32 = l1
		s2i32 = l2
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		l6 = s2i32
		s1i32 = s1i32 + s2i32
		l7 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l8 = s1f32
		s1i32 = l1
		s2i32 = l2
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		l2 = s2i32
		s3i32 = l3
		s2i32 = i32RemS(s2i32, s3i32)
		s3i32 = 3
		s2i32 = s2i32 << (uint32(s3i32) & 31)
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		l11 = s1f32
		s1i32 = l7
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l10 = s1f32
		s1i32 = l1
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l12 = s1f32
		s1i32 = l5
		s1i32 = f45(ctx, s1i32)
		l1 = s1i32
		s2f32 = l12
		s3f32 = l10
		s2f32 = s2f32 - s3f32
		l10 = s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = s2f32
		s1i32 = l1
		s2f32 = l11
		s3f32 = l8
		s2f32 = s2f32 - s3f32
		l8 = s2f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = s2f32
		s1i32 = l4
		s2f32 = l9
		s3f32 = l8
		s4i32 = l0
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+120)]))
		s5i32 = l6
		s6i32 = l0
		s6i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s6i32+76)]))
		l1 = s6i32
		s5i32 = s5i32 + s6i32
		l3 = s5i32
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+4)]))
		s4f32 = s4f32 - s5f32
		s3f32 = s3f32 * s4f32
		s4f32 = l10
		s5i32 = l0
		s5f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s5i32+116)]))
		s6i32 = l3
		s6f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s6i32+0)]))
		s5f32 = s5f32 - s6f32
		s4f32 = s4f32 * s5f32
		s3f32 = s3f32 - s4f32
		s2f32 = s2f32 * s3f32
		s3f32 = 0
		if s2f32 <= s3f32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		s3i32 = 1
		s2i32 = s2i32 ^ s3i32
		s1i32 = s1i32 & s2i32
		l4 = s1i32
		s1i32 = l2
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+84)]))
		l3 = s2i32
		if s1i32 < s2i32 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl1
		}
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 ^ s2i32
	} else {
		s1i32 = 0
	}
	s2i32 = l0
	s2i32 = int32(ctx.Mem[int(s2i32+157)])
	s3i32 = 0
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 | s2i32
	ctx.Mem[int(s0i32+157)] = uint8(s1i32)
}
