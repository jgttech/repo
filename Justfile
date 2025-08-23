set export
set dotenv-load

AQUA_ROOT_DIR := ".aqua"
PATH := ".aqua/bin:" + env_var("PATH")
REPO_TMP := "/tmp/code"
REPO_HOME := "$HOME/.code"

alias i := install
alias u := uninstall
alias up := update
alias a := add
alias r := remove
alias c := connect
alias x := exec

default: install

@_cli *args='':
  docker compose exec repo bash -c 'cd $(eval echo "$REPO_HOME"); go run main.go {{args}}'

[working-directory: 'bin/docker']
_docker *args='':
  -bash {{args}}

_runtime:
  #!/usr/bin/env bash
  set -e
  STATE="$(docker compose ps --format=json repo | jq ".State" | tr -d '\"')"

  if [[ "$STATE" != "running" ]]; then
    docker compose up repo -d
    echo ""
  fi

_sync:
  #!/usr/bin/env bash
  docker compose exec repo bash -c '
    vol="$REPO_TMP"
    base=$(eval echo "$REPO_HOME")

    # What we want to copy from the volume.
    sources=(\
      "bin" \
      "cmds" \
      "core" \
      "test" \
      "tui" \
      "go.mod" \
      "go.sum" \
      "main.go" \
    )

    # Single removal and creation
    [[ -d "$base" ]] && rm -rf "$base"; mkdir -p "$base"

    # Batch copy with fewer syscalls
    (cd "$vol" && tar cf - "${sources[@]}" 2>/dev/null) | (cd "$base" && tar xf -)
  '

@install:
  echo "Installing, please wait..."
  aqua install
  lefthook install
  go mod download
  go mod tidy
  just _runtime
  just _sync
  just _cli

@uninstall:
  echo "Uninstalling, please wait..."
  rm -rf {{AQUA_ROOT_DIR}}
  docker compose kill
  docker compose down --remove-orphans
  echo ""

update:
  @echo "Updating, please wait..."
  -aqua update
  -rm -rf {{AQUA_ROOT_DIR}}
  -aqua install

refresh:
  -rm -rf {{AQUA_ROOT_DIR}}
  aqua install

reinstall: uninstall install

add:
  aqua generate -i
  aqua install

remove:
  @echo "WIP"

connect:
  docker compose exec repo bash

@exec *args='':
  -docker compose exec repo bash -c "{{args}}"

@cli *args='':
  just _sync
  just _cli {{args}}
