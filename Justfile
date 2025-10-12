set quiet
set export

import 'bin/just/add.just'
import 'bin/just/clean.just'
import 'bin/just/install.just'
import 'bin/just/refresh.just'
import 'bin/just/repo.just'
import 'bin/just/ssh.just'
import 'bin/just/start.just'

default: install
