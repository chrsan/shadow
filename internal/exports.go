package internal

func (c *Context) SkiaCanvasClear(p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) {
	f1705(c, p0, p1, p2, p3, p4)
}

func (c *Context) SkiaCanvasTranslate(p0 int32, p1 float32, p2 float32) {
	f1349(c, p0, p1, p2)
}

func (c *Context) SkiaPaintSize() int32 {
	return f1299(c)
}

func (c *Context) SkiaPathGetPoint(p0 int32, p1 int32, p2 int32) {
	f1002(c, p0, p1, p2)
}

func (c *Context) SkiaSurfaceReadPixels(p0 int32, p1 int32) {
	f1979(c, p0, p1)
}

func (c *Context) SkiaCanvasSave(p0 int32) {
	f1573(c, p0)
}

func (c *Context) SkiaCanvasDrawPath(p0 int32, p1 int32, p2 int32) {
	f1312(c, p0, p1, p2)
}

func (c *Context) SkiaPathMoveTo(p0 int32, p1 float32, p2 float32) {
	f1078(c, p0, p1, p2)
}

func (c *Context) SkiaMatrixCreate() int32 {
	return f1280(c)
}

func (c *Context) SkiaSurfaceGetCanvas(p0 int32) int32 {
	return f1853(c, p0)
}

func (c *Context) SkiaCanvasScale(p0 int32, p1 float32, p2 float32) {
	f1401(c, p0, p1, p2)
}

func (c *Context) SkiaPaintCreate() int32 {
	return f1289(c)
}

func (c *Context) SkiaPaintGetAlpha(p0 int32) int32 {
	return f1257(c, p0)
}

func (c *Context) SkiaPaintReset(p0 int32) {
	f1143(c, p0)
}

func (c *Context) SkiaPathSize() int32 {
	return f1128(c)
}

func (c *Context) SkiaPathGetPoints(p0 int32, p1 int32, p2 int32) int32 {
	return f993(c, p0, p1, p2)
}

func (c *Context) StackAlloc(p0 int32) int32 {
	return f1343(c, p0)
}

func (c *Context) StackRestore(p0 int32) {
	f1342(c, p0)
}

func (c *Context) SkiaMatrixCreateFrom(p0 float32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32) int32 {
	return f1118(c, p0, p1, p2, p3, p4, p5)
}

func (c *Context) SkiaCanvasFlush(p0 int32) {
	f1647(c, p0)
}

func (c *Context) SkiaPaintSetStrokeJoin(p0 int32, p1 int32) {
	f1169(c, p0, p1)
}

func (c *Context) SkiaPathGetFillType(p0 int32) int32 {
	return f1094(c, p0)
}

func (c *Context) SkiaPathCountVerbs(p0 int32) int32 {
	return f983(c, p0)
}

func (c *Context) SkiaMatrixInit(p0 int32) {
	f557(c, p0)
}

func (c *Context) SkiaIsSurfaceBgra() int32 {
	return f1781(c)
}

func (c *Context) SkiaPaintDestroy(p0 int32) {
	f1279(c, p0)
}

func (c *Context) SkiaPaintSetAntiAlias(p0 int32, p1 int32) {
	f1225(c, p0, p1)
}

func (c *Context) SkiaPaintSetStyle(p0 int32, p1 int32) {
	f1209(c, p0, p1)
}

func (c *Context) SkiaPaintGetStrokeWidth(p0 int32) float32 {
	return f1199(c, p0)
}

func (c *Context) SkiaPathLineTo(p0 int32, p1 float32, p2 float32) {
	f1071(c, p0, p1, p2)
}

func (c *Context) SkiaPathClose(p0 int32) {
	f1046(c, p0)
}

func (c *Context) SkiaPathGetBounds(p0 int32, p1 int32) {
	f977(c, p0, p1)
}

func (c *Context) Malloc(p0 int32) int32 {
	return f164(c, p0)
}

func (c *Context) SkiaSurfaceDestroy(p0 int32) {
	f1914(c, p0)
}

func (c *Context) SkiaPaintIsAntiAlias(p0 int32) int32 {
	return f1236(c, p0)
}

func (c *Context) SkiaPaintGetStrokeMiter(p0 int32) float32 {
	return f1160(c, p0)
}

func (c *Context) SkiaCanvasGetTotalMatrix(p0 int32, p1 int32) {
	f1318(c, p0, p1)
}

func (c *Context) SkiaPaintSetAlpha(p0 int32, p1 int32) {
	f1247(c, p0, p1)
}

