# Run the built binary.
[group('project')]
repo *args='':
  just docker cli {{args}}
