update:
  echo "Updating, please wait..."
  aqua update
  rm -rf {{AQUA_ROOT_DIR}}
  aqua install
