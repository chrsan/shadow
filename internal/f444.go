package internal

import (
	"unsafe"
)

func f444(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 float32
	_ = l4
	var l5 float32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
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
	var s7f32 float32
	_ = s7f32
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l2
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = l0
		s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
		s0i32 = l1
		s1i32 = 16
		s0i32 = s0i32 + s1i32
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l3
	f688(ctx, s0i32, s1i32)
	s0i32 = l0
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l5 = s0f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l4 = s1f32
	s0f32 = s0f32 - s1f32
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l6 = s1f32
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l5
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	l4 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = l6
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1f32 = l5
		s2f32 = l6
		s3f32 = l4
		s4f32 = l5
		s3f32 = s3f32 - s4f32
		l7 = s3f32
		s3f32 = -s3f32
		s4f32 = l7
		s5f32 = l7
		s6f32 = 0
		if s5f32 < s6f32 {
			s5i32 = 1
		} else {
			s5i32 = 0
		}
		if s5i32 != 0 {
			// s3f32 = s3f32
		} else {
			s3f32 = s4f32
		}
		s4f32 = l4
		s5f32 = l6
		s4f32 = s4f32 - s5f32
		l4 = s4f32
		s4f32 = -s4f32
		s5f32 = l4
		s6f32 = l4
		s7f32 = 0
		if s6f32 < s7f32 {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4f32 = s4f32
		} else {
			s4f32 = s5f32
		}
		if s3f32 < s4f32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1f32 = s1f32
		} else {
			s1f32 = s2f32
		}
		l4 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = s1f32
		s0i32 = l3
		s1f32 = l4
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = s1f32
	}
	s0f32 = l5
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	l7 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = l4
	s2f32 = l7
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1f32 = l5
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = s1f32
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
	l4 = s1f32
	s0f32 = s0f32 - s1f32
	s1f32 = l6
	s2f32 = l4
	s1f32 = s1f32 - s2f32
	s0f32 = s0f32 * s1f32
	s1f32 = 0
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s1f32 = l6
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
lbl2:
	s0i32 = l3
	s1i32 = 28
	s0i32 = s0i32 + s1i32
	s1i32 = l3
	s2i32 = l1
	s3i32 = l2
	s4i32 = -1
	s3i32 = s3i32 + s4i32
	l0 = s3i32
	s1i32 = f444(ctx, s1i32, s2i32, s3i32)
	s2i32 = l0
	s0i32 = f444(ctx, s0i32, s1i32, s2i32)
lbl0:
	l0 = s0i32
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
