# gotem - Dotfiles Management CLI

> gotem is a command-line interface (CLI) tool for managing and organizing your dotfiles in a clean and efficient way. It allows you to easily back up, restore, and sync your dotfiles across different systems using a configuration-driven approach.

### Features

- Backup & Restore: Easily back up and restore your dotfiles with configurable options.
- Profile Management: Store and manage different profiles for various setups (e.g., "default", "archlinux", "ubuntu").
- Customizable File Handling: Support for handling individual dotfiles with custom paths for local and stash storage.
- Symlink or Copy: Choose whether to symlink or copy dotfiles to their respective locations.
- Automatic Directory Creation: Automatically create missing directories or files for a seamless workflow.

### Installation

##### Prerequisites

- Go 1.18+: Make sure you have Go installed on your system. You can check your Go version with:

```bash
go version
```

- Git: You will need Git to clone the repository and manage your code.

##### Clone the Repository

```bash
git clone https://github.com/baolhq/gotem.git
cd gotem
```

##### Build the Application

```bash
go build -o gotem .
```

##### Install the Application

Alternatively, you can install the application globally:

```bash
go install
```

After the installation, you can use `gotem` from anywhere in your terminal.

### Configuration

The configuration for gotem is stored in a TOML file, typically named `config.toml`. This file contains settings for backing up dotfiles, specifying directories for file storage, and managing profiles. Here's an example of a basic configuration:

```toml
[stash]  # General configurations
backup = false  # Replace existing files without backup
create = true   # Automatically create missing directories or files
dotpath = "stash"  # Root directory for storing dotfiles
uselink = false  # Whether the program should copy or create symlinks
keepdot = false  # Remove leading '.' when storing in the stash

[stash.default]  # Profile-specific configurations, overwrite general configs
backup = true
dotpath = "default"

[stash.default.nvim]
local_path = "~/.config/nvim"
stash_path = "config/nvim"

[stash.default.zsh]
local_path = "~/.zshrc"
stash_path = "zshrc"
```

Key Configuration Fields:

- [stash]: Contains the general configuration for dotfile management.
  - backup: Set to `true` to enable file backups.
  - create: Set to `true` to create missing directories and files.
  - dotpath: Defines the root directory for storing dotfiles.
  - uselink: Set to `true` to use symlinks for dotfiles, or `false` to copy them.
  - keepdot: Set to `true` to store dotfiles without the leading `.`.

- [stash.default]: Profile-specific settings. Override general configurations here.
  - backup: Enable or disable backup for this profile.
  - dotpath: Define the directory path for this profile’s dotfiles.

- [stash.<profile>]: Each profile (e.g., `default`, `archlinux`, `ubuntu`) can have its own dotfile settings. Each profile can contain settings for individual dotfiles.

Example:

```toml
[stash.archlinux]
dotpath = "archlinux"

[stash.archlinux.nvim]
local_path = "nvim"
stash_path = "nvim"
```

This configuration specifies a profile for `archlinux`, with a file mapping for `nvim`.

### Usage

Once you’ve configured your `config.toml` file, you can use `gotem` to manage your dotfiles. Here are some example commands:

Update dotfiles from local machine into stash:

```bash
gotem update
```

Install dotfiles from stash into local machine:

```bash
gotem install
```

You can also specify a profile when running the commands. For example, to use the "archlinux" profile:

```bash
gotem update --profile archlinux
```

For more advanced options, run:

```bash
gotem --help
```
