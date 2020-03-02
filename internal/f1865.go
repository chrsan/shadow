package internal

func f1865(ctx *Context, l0 int32, l1 float32, l2 float32) float32 {
	var s2i32 int32
	_ = s2i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0f32 = l2
	s1f32 = 12.92
	s0f32 = s0f32 * s1f32
	s1f32 = l2
	s2f32 = 0.41666666
	s1f32 = f251(ctx, s1f32, s2f32)
	s2f32 = 1.055
	s1f32 = s1f32 * s2f32
	s2f32 = -0.055
	s1f32 = s1f32 + s2f32
	s2f32 = l2
	s3f32 = 0.0031308
	if s2f32 <= s3f32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	if s2i32 != 0 {
		// s0f32 = s0f32
	} else {
		s0f32 = s1f32
	}
	return s0f32
}
