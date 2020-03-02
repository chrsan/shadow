package internal

import (
	"unsafe"
)

func f1688(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 float32
	_ = l6
	var l7 float32
	_ = l7
	var l8 float32
	_ = l8
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
	s0i32 = ctx.g0
	s1i32 = 160
	s0i32 = s0i32 - s1i32
	l4 = s0i32
	ctx.g0 = s0i32
	s0i32 = l4
	s1i32 = l0
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+60)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+128)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+136)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l2
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+144)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+152)])) = uint64(s1i64)
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 96
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 92
	s2i32 = s2i32 + s3i32
	s0i32 = f1699(ctx, s0i32, s1i32, s2i32)
	l1 = s0i32
	s1i32 = 1
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = 1
		s0i32 = s0i32 - s1i32
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l3
			s2i32 = 0
			f78(ctx, s0i32, s1i32, s2i32)
			goto lbl0
		}
		s0i32 = l0
		s1i32 = l3
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l1
	s1i32 = -3
	s0i32 = s0i32 + s1i32
	s1i32 = 2
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l4
		s2i32 = 96
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)]))
		l2 = s0i32
		s0i32 = l0
		s1i32 = 25480
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 4
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l0
		s1i32 = l4
		s2i32 = 96
		s1i32 = s1i32 + s2i32
		s2i32 = 8
		s1i32 = s1i32 | s2i32
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		s1i32 = 5
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl4
		}
		s0i32 = l0
		s1i32 = l4
		s2i32 = 112
		s1i32 = s1i32 + s2i32
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
	lbl4:
		s0i32 = l0
		s1i32 = l3
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		goto lbl0
	}
	s0i32 = l0
	s1i32 = l4
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+92)]))
	s2i32 = l4
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l4
	s4i32 = 72
	s3i32 = s3i32 + s4i32
	s4i32 = 0
	s0i32 = f313(ctx, s0i32, s1i32, s2i32, s3i32, s4i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s2i32 = 0
		f78(ctx, s0i32, s1i32, s2i32)
		goto lbl0
	}
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s1i32 = l4
	s2i32 = 56
	s1i32 = s1i32 + s2i32
	s0i32 = f2112(ctx, s0i32, s1i32)
	l1 = s0i32
	s1i32 = 0
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 0
		l2 = s0i32
	lbl7:
		s0f32 = 1
		l6 = s0f32
		s0i32 = l2
		s1i32 = l1
		if s0i32 < s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l4
			s1i32 = 56
			s0i32 = s0i32 + s1i32
			s1i32 = l2
			s2i32 = 2
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			s0i32 = s0i32 + s1i32
			s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
			l6 = s0f32
		}
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+140)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = 1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint32(s1i32)
		s0i32 = l4
		s1f32 = l6
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
		s0i32 = l4
		s1f32 = l7
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l4
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint16(s1i32)
		s0i32 = l4
		s1f32 = l7
		s2f32 = l6
		s1f32 = s1f32 + s2f32
		s2f32 = 0.5
		s1f32 = s1f32 * s2f32
		l8 = s1f32
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
		s0i32 = l0
		s1i32 = l4
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s0i32 = f309(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = 0
		ctx.Mem[int(s0i32+140)] = uint8(s1i32)
		s0i32 = l0
		s1i32 = -1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+132)])) = uint32(s1i32)
		s0i32 = l4
		s1f32 = l6
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+48)])) = s1f32
		s0i32 = l4
		s1f32 = l8
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = s1f32
		s0i32 = l4
		s1f32 = l7
		*(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = s1f32
		s0i32 = l4
		s1i32 = 0
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)])) = uint16(s1i32)
		s0i32 = l0
		s1i32 = l4
		s2i32 = 128
		s1i32 = s1i32 + s2i32
		s2i32 = l4
		s0i32 = f309(ctx, s0i32, s1i32, s2i32)
		s0i32 = l1
		s1i32 = l2
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l5 = s0i32
		s0f32 = l6
		l7 = s0f32
		s0i32 = l2
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l5
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl7
		}
	}
	s0i32 = l4
	s1i32 = 128
	s0i32 = s0i32 + s1i32
	s0f32 = f2110(ctx, s0i32)
	l6 = s0f32
	s1f32 = 0
	if s0f32 > s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = 128
		s0i32 = s0i32 + s1i32
		s1f32 = l6
		s2i32 = l4
		s3i32 = 0
		f337(ctx, s0i32, s1f32, s2i32, s3i32)
		s0i32 = l0
		s1i32 = 120
		s0i32 = s0i32 + s1i32
		s1i32 = l4
		s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
		s2i32 = l4
		s2f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		s3i32 = l0
		s3f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)]))
		f677(ctx, s0i32, s1f32, s2f32, s3f32)
	}
	s0i32 = l0
	s1i32 = l4
	s2i32 = 128
	s1i32 = s1i32 + s2i32
	s2i32 = l4
	s3i32 = 80
	s2i32 = s2i32 + s3i32
	s3i32 = l4
	s4i32 = 72
	s3i32 = s3i32 + s4i32
	s4i32 = l4
	s5i32 = l4
	s6i32 = -64
	s5i32 = s5i32 - s6i32
	f1700(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	s0i32 = l0
	s1i32 = 1
	ctx.Mem[int(s0i32+141)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l3
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+60)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+64)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+44)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l4
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+28)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
lbl0:
	s0i32 = l4
	s1i32 = 160
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
