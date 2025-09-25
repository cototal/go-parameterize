package parameterize_test

import (
	"testing"

	"github.com/cototal/go-parameterize"
)

type TestRecord struct {
	input    string
	expected string
}

func TestParameterize(t *testing.T) {
	tests := []TestRecord{
		{"Hello, World!", "hello-world"},
		{"123abcDEF", "123abcdef"},
		{"--abc--", "abc"},
		{"abc123", "abc123"},
		{"abc-def!@#", "abc def"},
		// Notice it doesn't work well with things already in PascalCase or camelCase
		{"TaskUpdateID", "taskupdateid"},
	}
	extraParams := []rune{
		'-',
		'_',
		'.',
		'+',
		' ',
		'_',
	}

	for idx, test := range tests {
		seprune := extraParams[idx]
		result := parameterize.Parameterize(test.input, seprune)
		if result != test.expected {
			t.Errorf("Parameterize(%q, %c) = %q; want %q", test.input, seprune, result, test.expected)
		}
	}
}

func TestToPasCamelCase(t *testing.T) {
	tests := []TestRecord{
		{"one two three", "OneTwoThree"},
		{"id", "ID"},
		{"one two three", "oneTwoThree"},
		{"id", "ID"},
		{"The - And the !.other", "TheAndTheOther"},
	}
	extraParams := []bool{
		false,
		false,
		true,
		true,
		false,
	}

	for idx, test := range tests {
		camel := extraParams[idx]
		result := parameterize.ToPasCamelCase(test.input, camel)
		if result != test.expected {
			t.Errorf("ToPasCamelCase(%q, %v) = %q; want %q", test.input, camel, result, test.expected)
		}
	}
}

func TestFromPasCamelCase(t *testing.T) {
	tests := []TestRecord{
		{"PascalCase", "pascal-case"},
		{"ID", "id"},
		{"camelCase", "camel-case"},
		{"testID", "test_id"},
		{"exampleString", "example+string"},
		{"alreadyHyphenated", "already hyphenated"},
		{"noChangesNeeded", "no.changes.needed"},
		{"leadingUpperCase", "leading_upper_case"},
		{"trailingUpperCase", "trailing_upper_case"},
	}

	extraParams := []rune{
		'-',
		'-',
		'-',
		'_',
		'+',
		' ',
		'.',
		'_',
		'_',
	}

	for idx, test := range tests {
		seprune := extraParams[idx]
		result := parameterize.FromPasCamelCase(test.input, seprune)
		if result != test.expected {
			t.Errorf("FromCamelCase(%q, %c) = %q; want %q", test.input, seprune, result, test.expected)
		}
	}
}

func TestP8ize(t *testing.T) {
	tests := [][]string{
		{"PascalCase", "pascal-case", "pascal", "kebab"},
		{"kebab-case", "KebabCase", "kebab", "pascal"},
		{"a random string", "aRandomString", "kebab", "camel"},
		{"A2#String()withSom%weRidStuf", "A2StringWithsomWeridstuf", "any", "pascal"},
		{"TaskUpdateID", "task_update_id", "pascal", "snake"},
	}

	for _, test := range tests {
		result := parameterize.P8ize(test[0], test[2], test[3])
		if result != test[1] {
			t.Errorf("P8ize(%q, %q, %q) = %q; want %q", test[0], test[2], test[3], result, test[1])
		}
	}
}
