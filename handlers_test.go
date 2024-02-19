package parser_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/smikelson75/parser/handlers"
	"github.com/smikelson75/parser/readers"
	"github.com/smikelson75/parser/tokens"
)

type Scenerio struct {
	input    string
	expected struct {
		tokenType tokens.TokenType
		value     string
	}
}

type ScenerioFunction func(t *testing.T) []Scenerio

func TestingScenerios(t *testing.T) []ScenerioFunction {
	t.Helper()
	functions := []ScenerioFunction{
		LoadIdentifierScenerios,
	}

	return functions
}

func LoadIdentifierScenerios(t *testing.T) []Scenerio {
	t.Helper()

	return []Scenerio{
		{
			input: "Hello",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.IDENTIFIER,
				value:     "Hello",
			},
		},
		{
			input: "a",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.IDENTIFIER,
				value:     "a",
			},
		},
		{
			input: "A",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.IDENTIFIER,
				value:     "A",
			},
		},
		{
			input: "a1",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.IDENTIFIER,
				value:     "a1",
			},
		},
		{
			input: "a1b",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.IDENTIFIER,
				value:     "a1b",
			},
		},
		{
			input: "a1_",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.IDENTIFIER,
				value:     "a1",
			},
		},
		{
			input: " ",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.WHITESPACE,
				value:     " ",
			},
		},
		{
			input: "\t",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.WHITESPACE,
				value:     "\t",
			},
		},
		{
			input: "  \t",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.WHITESPACE,
				value:     "  \t",
			},
		},
		{
			input: "  \t\t  ",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.WHITESPACE,
				value:     "  \t\t  ",
			},
		},
		{
			input: "\n",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.NEWLINE,
				value:     "\n",
			},
		},
		{
			input: "1",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.NUMBER,
				value:     "1",
			},
		},
		{
			input: "1.1",
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.NUMBER,
				value:     "1.1",
			},
		},
		{
			input: `"Some String"`,
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.COMPLEXIDENTIFIER,
				value:     `"Some String"`,
			},
		},
		{
			input: `"Some ""String"`,
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.COMPLEXIDENTIFIER,
				value:     `"Some ""String"`,
			},
		},
		{
			input: `"Some ""String"""""`,
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.COMPLEXIDENTIFIER,
				value:     `"Some ""String"""""`,
			},
		},
		{
			input: `'Some String'`,
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.STRING,
				value:     `'Some String'`,
			},
		},
		{
			input: `'Some ''String'`,
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.STRING,
				value:     `'Some ''String'`,
			},
		},
		{
			input: `'Some ''String'''''`,
			expected: struct {
				tokenType tokens.TokenType
				value     string
			}{
				tokenType: tokens.STRING,
				value:     `'Some ''String'''''`,
			},
		},
		{

		},
	}
}

func TestHandlers(t *testing.T) {
	repository := handlers.Create()
	for _, factory := range TestingScenerios(t) {
		fmt.Println("Running Scenerio...")
		scenerios := factory(t)
		for _, scenerio := range scenerios {
			fmt.Printf("Input: %s; ", scenerio.input)
			reader := readers.FromString(scenerio.input, 80)
			err := reader.ReadTo('\n')

			if err != nil && err != io.EOF {
				t.Errorf("returned error: %v", err)
			}

			token, err := repository.Get(reader)
			if err != nil && err != io.EOF {
				t.Errorf("returned error: %v", err)
			}

			if token == nil {
				t.Errorf("expected return of Token, got nil")
				continue
			}

			if token.Value() != scenerio.expected.value || token.Type() != scenerio.expected.tokenType {
				t.Errorf("expected value %s of type %s, got %s of type %s", 
					scenerio.expected.value,
					scenerio.expected.tokenType.String(), 
					scenerio.input,
					token.Type().String())
			}

			fmt.Printf("Retrieved Token: %s; Type %s\n", token.Value(), token.Type().String())
		}

		fmt.Println("Scenerios Complete.")
	}
}
