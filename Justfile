# Settings
set quiet
set export
set dotenv-load

# Enviroment
AQUA_ROOT_DIR := ".aqua"
PATH := ".aqua/bin:" + env_var("PATH")
REPO_TMP := "/tmp/code"
REPO_HOME := "$HOME/.code"

# Commands
import 'bin/just/app.just'
import 'bin/just/sync.just'
import 'bin/just/runtime.just'
import 'bin/just/install.just'
import 'bin/just/uninstall.just'
import 'bin/just/update.just'
import 'bin/just/refresh.just'
import 'bin/just/add.just'
import 'bin/just/remove.just'

default: install

reinstall: uninstall install

connect:
  docker compose exec repo bash

exec *args='':
  docker compose exec repo bash -c "{{args}}"

cli *args='':
  just sync
  just app {{args}}
