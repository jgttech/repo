[private]
app *args='':
  #!/usr/bin/env bash
  docker compose exec repo bash -c '
    cd $(eval echo "$REPO_HOME")
    go run main.go {{args}}
  '
