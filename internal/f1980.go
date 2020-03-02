package internal

import (
	"math"
	"unsafe"
)

func f1980(ctx *Context, l0 int32, l1 int32, l2 int32, l3 int32) {
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int64
	_ = l7
	var l8 int64
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
	var s5i32 int32
	_ = s5i32
	var s6i32 int32
	_ = s6i32
	var s7i32 int32
	_ = s7i32
	var s8i32 int32
	_ = s8i32
	var s0i64 int64
	_ = s0i64
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	var s3i64 int64
	_ = s3i64
	var s4i64 int64
	_ = s4i64
	var s5i64 int64
	_ = s5i64
	var s6i64 int64
	_ = s6i64
	var s7i64 int64
	_ = s7i64
	var s8i64 int64
	_ = s8i64
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s3f32 float32
	_ = s3f32
	var s4f32 float32
	_ = s4f32
	var s5f32 float32
	_ = s5f32
	s0i32 = l3
	s1i32 = 1
	if s0i32 >= s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0i64 = 2147475456
		s1i64 = 0
		s2i32 = l1
		s2i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s2i32+0)])))
		l7 = s2i64
		s3i64 = 32767
		s2i64 = s2i64 & s3i64
		l8 = s2i64
		s2i32 = int32(s2i64)
		s3i32 = 1023
		if uint32(s2i32) > uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		if s2i32 != 0 {
			// s0i64 = s0i64
		} else {
			s0i64 = s1i64
		}
		s1i64 = l8
		s2i64 = 13
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 939524096
		s1i64 = s1i64 + s2i64
		s0i64 = s0i64 & s1i64
		s1i64 = l7
		s2i64 = 16
		s1i64 = s1i64 << (uint64(s2i64) & 63)
		s2i64 = 2147483648
		s1i64 = s1i64 & s2i64
		s0i64 = s0i64 | s1i64
		s0i32 = int32(s0i64)
		s0f32 = math.Float32frombits(uint32(s0i32))
		s1i64 = 2147475456
		s2i64 = 0
		s3i32 = l1
		s4i32 = l2
		s3i32 = s3i32 + s4i32
		l4 = s3i32
		s3i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)])))
		l7 = s3i64
		s4i64 = 32767
		s3i64 = s3i64 & s4i64
		l8 = s3i64
		s3i32 = int32(s3i64)
		s4i32 = 1023
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		s2i64 = l8
		s3i64 = 13
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 939524096
		s2i64 = s2i64 + s3i64
		s1i64 = s1i64 & s2i64
		s2i64 = l7
		s3i64 = 16
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 2147483648
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s1i32 = int32(s1i64)
		s1f32 = math.Float32frombits(uint32(s1i32))
		l9 = s1f32
		s0f32 = s0f32 + s1f32
		s1f32 = l9
		s0f32 = s0f32 + s1f32
		s1i64 = 2147475456
		s2i64 = 0
		s3i32 = l2
		s4i32 = l4
		s3i32 = s3i32 + s4i32
		l2 = s3i32
		s3i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s3i32+0)])))
		l7 = s3i64
		s4i64 = 32767
		s3i64 = s3i64 & s4i64
		l8 = s3i64
		s3i32 = int32(s3i64)
		s4i32 = 1023
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i64 = s1i64
		} else {
			s1i64 = s2i64
		}
		s2i64 = l8
		s3i64 = 13
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 939524096
		s2i64 = s2i64 + s3i64
		s1i64 = s1i64 & s2i64
		s2i64 = l7
		s3i64 = 16
		s2i64 = s2i64 << (uint64(s3i64) & 63)
		s3i64 = 2147483648
		s2i64 = s2i64 & s3i64
		s1i64 = s1i64 | s2i64
		s1i32 = int32(s1i64)
		s1f32 = math.Float32frombits(uint32(s1i32))
		s0f32 = s0f32 + s1f32
		l9 = s0f32
	lbl1:
		s0i32 = l0
		s1i32 = l5
		s2i32 = 1
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		s0i32 = s0i32 + s1i32
		s1i32 = 65535
		s2i32 = 0
		s3f32 = l9
		s4i64 = 2147475456
		s5i64 = 0
		s6i32 = l1
		s6i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s6i32+2)])))
		l7 = s6i64
		s7i64 = 32767
		s6i64 = s6i64 & s7i64
		l8 = s6i64
		s6i32 = int32(s6i64)
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i64 = s4i64
		} else {
			s4i64 = s5i64
		}
		s5i64 = l8
		s6i64 = 13
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 939524096
		s5i64 = s5i64 + s6i64
		s4i64 = s4i64 & s5i64
		s5i64 = l7
		s6i64 = 16
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 2147483648
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s4i32 = int32(s4i64)
		s4f32 = math.Float32frombits(uint32(s4i32))
		s5i64 = 2147475456
		s6i64 = 0
		s7i32 = l4
		s7i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s7i32+2)])))
		l7 = s7i64
		s8i64 = 32767
		s7i64 = s7i64 & s8i64
		l8 = s7i64
		s7i32 = int32(s7i64)
		s8i32 = 1023
		if uint32(s7i32) > uint32(s8i32) {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i64 = s5i64
		} else {
			s5i64 = s6i64
		}
		s6i64 = l8
		s7i64 = 13
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 939524096
		s6i64 = s6i64 + s7i64
		s5i64 = s5i64 & s6i64
		s6i64 = l7
		s7i64 = 16
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 2147483648
		s6i64 = s6i64 & s7i64
		s5i64 = s5i64 | s6i64
		s5i32 = int32(s5i64)
		s5f32 = math.Float32frombits(uint32(s5i32))
		l9 = s5f32
		s4f32 = s4f32 + s5f32
		s5f32 = l9
		s4f32 = s4f32 + s5f32
		s5i64 = 2147475456
		s6i64 = 0
		s7i32 = l2
		s7i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s7i32+2)])))
		l7 = s7i64
		s8i64 = 32767
		s7i64 = s7i64 & s8i64
		l8 = s7i64
		s7i32 = int32(s7i64)
		s8i32 = 1023
		if uint32(s7i32) > uint32(s8i32) {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i64 = s5i64
		} else {
			s5i64 = s6i64
		}
		s6i64 = l8
		s7i64 = 13
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 939524096
		s6i64 = s6i64 + s7i64
		s5i64 = s5i64 & s6i64
		s6i64 = l7
		s7i64 = 16
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 2147483648
		s6i64 = s6i64 & s7i64
		s5i64 = s5i64 | s6i64
		s5i32 = int32(s5i64)
		s5f32 = math.Float32frombits(uint32(s5i32))
		s4f32 = s4f32 + s5f32
		l9 = s4f32
		s5f32 = l9
		s4f32 = s4f32 + s5f32
		s3f32 = s3f32 + s4f32
		s4i64 = 2147475456
		s5i64 = 0
		s6i32 = l1
		s6i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s6i32+4)])))
		l7 = s6i64
		s7i64 = 32767
		s6i64 = s6i64 & s7i64
		l8 = s6i64
		s6i32 = int32(s6i64)
		s7i32 = 1023
		if uint32(s6i32) > uint32(s7i32) {
			s6i32 = 1
		} else {
			s6i32 = 0
		}
		if s6i32 != 0 {
			// s4i64 = s4i64
		} else {
			s4i64 = s5i64
		}
		s5i64 = l8
		s6i64 = 13
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 939524096
		s5i64 = s5i64 + s6i64
		s4i64 = s4i64 & s5i64
		s5i64 = l7
		s6i64 = 16
		s5i64 = s5i64 << (uint64(s6i64) & 63)
		s6i64 = 2147483648
		s5i64 = s5i64 & s6i64
		s4i64 = s4i64 | s5i64
		s4i32 = int32(s4i64)
		s4f32 = math.Float32frombits(uint32(s4i32))
		s5i64 = 2147475456
		s6i64 = 0
		s7i32 = l4
		s7i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s7i32+4)])))
		l7 = s7i64
		s8i64 = 32767
		s7i64 = s7i64 & s8i64
		l8 = s7i64
		s7i32 = int32(s7i64)
		s8i32 = 1023
		if uint32(s7i32) > uint32(s8i32) {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i64 = s5i64
		} else {
			s5i64 = s6i64
		}
		s6i64 = l8
		s7i64 = 13
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 939524096
		s6i64 = s6i64 + s7i64
		s5i64 = s5i64 & s6i64
		s6i64 = l7
		s7i64 = 16
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 2147483648
		s6i64 = s6i64 & s7i64
		s5i64 = s5i64 | s6i64
		s5i32 = int32(s5i64)
		s5f32 = math.Float32frombits(uint32(s5i32))
		l9 = s5f32
		s4f32 = s4f32 + s5f32
		s5f32 = l9
		s4f32 = s4f32 + s5f32
		s5i64 = 2147475456
		s6i64 = 0
		s7i32 = l2
		s7i64 = int64(*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s7i32+4)])))
		l7 = s7i64
		s8i64 = 32767
		s7i64 = s7i64 & s8i64
		l8 = s7i64
		s7i32 = int32(s7i64)
		s8i32 = 1023
		if uint32(s7i32) > uint32(s8i32) {
			s7i32 = 1
		} else {
			s7i32 = 0
		}
		if s7i32 != 0 {
			// s5i64 = s5i64
		} else {
			s5i64 = s6i64
		}
		s6i64 = l8
		s7i64 = 13
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 939524096
		s6i64 = s6i64 + s7i64
		s5i64 = s5i64 & s6i64
		s6i64 = l7
		s7i64 = 16
		s6i64 = s6i64 << (uint64(s7i64) & 63)
		s7i64 = 2147483648
		s6i64 = s6i64 & s7i64
		s5i64 = s5i64 | s6i64
		s5i32 = int32(s5i64)
		s5f32 = math.Float32frombits(uint32(s5i32))
		s4f32 = s4f32 + s5f32
		l9 = s4f32
		s3f32 = s3f32 + s4f32
		s4f32 = 0.0625
		s3f32 = s3f32 * s4f32
		s3i32 = int32(math.Float32bits(s3f32))
		l6 = s3i32
		s4i32 = 2147475456
		s3i32 = s3i32 & s4i32
		s4i32 = 947904511
		if uint32(s3i32) > uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		s2i32 = l6
		s3i32 = -134217728
		s2i32 = s2i32 - s3i32
		s3i32 = 13
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s1i32 = s1i32 & s2i32
		s2i32 = l6
		s3i32 = 16
		s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
		s3i32 = 32768
		s2i32 = s2i32 & s3i32
		s1i32 = s1i32 | s2i32
		*(*uint16)(unsafe.Pointer(&ctx.Mem[int(s0i32+0)])) = uint16(s1i32)
		s0i32 = l2
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l2 = s0i32
		s0i32 = l4
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l4 = s0i32
		s0i32 = l1
		s1i32 = 4
		s0i32 = s0i32 + s1i32
		l1 = s0i32
		s0i32 = l5
		s1i32 = 1
		s0i32 = s0i32 + s1i32
		l5 = s0i32
		s1i32 = l3
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		if s0i32 != 0 {
			goto lbl1
		}
	}
}
