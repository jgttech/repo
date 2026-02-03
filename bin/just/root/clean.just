# Clean up repo and/or packages.
[group('project')]
clean:
  echo "Cleaning, please wait..."
  [[ -d ".devbox" ]] && rm -rf ".devbox"
  devbox run -- nix store gc --extra-experimental-features nix-command
  echo "Done"
