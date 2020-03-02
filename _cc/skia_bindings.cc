#include "skia_bindings.hh"

#include <include/core/SkCanvas.h>
#include <include/core/SkColor.h>
#include <include/core/SkGraphics.h>
#include <include/core/SkMatrix.h>
#include <include/core/SkPaint.h>
#include <include/core/SkPath.h>
#include <include/core/SkPathEffect.h>
#include <include/core/SkPixmap.h>
#include <include/core/SkPoint.h>
#include <include/core/SkRect.h>
#include <include/core/SkSurface.h>

#include <memory>

namespace {
#ifdef WASM_BUILD
SkSurface* SkiaSurfaceCreate(int width, int height, SkAlphaType alpha_type) {
  // SkGraphics::Init is idempotent.
  SkGraphics::Init();

  SkImageInfo info =
      SkImageInfo::Make(width, height, kN32_SkColorType, alpha_type);

  sk_sp<SkSurface> surface = SkSurface::MakeRaster(info);

  // The surface ref count will equal one after the pointer is returned.
  return surface.release();
}
#endif

#ifndef WASM_BUILD
SkSurface* SkiaSurfaceCreateForPixels(int width,
                                      int height,
                                      SkAlphaType alpha_type,
                                      void* pixels) {
  // SkGraphics::Init is idempotent.
  SkGraphics::Init();

  SkImageInfo info =
      SkImageInfo::Make(width, height, kN32_SkColorType, alpha_type);

  sk_sp<SkSurface> surface =
      SkSurface::MakeRasterDirect(info, pixels, info.minRowBytes());

  // The surface ref count will equal one after the pointer is returned.
  return surface.release();
}
#endif
}  // namespace

extern "C" {

// Matrix
#ifdef WASM_BUILD
size_t skia_matrix_size() {
  return sizeof(SkMatrix);
}
#endif

skia_matrix* skia_matrix_create() {
  return reinterpret_cast<skia_matrix*>(new SkMatrix());
}

void skia_matrix_destroy(skia_matrix* matrix) {
  delete reinterpret_cast<SkMatrix*>(matrix);
}

skia_matrix* skia_matrix_create_from(float scale_x,
                                     float skew_x,
                                     float trans_x,
                                     float skew_y,
                                     float scale_y,
                                     float trans_y) {
  SkMatrix* m = new SkMatrix();
  m->setAll(scale_x, skew_x, trans_x, skew_y, scale_y, trans_y, 0, 0, 1);
  return reinterpret_cast<skia_matrix*>(m);
}

uint8_t skia_matrix_create_inverse(const skia_matrix* matrix,
                                   skia_matrix* inverse) {
  SkMatrix* m = reinterpret_cast<SkMatrix*>(inverse);
  if (reinterpret_cast<const SkMatrix*>(matrix)->invert(m)) {
    return 1;
  }

  return 0;
}

void skia_matrix_set(skia_matrix* matrix,
                     float scale_x,
                     float skew_x,
                     float trans_x,
                     float skew_y,
                     float scale_y,
                     float trans_y) {
  SkMatrix* m = reinterpret_cast<SkMatrix*>(matrix);
  m->setAll(scale_x, skew_x, trans_x, skew_y, scale_y, trans_y, 0, 0, 1);
}

void skia_matrix_reset(skia_matrix* matrix) {
  reinterpret_cast<SkMatrix*>(matrix)->reset();
}

// Surface

#ifdef WASM_BUILD
size_t skia_surface_data_size() {
  return sizeof(skia_surface_data);
}

skia_surface* skia_surface_create_rgba(int width, int height) {
  return reinterpret_cast<skia_surface*>(
      SkiaSurfaceCreate(width, height, kUnpremul_SkAlphaType));
}

skia_surface* skia_surface_create_rgba_premultiplied(int width, int height) {
  return reinterpret_cast<skia_surface*>(
      SkiaSurfaceCreate(width, height, kPremul_SkAlphaType));
}

void skia_surface_read_pixels(skia_surface* surface, skia_surface_data* data) {
  SkSurface* s = reinterpret_cast<SkSurface*>(surface);

  data->ptr = nullptr;
  data->size = 0;

  SkPixmap pixmap;
  if (s->peekPixels(&pixmap)) {
    data->ptr = static_cast<uint8_t*>(pixmap.writable_addr());
    data->size = static_cast<uint32_t>(pixmap.computeByteSize());
  }
}
#else
skia_surface* skia_surface_create_rgba(int width, int height, void* pixels) {
  return reinterpret_cast<skia_surface*>(
      SkiaSurfaceCreateForPixels(width, height, kUnpremul_SkAlphaType, pixels));
}

skia_surface* skia_surface_create_rgba_premultiplied(int width,
                                                     int height,
                                                     void* pixels) {
  return reinterpret_cast<skia_surface*>(
      SkiaSurfaceCreateForPixels(width, height, kPremul_SkAlphaType, pixels));
}
#endif

void skia_surface_destroy(skia_surface* surface) {
  if (surface) {
    // SkSurface is ref counted.
    SkSurface* s = reinterpret_cast<SkSurface*>(surface);
    SkSafeUnref(s);
  }
}

skia_canvas* skia_surface_get_canvas(skia_surface* surface) {
  SkSurface* s = reinterpret_cast<SkSurface*>(surface);
  return reinterpret_cast<skia_canvas*>(s->getCanvas());
}

uint8_t skia_is_surface_bgra() {
  return kN32_SkColorType == kBGRA_8888_SkColorType;
}

// Canvas

void skia_canvas_clear(skia_canvas* canvas,
                       uint8_t r,
                       uint8_t g,
                       uint8_t b,
                       uint8_t a) {
  SkColor color = SkColorSetARGB(a, r, g, b);
  reinterpret_cast<SkCanvas*>(canvas)->clear(color);
}

void skia_canvas_flush(skia_canvas* canvas) {
  reinterpret_cast<SkCanvas*>(canvas)->flush();
}

void skia_canvas_save(skia_canvas* canvas) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->save();
}

