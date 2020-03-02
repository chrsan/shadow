package internal

import (
	"unsafe"
)

func f1930(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	s0i32 = l3
	if s0i32 != 0 {
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+41)])
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = 1
			ctx.Mem[int(s0i32+40)] = uint8(s1i32)
			s0i32 = l3
			s0i32 = f129(ctx, s0i32)
			s0i32 = l3
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			f143(ctx, s0i32)
			s0i32 = l3
			s1i32 = 1
			*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+41)])) = uint16(s1i32)
			return
		}
		s0i32 = l1
		s1i32 = l2
		s0i32 = s0i32 | s1i32
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			s0i32 = l3
			s1i32 = l0
			f223(ctx, s0i32, s1i32)
			s0i32 = l3
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			s1i32 = l0
			s2i32 = 20
			s1i32 = s1i32 + s2i32
			f512(ctx, s0i32, s1i32)
			s0i32 = l3
			s1i32 = l0
			s1i64 = *(*int64)(unsafe.Pointer(&ctx.Mem[int(s1i32+40)]))
			*(*uint64)(unsafe.Pointer(&ctx.Mem[int(s0i32+40)])) = uint64(s1i64)
			return
		}
		s0i32 = l3
		s1i32 = l0
		s1i32 = int32(ctx.Mem[int(s1i32+40)])
		ctx.Mem[int(s0i32+40)] = uint8(s1i32)
		s0i32 = l0
		s0i32 = int32(ctx.Mem[int(s0i32+40)])
		if s0i32 != 0 {
			s0i32 = l0
			s1i32 = l1
			s2i32 = l2
			s3i32 = l3
			f656(ctx, s0i32, s1i32, s2i32, s3i32)
			s0i32 = l3
			s1i32 = 20
			s0i32 = s0i32 + s1i32
			f143(ctx, s0i32)
			goto lbl3
		}
		s0i32 = l0
		s1i32 = 20
		s0i32 = s0i32 + s1i32
		s1i32 = l1
		s2i32 = l2
		s3i32 = l3
		s4i32 = 20
		s3i32 = s3i32 + s4i32
		f1183(ctx, s0i32, s1i32, s2i32, s3i32)
		s0i32 = l3
		s0i32 = f129(ctx, s0i32)
	lbl3:
		s0i32 = l3
		s1i32 = l3
		s1i32 = int32(ctx.Mem[int(s1i32+40)])
		if s1i32 != 0 {
			s1i32 = l3
			s2i32 = l3
			s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+16)]))
			s3i32 = -1
			if s2i32 == s3i32 {
				s2i32 = 1
			} else {
				s2i32 = 0
			}
			ctx.Mem[int(s1i32+41)] = uint8(s2i32)
			goto lbl7
		}
		s1i32 = l3
		s2i32 = l3
		s2i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s2i32+36)]))
		l0 = s2i32
		if s2i32 == 0 {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		ctx.Mem[int(s1i32+41)] = uint8(s2i32)
		s1i32 = l0
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
		s1i32 = l3
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		l0 = s1i32
		s1i32 = f92(ctx, s1i32)
		if s1i32 != 0 {
			s1i32 = l3
			s2i32 = l0
			s1i32 = f81(ctx, s1i32, s2i32)
			s1i32 = l0
			f143(ctx, s1i32)
			s1i32 = l3
			s2i32 = 1
			ctx.Mem[int(s1i32+40)] = uint8(s2i32)
			goto lbl7
		}
		s1i32 = l3
		s1i32 = int32(ctx.Mem[int(s1i32+40)])
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		if s1i32 != 0 {
			goto lbl6
		}
	lbl7:
		s1i32 = l3
		s1i32 = *(*int32)(unsafe.Pointer(&ctx.Mem[int(s1i32+16)]))
		if s1i32 == 0 {
			s1i32 = 1
		} else {
			s1i32 = 0
		}
		goto lbl5
	lbl6:
		s1i32 = l3
		s2i32 = 20
		s1i32 = s1i32 + s2i32
		s1i32 = f92(ctx, s1i32)
	lbl5:
		ctx.Mem[int(s0i32+42)] = uint8(s1i32)
	}
}
