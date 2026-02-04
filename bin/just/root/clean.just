# Clean up repo and/or packages.
[group('project')]
clean:
  #!/usr/bin/env bash
  set -euo pipefail
  echo "Cleaning, please wait..."

  # Clear the nix store
  devbox run -- nix store gc --extra-experimental-features nix-command

  # Delete the .devbox dir
  [[ -d ".devbox" ]] && rm -rf ".devbox"

  # Remove the development container
  just docker down

  echo "Done"
  exit 0
