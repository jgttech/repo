# Handle starting the docker services for development.
[group('docker')]
up:
  echo "BUILD_DIR: $BUILD_DIR"
  echo "Docker up!!"
