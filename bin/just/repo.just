repo *args='':
  docker compose exec repo bash -c "cd /tmp/repo; go run main.go {{args}}"
