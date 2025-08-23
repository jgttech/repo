remove:
  #!/usr/bin/env bash
  pkg=$(cat aqua.yml | yq ".packages[].name" | fzf)
  yq eval -i "del(.packages[] | select(.name == \"$pkg\"))" aqua.yml

  if [ -z "$pkg" ]; then
    echo "No package selected. Cancelled."
    exit 0
  fi

  just refresh
