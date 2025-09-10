package parameterize_test

import (
	"testing"

	"github.com/cototal/go-parameterize"
)

func TestParameterize(t *testing.T) {
	tests := []struct {
		input    string
		seprune  rune
		expected string
	}{
		{"Hello, World!", '-', "hello-world"},
		{"123abcDEF", '_', "123abcdef"},
		{"--abc--", '.', "abc"},
		{"abc123", '+', "abc123"},
		{"abc-def!@#", ' ', "abc def"},
	}

	for _, test := range tests {
		result := parameterize.Parameterize(test.input, test.seprune)
		if result != test.expected {
			t.Errorf("Parameterize(%q, %c) = %q; want %q", test.input, test.seprune, result, test.expected)
		}
	}
}

func TestFromCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		seprune  rune
		expected string
	}{
		{"PascalCase", '-', "pascal-case"},
		{"ID", '-', "id"},
		{"camelCase", '-', "camel-case"},
		{"testID", '_', "test_id"},
		{"exampleString", '+', "example+string"},
		{"alreadyHyphenated", ' ', "already hyphenated"},
		{"noChangesNeeded", '.', "no.changes.needed"},
		{"leadingUpperCase", '_', "leading_upper_case"},
		{"trailingUpperCase", '_', "trailing_upper_case"},
	}

	for _, test := range tests {
		result := parameterize.FromPasCamelCase(test.input, test.seprune)
		if result != test.expected {
			t.Errorf("FromCamelCase(%q, %c) = %q; want %q", test.input, test.seprune, result, test.expected)
		}
	}
}
