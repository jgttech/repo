# Use this to easily invoke the README.md files from
# the docs/prompts the repo uses.
[private]
[group('claude')]
prompt directory context="Run the prompt against the entire repo.":
  #!/usr/bin/env bash
  prompt="{{context}} "
  prompt+="@docs/prompts/{{directory}}/README.md "
  prompt+="Do NOT add AI signatures or co-author attributions."

  claude -p "$prompt" --allowedTools "Bash,Read,Edit,Write"
