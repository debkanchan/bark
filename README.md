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
  - [Architecture](#architecture)
    - [Key Components](#key-components)
  - [Development](#development)
    - [Running Tests](#running-tests)
    - [Testing Git Hooks Locally](#testing-git-hooks-locally)
    - [Adding a New Language](#adding-a-new-language)
    - [Build Commands](#build-commands)
  - [Contributing](#contributing)
    - [Contribution Ideas](#contribution-ideas)
  - [License](#license)
  - [Why "Bark"?](#why-bark)

## Features

- 🌍 **Cross-platform**: Works on Windows, macOS, and Linux
- ⚡ **Fast**: Concurrent file processing using goroutines
- 🌳 **Tree-sitter powered**: Accurate parsing using tree-sitter grammars
- 🔧 **Modular architecture**: Separated core logic for easy integration
- 📋 **Multiple output formats**: Text for CLI, JSON for CI/CD pipelines
- 🎯 **Wide language support**: 18 languages including Go, JavaScript, TypeScript, Python, Java, Kotlin, C, C++, Bash, Rust, Zig, Lua, HCL, YAML, Docker, XML, TOML, JSON
- 🎬 **GitHub Action**: One-line integration for your CI/CD pipeline
- 🪝 **Git Hooks**: Automatic pre-push hook installation

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

The easiest way to use Bark is to install it as a git pre-push hook. This automatically prevents you from pushing code with BARK comments:

```bash
# Install bark
go install github.com/debkanchan/bark/cmd/bark@latest

# Install the git hook (one-time setup)
bark git-hook install
```

**That's it!** Now bark runs automatically before every `git push` and blocks the push if BARK comments are found.

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

Bark can automatically install a git pre-push hook to prevent BARK comments from being pushed:

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

This will:
- ✅ Remove only the bark section (preserves other hooks)
- ✅ Back up the hook before modification
- ✅ Delete the file if empty after removal

**How it works:**

Bark uses markers (`# BEGIN bark hook` / `# END bark hook`) to identify its section, allowing it to coexist with other git hooks safely.

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
    path: './src'
```

**JSON output:**

```yaml
- uses: debkanchan/bark@v1
  with:
    format: 'json'
```

**Report only (don't fail the build):**

```yaml
- uses: debkanchan/bark@v1
  with:
    fail-on-findings: 'false'
```

**Specific version:**

```yaml
- uses: debkanchan/bark@v1
  with:
    version: 'v1.0.0'
```

### Action Inputs

| Input | Description | Default | Required |
|-------|-------------|---------|----------|
| `path` | Path to scan for BARK comments | `.` | No |
| `format` | Output format (`text` or `json`) | `text` | No |
| `fail-on-findings` | Fail the build if BARK comments found | `true` | No |
| `version` | Bark version to install (`latest` or `v1.0.0`) | `latest` | No |

### Examples

**Complete workflow with multiple jobs:**

```yaml
name: Code Quality

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  bark-check:
    name: Check for BARK Comments
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: debkanchan/bark@v1
        with:
          path: '.'
          format: 'text'
```

**Matrix strategy - scan multiple directories:**

```yaml
jobs:
  bark-matrix:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        directory: ['./frontend', './backend', './shared']
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
| ---------- | -----------------------------------------------------------|
| Go         | `.go`                                                      |
| JavaScript | `.js`, `.jsx`, `.mjs`, `.cjs`                              |
| TypeScript | `.ts`, `.tsx`                                              |
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

## Architecture

Bark follows standard Go project layout with a modular architecture:

```text
bark/
├── cmd/bark/              # CLI entry point with git hook commands
├── internal/
│   ├── parser/            # Tree-sitter integration
│   │   ├── registry.go    # Language registry
│   │   ├── parser.go      # Comment extraction
│   │   └── languages/     # Individual language configs
│   ├── scanner/           # Concurrent file scanner with worker pool
│   └── results/           # Result types and formatters (text, JSON)
├── action.yml             # GitHub Action definition (composite)
└── .github/workflows/     # CI/CD workflows
```

### Key Components

- **Language Registry**: Extensible registry mapping file extensions to tree-sitter parsers
- **Parser**: Uses tree-sitter queries to extract comments from source files
- **Scanner**: Concurrent file processing with worker pool pattern
- **Formatters**: Interface-based output formatting (text, JSON)
- **Git Hooks**: Smart installation with marker-based merging
- **GitHub Action**: Composite action using preinstalled Go (no Docker!)

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package
go test ./internal/parser
go test ./internal/scanner
go test ./internal/results

# Using Makefile
make test
make test-coverage
```

### Testing Git Hooks Locally

```bash
# Install the hook in your bark repository
bark git-hook install

# Create a test commit with BARK comments
echo "// BARK test" >> test.go
git add test.go
git commit -m "test"

# Try to push (should be blocked)
git push

# Uninstall when done testing
bark git-hook uninstall
```

### Adding a New Language

1. **Install the tree-sitter parser binding:**

   ```bash
   go get github.com/tree-sitter/tree-sitter-{language}/bindings/go
   ```

2. **Create a new file** `internal/parser/languages/{language}.go`:

   ```go
   package languages

   import (
       sitter "github.com/tree-sitter/go-tree-sitter"
       tree_sitter "github.com/tree-sitter/tree-sitter-{language}/bindings/go"
   )

   func YourLanguage() Language {
       return Language{
           Name:       "YourLanguage",
           Extensions: []string{".ext"},
           Parser:     sitter.NewLanguage(tree_sitter.Language()),
           Query:      "((comment) @comment)",
       }
   }
   ```

3. **Add to the registry** in `internal/parser/registry.go`:

   ```go
   languageList := []languages.Language{
       // ... existing languages
       languages.YourLanguage(),
   }
   ```

4. **Test it:**

   ```bash
   go build -o bark ./cmd/bark
   ./bark path/to/file.ext
   ```

### Build Commands

Using the Makefile:

```bash
make build              # Build the binary
make test               # Run all tests
make test-coverage      # Run tests with coverage report
make run-testdata       # Test on sample files
make install            # Install globally
make clean              # Clean build artifacts
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Contribution Ideas

- Add support for more programming languages
- Improve error messages
- Add configuration file support
- Enhance test coverage
- Improve documentation

## License

GNU Affero General Public License - see [LICENSE](LICENSE) file for details

## Why "Bark"?

Like a faithful dog that barks to alert you, Bark helps you catch those temporary comments and debug code before they make it into your repository! 🐕

---

**Made with ❤️ to prevent embarrassments**
