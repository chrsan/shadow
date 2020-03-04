package internal

import (
	"math"
	"math/bits"
	"unsafe"
)

func f147(ctx *Context, l0 int32, l1 int32, l2 int32) {
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
	var l8 int32
	_ = l8
	var l9 int64
	_ = l9
	var l10 int64
	_ = l10
	var l11 float32
	_ = l11
	var l12 float32
	_ = l12
	var l13 float32
	_ = l13
	var l14 float32
	_ = l14
	var l15 float64
	_ = l15
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
	var s1f64 float64
	_ = s1f64
	var s2f64 float64
	_ = s2f64
	var s3f64 float64
	_ = s3f64
	var s4f64 float64
	_ = s4f64
	s0i32 = ctx.g0
	s1i32 = 1376
	s0i32 = s0i32 - s1i32
	l3 = s0i32
	ctx.g0 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+41)])
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = 1192
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = 3
	s0i32 = f28(ctx, s0i32, s1i32, s2i32)
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1192)]))
	l11 = s0f32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1200)]))
	l12 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l11
	s1f32 = -16383
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0i32 = l3
	s0f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1196)]))
	l13 = s0f32
	s1i32 = l3
	s1f32 = *(*float32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1204)]))
	l14 = s1f32
	if s0f32 < s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 != 0 {
		goto lbl2
	}
	s0f32 = l14
	s1f32 = 16383
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
	s0f32 = l12
	s1f32 = 16383
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
	s0f32 = l13
	s1f32 = -16383
	if s0f32 >= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl1
	}
lbl2:
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s0i32 = f37(ctx, s0i32)
	l4 = s0i32
	s1i32 = l0
	s2i32 = 3
	s3i32 = 0
	f277(ctx, s0i32, s1i32, s2i32, s3i32)
	s0i32 = l4
	s1i32 = l1
	s2i32 = l2
	f416(ctx, s0i32, s1i32, s2i32)
	s0i32 = l4
	f34(ctx, s0i32)
	goto lbl0
lbl1:
	s0i32 = l3
	s1f32 = l11
	s1f64 = float64(s1f32)
	s2f64 = -0.5234375
	s1f64 = s1f64 + s2f64
	s1f64 = math.Ceil(s1f64)
	l15 = s1f64
	s2f64 = 2.147483647e+09
	s3f64 = l15
	s4f64 = 2.147483647e+09
	if s3f64 < s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s2f64 = -2.147483647e+09
	s3f64 = l15
	s4f64 = -2.147483647e+09
	if s3f64 > s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s1f64 = math.Abs(s1f64)
	s2f64 = 2.147483648e+09
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f64 = l15
		s1i32 = int32(math.Trunc(s1f64))
		goto lbl3
	}
	s1i32 = -2147483648
lbl3:
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1176)])) = uint32(s1i32)
	s0i32 = l4
	s0i64 = int64(s0i32)
	l9 = s0i64
	s0i32 = l3
	s1f32 = l12
	s1f64 = float64(s1f32)
	s2f64 = 0.5234375
	s1f64 = s1f64 + s2f64
	s1f64 = math.Floor(s1f64)
	l15 = s1f64
	s2f64 = 2.147483647e+09
	s3f64 = l15
	s4f64 = 2.147483647e+09
	if s3f64 < s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s2f64 = -2.147483647e+09
	s3f64 = l15
	s4f64 = -2.147483647e+09
	if s3f64 > s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s1f64 = math.Abs(s1f64)
	s2f64 = 2.147483648e+09
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f64 = l15
		s1i32 = int32(math.Trunc(s1f64))
		goto lbl5
	}
	s1i32 = -2147483648
lbl5:
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1184)])) = uint32(s1i32)
	s0i32 = l4
	s0i64 = int64(s0i32)
	s1i64 = l9
	s0i64 = s0i64 - s1i64
	l9 = s0i64
	s0i32 = l3
	s1f32 = l13
	s1f64 = float64(s1f32)
	s2f64 = -0.5234375
	s1f64 = s1f64 + s2f64
	s1f64 = math.Ceil(s1f64)
	l15 = s1f64
	s2f64 = 2.147483647e+09
	s3f64 = l15
	s4f64 = 2.147483647e+09
	if s3f64 < s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s2f64 = -2.147483647e+09
	s3f64 = l15
	s4f64 = -2.147483647e+09
	if s3f64 > s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s1f64 = math.Abs(s1f64)
	s2f64 = 2.147483648e+09
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f64 = l15
		s1i32 = int32(math.Trunc(s1f64))
		goto lbl7
	}
	s1i32 = -2147483648
