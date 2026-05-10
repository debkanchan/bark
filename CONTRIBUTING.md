# Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Contribution Ideas

- Add support for more programming languages
- Improve error messages
- Add configuration file support
- Enhance test coverage
- Improve documentation

## Architecture

Bark follows standard Go project layout with a modular architecture:

```text
bark/
├── cmd/bark/
│   ├── main.go            # CLI entry point, flag parsing, scan logic
│   └── git-hooks.go       # Git hook install/uninstall logic and hook scripts
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
- **Git Hooks** (`git-hooks.go`): Hook scripts, install/uninstall logic with marker-based merging
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
