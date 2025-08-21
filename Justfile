set export
set dotenv-load

# This is where system dependencies will go.
AQUA_ROOT_DIR := ".aqua"

# Augment PATH to include the managed
# system dependencies.
PATH := ".aqua/bin:" + env_var("PATH")

# Development environment variables
REPO_TMP := "/tmp/code"
REPO_HOME := "$HOME/.code"

default: install

# Ensures that the container is running.
run:
  -bash bin/docker/run

# Copies the code, in the container, to
# a location it can be run, from the volume.
@sync:
  -docker compose exec repo bash -c "bash $REPO_TMP/bin/docker/sync"

# All the setup work required to get things working.
install:
  @echo "Installing, please wait..."
  aqua install
  lefthook install
  go mod tidy
  just run
  just sync
  docker compose exec repo bash -c "bash $REPO_TMP/bin/docker/cli"

# Updates the aqua CLI packages
update:
  @echo "Updating system packages, please wait..."
  aqua update

@clean:
  -rm -rf {{AQUA_ROOT_DIR}}
  -docker compose kill
  -docker compose down --remove-orphans

connect:
  docker compose exec repo bash

@cli *args='':
  just sync
  -docker compose exec repo bash -c "bash $REPO_TMP/bin/docker/cli {{args}}"
