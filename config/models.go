package config

import "strings"

type Definition struct {
	ClientAccessToken string `yaml:"clientAccessToken"`
	DefaultLanguage string `yaml:"defaultLanguage"`
	DefaultLocale string `yaml:"defaultLocale"`
	Tests []Test `yaml:"tests"`
	Environment Environment
}

type Environment struct {
	Debug bool
}

type Test struct {
	Condition Condition `yaml:"condition"`
	Expect Expect `yaml:"expect"`
}

func (test *Test) CreatePrefix() string {
	var contexts string
	if test.Condition.Contexts != nil {
		contexts = " (" + strings.Join(test.Condition.Contexts, ",") + ")"
	} else {
		contexts = ""
	}
	if test.Condition.Query != "" {
		return "Query: " + test.Condition.Query + contexts
	} else {
		return "Event: " + test.Condition.EventName + contexts
	}
}

type Condition struct {
	Contexts []string `yaml:"contexts"`
	Language string `yaml:"language"`
	Locale string `yaml:"locale"`
	Query string `yaml:"query"`
	EventName string `yaml:"eventName"`
	SessionId string `yaml:"sessionId"`
}

type Expect struct {
	Action string `yaml:"action"`
	IntentName string `yaml:"intentName"`
	Parameters Parameter `yaml:"parameters"`
	Contexts []string `yaml:"contexts"`
	Speech string `yaml:"speech"`
	Speeches []string `yaml:"speeches"`
	EndConversation string `yaml:"endConversation"`
}

type Parameter struct {
	Date string `yaml:"date"`
	Prefecture string `yaml:"prefecture"`
	Keyword string `yaml:"keyword"`
	Event string `yaml:"event"`
}