void skia_canvas_restore(skia_canvas* canvas) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->restore();
}

void skia_canvas_scale(skia_canvas* canvas, float sx, float sy) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->scale(sx, sy);
}

void skia_canvas_translate(skia_canvas* canvas, float dx, float dy) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->translate(dx, dy);
}

void skia_canvas_set_matrix(skia_canvas* canvas, const skia_matrix* matrix) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->setMatrix(*reinterpret_cast<const SkMatrix*>(matrix));
}

void skia_canvas_concat(skia_canvas* canvas, const skia_matrix* matrix) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->concat(*reinterpret_cast<const SkMatrix*>(matrix));
}

void skia_canvas_get_total_matrix(const skia_canvas* canvas,
                                  skia_matrix* matrix) {
  const SkCanvas* c = reinterpret_cast<const SkCanvas*>(canvas);
  *reinterpret_cast<SkMatrix*>(matrix) = c->getTotalMatrix();
}

void skia_canvas_draw_path(skia_canvas* canvas,
                           const skia_path* path,
                           const skia_paint* paint) {
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->drawPath(*reinterpret_cast<const SkPath*>(path),
              *reinterpret_cast<const SkPaint*>(paint));
}

void skia_canvas_draw_rect(skia_canvas* canvas,
                           float x,
                           float y,
                           float w,
                           float h,
                           const skia_paint* paint) {
  SkRect rect = SkRect::MakeXYWH(x, y, w, h);
  SkCanvas* c = reinterpret_cast<SkCanvas*>(canvas);
  c->drawRect(rect, *reinterpret_cast<const SkPaint*>(paint));
}

// Paint

#ifdef WASM_BUILD
size_t skia_paint_size() {
  return sizeof(SkPaint);
}
#endif

skia_paint* skia_paint_create() {
  return reinterpret_cast<skia_paint*>(new SkPaint());
}

void skia_paint_destroy(skia_paint* paint) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);

  // Setting these references to nullptr should decrement their ref count.
  p->setShader(nullptr);
  p->setPathEffect(nullptr);

  // SkPaint is not ref counted, so explicitly delete.
  delete p;
}

uint32_t skia_paint_get_color(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return p->getColor();
}

void skia_paint_set_color(skia_paint* paint,
                          uint8_t r,
                          uint8_t g,
                          uint8_t b,
                          uint8_t a) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setARGB(a, r, g, b);
}

uint8_t skia_paint_get_alpha(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return p->getAlpha();
}

void skia_paint_set_alpha(skia_paint* paint, uint8_t a) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setAlpha(a);
}

uint8_t skia_paint_is_anti_alias(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return p->isAntiAlias();
}

void skia_paint_set_anti_alias(skia_paint* paint, uint8_t aa) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setAntiAlias(aa);
}

PaintStyle skia_paint_get_style(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return (PaintStyle)p->getStyle();
}

void skia_paint_set_style(skia_paint* paint, PaintStyle style) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setStyle((SkPaint::Style)style);
}

float skia_paint_get_stroke_width(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return p->getStrokeWidth();
}

void skia_paint_set_stroke_width(skia_paint* paint, float width) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setStrokeWidth(width);
}

StrokeCap skia_paint_get_stroke_cap(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return (StrokeCap)p->getStrokeCap();
}

void skia_paint_set_stroke_cap(skia_paint* paint, StrokeCap cap) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setStrokeCap((SkPaint::Cap)cap);
}

StrokeJoin skia_paint_get_stroke_join(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return (StrokeJoin)p->getStrokeJoin();
}

