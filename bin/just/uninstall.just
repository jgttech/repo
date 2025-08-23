uninstall:
  #!/usr/bin/env bash
  echo "Uninstalling, please wait..."

  rm -rf {{AQUA_ROOT_DIR}}
  docker compose kill

  out=$(docker compose kill 2>&1)

  # If output doesn't end with newline and matches the message
  if [[ "$output" == "no container to kill" && "$output" != *$'\n' ]]; then
    echo ""
  fi

  docker compose down --remove-orphans
