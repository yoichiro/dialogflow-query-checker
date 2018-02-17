package config

import "strings"

type Definition struct {
	ClientAccessToken string `yaml:"clientAccessToken"`
	DefaultLanguage string `yaml:"defaultLanguage"`
	Tests []Test `yaml:"tests"`
}

type Test struct {
	Condition Condition `yaml:"condition"`
	Expect Expect `yaml:"expect"`
}

func (test *Test) CreatePrefix() string {
	if test.Condition.Contexts != nil {
		return strings.Join(test.Condition.Contexts, ",") + " " + test.Condition.Query
	} else {
		return test.Condition.Query
	}
}

type Condition struct {
	Contexts []string `yaml:"contexts"`
	Language string `yaml:"language"`
	Query string `yaml:"query"`
}

type Expect struct {
	Action string `yaml:"action"`
	IntentName string `yaml:"intentName"`
	Parameters Parameter `yaml:"parameters"`
	Contexts []string `yaml:"contexts"`
	Speech string `yaml:"speech"`
	Speeches []string `yaml:"speeches"`
}

type Parameter struct {
	Date string `yaml:"date"`
	Prefecture string `yaml:"prefecture"`
	Keyword string `yaml:"keyword"`
	Event string `yaml:"event"`
}
