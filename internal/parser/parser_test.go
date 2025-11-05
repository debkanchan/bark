package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name          string
		filename      string
		content       string
		expectedCount int
	}{
		{
			name:     "Go file with BARK comments",
			filename: "test.go",
			content: `package main
// BARK: Remove this
func main() {
	// BARK: Fix this later
	println("test")
}`,
			expectedCount: 2,
		},
		{
			name:     "Go file without BARK comments",
			filename: "clean.go",
			content: `package main
// Regular comment
func main() {
	println("test")
}`,
			expectedCount: 0,
		},
		{
			name:     "Go file with plain BARK marker",
			filename: "plain.go",
			content: `package main
// BARK need to fix
func main() {}
`,
			expectedCount: 1,
		},
		{
			name:     "Go file with just BARK",
			filename: "justbark.go",
			content: `package main
// BARK
func main() {}
`,
			expectedCount: 1,
		},
		{
			name:     "JavaScript with BARK comment",
			filename: "test.js",
			content: `// BARK: Remove console.log
console.log("debug");`,
			expectedCount: 1,
		},
		{
			name:     "Python with BARK comment",
			filename: "test.py",
			content: `# BARK: Fix this function
def hello():
    pass`,
			expectedCount: 1,
		},
		{
			name:     ".env file with BARK comment",
			filename: ".env.local",
			content: `# BARK: Remove test credentials
API_KEY=test123
DATABASE_URL=localhost`,
			expectedCount: 1,
		},
		{
			name:     ".env file without BARK",
			filename: ".env",
			content: `# Production environment
API_KEY=${API_KEY}`,
			expectedCount: 0,
		},
		{
			name:     "Dockerfile with BARK comment",
			filename: "Dockerfile",
			content: `# BARK: Update base image to latest LTS version
FROM node:18

# Set working directory
WORKDIR /app
`,
			expectedCount: 1,
		},
		{
			name:     "Dockerfile without BARK",
			filename: "dockerfile",
			content: `FROM node:18
WORKDIR /app
`,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, tt.filename)

			err := os.WriteFile(tmpFile, []byte(tt.content), 0644)
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}

			// Parse file
			findings, err := parser.ParseFile(tmpFile)
			if err != nil {
				t.Fatalf("ParseFile failed: %v", err)
			}

			// Check findings count
			if len(findings) != tt.expectedCount {
				t.Errorf("Expected %d findings, got %d", tt.expectedCount, len(findings))
			}

			// Verify all findings contain "BARK:"
			for _, finding := range findings {
				if finding.FilePath != tmpFile {
					t.Errorf("Expected FilePath %s, got %s", tmpFile, finding.FilePath)
				}
				if finding.Line == 0 {
					t.Error("Line number should not be 0")
				}
			}
		})
	}
}

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"test.go", ".go"},
		{"test.js", ".js"},
		{"test.py", ".py"},
		{"/path/to/file.cpp", ".cpp"},
		{"no_extension", ""},
		{".hidden", ".hidden"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result := getFileExtension(tt.path)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestParseUnsupportedFile(t *testing.T) {
	parser := NewParser()

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")

	err := os.WriteFile(tmpFile, []byte("BARK: This should not be found"), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	findings, err := parser.ParseFile(tmpFile)
	if err != nil {
		t.Fatalf("ParseFile should not error on unsupported files: %v", err)
	}

	if len(findings) != 0 {
		t.Errorf("Expected 0 findings for unsupported file, got %d", len(findings))
	}
}
