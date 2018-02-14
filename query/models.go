package query

type Response struct {
	Result Result `json:"result"`
	Status Status `json:"status"`
}

type Result struct {
	Action string `json:"action"`
	Parameters Parameter `json:"parameters"`
	Metadata Metadata `json:"metadata"`
	Contexts []Context `json:"contexts"`
	Fulfillment Fulfillment `json:"fulfillment"`
}

type Status struct {
	Code int `json:"code"`
}

type Context struct {
	Name string `json:"name"`
}

type Parameter struct {
	Date string `json:"date"`
	Prefecture string `json:"prefecture"`
	Keyword string `json:"keyword"`
	Event string `json:"event"`
}

type Metadata struct {
	IntentName string `json:"intentName"`
}

type Fulfillment struct {
	Speech string `json:"speech"`
}
