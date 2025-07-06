# Repo CLI

This is a CLI uility that adds a thin opinionated layer around the `git` and `gh` interfaces. The goal of this is to create a common interface that allows you to manage multiple users across different repositories with different tools, through a single CLI interface.

# Commands

These commands are what you can do with the `repo` CLI.

---

```bash
repo install
```

> Performs a system installation. This does all the initial setup work to use and maange the `repo` CLI. This gets automatically run by the installer if you installed `repo` using the copy-and-paste command from these docs.

---

```bash
repo add <optional_path_to_git_repo> --alias <name>
```

> This registers a repo with the CLI. Once registered, you can optionally assign it an `--alias` when adding it if you want to keep the repo under another name that is different than the name of the directory.

**_Args:_**

This command can optionally accept a path. The path passed in is checked to see if it is a `git` repo. If it is, it will be registered with the `repo` CLI.

**_Flags:_**

- `--alias -a`: Allows assigning another name to this repo as a reference within the `repo` CLI. The name (or alias) of registered repos has an effect on using how you can refer to that given repo within the `repo` CLI.
