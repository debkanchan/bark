package parser

import (
	"path/filepath"
	"regexp"

	"github.com/debkanchan/bark/internal/parser/languages"
)

// Registry holds all supported languages
type Registry struct {
	languages       []languages.Language
	extensionLookup map[string]*languages.Language
}

// NewRegistry creates a new language registry with all supported languages
func NewRegistry() *Registry {
	languageList := []languages.Language{
		languages.Go(),
		languages.JavaScript(),
		languages.TypeScript(),
		languages.Python(),
		languages.Java(),
		languages.C(),
		languages.Cpp(),
		languages.JSON(),
		languages.Bash(),
		languages.Lua(),
		languages.HCL(),
		languages.YAML(),
		languages.XML(),
		languages.TOML(),
		languages.Rust(),
		languages.Zig(),
		languages.Kotlin(),
		languages.Docker(),
	}

	// Build extension lookup map
	extensionLookup := make(map[string]*languages.Language)
	for i := range languageList {
		for _, ext := range languageList[i].Extensions {
			extensionLookup[ext] = &languageList[i]
		}
	}

	return &Registry{
		languages:       languageList,
		extensionLookup: extensionLookup,
	}
}

// GetLanguageByExtension returns the language for a given file extension
func (r *Registry) GetLanguageByExtension(ext string) (*languages.Language, bool) {
	lang, found := r.extensionLookup[ext]
	return lang, found
}

// GetSupportedExtensions returns all supported file extensions
func (r *Registry) GetSupportedExtensions() []string {
	extensions := make([]string, 0, len(r.extensionLookup))
	for ext := range r.extensionLookup {
		extensions = append(extensions, ext)
	}
	return extensions
}

// GetLanguages returns all supported languages
func (r *Registry) GetLanguages() []languages.Language {
	return r.languages
}

// getFileExtension extracts the file extension from a path
func getFileExtension(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[i:]
		}
		if path[i] == '/' || path[i] == '\\' {
			return ""
		}
	}
	return ""
}

// GetLanguageByFilename returns the language for a given file path
// It first tries extension lookup (fast), then falls back to filename pattern matching
func (r *Registry) GetLanguageByFilename(filePath string) (*languages.Language, bool) {
	// Extract extension
	ext := getFileExtension(filePath)

	// Fast path: try extension lookup first (O(1) hashmap)
	if lang, found := r.extensionLookup[ext]; found {
		return lang, true
	}

	// Slow path: check filename patterns using regex (only if no extension match)
	basename := filepath.Base(filePath)
	for i := range r.languages {
		for _, pattern := range r.languages[i].FilenamePatterns {
			// Compile and match regex
			matched, err := regexp.MatchString(pattern, basename)
			if err == nil && matched {
				return &r.languages[i], true
			}
		}
	}

	return nil, false
}
