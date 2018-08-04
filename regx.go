package textbot

import (
	re "regexp"
	"strings"
	"unicode"
)

type Regx struct {
	*re.Regexp
}

// Regx is just shorthand for regexp.MustCompile. It's worth nothing,
// however, that all text input (args or repl) has white space trimmed
// before passed to any comparison, which simplifies the expressions
// needed. Case, however, has to be explicitly ignored if that is what
// you want (to preserve names and such in some responders).
//
// By the way, you REALLY want to use/learn regular expressions for
// creating the most "natural" responders because they allow so much
// more variation in your interactions.
func X(s string) *Regx {
	return &Regx{re.MustCompile(SpaceToRegx(s))}
}

func (r *Regx) M(s string) bool {
	return r.MatchString(s)
}

// CrunchSpace is the fastest possible method to crunch all unicode
// spaces into a single space, the first one detected.
func CrunchSpace(s string) string {
	sawother := true
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			if sawother {
				sawother = false
				return r
			}
			return -1
		}
		sawother = true
		return r
	}, s)
}

// SpaceToRegx converts any whitespace between the fields of the string
// into \s+ suitable for use to compile as a regular expression.
func SpaceToRegx(s string) string {
	return strings.Join(strings.Fields(s), "\\s+")
}
