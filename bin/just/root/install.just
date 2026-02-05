# Handle initial installation.
[group('project')]
install:
  echo "Running 'devbox install', please wait..."
  devbox install
  echo "Running 'go mod download', please wait..."
  go mod download
  echo "Running 'just docker up', please wait..."
  just docker up
  echo "Running 'just docker exec go mod tidy', please wait..."
  just docker exec go mod tidy
  echo "Done"
