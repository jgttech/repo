# Handle initial installation.
[group('project')]
install:
  devbox install
  go mod download
