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
import 'bin/just/add.just'
import 'bin/just/app.just'
import 'bin/just/cli.just'
import 'bin/just/connect.just'
import 'bin/just/exec.just'
import 'bin/just/install.just'
import 'bin/just/ps.just'
import 'bin/just/refresh.just'
import 'bin/just/reinstall.just'
import 'bin/just/remove.just'
import 'bin/just/runtime.just'
import 'bin/just/sync.just'
import 'bin/just/uninstall.just'
import 'bin/just/update.just'

default: install
