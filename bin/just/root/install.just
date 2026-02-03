# Handle initial installation.
[group('project')]
install:
  devbox install
  go mod download
  just docker up
  just docker exec go mod download
