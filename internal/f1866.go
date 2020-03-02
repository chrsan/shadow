package internal

func f1866(ctx *Context, l0 int32, l1 float32, l2 float32) float32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	s0f32 = l2
	s1f32 = 0.04045
	if s0f32 <= s1f32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	s1i32 = 1
	s0i32 = s0i32 ^ s1i32
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	if s0i32 != 0 {
		s0f32 = l2
		s1f32 = 12.92
		s0f32 = s0f32 / s1f32
		return s0f32
	}
	s0f32 = l2
	s1f32 = 0.055
	s0f32 = s0f32 + s1f32
	s1f32 = 1.055
	s0f32 = s0f32 / s1f32
	s1f32 = 2.4
	s0f32 = f251(ctx, s0f32, s1f32)
	return s0f32
}
