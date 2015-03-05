package server

import (
	"encoding/json"
	"regexp"
	"strings"
)

type Rule struct {
	Regexp   string
	Parsed   *regexp.Regexp
	Into     string
	Category string
}

type Rules []*Rule

func LoadRules(definition string) (*Rules, error) {
	rules := &Rules{}
	if definition == "" {
		return rules, nil
	}

	if err := json.Unmarshal([]byte(definition), rules); err != nil {
		return &Rules{}, err
	}
	for _, rule := range *rules {
		if compiled, err := regexp.Compile(rule.Regexp); err != nil {
			return &Rules{}, err
		} else {
			rule.Parsed = compiled
		}
	}
	return rules, nil
}

func ApplyRules(rules *Rules, name string) (prefix string, category string) {
	prefix = strings.TrimPrefix(name, "/")
	for _, rule := range *rules {
		if rule.Parsed == nil {
			return "not parsed", ""
		}
		if rule.Parsed.MatchString(prefix) {
			if rule.Into != "" {
				prefix = rule.Parsed.ReplaceAllString(prefix, rule.Into)
			}
			return prefix, rule.Category
		}
	}
	return prefix, ""
}
