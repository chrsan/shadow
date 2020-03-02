package internal

import (
	"unsafe"
)

func f311(ctx *Context, l0 int32, l1 int32, l2 int32) int32 {
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int64
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
	var l9 float32
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
	var s0i64 int64
	_ = s0i64
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
	s0i32 = ctx.g0
	s1i32 = -64
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s2i32 = l2
	s0i32 = f1692(ctx, s0i32, s1i32, s2i32)
	s1i32 = -1
	s0i32 = s0i32 + s1i32
	l4 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = 1
			l4 = s0i32
			s0i32 = l0
			s1i32 = 108
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = 96
			s1i32 = s1i32 + s2i32
			s2i32 = l0
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
			s3i32 = 1
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			if s2i32 != 0 {
				// s0i32 = s0i32
			} else {
				s0i32 = s1i32
			}
			s1i32 = l2
			s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
			s2i32 = l2
			s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
			s3i32 = l2
			s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+16)]))
			s4i32 = l2
			s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+20)]))
			f130(ctx, s0i32, s1f32, s2f32, s3f32, s4f32)
			goto lbl0
		}
		s0i32 = 1
		l4 = s0i32
		s0i32 = l0
		s1i32 = 108
		s0i32 = s0i32 + s1i32
		s1i32 = l0
		s2i32 = 96
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+132)]))
		s3i32 = 1
		if s2i32 == s3i32 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		s1i32 = l2
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		s2i32 = l2
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		s0i32 = f30(ctx, s0i32, s1f32, s2f32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+136)]))
	l5 = s1i32
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
	s0i32 = 0
	l4 = s0i32
	s0i32 = l5
	s1i32 = 32
	if s0i32 > s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)]))
	l7 = s0f32
	s0i32 = l3
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = s1f32
	s0i32 = l3
	s1f32 = l7
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
	s0i32 = l3
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint16(s1i32)
	s0i32 = l3
	s1f32 = l7
	s2f32 = l8
	s1f32 = s1f32 + s2f32
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0f32 = l9
	s1f32 = l7
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0f32 = l9
	s1f32 = l8
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl3
	}
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)]))
	l6 = s0i64
	s0i32 = l3
	s1i32 = 1
	ctx.Mem[int(s0i32+60)] = uint8(s1i32)
	s0i32 = l3
	s1i64 = l6
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
lbl3:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l3
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0i32 = f311(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l2
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)]))
	l7 = s0f32
	s0i32 = l3
	s1i32 = l2
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+48)]))
	l8 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+56)])) = s1f32
	s0i32 = l3
	s1f32 = l7
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
	s0i32 = l3
	s1i32 = 0
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint16(s1i32)
	s0i32 = l3
	s1f32 = l7
	s2f32 = l8
	s1f32 = s1f32 + s2f32
	s2f32 = 0.5
	s1f32 = s1f32 * s2f32
	l9 = s1f32
	*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = s1f32
	s0f32 = l9
	s1f32 = l7
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl4
	}
	s0f32 = l9
	s1f32 = l8
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl4
	}
	s0i32 = l3
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l2
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)]))
	l6 = s0i64
	s0i32 = l3
	s1i32 = 1
	ctx.Mem[int(s0i32+61)] = uint8(s1i32)
	s0i32 = l3
	s1i64 = l6
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
lbl4:
	s0i32 = l0
	s1i32 = l1
	s2i32 = l3
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	s0i32 = f311(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+136)]))
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint32(s1i32)
	s0i32 = 1
	l4 = s0i32
lbl0:
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	ctx.g0 = s0i32
	s0i32 = l4
	return s0i32
}
