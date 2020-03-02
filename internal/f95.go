package internal

import (
	"unsafe"
)

func f95(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	s0i32 = ctx.g0
	s1i32 = 48
	s0i32 = s0i32 - s1i32
	l5 = s0i32
	ctx.g0 = s0i32
	s0i32 = l0
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l3
	s1i32 = l1
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1i32 = f105(ctx)
		l1 = s1i32
	}
	s1i32 = l1
	s2i32 = l3
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l3 = s0i32
	s0i32 = l2
	s1i32 = l4
	s2i32 = l4
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
	l6 = s0i32
	s1i32 = l2
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l1
		s0i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])))
		s1i32 = l1
		s1i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)])))
		s2i64 = 32
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s0i64 = s0i64 | s1i64
		s1i32 = l3
		s1i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)])))
		s2i32 = l3
		s2i64 = int64(*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s2i32+4)])))
		s3i64 = 32
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s1i64 = s1i64 | s2i64
		if s0i64 == s1i64 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl2
		}
	}
	s0i32 = l0
	s1i32 = l2
	s2i32 = 2
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
	s0i32 = 1
	l4 = s0i32
	s0i32 = l0
	s1i32 = l1
	s1i32 = f297(ctx, s1i32)
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
	s2i32 = l3
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+8)]))
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	ctx.Mem[int(s0i32+2)] = uint8(s1i32)
	s0i32 = l3
	s0i32 = f297(ctx, s0i32)
	l7 = s0i32
	s0i32 = l0
	s1i32 = l6
	s2i32 = 2
	if s1i32 == s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	s2i32 = l2
	s3i32 = 1
	if s2i32 != s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = s1i32 & s2i32
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l7
	s2i32 = 1
	s1i32 = s1i32 ^ s2i32
	ctx.Mem[int(s0i32+3)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l3
		s2i32 = l5
		s3i32 = 8
		s2i32 = s2i32 + s3i32
		f1300(ctx, s0i32, s1i32, s2i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+8)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+64)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+72)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+76)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+24)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+80)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+88)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+92)])) = uint32(s1i32)
		s0i32 = l0
		s1i32 = l5
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+96)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+12)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+8)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+36)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+32)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+28)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+24)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l1
	s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+20)]))
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+16)])) = uint64(s1i64)
	s0i32 = l3
	s1i32 = l0
	s2i32 = 36
	s1i32 = s1i32 + s2i32
	f1301(ctx, s0i32, s1i32)
	s0i32 = l0
	s1i32 = l1
	s1i32 = f525(ctx, s1i32)
	ctx.Mem[int(s0i32+5)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = l3
	s1i32 = f525(ctx, s1i32)
	ctx.Mem[int(s0i32+6)] = uint8(s1i32)
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+1)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = 0
	l4 = s0i32
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+2)])
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l1
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+4)]))
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl5
	}
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+3)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+1)] = uint8(s1i32)
	s0i32 = 1
	l4 = s0i32
lbl5:
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l4
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+3)])
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s0i32 = int32(ctx.Mem[int(s0i32+4)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+4)] = uint8(s1i32)
	s0i32 = l0
	s1i32 = 0
	ctx.Mem[int(s0i32+0)] = uint8(s1i32)
lbl2:
	s0i32 = l5
	s1i32 = 48
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
	s0i32 = l0
	return s0i32
}
