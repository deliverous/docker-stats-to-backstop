package main

import (
	"encoding/json"
	"github.com/deliverous/docker-stats-to-backstop/translate/docker"
	"regexp"
	"strings"
)

type prefixRule struct {
	Regexp string
	Into   string
	parsed *regexp.Regexp
}

func loadPrefixRule(definition string) (*prefixRule, error) {
	if definition == "" {
		return nil, nil
	}

	s := prefixRule{}
	if err := json.Unmarshal([]byte(definition), &s); err != nil {
		return nil, err
	}
	if r, err := regexp.Compile(s.Regexp); err != nil {
		return nil, err
	} else {
		s.parsed = r
	}
	return &s, nil
}

func computePrefix(container *docker.Container, rule *prefixRule) string {
	prefix := container.Id
	if len(container.Names) > 0 {
		prefix = strings.TrimPrefix(container.Names[0], "/")
		if rule != nil {
			return rule.parsed.ReplaceAllString(prefix, rule.Into)
		}
	}
	return prefix
}
