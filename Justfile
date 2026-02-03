mod docker "bin/just/docker/.mod.just"
mod claude "bin/just/claude/.mod.just"

import "bin/just/root/.mod.just"

default:
  just -l
  echo ""
  just -l docker
  echo ""
  just -l claude
