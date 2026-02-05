# Run the built binary.
[group('project')]
sync *args='':
  echo "Taking down container, please wait..."
  just docker down
  echo "Starting new container, please wait..."
  just docker up
  echo "Running 'just docker exec go mod tidy', please wait..."
  just docker exec go mod tidy