lbl7:
	l4 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1180)])) = uint32(s1i32)
	s0i32 = l3
	s1f32 = l14
	s1f64 = float64(s1f32)
	s2f64 = 0.5234375
	s1f64 = s1f64 + s2f64
	s1f64 = math.Floor(s1f64)
	l15 = s1f64
	s2f64 = 2.147483647e+09
	s3f64 = l15
	s4f64 = 2.147483647e+09
	if s3f64 < s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s2f64 = -2.147483647e+09
	s3f64 = l15
	s4f64 = -2.147483647e+09
	if s3f64 > s4f64 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	if s3i32 != 0 {
		// s1f64 = s1f64
	} else {
		s1f64 = s2f64
	}
	l15 = s1f64
	s1f64 = math.Abs(s1f64)
	s2f64 = 2.147483648e+09
	if s1f64 < s2f64 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	if s1i32 != 0 {
		s1f64 = l15
		s1i32 = int32(math.Trunc(s1f64))
		goto lbl9
	}
	s1i32 = -2147483648
lbl9:
	l5 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1188)])) = uint32(s1i32)
	s0i64 = l9
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l5
	s0i64 = int64(s0i32)
	s1i32 = l4
	s1i64 = int64(s1i32)
	s0i64 = s0i64 - s1i64
	l10 = s0i64
	s1i64 = 1
	if s0i64 < s1i64 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i64 = l9
	s1i64 = l10
	s0i64 = s0i64 | s1i64
	s1i64 = 2147483648
	s0i64 = s0i64 + s1i64
	s1i64 = 4294967295
	if uint64(s0i64) > uint64(s1i64) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s1i32 = l3
	s2i32 = 1176
	s1i32 = s1i32 + s2i32
	s2i32 = l1
	s3i32 = l1
	s4i32 = 20
	s3i32 = s3i32 + s4i32
	s4i32 = l1
	s4i32 = int32(ctx.Mem[int(s4i32+40)])
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	s0i32 = f25(ctx, s0i32, s1i32, s2i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl0
	}
	s0i32 = l3
	s1i32 = -64
	s0i32 = s0i32 - s1i32
	s0i32 = f152(ctx, s0i32)
	l4 = s0i32
	s0i32 = l1
	s0i32 = int32(ctx.Mem[int(s0i32+40)])
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l4
		s1i32 = l1
		s2i32 = l2
		f150(ctx, s0i32, s1i32, s2i32)
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1104)]))
		l2 = s0i32
		s0i32 = l4
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1100)]))
		l1 = s0i32
	}
	s0i32 = l3
	s1i32 = l2
	s2i32 = l1
	s3i32 = l3
	s4i32 = 1176
	s3i32 = s3i32 + s4i32
	s4i32 = 0
	s5i32 = 0
	s0i32 = f407(ctx, s0i32, s1i32, s2i32, s3i32, s4i32, s5i32)
	l7 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+52)]))
	l8 = s0i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l3
	s1i32 = 1280
	s0i32 = s0i32 + s1i32
	s1i32 = l0
	s2i32 = l0
	s3i32 = 8
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s3i32 = l7
	s3i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s3i32+56)]))
	l5 = s3i32
	s0i32 = f419(ctx, s0i32, s1i32, s2i32, s3i32)
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s1i32 = 1280
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l3
		s1i32 = 1268
		s0i32 = s0i32 + s1i32
		goto lbl13
	}
	s0i32 = l3
	s1i32 = l3
	s2i32 = 1280
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1268)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 1308
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s0i32 = l3
	s1i32 = 1272
	s0i32 = s0i32 + s1i32
