uninstall:
  #!/usr/bin/env bash
  echo "Uninstalling, please wait..."

  rm -rf {{AQUA_ROOT_DIR}}
  docker compose kill

  out=$(docker compose kill 2>&1)

  if [[ "$out" == "no container to kill" ]]; then
    echo ""
  fi

  docker compose down --remove-orphans
