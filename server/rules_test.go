package server

import (
	"testing"
)

func Test_ApplyRules_NoRule(t *testing.T) {
	rules := witRules(t, `[]`)
	ensure(t, rules, "name", "name", "")
}

func Test_ApplyRules_SingleRule(t *testing.T) {
	rules := witRules(t, `[{"regexp": "(.*)\\..*", "into": "prefix.$1", "category": "customer"}]`)
	ensure(t, rules, "name", "name", "")
	ensure(t, rules, "name.bla", "prefix.name", "customer")
}

func Test_ApplyRules_MultipleRules(t *testing.T) {
	rules := witRules(t, `[{"regexp": "first\\.(.*)", "into": "$1", "category": "first"}, {"regexp": "second\\.(.*)", "into": "$1", "category": "second"}]`)
	ensure(t, rules, "name", "name", "")
	ensure(t, rules, "first.name", "name", "first")
	ensure(t, rules, "second.name", "name", "second")
}

func witRules(t *testing.T, definitions string) *Rules {
	rules, err := loadRules(definitions)
	if err != nil {
		t.Fatalf("unexpected error on '%s': %#v", definitions, err)
	}
	return rules
}

func ensure(t *testing.T, rules *Rules, container, expectedPrefix, expectedCategory string) {
	prefix, category := applyRules(rules, container)
	if prefix != expectedPrefix {
		t.Errorf("applyRules failed:\nexpected %#v,\ngot      %#v", expectedPrefix, prefix)
	}
	if category != expectedCategory {
		t.Errorf("applyRules failed:\nexpected %#v,\ngot      %#v", expectedCategory, category)
	}
}
