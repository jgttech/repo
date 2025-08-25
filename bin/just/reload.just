reload:
  #!/usr/bin/env bash
  echo "Reloading docker container, please wait..."

  docker compose kill
  out=$(docker compose kill 2>&1)

  # If output doesn't end with newline and matches the message
  if [[ "$output" == "no container to kill" && "$output" != *$'\n' ]]; then
    echo ""
  fi

  docker compose down --remove-orphans
  just runtime
  just sync

  just exec "cd \$(echo \"\$REPO_HOME\"); go mod tidy"
