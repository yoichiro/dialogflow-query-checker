package config

type Definition struct {
	ClientAccessToken string `yaml:"clientAccessToken"`
	Tests []Test `yaml:"tests"`
}

type Test struct {
	Condition Condition `yaml:"condition"`
	Expect Expect `yaml:"expect"`
}

type Condition struct {
	Contexts []string `yaml:"contexts"`
	Query string `yaml:"query"`
}

type Expect struct {
	Action string `yaml:"action"`
	IntentName string `yaml:"intentName"`
	Parameters Parameter `yaml:"parameters"`
	Contexts []string `yaml:"contexts"`
	Speech string `yaml:"speech"`
}

type Parameter struct {
	Date string `yaml:"date"`
	Prefecture string `yaml:"prefecture"`
	Keyword string `yaml:"keyword"`
	Event string `yaml:"event"`
}
