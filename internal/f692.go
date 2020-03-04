package internal

func f692(ctx *Context, l0 float32, l1 float32, l2 float32, l3 float32, l4 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s3i32 int32
	_ = s3i32
	var s0f32 float32
	_ = s0f32
	var s1f32 float32
	_ = s1f32
	var s2f32 float32
	_ = s2f32
	var s3f32 float32
	_ = s3f32
	s0f32 = l3
	s1f32 = l0
	s0f32 = s0f32 - s1f32
	s1f32 = l1
	s2f32 = l2
	s1f32 = s1f32 - s2f32
	s2f32 = 3
	s1f32 = s1f32 * s2f32
	s0f32 = s0f32 + s1f32
	s1f32 = l0
	s2f32 = l1
	s1f32 = s1f32 - s2f32
	s2f32 = l1
	s1f32 = s1f32 - s2f32
	s2f32 = l2
	s1f32 = s1f32 + s2f32
	l2 = s1f32
	s2f32 = l2
	s1f32 = s1f32 + s2f32
	s2f32 = l1
	s3f32 = l0
	s2f32 = s2f32 - s3f32
	s3i32 = l4
	s0i32 = f122(ctx, s0f32, s1f32, s2f32, s3i32)
	return s0i32
}
