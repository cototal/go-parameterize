package parameterize

import (
	"strings"
)

func ToPasCamelCase(str string, camel bool) string {
	parameterized := Parameterize(str, '-')
	parts := strings.Split(parameterized, "-")
	runeList := make([]rune, 0, len(parameterized))
	for idx, part := range parts {
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

	return string(tokens)
}
