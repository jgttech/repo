# Repo CLI

This is a CLI uility that adds a thin opinionated layer around the `git` and `gh` interfaces. The goal of this is to create a common interface that allows you to manage multiple users across different repositories with different tools, through a single CLI interface.

# Install

Installs the CLI utility on your system.

```bash
<clone> | bash
```

# Uninstall

Completely removes (including configuration and data) everything from your system, purging it 100%.

```bash
repo purge
```

# Adding an Existing Repo

Running this from within a repo you aready have will register it with the CLI.

```bash
repo register
```

# Cloning a Repo

```bash
repo clone <user>/<repo> <optional_directory>
```
