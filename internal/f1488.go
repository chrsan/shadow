package internal

import (
	"math"
	"unsafe"
)

func f1488(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 float32
	_ = l4
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
	var s1i64 int64
	_ = s1i64
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	s0i32 = l0
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	l3 = s0i32
	s0i32 = 1
	l1 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	l2 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s1i32 = f24(ctx, s1i32)
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 128
		s0i32 = s0i32 & s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		l1 = s0i32
	}
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(ctx.Mem[int(s1i32+75)])
	s2i32 = 1
	s1i32 = s1i32 + s2i32
	*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+288)])) = uint16(s1i32)
	s0i32 = l0
	s1i32 = l2
	s2i32 = 15
	s1i32 = s1i32 & s2i32
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s2i32 = 21840
	s1i32 = s1i32 + s2i32
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+260)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+32)]))
	s2f32 = 4.2949673e+09
	s1f32 = s1f32 * s2f32
	l4 = s1f32
	s2f32 = 9.2233715e+18
	s3f32 = l4
	s4f32 = 9.2233715e+18
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
	s2f32 = -9.2233715e+18
	s3f32 = l4
	s4f32 = -9.2233715e+18
	if s3f32 > s4f32 {
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
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 9.223372e+18
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l4
		s1i64 = int64(math.Trunc(float64(s1f32)))
		goto lbl1
	}
	s1i64 = -9223372036854775808
lbl1:
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+264)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = l0
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+44)]))
	s2f32 = 4.2949673e+09
	s1f32 = s1f32 * s2f32
	l4 = s1f32
	s2f32 = 9.2233715e+18
	s3f32 = l4
	s4f32 = 9.2233715e+18
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
	s2f32 = -9.2233715e+18
	s3f32 = l4
	s4f32 = -9.2233715e+18
	if s3f32 > s4f32 {
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
	s1f32 = float32(math.Abs(float64(s1f32)))
	s2f32 = 9.223372e+18
	if s1f32 < s2f32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f32 = l4
		s1i64 = int64(math.Trunc(float64(s1f32)))
		goto lbl3
	}
	s1i64 = -9223372036854775808
lbl3:
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+272)])) = uint64(s1i64)
	s0i32 = l1
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s1i32 = f24(ctx, s1i32)
		l2 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = l0
	s2i32 = l2
	s3i32 = 14
	s2i32 = s2i32 & s3i32
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	s1i32 = f1508(ctx, s1i32, s2i32)
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+296)])) = uint32(s1i32)
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+84)]))
	l2 = s0i32
	s0i32 = l0
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)]))
	l1 = s0i32
	s1i32 = 128
	s0i32 = s0i32 & s1i32
	if s0i32 != 0 {
		s0i32 = l0
		s1i32 = l3
		s1i32 = f24(ctx, s1i32)
		l1 = s1i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+68)])) = uint32(s1i32)
	}
	s0i32 = l0
	s1i32 = 28028
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	s2i32 = 1225
	s3i32 = l2
	s4i32 = 0
	if s3i32 > s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	l2 = s3i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	s2i32 = 22144
	s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)]))
	s3i32 = 1226
	s4i32 = l2
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s3i32 = l1
	s4i32 = 12
	s3i32 = s3i32 & s4i32
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+300)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l0
	s1i32 = int32(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s1i32+288)])))
	s2i32 = 256
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		goto lbl8
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+84)]))
	if s1i32 != 0 {
		goto lbl8
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+76)]))
	if s1i32 != 0 {
		goto lbl8
	}
	s1i32 = l0
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+80)]))
	if s1i32 != 0 {
		goto lbl8
	}
	s1i32 = l1
	s2i32 = 128
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		s1i32 = l0
		s2i32 = l3
		s2i32 = f24(ctx, s2i32)
		l1 = s2i32
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s1i32+68)])) = uint32(s2i32)
	}
	s1i32 = l1
	s2i32 = 12
	s1i32 = s1i32 & s2i32
	if s1i32 != 0 {
		goto lbl8
	}
	s1i32 = 1227
	goto lbl7
lbl8:
	s1i32 = l0
	s1i32 = f1484(ctx, s1i32)
lbl7:
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+292)])) = uint32(s1i32)
	s0i32 = 1
	return s0i32
}
