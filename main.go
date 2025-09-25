package parameterize

import (
	"strings"
)

// Allows specification of case to case. Defaults to kebab
func P8ize(str string, from string, to string) string {
	var parameterized string
	switch from {
	case "pascal":
		parameterized = FromPasCamelCase(str, '-')
	case "camel":
		parameterized = FromPasCamelCase(str, '-')
	case "snake":
		parameterized = Parameterize(str, '_')
	default:
		parameterized = Parameterize(str, '-')
	}

	switch to {
	case "pascal":
		return ToPascalCase(parameterized)
	case "camel":
		return ToCamelCase(parameterized)
	case "snake":
		return ToSnakeCase(parameterized)
	default:
		return ToKebabCase(parameterized)
	}
}

// Converts a string to PascalCase or camelCase
func ToPasCamelCase(str string, camel bool) string {
	parameterized := Parameterize(str, '-')
	parts := strings.Split(parameterized, "-")
	runeList := make([]rune, 0, len(parameterized))
	for idx, part := range parts {
		// ID seems to generally be treated differently by convention
		if part == "id" {
			runeList = append(runeList, 'I')
			runeList = append(runeList, 'D')
			continue
		}
		for jdx, run := range part {
			if camel && idx == 0 {
				runeList = append(runeList, run)
				continue
			}
			if jdx == 0 {
				runeList = append(runeList, run-32)
			} else {
				runeList = append(runeList, run)
			}
		}
	}
	return string(runeList)

}

func IsCapitalRune(run rune) bool {
	return run >= 'A' && run <= 'Z'
}

func FromPasCamelCase(str string, seprune rune) string {
	tokens := make([]rune, 0, len(str))
	if strings.ToLower(str) == "id" {
		return "id"
	}
	for idx, run := range str {
		if idx == 0 && IsCapitalRune(run) {
			tokens = append(tokens, run+32)
			continue
		}
		if IsCapitalRune(run) {
			if run == 'I' && str[idx+1] == 'D' && idx == len(str)-2 {
				tokens = append(tokens, seprune)
				tokens = append(tokens, 'i')
				tokens = append(tokens, 'd')
				break
			}
			tokens = append(tokens, seprune)
			tokens = append(tokens, run+32)
			continue
		}
		tokens = append(tokens, run)
	}

	return string(tokens)
}

func ToPascalCase(str string) string {
	return ToPasCamelCase(str, false)
}

func ToCamelCase(str string) string {
	return ToPasCamelCase(str, true)
}

func ToKebabCase(str string) string {
	return Parameterize(str, '-')
}

func ToSnakeCase(str string) string {
	return Parameterize(str, '_')
}

// Removes special characters and separates each part of a string with seprune
// NOTE: Use FromPasCamelCase on strings that are already PascalCase or camelCase
func Parameterize(str string, seprune rune) string {
	tokens := make([]rune, 0, len(str))
	for idx, run := range str {
		if run >= '0' && run <= '9' {
			tokens = append(tokens, run)
			continue
		}
		if run >= 'a' && run <= 'z' {
			tokens = append(tokens, run)
			continue
		}
		if run >= 'A' && run <= 'Z' {
			tokens = append(tokens, run+32)
			continue
		}
		// Don't start string with a separator
		if idx == 0 || len(tokens) == 0 {
			continue
		}
		// Don't add another separator if the last rune was a separator
		if tokens[len(tokens)-1] == seprune {
			continue
		}
		tokens = append(tokens, seprune)
	}
	// This for loop uses the conditional to continuously check until 'tokens' doesn't end with a seprune
	for len(tokens) > 0 && tokens[len(tokens)-1] == seprune {
		tokens = tokens[:len(tokens)-1]
	}

	return string(tokens)
}
