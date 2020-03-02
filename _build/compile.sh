#!/bin/bash

set -ex

BASE_DIR=`cd $(dirname ${BASH_SOURCE[0]}) && pwd`
BUILD_DIR=${BUILD_DIR:="${BASE_DIR}/out/wasm"}
mkdir -p $BUILD_DIR

if [[ ! -d $EMSDK ]]; then
  cat >&2 << "EOF"
Be sure to set the EMSDK environment variable to the location of Emscripten SDK:
    https://emscripten.org/docs/getting_started/downloads.html
EOF
  exit 1
fi

source ${EMSDK}/emsdk_env.sh

EMCC=`which emcc`
if [[ -z $EMCC ]]; then
  echo "Could not find emcc"
  exit 1
fi

EMCXX=`which em++`
if [[ -z $EMCXX ]]; then
  echo "Could not find em++"
  exit 1
fi

EMAR=`which emar`
if [[ -z $EMAR ]]; then
  echo "Could not find emar"
  exit 1
fi

NINJA=`which ninja`
if [[ -z $NINJA ]]; then
  echo "Could not find ninja"
  exit 1
fi

echo "Compiling bitcode"

pushd ${BASE_DIR}/../_deps/skia

./bin/gn gen ${BUILD_DIR} \
  --args="cc=\"${EMCC}\" \
  cxx=\"${EMCXX}\" \
  ar=\"${EMAR}\" \
  extra_cflags_cc=[\"-Wdeprecated\", \"-fno-exceptions\"] \
  extra_cflags=[\"-DSK_DISABLE_READBUFFER=1\", \"-s\", \"WARN_UNALIGNED=1\"] \
  is_debug=false \
  is_official_build=true \
  is_component_build=false \
  target_cpu=\"wasm\" \
  skia_use_expat=false \
  skia_use_icu=false \
  skia_use_libjpeg_turbo=false \
  skia_use_libwebp=false \
  skia_use_lua=false \
  skia_use_piex=false \
  skia_use_libheif=false \
  skia_enable_discrete_gpu=false \
  skia_enable_gpu=false \
  skia_enable_pdf=false \
  skia_enable_skottie=false \
  skia_enable_tools=false \
  skia_use_libpng=false \
  skia_use_harfbuzz=false \
  skia_use_icu=false \
  skia_enable_skshaper=false \
  skia_use_zlib=false \
  skia_use_libgifcodec=false \
  skia_use_xps=false \
  skia_use_fontconfig=false \
  skia_use_fonthost_mac=false \
  skia_use_freetype=false \
  skia_enable_fontmgr_empty=true \
  skia_enable_fontmgr_fuchsia=false \
  skia_enable_fontmgr_win=false \
  skia_enable_fontmgr_win_gdi=false \
  skia_enable_skshaper=false"

${NINJA} -C ${BUILD_DIR} skia

popd

${EMCXX} -O3 -s EVAL_CTORS=1 --llvm-lto 3 -std=c++17 \
  -DWASM_BUILD=1 \
  -Wall \
  -Wdeprecated \
  -fno-rtti -fno-exceptions \
  -I${BASE_DIR}/../_cc \
  -I${BASE_DIR}/../_deps/skia \
  -s WASM=1 \
  -s STANDALONE_WASM=1 \
  -s ALLOW_MEMORY_GROWTH=1 \
  -s NO_EXIT_RUNTIME=1 \
  -s NO_FILESYSTEM=1 \
  -s STRICT=1 \
  -s ENVIRONMENT=node \
  -s EXPORTED_FUNCTIONS=@${BASE_DIR}/exported_functions.json \
  -s EXTRA_EXPORTED_RUNTIME_METHODS=@${BASE_DIR}/exported_runtime_methods.json \
  -o ${BUILD_DIR}/shadow.js \
  ${BASE_DIR}/../_cc/skia_bindings.cc \
  ${BUILD_DIR}/libskia.a \

echo "Converting WASM to wagu IR"

pushd ${BASE_DIR}/../_deps/wagu

go run cmd/*.go ir -o ${BASE_DIR}/shadow.ir ${BUILD_DIR}/shadow.wasm

echo "Generating Go source from wagu IR"

go run cmd/*.go gen -c -e -m -u -p internal -d ${BASE_DIR}/../internal ${BASE_DIR}/shadow.ir
