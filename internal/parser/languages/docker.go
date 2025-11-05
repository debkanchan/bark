package languages

import (
	tree_sitter "github.com/debkanchan/tree-sitter-dockerfile/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func Docker() Language {
	return Language{
		Name:             "Docker",
		Extensions:       []string{".dockerfile", ".Dockerfile"},
		FilenamePatterns: []string{`^[Dd]ockerfile$`}, // Matches Dockerfile or dockerfile
		Parser:           sitter.NewLanguage(tree_sitter.Language()),
		Query:            "((comment) @comment)",
	}
}
