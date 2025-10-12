install:
  #!/usr/bin/env bash
  aqua install
  lefthook install
  just start

  for file in "$(pwd)/bin/dev"/*; do
    bin="$(basename $file)"

    [[ ! -e "$PROJECT_BIN/$bin" ]] && \
      ln -s "../../../bin/dev/$bin" "$PROJECT_BIN/$bin"
  done
