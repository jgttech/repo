# Repo CLI, A "Git Account Multiplexer"

This CLI utility is what I like to call a "Git Account Mulitplexer". The idea is that anywhere a `git` configuration is used this tool can sit on top of that and manage the account contexts between an endless number of `git` configurations and their, respective, SSH keys for each account from a centralized and state managed solution that removes the need to remember how to maintain multiple sets of configurations and SSH keys. It also enabled portability of the configurations built up for this CLI that can be committed to your own dotfiles or otherwise exported and imports within this CLI tool.

> The was inspired off of the GitHub `gh` CLI.

# WIP Commands

These commands are still a work in progress but there is an embedded Obsidian docs within this repo, please feel free to reivew it, if need be. There I can go over the technical considerations of the tool and its use cases in greater depth. Here is just an overview of its use cases.

---

> This is run for you, automatically.

```bash
repo install
```

---

> Uninstalls everything from your system, except the CLI itself.

```bash
repo uninstall
```

---
