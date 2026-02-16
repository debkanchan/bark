package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-kotlin/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func Kotlin() Language {
	return Language{
		Name:       "Kotlin",
		Extensions: []string{".kt", ".kts"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "([(line_comment) (block_comment)] @comment)",
	}
}