lbl13:
	l1 = s0i32
	s0i32 = l2
	s1i32 = l6
	s2i32 = l0
	s3i32 = 16
	s2i32 = s2i32 + s3i32
	l6 = s2i32
	s3i32 = l5
	s0i32 = f419(ctx, s0i32, s1i32, s2i32, s3i32)
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l2
		s1i32 = 28
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	}
	s0i32 = l2
	s1i32 = l6
	s2i32 = l0
	s3i32 = l5
	s0i32 = f419(ctx, s0i32, s1i32, s2i32, s3i32)
	if s0i32 != 0 {
		s0i32 = l1
		s1i32 = l2
		*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
	}
	s0i32 = l1
	s1i32 = l3
	s2i32 = 1268
	s1i32 = s1i32 + s2i32
	s0i32 = s0i32 - s1i32
	l2 = s0i32
	s1i32 = 8
	if s0i32 < s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl12
	}
	s0i32 = l1
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l0 = s0i32
	s1i32 = l3
	s2i32 = 1268
	s1i32 = s1i32 + s2i32
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = 64
		s1i32 = l0
		s2i32 = l3
		s3i32 = 1268
		s2i32 = s2i32 + s3i32
		s1i32 = s1i32 - s2i32
		s2i32 = 2
		s1i32 = s1i32 >> (uint32(s2i32) & 31)
		s2i32 = -1
		s1i32 = s1i32 + s2i32
		s1i32 = int32(bits.LeadingZeros32(uint32(s1i32)))
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 - s1i32
		s1i32 = l3
		s2i32 = 1268
		s1i32 = s1i32 + s2i32
		s2i32 = l0
		f408(ctx, s0i32, s1i32, s2i32)
	}
	s0i32 = l2
	s1i32 = 2
	s0i32 = s0i32 >> (uint32(s1i32) & 31)
	l0 = s0i32
	s0i32 = 1
	l1 = s0i32
lbl18:
	s0i32 = l3
	s1i32 = 1268
	s0i32 = s0i32 + s1i32
	s1i32 = l1
	s2i32 = 2
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	s0i32 = s0i32 + s1i32
	l2 = s0i32
	s1i32 = -4
	s0i32 = s0i32 + s1i32
	l6 = s0i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l2
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l2
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	s1i32 = l6
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+0)]))
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l1
	s1i32 = 1
	s0i32 = s0i32 + s1i32
	l1 = s0i32
	s1i32 = l0
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		goto lbl18
	}
	s0i32 = l3
	s1i32 = -2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1256)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1268)]))
	l1 = s1i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1240)])) = uint32(s1i32)
	s0i32 = l3
	s1i64 = -9223372032559808512
	*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+1244)])) = uint64(s1i64)
	s0i32 = l0
	s1i32 = 2
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	s1i32 = l3
	s0i32 = s0i32 + s1i32
	s1i32 = 1264
	s0i32 = s0i32 + s1i32
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)]))
	l0 = s0i32
	s0i32 = l1
	s1i32 = l3
	s2i32 = 1240
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 2147483647
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1224)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = 0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1208)])) = uint32(s1i32)
	s0i32 = l3
	s1i32 = l0
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1212)])) = uint32(s1i32)
	s0i32 = l0
	s1i32 = l3
	s2i32 = 1208
	s1i32 = s1i32 + s2i32
	*(*uint32)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint32(s1i32)
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1188)]))
	l1 = s0i32
	s0i32 = l5
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i32 = l3
		s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1180)]))
		l0 = s0i32
		goto lbl19
	}
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+4)]))
	l0 = s0i32
	s1i32 = l3
	s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+1180)]))
	l2 = s1i32
	s2i32 = l2
	s3i32 = l0
	if s2i32 < s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l0 = s0i32
	s0i32 = l5
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+12)]))
	l2 = s0i32
	s1i32 = l1
	s2i32 = l1
	s3i32 = l2
	if s2i32 > s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	l1 = s0i32
lbl19:
	s0i32 = l3
	s0i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s0i32+1240)]))
	s1i32 = l8
	s2i32 = l0
	s3i32 = l1
	f624(ctx, s0i32, s1i32, s2i32, s3i32)
lbl12:
	s0i32 = l7
	s1i32 = 32
	s0i32 = s0i32 + s1i32
	s0i32 = f35(ctx, s0i32)
	s0i32 = l7
	s0i32 = f35(ctx, s0i32)
	s0i32 = l4
	s1i32 = 20
	s0i32 = s0i32 + s1i32
	s0i32 = f74(ctx, s0i32)
	s0i32 = l4
	f40(ctx, s0i32)
lbl0:
	s0i32 = l3
	s1i32 = 1376
	s0i32 = s0i32 + s1i32
	ctx.g0 = s0i32
}
