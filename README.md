<div align="center">
  <img src=".docs/images/logo.svg#gh-dark-mode-only" alt="Repo CLI" width="200" />
  <img src=".docs/images/logo.svg#gh-light-mode-only" alt="Repo CLI" width="200" />
</div>

# Repo CLI

> [!NOTE]
> Inspired by the GitHub CLI.

The `repo` CLI project is a management tool for git repositories. I use the term **_Git Repository Account Multiplexer_** (GRAM). This is my own term for a repository manager that integrates tools and utilities over existing tools with added configuration, features, and improvemnts to the way repositories are managed using the CLI.

## Development Dependencies

> This section is only for development dependencies.

- [just](https://just.systems/man/en): Command runner for the project.
- [aqua](https://aquaproj.github.io/docs/install#homebrew): Manages all development depenencies.

## Getting Started

Run the `just` command to get everything setup and configured for development.

```bash
just
```

## CLI Commands

You can invoke the CLI commands from within the container using `just`.

```bash
just cli <args>
```

> [!TIP]
> Example command: `just cli version`
