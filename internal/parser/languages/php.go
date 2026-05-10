package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_php "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

func PHP() Language {
	return Language{
		Name:       "PHP",
		Extensions: []string{".php"},
		Parser:     sitter.NewLanguage(tree_sitter_php.LanguagePHP()),
		Query:      "((comment) @comment)",
	}
}