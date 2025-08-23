install:
  echo "Installing, please wait..."
  aqua install
  lefthook install
  go mod download
  go mod tidy
  just runtime
  just sync
  just app
