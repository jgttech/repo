clean:
  echo "Deleting, please wait..."
  rm -rf "$PROJECT_DIR"

  echo "Stopping services, please wait..."
  docker compose kill --remove-orphans
  docker compose down --rmi all

  echo "Done"
