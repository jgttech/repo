# Run the built binary.
[group('project')]
sync *args='':
  just docker down
  just docker up
