#!/bin/bash

BASE_DIR=$(cd $(dirname ${BASH_SOURCE[0]}) && pwd)

git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git ${BASE_DIR}/_deps/depot_tools
git clone https://skia.googlesource.com/skia.git ${BASE_DIR}/_deps/skia
git clone https://github.com/chrsan/wagu.git ${BASE_DIR}/_deps/wagu

pushd ${BASE_DIR}/_deps/skia >/dev/null

python2 tools/git-sync-deps
