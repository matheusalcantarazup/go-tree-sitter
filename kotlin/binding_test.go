package kotlin_test

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/kotlin"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte(`
fun main() {
    println("Hello, world!!!")
}
	`), kotlin.GetLanguage())
	assert.Equal("(source_file (function_declaration (simple_identifier) (function_body (statements (call_expression (simple_identifier) (call_suffix (value_arguments (value_argument (line_string_literal)))))))))",
		n.String(),
	)
}
