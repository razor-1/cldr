package main

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// PluralGroup is a group of locales with the same plural rules.
type PluralGroup struct {
	Locales     string
	PluralRules []PluralRule
}

var titleCaser = cases.Title(language.AmericanEnglish)

// Name returns a unique name for this plural group.
func (pg *PluralGroup) Name() string {
	n := titleCaser.String(pg.Locales)
	return strings.ReplaceAll(n, " ", "")
}

// SplitLocales returns all the locales in the PluralGroup as a slice.
func (pg *PluralGroup) SplitLocales() []string {
	return strings.Split(pg.Locales, " ")
}

// PluralRule is the rule for a single plural form.
type PluralRule struct {
	Count string
	Rule  string
}

// CountTitle returns the title case of the PluralRule's count.
func (pr *PluralRule) CountTitle() string {
	return titleCaser.String(pr.Count)
}

// Condition returns the condition where the PluralRule applies.
func (pr *PluralRule) Condition() string {
	i := strings.Index(pr.Rule, "@")
	return pr.Rule[:i]
}

// Examples returns the integer and decimal exmaples for the PluralRule.
func (pr *PluralRule) Examples() (integer []string, decimal []string) {
	ex := strings.ReplaceAll(pr.Rule, ", …", "")
	ddelim := "@decimal"
	if i := strings.Index(ex, ddelim); i > 0 {
		dex := strings.TrimSpace(ex[i+len(ddelim):])
		decimal = strings.Split(dex, ", ")
		ex = ex[:i]
	}
	idelim := "@integer"
	if i := strings.Index(ex, idelim); i > 0 {
		iex := strings.TrimSpace(ex[i+len(idelim):])
		integer = strings.Split(iex, ", ")
	}
	return integer, decimal
}

// IntegerExamples returns the integer exmaples for the PluralRule.
func (pr *PluralRule) IntegerExamples() []string {
	integer, _ := pr.Examples()
	return integer
}

// DecimalExamples returns the decimal exmaples for the PluralRule.
func (pr *PluralRule) DecimalExamples() []string {
	_, decimal := pr.Examples()
	return decimal
}

var relationRegexp = regexp.MustCompile(`([nivwftce])(?:\s*%\s*([0-9]+))?\s*(!=|=)(.*)`)

// GoCondition converts the XML condition to valid Go code.
func (pr *PluralRule) GoCondition() string {
	ors := make([]string, 0)
	for _, and := range strings.Split(pr.Condition(), "or") {
		var ands []string
		for _, relation := range strings.Split(and, "and") {
			parts := relationRegexp.FindStringSubmatch(relation)
			if parts == nil {
				continue
			}
			lvar, lmod, op, rhs := titleCaser.String(parts[1]), parts[2], parts[3], strings.TrimSpace(parts[4])
			if op == "=" {
				op = "=="
			}
			// e is the deprecated form of c, but appears in the rules as of CLDR 46.
			if lvar == "E" {
				lvar = "C"
			}
			lvar = "ops." + lvar
			var rhor []string
			var rany []string
			for _, rh := range strings.Split(rhs, ",") {
				if parts := strings.Split(rh, ".."); len(parts) == 2 {
					from, to := parts[0], parts[1]
					//nolint:gocritic //ifElseChain is more readable here
					if lvar == "ops.N" {
						if lmod != "" {
							rhor = append(rhor, fmt.Sprintf("ops.NModInRange(%s, %s, %s)", lmod, from, to))
						} else {
							rhor = append(rhor, fmt.Sprintf("ops.NInRange(%s, %s)", from, to))
						}
					} else if lmod != "" {
						rhor = append(rhor, fmt.Sprintf("intInRange(%s %% %s, %s, %s)", lvar, lmod, from, to))
					} else {
						rhor = append(rhor, fmt.Sprintf("intInRange(%s, %s, %s)", lvar, from, to))
					}
				} else {
					rany = append(rany, rh)
				}
			}

			if len(rany) > 0 {
				rh := strings.Join(rany, ",")
				//nolint:gocritic //ifElseChain is more readable here
				if lvar == "ops.N" {
					if lmod != "" {
						rhor = append(rhor, fmt.Sprintf("ops.NModEqualsAny(%s, %s)", lmod, rh))
					} else {
						rhor = append(rhor, fmt.Sprintf("ops.NEqualsAny(%s)", rh))
					}
				} else if lmod != "" {
					rhor = append(rhor, fmt.Sprintf("intEqualsAny(%s %% %s, %s)", lvar, lmod, rh))
				} else {
					rhor = append(rhor, fmt.Sprintf("intEqualsAny(%s, %s)", lvar, rh))
				}
			}
			r := strings.Join(rhor, " || ")
			if len(rhor) > 1 {
				r = "(" + r + ")"
			}
			if op == "!=" {
				r = "!" + r
			}
			ands = append(ands, r)
		}
		ors = append(ors, strings.Join(ands, " && "))
	}
	return strings.Join(ors, " ||\n")
}