void skia_paint_set_stroke_join(skia_paint* paint, StrokeJoin join) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setStrokeJoin((SkPaint::Join)join);
}

float skia_paint_get_stroke_miter(const skia_paint* paint) {
  const SkPaint* p = reinterpret_cast<const SkPaint*>(paint);
  return p->getStrokeMiter();
}

void skia_paint_set_stroke_miter(skia_paint* paint, float miter) {
  SkPaint* p = reinterpret_cast<SkPaint*>(paint);
  p->setStrokeMiter(miter);
}

void skia_paint_reset(skia_paint* paint) {
  reinterpret_cast<SkPaint*>(paint)->reset();
}

// Point and rect

#ifdef WASM_BUILD
size_t skia_point_size() {
  return sizeof(skia_point);
}

size_t skia_rect_size() {
  return sizeof(skia_rect);
}
#endif

// Path

#ifdef WASM_BUILD
size_t skia_path_size() {
  return sizeof(SkPath);
}
#endif

skia_path* skia_path_create() {
  return reinterpret_cast<skia_path*>(new SkPath());
}

void skia_path_destroy(skia_path* path) {
  delete reinterpret_cast<SkPath*>(path);
}

PathFillType skia_path_get_fill_type(const skia_path* path) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  return (PathFillType)p->getFillType();
}

void skia_path_set_fill_type(skia_path* path, PathFillType type) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->setFillType((SkPathFillType)type);
}

void skia_path_move_to(skia_path* path, float x, float y) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->moveTo(x, y);
}

void skia_path_line_to(skia_path* path, float x, float y) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->lineTo(x, y);
}

void skia_path_quad_to(skia_path* path,
                       float x1,
                       float y1,
                       float x2,
                       float y2) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->quadTo(x1, y1, x2, y2);
}

void skia_path_conic_to(skia_path* path,
                        float x1,
                        float y1,
                        float x2,
                        float y2,
                        float w) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->conicTo(x1, y1, x2, y2, w);
}

void skia_path_cubic_to(skia_path* path,
                        float x1,
                        float y1,
                        float x2,
                        float y2,
                        float x3,
                        float y3) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->cubicTo(x1, y1, x2, y2, x3, y3);
}

void skia_path_close(skia_path* path) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->close();
}

uint8_t skia_path_apply(skia_path* path,
                        const uint8_t* verbs,
                        int num_verbs,
                        const float* values) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  const float* v = values;
  for (int i = 0; i < num_verbs; ++i) {
    uint8_t verb = verbs[i];
    if (verb > static_cast<uint8_t>(SkPath::kClose_Verb)) {
      return 0;
    }

    switch (static_cast<SkPath::Verb>(verbs[i])) {
      case SkPath::kMove_Verb: {
        p->moveTo(*v, *(v + 1));
        v += 2;
        break;
      }
      case SkPath::kLine_Verb: {
        p->lineTo(*v, *(v + 1));
        v += 2;
        break;
      }
      case SkPath::kQuad_Verb:
        p->quadTo(*v, *(v + 1), *(v + 2), *(v + 3));
        v += 4;
        break;
      case SkPath::kConic_Verb: {
        p->conicTo(*v, *(v + 1), *(v + 2), *(v + 3), *(v + 4));
        v += 4;
        break;
      }
      case SkPath::kCubic_Verb: {
        p->cubicTo(*v, *(v + 1), *(v + 2), *(v + 3), *(v + 4), *(v + 5));
        v += 6;
        break;
      }
      case SkPath::kClose_Verb:
        p->close();
        break;
      default:
        return 0;
    }
  }

  return 1;
}

void skia_path_reset(skia_path* path) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->reset();
}

void skia_path_rewind(skia_path* path) {
  SkPath* p = reinterpret_cast<SkPath*>(path);
  p->rewind();
}

int skia_path_count_points(const skia_path* path) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  return p->countPoints();
}

skia_point skia_path_get_point(const skia_path* path, int index) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  SkPoint pt = p->getPoint(index);
  return *reinterpret_cast<skia_point*>(&pt);
}

int skia_path_get_points(const skia_path* path, skia_point* points, int max) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  SkPoint* pts = reinterpret_cast<SkPoint*>(points);
  return p->getPoints(pts, max);
}

int skia_path_count_verbs(const skia_path* path) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  return p->countVerbs();
}

int skia_path_get_verbs(const skia_path* path, uint8_t* verbs, int max) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  return p->getVerbs(verbs, max);
}

skia_rect skia_path_get_bounds(const skia_path* path) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  SkRect r = p->getBounds();
  return *reinterpret_cast<skia_rect*>(&r);
}

skia_rect skia_path_compute_tight_bounds(const skia_path* path) {
  const SkPath* p = reinterpret_cast<const SkPath*>(path);
  SkRect r = p->computeTightBounds();
  return *reinterpret_cast<skia_rect*>(&r);
}
}
