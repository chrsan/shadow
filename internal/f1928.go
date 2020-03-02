package internal

import (
	"unsafe"
)

func f1928(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var s2i64 int64
	_ = s2i64
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)]))
	l4 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l2
		s1i32 = l2
		s1i32 = f24(ctx, s1i32)
		l4 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+36)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l4
	s2i32 = 15
	s1i32 = s1i32 & s2i32
	l4 = s1i32
	s2i32 = 1
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l4
		s2i32 = 1
		s1i32 = s1i32 - s2i32
		if s1i32 != 0 {
			goto lbl1
		}
		s1i32 = 0
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l4 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l5 = s1i32
		s2i32 = 8
		s1i32 = s1i32 | s2i32
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l4
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l1
			s2i32 = 8
			s3i32 = 4
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = 0
			s2i32 = l1
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l4 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			l5 = s1i32
		}
		s1i32 = l1
		s2i32 = l4
		s3i32 = l5
		s2i32 = s2i32 + s3i32
		l5 = s2i32
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l5
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l5
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l6 = s1i32
		s1i32 = 0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l1 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l2 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l3 = s1i32
		s2i32 = 12
		s1i32 = s1i32 | s2i32
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l1
			s2i32 = 12
			s3i32 = 4
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = 0
			s2i32 = l1
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l2 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			l3 = s1i32
		}
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		s2i32 = s2i32 + s3i32
		l1 = s2i32
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l5
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = 125
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = 2
		s2i32 = 1
		s3i32 = l4
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		goto lbl2
	}
	s1i32 = 0
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l5 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l3 = s1i32
	s1i32 = l4
	s2i32 = 3
	if uint32(s1i32) <= uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 16
		s1i32 = s1i32 | s2i32
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l5
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l1
			s2i32 = 16
			s3i32 = 4
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = 0
			s2i32 = l1
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l5 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			l3 = s1i32
		}
		s1i32 = l1
		s2i32 = l3
		s3i32 = l5
		s2i32 = s2i32 + s3i32
		l1 = s2i32
		s3i32 = 16
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l1
		s2i32 = l2
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)])) = uint32(s2i32)
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l6 = s1i32
		s1i32 = 0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l4 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l2 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l3 = s1i32
		s2i32 = 12
		s1i32 = s1i32 | s2i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l4
			s2i32 = 12
			s3i32 = 4
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = 0
			s2i32 = l4
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l2 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			l3 = s1i32
		}
		s1i32 = l4
		s2i32 = l2
		s3i32 = l3
		s2i32 = s2i32 + s3i32
		l2 = s2i32
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = 126
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = 2
		s2i32 = 1
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		goto lbl2
	}
	s1i32 = l3
	s2i32 = 36
	s1i32 = s1i32 | s2i32
	s2i32 = l1
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l5
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l1
		s2i32 = 36
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l1
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l5 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l3 = s1i32
	}
	s1i32 = l1
	s2i32 = l3
	s3i32 = l5
	s2i32 = s2i32 + s3i32
	l1 = s2i32
	s3i32 = 36
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l1
	s1i32 = f474(ctx, s1i32, s2i32)
	if s1i32 != 0 {
		s1i32 = l0
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
		l6 = s1i32
		s1i32 = 0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
		l4 = s2i32
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l2 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l3 = s1i32
		s2i32 = 12
		s1i32 = s1i32 | s2i32
		s2i32 = l4
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = l2
		s2i32 = s2i32 - s3i32
		if uint32(s1i32) > uint32(s2i32) {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			s1i32 = l4
			s2i32 = 12
			s3i32 = 4
			f18(ctx, s1i32, s2i32, s3i32)
			s1i32 = 0
			s2i32 = l4
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
			l2 = s2i32
			s1i32 = s1i32 - s2i32
			s2i32 = 3
			s1i32 = s1i32 & s2i32
			l3 = s1i32
		}
		s1i32 = l4
		s2i32 = l2
		s3i32 = l3
		s2i32 = s2i32 + s3i32
		l2 = s2i32
		s3i32 = 12
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = l1
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = 127
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l2
		s2i32 = l6
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
		s1i32 = l0
		s2i32 = l0
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
		s3i32 = 1
		s2i32 = s2i32 + s3i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
		s1i32 = 2
		s2i32 = 1
		s3i32 = l5
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		goto lbl2
	}
	s1i32 = l1
	s2i32 = l2
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint64(s2i64)
	s1i32 = l1
	s2i32 = l2
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+32)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)])) = uint32(s2i32)
	s1i32 = l1
	s2i32 = l2
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+24)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)])) = uint64(s2i64)
	s1i32 = l1
	s2i32 = l2
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)])) = uint64(s2i64)
	s1i32 = l1
	s2i32 = l2
	s2i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint64(s2i64)
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	l4 = s1i32
	s1i32 = 0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	l3 = s2i32
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
	l2 = s2i32
	s1i32 = s1i32 - s2i32
	s2i32 = 3
	s1i32 = s1i32 & s2i32
	l5 = s1i32
	s2i32 = 12
	s1i32 = s1i32 | s2i32
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = l2
	s2i32 = s2i32 - s3i32
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = l3
		s2i32 = 12
		s3i32 = 4
		f18(ctx, s1i32, s2i32, s3i32)
		s1i32 = 0
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)]))
		l2 = s2i32
		s1i32 = s1i32 - s2i32
		s2i32 = 3
		s1i32 = s1i32 & s2i32
		l5 = s1i32
	}
	s1i32 = l3
	s2i32 = l2
	s3i32 = l5
	s2i32 = s2i32 + s3i32
	l2 = s2i32
	s3i32 = 12
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l1
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = 132
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l2
	s2i32 = l4
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l2
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])) = uint32(s2i32)
	s1i32 = l0
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	s3i32 = 1
	s2i32 = s2i32 + s3i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])) = uint32(s2i32)
	s1i32 = 2
lbl2:
	s2i32 = l0
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+12)]))
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)])) = uint32(s1i32)
lbl1:
}
