package testing

import (
	"os"
	"strings"
	"testing"

	"github.com/pittfit/ortho/ast"
	"github.com/pittfit/ortho/lexer"
	"github.com/pittfit/ortho/parser"
	"github.com/pittfit/ortho/token"
	"github.com/pittfit/ortho/tracing"
	"github.com/stretchr/testify/assert"
)

var nodeFormatter = ast.NewNodeFormatter("  ")

func Test(t *testing.T) {
	fixtures, err := loadFixtures()
	if err != nil {
		t.Error(err)
	}

	for _, fx := range fixtures {
		t.Run(fx.Input, func(t *testing.T) {
			// Test the lexer
			l := lexer.NewLexer([]byte(fx.Input))
			assert.Equal(t, fx.Tokens, tokenTypes(l.All()), "Unexpected tokens")

			// Test the AST
			tracing.EnableIf(os.Getenv("TRACE") == "1")
			parsedAst := parser.NewParser([]byte(fx.Input)).Parse()
			tracing.Disable()

			actual := strings.Split(strings.TrimSpace(nodeFormatter.ToString(parsedAst.Root)), "\n")
			expected := strings.Split(strings.TrimSpace(fx.Ast), "\n")

			assert.Equal(t, expected, actual, "Unexpected generated AST")

			for _, output := range fx.Output {
				// Verify regex
				actualPattern, err := parsedAst.ToPattern(ast.PatternOpts{Anchored: false})
				if err != nil {
					t.Error(err)
				}
				expectedPattern := output.Pattern
				assert.Equal(t, expectedPattern, actualPattern, "Unexpected regex")

				// Verify produced strings
				actualStrings, err := parsedAst.ToStrings()
				if err != nil {
					t.Error(err)
				}
				expectedStrings := output.Strings
				assert.Equal(t, expectedStrings, actualStrings, "Unexpected strings")
			}
		})
	}
}

func Benchmark(b *testing.B) {
	fixtures, err := loadFixtures()
	if err != nil {
		b.Error(err)
	}

	for _, fx := range fixtures {
		// Benchmark the lexer
		b.Run(fx.Input+"/Lexer", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lexer.NewLexer([]byte(fx.Input)).All()
			}
		})

		// Benchmark the parser
		b.Run(fx.Input+"/Parser", func(b *testing.B) {
			// Test the AST
			for i := 0; i < b.N; i++ {
				parser.NewParser([]byte(fx.Input)).Parse()
			}
		})

		// Pattern production
		b.Run(fx.Input+"/Pattern", func(b *testing.B) {
			parsedAst := parser.NewParser([]byte(fx.Input)).Parse()

			for range fx.Output {
				for i := 0; i < b.N; i++ {
					parsedAst.ToPattern(ast.PatternOpts{Anchored: false})
				}
			}
		})

		// Strings production
		b.Run(fx.Input+"/Strings", func(b *testing.B) {
			parsedAst := parser.NewParser([]byte(fx.Input)).Parse()

			for range fx.Output {
				for i := 0; i < b.N; i++ {
					parsedAst.ToStrings()
				}
			}
		})
	}
}

func tokenTypes(toks []token.Token) []string {
	types := []string{}
	for _, tok := range toks {
		types = append(types, tok.Type.String())
	}
	return types
}
