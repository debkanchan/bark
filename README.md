# 🐕 Bark! - Save yourself from embarrassments

**Ever changed something temporarily to test something locally that should never be in production/in your PR but forgot to revert the change before pushing? Me too.**

Bark is an **"embarrassment linter"** that detects `BARK` comments in your code. Add `BARK` comments to temporary code and Bark! will stop you from pushing it to version control.

## Table of Contents

- [🐕 Bark! - Save yourself from embarrassments](#-bark---save-yourself-from-embarrassments)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
    - [Homebrew (macOS / Linux)](#homebrew-macos--linux)
    - [npm (any platform with Node.js)](#npm-any-platform-with-nodejs)
    - [Prebuilt binary (manual)](#prebuilt-binary-manual)
    - [From source](#from-source)
  - [Usage](#usage)
    - [Recommended: Install Git Hook (Set and Forget!)](#recommended-install-git-hook-set-and-forget)
    - [Add BARK Comments to Your Code](#add-bark-comments-to-your-code)
    - [Manual Scanning (CLI)](#manual-scanning-cli)
    - [Output Formats](#output-formats)
    - [Git Hook Commands](#git-hook-commands)
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
- 🪝 **Git Hooks**: Automatic git hook installation; pre-commit hook works with GitHub Desktop

## Installation

Pick whichever fits your setup — both install a prebuilt binary, no compiler required.

### Homebrew (macOS / Linux)

```bash
brew install debkanchan/tap/bark
```

### npm (any platform with Node.js)

```bash
npm install -g @debkanchan/bark
```

The npm package ships a tiny launcher plus a per-platform binary selected
automatically (linux/macOS x64 & arm64, Windows x64).

### Prebuilt binary (manual)

Download the archive for your OS/arch from the
[latest release](https://github.com/debkanchan/bark/releases/latest), extract it,
and move the binary onto your `PATH`.

### From source

> **Note:** building from source requires a C compiler (gcc/clang). `bark` uses
> tree-sitter via CGO, so `CGO_ENABLED=1` and a C toolchain are needed. The
> Homebrew and npm installs above avoid this entirely.

```bash
go install github.com/debkanchan/bark/cmd/bark@latest
```

Or build a local checkout:

```bash
git clone https://github.com/debkanchan/bark.git
cd bark
go build -o bark ./cmd/bark
```

## Usage

### Recommended: Install Git Hook (Set and Forget!)

The easiest way to use Bark is to install it as a git hook. This automatically prevents you from committing or pushing code with BARK comments:

```bash
# Install bark
go install github.com/debkanchan/bark/cmd/bark@latest

# Install the hook (default: pre-commit)
bark git-hook install

# Or install the pre-push hook instead
bark git-hook install pre-push
```

**That's it!** Bark runs automatically and blocks the operation if BARK comments are found.

- **pre-commit** (default): catches issues at commit time, works with all Git clients
- **pre-push**: catches issues at push time, allows local commits with BARK comments for WIP

> **Note:** GitHub Desktop does not trigger pre-push hooks. Use the default pre-commit hook if you use GitHub Desktop.

**You can install both hooks** for double protection!

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

**Scan one or more paths:**

```bash
bark ./src
bark ./src ./lib ./tests
```

**Using flag syntax (single path only):**

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

```bash
bark git-hook install [pre-commit|pre-push]   # default: pre-commit
bark git-hook uninstall [pre-commit|pre-push] # default: pre-commit
```

Each install will:

- ✅ Create or update the hook file in `.git/hooks/`
- ✅ Safely merge with existing hooks using markers
- ✅ Back up any existing hook before modification

Bark uses markers (`# BEGIN bark hook` / `# END bark hook`) to identify its section, so it coexists safely with other hooks. Both hooks can be installed simultaneously for double protection!

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
| `path`             | Path(s) to scan — space-separated for multiple | `.`      | No       |
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
