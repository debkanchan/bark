# 🐕 Bark! - Save yourself from embarrassments

**Ever changed something temporarily to test something locally that should never be in production/in your PR but forgot to revert the change before pushing? Me too.**

Bark is an **"embarrassment linter"** that detects `BARK` comments in your code. Add `BARK` comments to temporary code and Bark! will stop you from pushing it to version control.

## Table of Contents

- [🐕 Bark! - Save yourself from embarrassments](#-bark---save-yourself-from-embarrassments)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
    - [Install from Source](#install-from-source)
    - [Build Locally](#build-locally)
  - [Usage](#usage)
    - [Recommended: Install Git Hook (Set and Forget!)](#recommended-install-git-hook-set-and-forget)
      - [For GitHub Desktop Users (Pre-Commit Hook)](#for-github-desktop-users-pre-commit-hook)
      - [For CLI Git Users (Pre-Push Hook)](#for-cli-git-users-pre-push-hook)
      - [Why Two Options?](#why-two-options)
    - [Add BARK Comments to Your Code](#add-bark-comments-to-your-code)
    - [Manual Scanning (CLI)](#manual-scanning-cli)
    - [Output Formats](#output-formats)
    - [Git Hook Commands](#git-hook-commands)
      - [Pre-Commit Hook (Recommended for GitHub Desktop)](#pre-commit-hook-recommended-for-github-desktop)
      - [Pre-Push Hook (For CLI Git Users)](#pre-push-hook-for-cli-git-users)
      - [How it works](#how-it-works)
    - [Exit Codes](#exit-codes)
  - [GitHub Action](#github-action)
    - [Basic Usage](#basic-usage)
    - [Configuration Options](#configuration-options)
    - [Action Inputs](#action-inputs)
    - [Examples](#examples)
  - [Supported Languages](#supported-languages)
  - [Contributing](#contributing)
  - [Why "Bark"?](#why-bark)

## Features

- 🌍 **Cross-platform**: Works on Windows, macOS, and Linux
- ⚡ **Fast**: Concurrent file processing using goroutines
- 🌳 **Tree-sitter powered**: Accurate parsing using tree-sitter grammars
- 🔧 **Modular architecture**: Separated core logic for easy integration
- 📋 **Multiple output formats**: Text for CLI, JSON for CI/CD pipelines
- 🎯**Wide language support**: 19 languages including Go, JavaScript, TypeScript, PHP, Python, Java, Kotlin, C, C++, Bash, Rust, Zig, Lua, HCL, YAML, Docker, XML, TOML, JSON
- 🎬 **GitHub Action**: One-line integration for your CI/CD pipeline
- 🪝 **Git Hooks**: Automatic pre-push and pre-commit hook installation (works with GitHub Desktop!)

## Installation

### Install from Source

```bash
go install github.com/debkanchan/bark/cmd/bark@latest
```

### Build Locally

```bash
git clone https://github.com/debkanchan/bark.git
cd bark
go build -o bark ./cmd/bark
```

## Usage

### Recommended: Install Git Hook (Set and Forget!)

The easiest way to use Bark is to install it as a git hook. This automatically prevents you from committing or pushing code with BARK comments:

#### For GitHub Desktop Users (Pre-Commit Hook)

If you use GitHub Desktop or want to catch BARK comments at commit time:

```bash
# Install bark
go install github.com/debkanchan/bark/cmd/bark@latest

# Install the pre-commit hook (works with GitHub Desktop!)
bark git-hook install-commit
```

**That's it!** Now bark runs automatically before every commit and blocks commits if BARK comments are found. This works with GitHub Desktop, VS Code, and all other Git clients!

#### For CLI Git Users (Pre-Push Hook)

If you primarily use command-line Git and prefer to catch BARK comments at push time:

```bash
# Install bark
go install github.com/debkanchan/bark/cmd/bark@latest

# Install the pre-push hook (one-time setup)
bark git-hook install
```

**That's it!** Now bark runs automatically before every `git push` and blocks the push if BARK comments are found.

#### Why Two Options?

- **Pre-commit hook** (`install-commit`): Catches issues earlier (at commit time). Works with GitHub Desktop and all Git clients.
- **Pre-push hook** (`install`): Only blocks at push time. Allows you to make local commits with BARK comments for work-in-progress.

**You can install both hooks** if you want double protection!

### Add BARK Comments to Your Code

Use BARK comments as reminders for things that need to be fixed before pushing:

```go
package main

import "fmt"

// BARK: Remove debug code before commit
func main() {
    fmt.Println("Debug mode enabled")
    // BARK: Replace with proper configuration
    apiKey := "test-key-123"
}
```

When you try to push:

```bash
$ git push
🐕 Running bark to check for BARK comments...
Found 2 BARK comment(s):

main.go:4:1: // BARK: Remove debug code before commit
main.go:7:5: // BARK: Replace with proper configuration

❌ Push blocked: BARK comments found
Please remove BARK comments before pushing
```

Fix the issues, and push successfully! ✅

### Manual Scanning (CLI)

If you need to manually scan your code without git hooks:

**Scan current directory:**

```bash
bark
# or explicitly
bark .
```

**Scan specific path:**

```bash
bark ./src
bark ./path/to/code
```

**Using flag syntax (alternative):**

```bash
bark -path ./src
bark -p ./src
```

### Output Formats

**Text format (default)** - Human-readable output:

```bash
bark ./src
bark -format text .
```

**JSON format** - For CI/CD integration and parsing:

```bash
bark -format json .
bark -f json ./testdata
```

JSON output example:

```json
{
  "findings": [
    {
      "file_path": "main.go",
      "line": 4,
      "column": 1,
      "comment": "// BARK: Remove debug code"
    }
  ],
  "count": 1
}
```

### Git Hook Commands

Bark supports both pre-commit and pre-push hooks:

#### Pre-Commit Hook (Recommended for GitHub Desktop)

**Install pre-commit hook:**

```bash
bark git-hook install-commit
```

This will:

- ✅ Create or update `.git/hooks/pre-commit`
- ✅ Safely merge with existing hooks using markers
- ✅ Back up any existing hook before modification
- ✅ Run bark automatically before each commit
- ✅ Block commits if BARK comments are found
- ✅ **Works with GitHub Desktop, VS Code, and all Git clients!**

**Uninstall pre-commit hook:**

```bash
bark git-hook uninstall-commit
```

#### Pre-Push Hook (For CLI Git Users)

**Install pre-push hook:**

```bash
bark git-hook install
```

This will:

- ✅ Create or update `.git/hooks/pre-push`
- ✅ Safely merge with existing hooks using markers
- ✅ Back up any existing hook before modification
- ✅ Run bark automatically before each `git push`
- ✅ Block pushes if BARK comments are found

**Uninstall pre-push hook:**

```bash
bark git-hook uninstall
```

#### How it works

Bark uses markers (`# BEGIN bark hook` / `# END bark hook`) to identify its section, allowing it to coexist with other git hooks safely. Both hooks can be installed simultaneously for double protection!

### Exit Codes

- `0` - No BARK comments found (clean)
- `1` - BARK comments found
- `2` - Error occurred during scanning

## GitHub Action

### Basic Usage

Add Bark to your GitHub Actions workflow with a single line:

```yaml
name: Check for BARK comments

on: [push, pull_request]

jobs:
  bark:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: debkanchan/bark@v1
```

### Configuration Options

**Scan specific directory:**

```yaml
- uses: debkanchan/bark@v1
  with:
    path: "./src"
```

**JSON output:**

```yaml
- uses: debkanchan/bark@v1
  with:
    format: "json"
```

**Report only (don't fail the build):**

```yaml
- uses: debkanchan/bark@v1
  with:
    fail-on-findings: "false"
```

**Specific version:**

```yaml
- uses: debkanchan/bark@v1
  with:
    version: "v1.0.0"
```

### Action Inputs

| Input              | Description                                    | Default  | Required |
| ------------------ | ---------------------------------------------- | -------- | -------- |
| `path`             | Path to scan for BARK comments                 | `.`      | No       |
| `format`           | Output format (`text` or `json`)               | `text`   | No       |
| `fail-on-findings` | Fail the build if BARK comments found          | `true`   | No       |
| `version`          | Bark version to install (`latest` or `v1.0.0`) | `latest` | No       |

### Examples

**Complete workflow with multiple jobs:**

```yaml
name: Code Quality

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  bark-check:
    name: Check for BARK Comments
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: debkanchan/bark@v1
        with:
          path: "."
          format: "text"
```

**Matrix strategy - scan multiple directories:**

```yaml
jobs:
  bark-matrix:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        directory: ["./frontend", "./backend", "./shared"]
    steps:
      - uses: actions/checkout@v4
      - uses: debkanchan/bark@v1
        with:
          path: ${{ matrix.directory }}
```

**Manual installation (if you need more control):**

```yaml
steps:
  - uses: actions/checkout@v4
  - uses: actions/setup-go@v5
    with:
      go-version: "1.21"
  - name: Install Bark
    run: go install github.com/debkanchan/bark/cmd/bark@latest
  - name: Run Bark
    run: bark -format json .
```

## Supported Languages

| Language   | Extensions                                                 |
| ---------- | ---------------------------------------------------------- |
| Go         | `.go`                                                      |
| JavaScript | `.js`, `.jsx`, `.mjs`, `.cjs`                              |
| TypeScript | `.ts`, `.tsx`                                              |
| PHP        | `.php`, `.phtml`, `.php3`, `.php4`, `.php5`                |
| Python     | `.py`, `.pyw`                                              |
| Java       | `.java`                                                    |
| Kotlin     | `.kt`, `.kts`                                              |
| C          | `.c`, `.h`                                                 |
| C++        | `.cpp`, `.cc`, `.cxx`, `.hpp`, `.hh`, `.hxx`               |
| Bash       | `.sh`, `.bash`, `.env`, `.env.*`                           |
| Rust       | `.rs`                                                      |
| Zig        | `.zig`                                                     |
| Lua        | `.lua`                                                     |
| HCL        | `.hcl`, `.tf`, `.tfvars`                                   |
| YAML       | `.yml`, `.yaml`                                            |
| Docker     | `dockerfile`, `Dockerfile`, `*.dockerfile`, `*.Dockerfile` |
| XML        | `.xml`                                                     |
| TOML       | `.toml`                                                    |
| JSON       | `.json`, `.jsonc`                                          |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## Why "Bark"?

Like a faithful dog that barks to alert you, Bark helps you catch those temporary comments and debug code before they make it into your repository! 🐕

---

**Made with ❤️ to prevent embarrassments**
