# Build the CLI application into a .build directory.
[group('project')]
build:
  [[ -d "$BUILD_DIR" ]] && rm -rf "$BUILD_DIR"
  go build -o "$BUILD_DIR/$BUILD_BIN" main.go