func (c *Context) SkiaPaintGetStrokeCap(p0 int32) int32 {
	return f1188(c, p0)
}

func (c *Context) SkiaPathCreate() int32 {
	return f1108(c)
}

func (c *Context) SkiaPathConicTo(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32) {
	f1057(c, p0, p1, p2, p3, p4, p5)
}

func (c *Context) SkiaPathCountPoints(p0 int32) int32 {
	return f1012(c, p0)
}

func (c *Context) SkiaPathInit(p0 int32) {
	f1117(c, p0)
}

func (c *Context) SkiaPathCubicTo(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32, p6 float32) {
	f1050(c, p0, p1, p2, p3, p4, p5, p6)
}

func (c *Context) SkiaPathApply(p0 int32, p1 int32, p2 int32, p3 int32) int32 {
	return f1040(c, p0, p1, p2, p3)
}

func (c *Context) Free(p0 int32) {
	f12(c, p0)
}

func (c *Context) SkiaMatrixCreateInverse(p0 int32, p1 int32) int32 {
	return f1047(c, p0, p1)
}

func (c *Context) SkiaSurfaceDataSize() int32 {
	return f464(c)
}

func (c *Context) SkiaCanvasRestore(p0 int32) {
	f1473(c, p0)
}

func (c *Context) SkiaPaintInit(p0 int32) {
	f1295(c, p0)
}

func (c *Context) SkiaPaintGetColor(p0 int32) int32 {
	return f1273(c, p0)
}

func (c *Context) SkiaPaintGetStyle(p0 int32) int32 {
	return f1215(c, p0)
}

func (c *Context) SkiaPathDestroy(p0 int32) {
	f1101(c, p0)
}

func (c *Context) SkiaPathReset(p0 int32) {
	f1031(c, p0)
}

func (c *Context) SkiaCanvasSetMatrix(p0 int32, p1 int32) {
	f1334(c, p0, p1)
}

func (c *Context) SkiaPathQuadTo(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32) {
	f1064(c, p0, p1, p2, p3, p4)
}

func (c *Context) SkiaMatrixDestroy(p0 int32) {
	f13(c, p0)
}

func (c *Context) SkiaMatrixSet(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 float32, p6 float32) {
	f968(c, p0, p1, p2, p3, p4, p5, p6)
}

func (c *Context) SkiaPathGetVerbs(p0 int32, p1 int32, p2 int32) int32 {
	return f982(c, p0, p1, p2)
}

func (c *Context) StackSave() int32 {
	return f1344(c)
}

func (c *Context) SkiaMatrixSize() int32 {
	return f2121(c)
}

func (c *Context) SkiaCanvasConcat(p0 int32, p1 int32) {
	f1323(c, p0, p1)
}

func (c *Context) SkiaRectSize() int32 {
	return f1133(c)
}

func (c *Context) Start() {
	f451(c)
}

func (c *Context) SkiaMatrixReset(p0 int32) {
	f557(c, p0)
}

func (c *Context) SkiaPaintSetColor(p0 int32, p1 int32, p2 int32, p3 int32, p4 int32) {
	f1265(c, p0, p1, p2, p3, p4)
}

func (c *Context) SkiaPaintSetStrokeCap(p0 int32, p1 int32) {
	f1186(c, p0, p1)
}

func (c *Context) SkiaPointSize() int32 {
	return f464(c)
}

func (c *Context) SkiaPathRewind(p0 int32) {
	f1022(c, p0)
}

func (c *Context) SkiaSurfaceCreateRgba(p0 int32, p1 int32) int32 {
	return f778(c, p0, p1)
}

func (c *Context) SkiaCanvasDrawRect(p0 int32, p1 float32, p2 float32, p3 float32, p4 float32, p5 int32) {
	f1306(c, p0, p1, p2, p3, p4, p5)
}

func (c *Context) SkiaPathSetFillType(p0 int32, p1 int32) {
	f1086(c, p0, p1)
}

func (c *Context) SkiaSurfaceCreateRgbaPremultiplied(p0 int32, p1 int32) int32 {
	return f2079(c, p0, p1)
}

func (c *Context) SkiaPaintSetStrokeWidth(p0 int32, p1 float32) {
	f1192(c, p0, p1)
}

func (c *Context) SkiaPaintGetStrokeJoin(p0 int32) int32 {
	return f1178(c, p0)
}

func (c *Context) SkiaPaintSetStrokeMiter(p0 int32, p1 float32) {
	f1152(c, p0, p1)
}

func (c *Context) SkiaPathComputeTightBounds(p0 int32, p1 int32) {
	f959(c, p0, p1)
}
