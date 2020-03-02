# shadow

Using [Skia][1] from [Go][2] via [WASM][3] and [wagu][4].

## Building

Not needed unless you've changed something.

The generated [Go][2] is under the `internal` directory.

If you do want to rebuild, make sure that [emscripten][5] is installed.

Then do the following:

```console
bash dload_deps.sh
cd _build
bash compile.sh
```

## Sample

See `cmd/sample/main.go`. Nothing fancy there, but it's a start.

[1]: https://skia.org
[2]: https://golang.org
[3]: https://webassembly.org
[4]: https://github.com/chrsan/wagu
[5]: https://emscripten.org
