package internal

import (
	"unsafe"
)

func f1148(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var l10 int32
	_ = l10
	var l11 int64
	_ = l11
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
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = ctx.g0
	s1i32 = 16
	s0i32 = s0i32 - s1i32
	l7 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 12
	s0i32 = s0i32 + s1i32
	l10 = s0i32
	s1i32 = l1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)]))
	l6 = s0i32
	if s0i32 != 0 {
		s0i32 = l6
		s1i32 = l6
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = 1
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l5 = s0i32
	s0i32 = l0
	s1i32 = l6
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)])) = uint32(s1i32)
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s1i32 = l5
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	l6 = s1i32
	s2i32 = -1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l6
	s1i32 = 1
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	f12(ctx, s0i32)
lbl0:
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l1
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l11 = s0i64
	s0i32 = l0
	s1i32 = l3
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	s0i32 = l0
	s1i64 = l11
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint64(s1i64)
	s0i32 = l4
	s0i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l11 = s0i64
	s0i32 = l0
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+104)])) = uint64(s1i64)
	s0i32 = l0
	s1i64 = l11
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint64(s1i64)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)]))
	l1 = s0i32
	s0i32 = l7
	s1i32 = 1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
	s0i32 = l7
	s1i64 = 0
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint64(s1i64)
	s0i32 = l7
	s1i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l7
	l8 = s0i32
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+52)]))
	s2i32 = l0
	s3i32 = 88
	s2i32 = s2i32 + s3i32
	f433(ctx, s0i32, s1i32, s2i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	s1i32 = 1
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		s2i32 = l0
		s3i32 = 96
		s2i32 = s2i32 + s3i32
		f662(ctx, s0i32, s1i32, s2i32)
		s0i32 = l8
		s1i32 = 9
		s2i32 = 0
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+20)]))
	l6 = s0i32
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)]))
		l2 = s0i32
		if s0i32 != 0 {
			s0i32 = l0
			s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
			s1i32 = 1
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			if s0i32 != 0 {
				goto lbl5
			}
		}
		s0i32 = f105(ctx)
		l2 = s0i32
	lbl5:
		s0i32 = 1
		s1i32 = 2
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+56)]))
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
		l3 = s0i32
		s0i32 = 0
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
		l9 = s1i32
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l5 = s1i32
		s0i32 = s0i32 - s1i32
		s1i32 = 3
		s0i32 = s0i32 & s1i32
		l1 = s0i32
		s1i32 = 100
		s0i32 = s0i32 | s1i32
		s1i32 = l9
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		s2i32 = l5
		s1i32 = s1i32 - s2i32
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l9
			s1i32 = 100
			s2i32 = 4
			f18(ctx, s0i32, s1i32, s2i32)
			s0i32 = 0
			s1i32 = l9
			s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
			l5 = s1i32
			s0i32 = s0i32 - s1i32
			s1i32 = 3
			s0i32 = s0i32 & s1i32
			l1 = s0i32
		}
		s0i32 = l9
		s1i32 = l1
		s2i32 = l5
		s1i32 = s1i32 + s2i32
		l1 = s1i32
		s2i32 = 100
		s1i32 = s1i32 + s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = l2
		s2i32 = l3
		s3i32 = l6
		s4i32 = 2
		s0i32 = f95(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
		s1i32 = l8
		s2i32 = 1941503
		s3i32 = l0
		s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+52)]))
		s4i32 = 2097151
		s3i32 = s3i32 & s4i32
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 1
		s2i32 = s2i32 & s3i32
		f296(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s1i32 = 108
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1f32 = 1
	if s0f32 != s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l8
		s1i32 = 92
		s2i32 = l1
		f16(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l0
	s1i32 = l10
	s2i32 = l4
	s3i32 = l8
	s4i32 = l0
	s4i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s4i32+56)]))
	s5i32 = 1
	if s4i32 == s5i32 {
		s4i32 = 1
	} else {
		s4i32 = 0
	}
	if s4i32 != 0 {
		s4i32 = l1
		s4f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s4i32+0)]))
		s5f32 = 1
		if s4f32 == s5f32 {
			s4i32 = 1
		} else {
			s4i32 = 0
		}
	} else {
		s4i32 = 0
	}
	s5i32 = l0
	s5i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s5i32+80)]))
	s1i32 = f431(ctx, s1i32, s2i32, s3i32, s4i32, s5i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
	s0i32 = l7
	s1i32 = 16
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
